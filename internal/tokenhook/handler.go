// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tokenhook

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	OriginStargate, OriginZone string
	AddAzpClaim, TraceRequests bool
}

func (h *Handler) handleError(w http.ResponseWriter, msg string, status int, err error) {
	log.Printf("%s: %v", msg, err)
	http.Error(w, msg, status)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// should handle a POST request with schema req.json as body
	// should return either 200 OK with schema output.json or 204 NO_CONTENT

	defer func() {
		// probably not needed, body is closed automatically
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed to close request body: %v", err)
		}
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, "Failed to read request body", http.StatusInternalServerError, err)
		return
	}

	if h.TraceRequests {
		log.Printf("Received hydra request: \n%s", body)
	}

	var req TokenHookRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.handleError(w, "Failed to unmarshal TokenHookRequest", http.StatusBadRequest, err)
		return
	}

	if req.Session == nil {
		// This should not happen
		h.handleError(w, "No Session in TokenHookRequest", http.StatusBadRequest, nil)
		return
	}
	sessionExtra := req.Session.sessionExtra()

	// Generate response

	// Access token additions
	if h.AddAzpClaim {
		// populate azp claim with client_id
		sessionExtra["azp"] = req.Request.ClientID
	}
	if len(h.OriginZone) != 0 {
		sessionExtra["originZone"] = h.OriginZone
	}
	if len(h.OriginStargate) != 0 {
		sessionExtra["originStargate"] = h.OriginStargate
	}

	tokenHookResponse := TokenHookResponse{
		Session: AcceptOAuth2ConsentRequestSession{
			AccessToken: sessionExtra,
		},
	}

	data, err := json.Marshal(tokenHookResponse)
	if err != nil {
		h.handleError(w, "Failed to marshal TokenHookResponse", http.StatusInternalServerError, err)
		return
	}
	if h.TraceRequests {
		log.Printf("Hydra request trasformed by token hook: \n%s", data)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to send TokenHookResponse: %v", err)
		return
	}
}

// returns the extra session data or an empty map
func (s *Session) sessionExtra() map[string]interface{} {
	if s.Extra == nil {
		return make(map[string]interface{})
	}
	return s.Extra
}

// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tokenhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenAddition(t *testing.T) {

	handler := &Handler{
		OriginStargate: "test-stargate",
		OriginZone:     "test-zone",
		AddAzpClaim:    true,
		TraceRequests:  true,
	}

	// Create a mock request
	reqBody := TokenHookRequest{
		Session: &Session{
			Extra: map[string]interface{}{},
			DefaultSession: &DefaultSession{
				Claims: &IDTokenClaims{
					Extra: map[string]interface{}{},
				},
			},
		},
		Request: Request{
			ClientID: "test-client",
		},
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rec, req)

	// Verify the response
	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", rec.Code)
	}

	var resp TokenHookResponse
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify AccessToken additions

	if resp.Session.AccessToken["azp"] != "test-client" {
		t.Errorf("Expected AccessToken 'azp' to be 'test-client', got %v", resp.Session.AccessToken["azp"])
	}
	if resp.Session.AccessToken["originZone"] != "test-zone" {
		t.Errorf("Expected AccessToken 'originZone' to be 'test-zone', got %v", resp.Session.AccessToken["originZone"])
	}
	if resp.Session.AccessToken["originStargate"] != "test-stargate" {
		t.Errorf("Expected AccessToken 'originStargate' to be 'test-stargate', got %v", resp.Session.AccessToken["originStargate"])
	}

}

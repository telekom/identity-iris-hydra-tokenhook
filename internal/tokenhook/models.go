// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tokenhook

import "time"

// The following datatypes are used to represent the data that is passed to the token hook.
// Copied from hydra source

type TokenHookRequest struct {
	// Session is the request's session..
	Session *Session `json:"session"`
	// Requester is a token endpoint's request context.
	Request Request `json:"request"`
}

type Session struct {
	*DefaultSession       `json:"id_token"`
	Extra                 map[string]interface{} `json:"extra"` // This field is transferred from response
	KID                   string                 `json:"kid"`
	ClientID              string                 `json:"client_id"`
	ConsentChallenge      string                 `json:"consent_challenge"`
	ExcludeNotBeforeClaim bool                   `json:"exclude_not_before_claim"`
	AllowedTopLevelClaims []string               `json:"allowed_top_level_claims"`
	MirrorTopLevelClaims  bool                   `json:"mirror_top_level_claims"`
}

// IDTokenSession is a session container for the id token
type DefaultSession struct {
	Claims    *IDTokenClaims          `json:"id_token_claims"`
	Headers   *Headers                `json:"headers"`
	ExpiresAt map[TokenType]time.Time `json:"expires_at"`
	Username  string                  `json:"username"`
	Subject   string                  `json:"subject"`
}

type TokenType string

// Headers is the jwt headers
type Headers struct {
	Extra map[string]interface{} `json:"extra"`
}

// IDTokenClaims represent the claims used in open id connect requests
type IDTokenClaims struct {
	JTI                                 string                 `json:"jti"`
	Issuer                              string                 `json:"iss"`
	Subject                             string                 `json:"sub"`
	Audience                            []string               `json:"aud"`
	Nonce                               string                 `json:"nonce"`
	ExpiresAt                           time.Time              `json:"exp"`
	IssuedAt                            time.Time              `json:"iat"`
	RequestedAt                         time.Time              `json:"rat"`
	AuthTime                            time.Time              `json:"auth_time"`
	AccessTokenHash                     string                 `json:"at_hash"`
	AuthenticationContextClassReference string                 `json:"acr"`
	AuthenticationMethodsReferences     []string               `json:"amr"`
	CodeHash                            string                 `json:"c_hash"`
	Extra                               map[string]interface{} `json:"ext"`
}

// Request is a token endpoint's request context.
//
// swagger:ignore
type Request struct {
	// ClientID is the identifier of the OAuth 2.0 client.
	ClientID string `json:"client_id"`
	// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
	GrantedScopes []string `json:"granted_scopes"`
	// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
	GrantedAudience []string `json:"granted_audience"`
	// GrantTypes is the requests grant types.
	GrantTypes []string `json:"grant_types"`
	// Payload is the requests payload.
	Payload map[string][]string `json:"payload"`
}

// TokenHookResponse is the response body received from the token hook.
//
// swagger:ignore
type TokenHookResponse struct {
	// Session is the session data returned by the hook.
	Session AcceptOAuth2ConsentRequestSession `json:"session"`
}

// Pass session data to a consent request.
type AcceptOAuth2ConsentRequestSession struct {
	// AccessToken sets session data for the access and refresh token, as well as any future tokens issued by the
	// refresh grant. Keep in mind that this data will be available to anyone performing OAuth 2.0 Challenge Introspection.
	// If only your services can perform OAuth 2.0 Challenge Introspection, this is usually fine. But if third parties
	// can access that endpoint as well, sensitive data from the session might be exposed to them. Use with care!
	AccessToken map[string]interface{} `json:"access_token"`

	// IDToken sets session data for the OpenID Connect ID token. Keep in mind that the session'id payloads are readable
	// by anyone that has access to the ID Challenge. Use with care!
	IDToken map[string]interface{} `json:"id_token"`
}

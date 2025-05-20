// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ory/graceful"
	cfg "github.com/telekom/identity-iris-hydra-tokenhook/internal/config"
	"github.com/telekom/identity-iris-hydra-tokenhook/internal/tokenhook"
)

func main() {
	port, exists := os.LookupEnv(cfg.TokenHookPort)
	if !exists {
		port = cfg.DefaultPort
	}
	debug := strings.EqualFold(os.Getenv(cfg.EnableDebug), "true")
	if debug {
		log.Print("Debug mode enabled, sensitive data might be logged")
	}
	// Set according to env var, if empty default is true
	addAzpClaim := strings.EqualFold(os.Getenv(cfg.AddAzpClaim), "true")

	handler := &tokenhook.Handler{
		os.Getenv(cfg.ClaimOriginStargate),
		os.Getenv(cfg.ClaimOriginZone),
		addAzpClaim,
		debug}

	tokenHookServer := graceful.WithDefaults(&http.Server{Addr: ":" + cfg.DefaultPort, Handler: handler})

	log.Printf("Starting token-hook server on port %s", port)
	if err := graceful.Graceful(tokenHookServer.ListenAndServe, tokenHookServer.Shutdown); err != nil {
		log.Fatalf("main: Failed to gracefully shutdown. Reason: %v", err)
	}
}

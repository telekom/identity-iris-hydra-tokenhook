<!--
SPDX-FileCopyrightText: 2025 Deutsche Telekom AG

SPDX-License-Identifier: Apache-2.0  
-->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0]

### Added

- Initial implementation of the `token-hook` server.
- Support for adding custom claims (`azp`, `originStargate`, `originZone`) to access tokens.
- Environment variable support for configuration.
- Dockerfile for building and running the application.
- Example `quickstart.yml` for integration with `iris-hydra`.
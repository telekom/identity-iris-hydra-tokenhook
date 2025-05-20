<!--
SPDX-FileCopyrightText: 2025 Deutsche Telekom AG

SPDX-License-Identifier: Apache-2.0  
-->

# Token Hook

The `token-hook` is an Ory Hydra "webhook", allowing you to modify access tokens.

## Features

- Add custom claims (`azp`, `originStargate`, `originZone`) to access tokens.
- Enable or disable debug mode for detailed logging.

## Environmental Variables

The following environment variables can be used to configure the `token-hook` service:

| Variable                    | Description                                                       | Default Value |
|-----------------------------|-------------------------------------------------------------------|---------------|
| `TOKEN_HOOK_PORT`           | The port on which the token hook server will listen.              | `4445`        |
| `CLAIM_SET_ORIGIN_STARGATE` | Value used for the `originStargate` claim.                        | Not set       |
| `CLAIM_SET_ORIGIN_ZONE`     | Value used for the `originZone` claim.                            | Not set       |
| `CLAIM_ADD_AZP`             | If set to `true`, adds the `azp` claim to the token.              | `true`        |
| `DEBUG`                     | If set to `true`, enables debug logging, including sensitive data | `false`       |

## Usage

## Usage

### Running Locally

1. **Build the Docker Image**
   Use the provided `Dockerfile` to build the `token-hook` image with the name `token-hook:latest`:
   ```bash
   docker build -t token-hook:latest .

2. Use the [quickstart-iris-hydra](quickstart-token-hook.yml) docker-compose file combined with iris-hydra
   [quickstart files](https://github.com/telekom/identity-iris-hydra?tab=readme-ov-file#build-and-run-locally).

```shell
docker compose -f quickstart.yml -f quickstart-iris-hydra.yml -f quickstart-tokenhook.yml up -d
```

To test the token hook, you can use the iris-hydra provided
[example](https://github.com/telekom/identity-iris-hydra?tab=readme-ov-file#running-tests) and verify
the additional claims are present in the token.
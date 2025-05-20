# SPDX-FileCopyrightText: 2025 Deutsche Telekom AG  1
#
# SPDX-License-Identifier: Apache-2.0

# Uses distroless base image, with a statically linked executable.
# CGO is disabled to ensure that the binary is statically linked.

ARG GO_IMAGE=golang:1.22
ARG DISTROLESS_IMAGE=gcr.io/distroless/static-debian12:debug-nonroot

FROM ${GO_IMAGE} AS builder

WORKDIR /go/src/token-hook

COPY . .
ENV CGO_ENABLED=0
RUN go build -o token-hook
RUN go test -v .
#########################

FROM ${DISTROLESS_IMAGE} AS runner

COPY --from=builder /go/src/token-hook/token-hook /usr/bin/token-hook
EXPOSE 4475

ENTRYPOINT ["token-hook"]

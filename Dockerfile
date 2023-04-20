# Build node feature discovery
FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.19-openshift-4.14 AS builder

WORKDIR /go/node-feature-discovery
COPY . .

# Do actual build
ARG VERSION=v0.10.0
ARG HOSTMOUNT_PREFIX=/host-
RUN make install VERSION=${VERSION} HOSTMOUNT_PREFIX=${HOSTMOUNT_PREFIX}

# Create full variant of the production image
FROM registry.ci.openshift.org/ocp/4.14:base

# Use more verbose logging of gRPC
ENV GRPC_GO_LOG_SEVERITY_LEVEL="INFO"

COPY --from=builder /go/node-feature-discovery/deployment/components/worker-config/nfd-worker.conf.example /etc/kubernetes/node-feature-discovery/nfd-worker.conf
COPY --from=builder /go/bin/* /usr/bin/

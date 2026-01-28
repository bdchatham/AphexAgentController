# AphexAgentController

Kubernetes controller for managing Agent custom resources in the Aphex platform.

## Overview

The AgentController reconciles `Agent` CRDs, managing:
- Model server deployments (vLLM)
- Model cache PersistentVolumeClaims
- Orchestrator deployments and services
- Integration with KnowledgeBase resources

## Dependencies

This controller depends on `AphexControllerRuntime` for shared types and utilities.

## Building

```bash
go mod tidy
go build ./...
```

## Running Locally

```bash
go run . --disable-leader-elect --development
```

## Docker Build

Build from the workspace root:

```bash
docker build -f AphexAgentController/Dockerfile -t agent-controller:latest .
```

## Deployment

Apply the manifests:

```bash
kubectl apply -f manifests/rbac.yaml
kubectl apply -f manifests/deployment.yaml
```

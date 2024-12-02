# Theia Workspace Garbage Collector
This operator is responsible for cleaning up the workspace of theia-ide pods.
It is a Kubernetes operator that cleans unused workspaces periodically.

## Configuration
The following values can be configured via environment variables:
- `K8S_NAMESPACE`: The namespace in which theia-ide pods are running. Default is `theia-prod`.
- `WORKSPACE_TTL`: The time to live of a workspace. Default is 2 weeks (1209600 000 000 000).
- `CHECK_INTERVAL`: The interval at which the operator checks for unused workspaces. Default is 1 hour (1800 000 000 000).

## Deployment
`helm upgrade --install theia-workspace-garbage-collector ./helm -f ./helm/values.yaml`
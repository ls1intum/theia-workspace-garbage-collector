# Theia Workspace Garbage Collector
This operator is responsible for cleaning up the workspace of theia-ide pods.
It is a Kubernetes operator that cleans unused workspaces periodically.

## Configuration
The operator's timing can be configured in the upper parts of `main.go`.

## Deployment
`helm upgrade --install theia-workspace-garbage-collector ./helm -f ./helm/values.yaml`
# Go Micro Gorm
Generated with

```
micro new github.com/bradford-hamilton/go-micro-gorm --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.webapp
- Type: srv
- Alias: go

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

## Usage

A Makefile is included for convenience

Run your service for development:
```
make run
```

Build the binary
```
make build-mac
```

or
```
make build-linux
```

Start it up
```
./go-micro-gorm
```

Build, tag, and push docker images with sha and latest tags
```
make docker
```

## Useful commands I'll forget
Micro CLI:

List services:
```
micro list services
```

Call a service directly:
```
micro call go.micro.srv.webapp Messaging.Call '{"name": "bradford"}'
```

Get health and endpoint IP:PORT for each service:
```
micro health go.micro.srv.webapp
```

Build and run service:
```
micro run service
```

Kubernetes commands
```
kc get services
kc get pods
kc describe deployments
kubectl describe pods {my-pod}

RESTARTING A POD (scale down to 0 then back to 1):
    kubectl scale deployment go-micro-gorm --replicas=0
    kubectl scale deployment go-micro-gorm --replicas=1
```
More: https://kubernetes.io/docs/reference/kubectl/cheatsheet

## Local kubernetes deployment:
Install etcd:
```
helm install go-micro-gorm --set customResources.createEtcdClusterCRD=true stable/etcd-operator
```

To later uninstall etcd:
```
helm delete go-micro-gorm
```

Install NATS:
```
kubectl apply -f https://github.com/nats-io/nats-operator/releases/latest/download/00-prereqs.yaml
```
```
kubectl apply -f https://github.com/nats-io/nats-operator/releases/latest/download/10-deployment.yaml
```

Apply/deploy:
```
kubectl apply -f k8s/deploy.yaml
```
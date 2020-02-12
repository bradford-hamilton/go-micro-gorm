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

- FQDN: go.micro.srv.messaging
- Type: srv
- Alias: go

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./go-micro-gorm
```

Build a docker image
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
micro call go.micro.srv.messaging Messaging.Call '{"name": "bradford"}'
```

Get health and endpoint IP:PORT for each service:
```
micro health go.micro.srv.messaging
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
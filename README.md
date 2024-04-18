# Demo - Blockchain RPC Proxy

## Introduction
This repository contains a demo showcasing how to create a Golang server to proxy requests to blockchain RPC Nodes.

### Use-cases
- Provides unified interface to customers accessing our RPC APIs while controlling list of available RPC Providers via proxy layer
- Can be used to perform custom routing, and load-balancing between available RPC Provider endpoints
- Can be used to integrate custom healthchecks to remove unhealthy nodes from list of providers

## Local Setup & Testing

### Prerequisite
- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)

### Option 01: Build & Run

Add required envs in `.env` file

```
APP_PORT=3000
RPC_PROVIDER=<VALID-RPC-PROVIDER-ENDPOINT>
```

Run server

```bash
$ cd proxy
$ go mod tidy
$ go run cmd/server/main.go
```

### Option 02: Run via docker

```bash
$ cd proxy
$ docker build . -t proxy:latest
$ docker run -d \
    -p 3000:3000 \
    -e APP_PORT=3000 \
    -e RPC_PROVIDER=<VALID-RPC-PROVIDER-ENDPOINT> \
    proxy:latest
```

### Test

```bash
$ curl --location 'http://localhost:3000/' \
        --header 'Content-Type: application/json' \
        --data '{"method":"eth_getBlockByNumber","params":["0x6C59B1",false],"id":1,"jsonrpc":"2.0"}'
```

## IaC Setup - Terraform

The `deploy` contains terraform codebase and it includes support to create:
- VPC
- ECS Cluster (Fargate)
- Task Definition
- ECS Service
- ALB 


### Prequisite:

- Set ECR Repo in `image` section of `deploy/containers/task.tpl.json`
- Update required envs in `00-variables.tf`

- Run terraform
```bash
$ cd proxy/deploy/

$ source ./envs/sandbox/source.rc
$ terraform plan
$ terraform apply
```

## TODO / Future Improvements
- Support for multiple RPC Providers
- Support for multiple blockchains
- Add custom healthcheck to remove nodes with lag from available healthy RPC node pool
- Autoscaling in ECS Service
- Custom prometheus metrics to denote api metrics

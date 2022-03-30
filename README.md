# Jump App Golang Backend (*Events*)

## Introduction

_Jump App Golang Backend_ is one of a set of microservices, named Jumps, developed to generate a microservice communication test tool in order to support multi hands-on and webinars around microservices in Kubernetes. 

In this case, this microservice is designed to integrate a event-based microservice for working with this kind of application communication design model.

## Prerequisites

It is required to meet the following requisites:

- Kafka server up and running
- Kafka topic created 
- Export variables (KAFKA_HOST and KAFKA_TOPIC)

```$bash
export KAFKA_HOST=a50e9747c108647d6aa9095ddd503628-1581487172.us-east-2.elb.amazonaws.com:9092
export KAFKA_TOPIC=my-topic
```

NOTE: Please visit the following [link](.kafka) for more information about Kafka

## Quick Start Jump App Golang Backend (*Events*)

Once the _Jump App Golang Backend_ project has been uploaded, it is required to execute the following process:

- Install Golang

```$bash
# MacOS
$ brew install golang

# Fedora
$ dnf install golang
```

- Build the App from the root folder

```$bash
$ make build
go build -o bin/back-golang-events -race cmd/main.go
```

- Execute the App

```$bash
$ make run
bin/golang-demo
2020/11/30 22:07:25 Starting server on :8442
```

## Golang Test

Regarding test, it is required execute next command:

```$bash
$ make test
```

## Build container image locally

This repository includes a Dockerfile in order to be able to build a container image. If it is required to generate locally this container image, please execute the following command:
    
```$bash
$ podman build .
```

## Author

Asier Cidon (@RedHat)
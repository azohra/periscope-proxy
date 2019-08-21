[![Go Report](https://goreportcard.com/badge/github.com/azohra/periscope-proxy)](https://goreportcard.com/badge/github.com/azohra/periscope-proxy)

# Periscope Proxy

![icon](./assets/periscope.png)

## What is this?

Dynamic, on-demand docker-as-a service on any compatible K8S cluster with istio.

This is the **proxy** which can be used in subsitution for localhost docker. It is meant to be used in conjunction with https://github.com/azohra/periscope/. The proxy's job is to simply: 

* negotiate with your cluster to spin up images of your choice
* tags local traffic with unique `X-State-ID` header that points to:
    * Service with exactly one Pod (with 1 or more containers)
* forwards all local traffic on `[PORT]` to cluster

## Use Cases:

* ephemeral tasks where a docker image is needed but your vm (docker-in-docker) does not support nested virtualization (therefore no docker-in-docker). 


## Requirements

* K8S 1.12+ 
* Istio 
* GCR

## Instructions

### To Get
`go get github.com/azohra/periscope-proxy/...`

### To Install 
`go install ./...`      

### To Run
`periscope-proxy port=[PORT] addr=[ADDR]`      

## Todo
* Makefile & install bin on go get (be better™)
* 2-way websocket between `periscope` and `periscope-proxy`
    * Graceful k8s `SIGTERM`
    * Real-time state
    * Support websocket traffic
    * HTTPS + WSS traffic with istio
* Helm chart

Made with ❤️ in Toronto, Canada
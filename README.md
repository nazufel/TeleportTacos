# Teleport Tacos

Teleport Tacos is demo app I wrote for [Scylla Summit 2023](https://github.com/nazufel/scylla-summit-2023). However, I do have plans for it to be an app used in future demos. Shout out to [Katan Coding](https://t.co/QaN1cAzDmu?amp=1) for the inspiration on the Hex Architecture in Go.

## About

Teleport Tacos has a whole backstory that I'll write up some other time. The heart of this is a Go gRPC server and client. The server exposes two RPC methods that will result in either a read or a write on the configured database. Currently, that database is [ScyllaDB](https://scylladb.com), but that can change depending on the demo, hence I used Hex Architecture.

## Requirements

This repo requires some tools to run.

* [Docker](https://www.docker.com/)
* [Docker Compose V2](https://docs.docker.com/compose/compose-file/)
* [go-task](https://taskfile.dev/installation/)

## Building the Images

There are go-task targets for build the images. 

```sh
task build-client && task build-server
```

## Running the App

There's a [docker-compose](./docker-compose.yaml) that will stand up the databases and will mount this local file system into the `client` and `server` contianers for development within the containers.

```sh
task up
```
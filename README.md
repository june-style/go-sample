# Go-sample
Go + gRPC with Clean Architecture

## Overview
I was born to learn Golang and Clean Architecture.
Good luck with this encounter!

![The Clean Architecture](.images/the_clean_architecture.png "The Clean Architecture")

## At the first startup
1. make dotenv
2. make compose-up
3. make login-devcontainer
4. make init
5. make db
6. make grpc-list

## Detail
The Go-API directory structure is shown below.

### Enterprise Bussiness Rules
- domain/dconfig
- domain/dcontext
- domain/derrors
- domain/entities
- domain/services

### Application Bussiness Rules
- application/interactors
- application/usecases

### Interface Adapters
- interface/controllers
- interface/gateways
- interface/interceptors
- interface/logs
- interface/repositories (Implementation)

### Frameworks & Drivers
- framework/protocol (gRPC)
- framework/registry (google/wire)

### Etc
- cmd
- tools

## Sample cURL commands
1. sign up
```
TODO
```
2. sign in
```
TODO
```
3. home get
```
TODO
```

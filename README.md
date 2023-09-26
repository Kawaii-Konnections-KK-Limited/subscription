# Subscription Microservice version 1.0.0.02

This is a microservice written in Go using the Gin framework for handling subscriptions.

## Overview

The microservice provides an API endpoint for getting a user's subscription information based on an auth token. It handles verifying the token, retrieving the user's subscriptions, and returning them.

## Files

- `api/` - Contains the API code
  - `handlers/` - Handlers for the API routes
    - `handler.go` - Main subscription handler
  - `middleware/` - Middleware for authentication
  - `routes.go` - Route definitions
- `cmd/` - Main entrypoint code
- `foreignusage/` - Example usage from another service 
- `utils/` - Utility functions
  - `utils.go` - Token verification and profile lookup functions
- `go.mod` - Go module definition
- `go.sum` - Go module checksums

## Usage

The `foreignusage/usage.go` file shows an example of how another service can initialize and use this microservice. It provides callback functions for the token verification and profile lookup.

The main API endpoint is:

```
GET /:token
```

It expects a valid auth token as a parameter. If verified successfully, it will look up the profiles for the user and return them in a simple plaintext format.

## Running

```
go run cmd/main.go
```

Will start the service on default port 8080.

## Importing as external

The `foreignusage/usage.go` file shows how another service can import and initialize the API from the subscription microservice package.

It imports the `api` package:

```go
import "github.com/Kawaii-Konnections-KK-Limited/subscription/api"
```

It then calls `api.InitRoutes` while passing callback functions for token verification and profile lookup and subscription update sender:

```go
func InitService(gupFunc func(token *string) *[]string, vtFunc func(token *string) bool ,SubUpdateSender func(token *string), certFile, keyFile *string) {

  api.InitRoutes(gupFunc, vtFunc).Run()

}
```

This allows the importing service to provide custom implementations for these functions that the subscription API relies on.

For example, the service could have its own way of looking up user profiles based on a token:

```go
gupFunc := func(token *string) *[]string {

  // Lookup user profiles based on token
  
  return &userProfiles 
}

vtFunc := func(token *string) bool {

  // Verify token is valid

  return tokenIsValid
}

SubUpdateSender := func(token *string) {

  // Send users token who updated their subscription

}
```

Then the service can initialize the API with these functions:

```go

InitService(gupFunc, vtFunc ,SubUpdateSender, &certFile, &keyFile) 
```

This initializes the API with routes and middleware, but with the custom verification and profile functions provided.

So in summary, the `api` package provides a simple way to initialize the subscription API routes while allowing importing services to customize the key dependent logic.

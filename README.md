# Subscription Microservice

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
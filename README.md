# AWS Lambda JWT Authorizer

Lambda function written in Go for validating a JWT token signed with a KMS key.
It uses the [matelang/jwt-go-aws-kms](https://github.com/matelang/jwt-go-aws-kms) library that provides an AWS KMS(Key Management Service) adapter to be used with the popular GoLang JWT library [golang-jwt/jwt-go](https://github.com/golang-jwt/jwt).

This Lambda can be used with AWS API Gateway as a [Lambda authorizer](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html) for HTTP APIs.
The Lambda function needs an environment variable `KMS_JWT_KEY_ID` for the KMS key ID used to sign the JWT token.

## Build

To build the executable:

```shell
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
```

`CGO_ENABLED=0` is set for creating a stable binary package for standard C library (`libc`) 

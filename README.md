# CrabStash backend repository

## How to run in dev mode
1. Run surreal db container with following command 
    docker run --rm --pull always -p 8000:8000 surrealdb/surrealdb:latest start --log trace --user root --pass root
2. Make sure you have built all microservices that you wanna test by using makefile
3. Run your microservices binaries
4. Run FE api gateway

## Testing gRPC Servers
If you wanna test just the gRPC servers we recommend using [evans](https://github.com/ktr0731/evans)

## Makefile usage
Use ```make help``` to get info on how to build your microservice
# CrabStash backend repository

## How to run in dev mode
1. Run surreal db container with following command 
    docker run --rm --pull always -p 8000:8000 surrealdb/surrealdb:latest start --log trace --user root --pass root
2. Make sure you have built all microservices that you wanna test by using makefile
3. Run your microservices binaries
4. Run FE api gateway


## Build image
Use ```make <microservice_name<``` to build Docker image

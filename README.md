# CrabStash backend repository

## How to run in dev mode
```docker-compose -f docker/dev-compose.yaml up```

## Build image
Use ```make <microservice_name>``` to build Docker image or ```make all``` to build all microservices images

## Importing surrealdata
To import clean data (i.e just db schema) run ```make importClean```, ```make importDemo``` will import demo data, that coresponds to what you can find inside the ```mock``` folder. ```make exportData``` will make a dump of your current db state

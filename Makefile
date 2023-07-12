.PHONY: api auth
project := api auth

all: $(project)

auth: $@
api: $@

$(project):
	docker build -t crabstash-$@ --build-arg MSNAME=$@ . --file docker/Dockerfile
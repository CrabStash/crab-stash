.PHONY: api auth warehouse user
project := api auth warehouse user

all: $(project)

auth: $@
api: $@
user: $@
warehouse: $@

$(project):
	docker build -t crabstash-$@ --build-arg MSNAME=$@ . --file docker/Dockerfile
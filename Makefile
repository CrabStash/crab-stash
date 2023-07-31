.PHONY: api auth warehouse
project := api auth warehouse

all: $(project)

auth: $@
api: $@
warehouse: $@

$(project):
	docker build -t crabstash-$@ --build-arg MSNAME=$@ . --file docker/Dockerfile
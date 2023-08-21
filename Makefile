.PHONY: api auth user
project := api auth user

all: $(project)

auth: $@
api: $@
user: $@

$(project):
	docker build -t crabstash-$@ --build-arg MSNAME=$@ . --file docker/Dockerfile
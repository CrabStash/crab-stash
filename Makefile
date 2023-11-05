.PHONY: api auth warehouse user core import exportData
project := api auth warehouse user core

importDemo: 
	surreal import --conn http://localhost:8000 --user root --pass root --ns crabstash --db data db/demoData.surql

exportData: 
	surreal export --conn http://localhost:8000 --user root --pass root --ns crabstash --db data db/dump.surql

importClean:
	surreal import --conn http://localhost:8000 --user root --pass root --ns crabstash --db data db/emptyDB.surql

all: $(project)

auth: $@
api: $@
user: $@
warehouse: $@
core: $@

$(project):
	docker build -t crabstash-$@ --build-arg MSNAME=$@ . --file docker/Dockerfile
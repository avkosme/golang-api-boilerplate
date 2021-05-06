start:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/local.yml --limit local
	docker-compose up -d app
stop:
	docker-compose stop
test:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/test.yml --limit local
	docker-compose up -d mongo-test
	docker-compose up -d redis
	docker-compose run --rm app go test -v ./...
	docker-compose stop
build:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/test.yml --limit local
	docker-compose run --rm app go build  -o ./bin ./...


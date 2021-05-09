start:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/local.yml --limit local
	docker-compose up -d app
stop:
	docker-compose stop
test:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/test.yml --limit local
	docker-compose run --rm app go build  -o ./bin ./...
	docker-compose up -d mongo-test
	docker-compose up -d redis
	docker-compose up -d app
	docker-compose exec -T app go test -count=1 -v ./...
	docker-compose stop
build:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/test.yml --limit local
	docker-compose run --rm app go build  -o ./bin ./...

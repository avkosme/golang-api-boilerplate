init:
	docker run --rm -v `pwd`:`pwd` -w `pwd`/compose williamyeh/ansible:alpine3 ansible-playbook -i hosts.yml local.yml -e @vars/common/deployments.yml -e @vars/common/routes.yml  -e @vars/common/services.yml
build:
	docker-compose run --rm app go build  -o ./bin ./...
start:
	docker-compose up -d app
stop:
	docker-compose stop
restart:
	docker-compose stop && docker-compose up -d
test:
	cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/test.yml --limit local
	docker-compose run --rm app go build  -o ./bin ./...
	docker-compose up -d mongo-test
	docker-compose up -d redis
	docker-compose up -d app
	docker-compose exec -T app go test -count=1 -v ./...
	docker-compose stop

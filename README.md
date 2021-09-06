# GoLang Api Boilerplate

## Installation, start
```
$ make init
$ cd compose && ansible-playbook -i hosts.yml configure.yml -e @vars/local.yml --limit local
$ make build
$ make start
```

```
$ curl -I --insecure --location --request POST 'https://localhost:8080'

$ curl --insecure --location --request POST 'https://localhost:8080' \
--header 'Content-Type: application/json' \
--data-raw '{"message": {"text": "You!", "update_id": 1234567890}}'
```

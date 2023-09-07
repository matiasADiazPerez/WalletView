include .env

app-deploy:
	set -a
	go mod vendor
	docker compose -f deploy/app-compose.yaml up --force-recreate -d
down:
	set -a
	docker compose -f deploy/app-compose.yaml down
destroy:
	set -a
	docker compose -f deploy/app-compose.yaml down
	docker rmi wallet_app:latest
vendor:
	go mod vendor
test:
	go clean -testcache
	go test -v ./...


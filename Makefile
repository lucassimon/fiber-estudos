PWD := $(shell pwd)
DB_USER := estudos
DB_PASSWORD := teste123
DB_NAME := fiberestudos
DB_PORT := 25432

start_postgres:
	docker start postgis && docker exec postgis service postgresql restart

migrateup:
	migrate -path $(PWD)/src/databases/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path $(PWD)/src/databases/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migrateupto:
	migrate -path $(PWD)/src/databases/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up $(VERSION)


compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(PWD)/build/api .

compress:
	@upx $(PWD)/build/api

build: compile compress

dev:
	@air -c .air.linux.conf -d

.PHONY: compile

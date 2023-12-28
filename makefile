ifneq (,$(wildcard ./.env))
include .env
export
endif

default: clean format
.PHONY: default

clean:
	@rm -rf dist/ node_modules/ \
	public/*.js public/*.css \
	views/**/*_templ.go && \
	templ generate && \
	go mod tidy && \
	bun install
.PHONY: clean

update:
	@go get -u ./... & \
	bun update
.PHONY: update

commitlint:
	@bun run commitlint
.PHONY: commitlint

format:
	@gofmt -s -l -w . && \
	bun run format
.PHONY: format

build:
	@templ generate && \
	go build -o ./dist/app ./cmd/app/ && \
	bun run build:js && \
	bun run build:css
.PHONY: build

docker:
	@docker build .
.PHONY: docker

dev: export APP_ENV=development
dev: export PORT=9876
dev:
	@bun concurrently -rk \
	"bun run dev:js" \
	"bun run dev:css" \
	"templ generate -watch" \
	"wgo run ./cmd/app/main.go"
.PHONY: dev

compose-up:
	@docker compose -f docker-compose.yml up -d
.PHONY: compose-up

compose-down:
	@docker compose -f docker-compose.yml down
.PHONY: compose-down

migrate-create:
	@migrate create -dir ./postgres/migrations -ext sql $(name)
.PHONY: migrate-create

migrate-up:
	@migrate -path ./postgres/migrations -database ${DB_URL} up 1
.PHONY: migrate-up

migrate-down:
	@migrate -path ./postgres/migrations -database ${DB_URL} down 1
.PHONY: migrate-down

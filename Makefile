dev:
	go run ./cmd/main.go

build: clean deps
	CGO_ENABLED=0 go build -v -o ./bin/fight_club ./cmd/main.go

deps:
	go mod tidy

clean:
	rm -rf ./bin

docker:
	docker build -t eugeniocunha/fight-club .

docker-push:
	docker buildx build --platform linux/amd64 --push --tag eugeniocunha/fight-club .

docker-down:
	docker compose down -v --remove-orphans

docker-dev: docker-down
	docker compose -f docker-compose.yml up --build

docker-up: docker-down
	docker compose -f docker-compose.yml up --build -d
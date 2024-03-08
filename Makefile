dev:
	go run ./cmd/main.go

build: clean deps
	CGO_ENABLED=0 go build -v -o ./bin/fight_club ./cmd/main.go

deps:
	go mod tidy

clean:
	rm -rf ./bin

docker:
	docker build -t eugenio-cunha/fight_club .

docker-push:
	docker buildx build --push --platform linux/arm/v7,linux/arm64/v8,linux/amd64 --tag eugeniocunha/fight_club .

docker-down:
	docker compose down -v --remove-orphans

docker-dev: docker-down
	docker compose -f docker-compose.yml up --build -d
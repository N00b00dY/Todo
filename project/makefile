FRONT_END_BINARY=frontEndServiceApp
DISTRIBUTOR_BINARY=distributorServiceApp
DB_BINARY=dbServiceApp


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_front build_distributor build_db
	@echo "Stopping docker images (if running...)"
	
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	#cd ../broker-service && go test -timeout 30s  ./... &&
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"



## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	#cd ../front-end && go test -timeout 30s  ./...
	cd ../front-end && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## build_db: builds the frone end binary
build_db:
	@echo "Building db service binary..."
	#cd ../db-service && go test -timeout 30s  ./...
	cd ../db-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DB_BINARY} ./cmd/api
	@echo "Done!"


## build_distributor: builds the frone end binary
build_distributor:
	@echo "Building distributor binary..."
	#cd ../distributor-service && go test -timeout 60s  ./...
	cd ../distributor-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DISTRIBUTOR_BINARY} ./cmd/api
	@echo "Done!"
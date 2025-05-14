include .env

up:
	@echo "Starting containners..."
	docker-compose up --build -d --remove-orphans 

down:
	@echo "Stoping containners..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/

start:
	@env MONGO_DB_USERNAME=${MONGO_DB_USERNAME} MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} ./${BINARY}

restart: build start
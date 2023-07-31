img-create:
	docker build --tag meli-challenge:alpha .

cont-create:
	docker run -d -p 8080:8080 --name meli-container meli-challenge:alpha

cont-remove:
	docker container stop meli-container
	docker container remove meli-container

# Build images
dc-build:
	docker-compose build

# Create and start containers
dc-up:
	docker-compose up -d

# Stop and destroy containers
dc-down:
	docker-compose down

dc-restart:
	docker-compose down
	docker-compose up -d

# Rebuild all
dc-rebuild:
	docker-compose down
	docker-compose build
	docker-compose up -d

mongo-shell:
	docker exec -it meli-mongodb mongosh

npm-install:
	docker exec -it meli-frontend npm install

npm-build:
	docker exec -it meli-frontend npm run build
DOCKERFILE_DEV=compose.dev.yml

run:
	docker-compose -f DOCKERFILE_DEV up

re:
	docker-compose -f DOCKERFILE_DEV up --build

IMAGE_VERSION := 1.x

.PHONY: docker-build
docker-build:
	docker build -f env/Dockerfile \
	             -t kamilsk/click:$(IMAGE_VERSION) \
	             -t kamilsk/click:latest \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-push
docker-push:
	docker push kamilsk/click:$(IMAGE_VERSION)
	docker push kamilsk/click:latest

.PHONY: docker-refresh
docker-refresh:
	docker images --all \
	| grep '^kamilsk\/click\s\+' \
	| awk '{print $$3}' \
	| xargs docker rmi -f &>/dev/null || true
	docker pull kamilsk/click:$(IMAGE_VERSION)



.PHONY: publish
publish: docker-build docker-push



.PHONY: docker-start
docker-start:
	docker run --rm -d \
	           --env-file env/.env.example \
	           --name click-dev \
	           --publish 8080:8080 \
	           --publish 8090:8090 \
	           --publish 8091:8091 \
	           kamilsk/click:$(IMAGE_VERSION)

.PHONY: docker-logs
docker-logs:
	docker logs -f click-dev

.PHONY: docker-stop
docker-stop:
	docker stop click-dev

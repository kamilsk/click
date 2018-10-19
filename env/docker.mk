IMAGE_VERSION := 3.x
PACKAGE       := github.com/kamilsk/click


.PHONY: docker-build
docker-build:
	docker build -f env/Dockerfile \
	             -t kamilsk/click:$(IMAGE_VERSION) \
	             -t kamilsk/click:latest \
	             -t quay.io/kamilsk/click:$(IMAGE_VERSION) \
	             -t quay.io/kamilsk/click:latest \
	             --build-arg PACKAGE=$(PACKAGE) \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-push
docker-push:
	docker push kamilsk/click:$(IMAGE_VERSION)
	docker push kamilsk/click:latest
	docker push quay.io/kamilsk/click:$(IMAGE_VERSION)
	docker push quay.io/kamilsk/click:latest

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
	           -p 8080:8080 \
	           -p 8090:8090 \
	           -p 8091:8091 \
	           -p 8092:8092 \
	           -p 8093:8093 \
	           kamilsk/click:$(IMAGE_VERSION) run --with-profiling --with-monitoring --with-grpc-gateway

.PHONY: docker-logs
docker-logs:
	docker logs -f click-dev

.PHONY: docker-stop
docker-stop:
	docker stop click-dev

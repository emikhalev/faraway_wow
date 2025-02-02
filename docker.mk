SERVERDIR := $(CURDIR)/cmd/wow
CLIENTDIR := $(CURDIR)/cmd/client

.PHONY: composer-up
composer-up:
	docker-compose up --build


.PHONY: build-server
build-server:
	docker build -t server-image -f $(SERVERDIR)/Dockerfile ./

.PHONY: run-server
run-server:
	docker run -d -p 52345:52345 --name server-container server-image

.PHONY: up-server
up-server: build-server run-server

.PHONY: stop-server
stop-server:
	docker stop server-container

.PHONY: clean-server
clean-containers:
	docker rm server-container

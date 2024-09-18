.PHONY: build
build:
	docker compose build

.PHONY: run
run:
	docker compose up -d

.PHONY: stop
stop:
	docker compose stop

.PHONY: clean
clean:
	docker compose rm --force --stop --volumes

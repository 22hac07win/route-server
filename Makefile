.PHONY: dev
dev:
	docker compose down
	docker compose up -d --build
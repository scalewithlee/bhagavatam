.PHONY: db-start

db-start:
	podman run --name bhagavatam-db \
	  -e POSTGRES_USER=bhagavatam \
	  -e POSTGRES_PASSWORD=password \
	  -e POSTGRES_DB=bhagavatam \
	  -p 5432:5432 \
	  -d postgres:15

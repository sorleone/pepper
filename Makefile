PROJECT_NAME = pepper

image:
	docker build -t $(PROJECT_NAME) .

run:
	docker run --name $(PROJECT_NAME) --rm -p 8080:8080 $(PROJECT_NAME)

dev:
	cd src && go run main.go

image:
	docker build -t pepper .

run:
	docker run --rm -p 8080:8080 pepper

dev:
	cd src && go run main.go

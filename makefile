generate:
	@oapi-codegen -config api/config.yml api/openapi.yml

docker-build:
	@docker build -t dafaque/mentbot:latest .

docker-push:
	@docker push dafaque/mentbot:latest

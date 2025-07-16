run:
	source .env && go run cmd/http/main.go -debug

build:
	docker build --no-cache -t orchestator:2 . 

buildx:
	docker buildx -t orches:arm32 --platform linux/arm/v7 .
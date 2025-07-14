buildx:
	docker build --no-cache -t facerecognition --platform linux/arm/v7 .

build:
	docker build --no-cache -t facerecognition2:8 .


run:
	docker run -it --entrypoint=/bin/bash facerecognition

run2:
	docker run -it --entrypoint=/bin/bash -p 8000:8000 facerecognition2 
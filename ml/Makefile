buildx:
	docker build --no-cache -t facerecognition --platform linux/arm/v7 .

build:
	docker build --no-cache -t facerecognition2:8 .


run:
	docker run -it --entrypoint=/bin/bash facerecognition

run2:
	docker run -it --entrypoint=/bin/bash -p 8000:8000 facerecognition2 


docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
docker buildx create --name arm-builder --use
docker buildx inspect --bootstrap
docker buildx build --platform linux/arm/v7 -t yourname/yourimage:arm32 --load .
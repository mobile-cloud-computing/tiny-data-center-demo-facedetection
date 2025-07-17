set-up:
	docker run --rm --privileged multiarch/qemu-user-static --reset -p yes \ 
	&& docker buildx create --name arm-builder --use \ 
	&& docker buildx inspect --bootstrap \ 

buildx-39-na:
	docker buildx build -f Dockerfile.py39 --platform linux/arm/v7 -t mlna:arm32 --load .

buildx-39-a:
	docker buildx build -f Dockerfile.py39-alpine --platform linux/arm/v7 -t mla:arm32 --load .

buildx-38-na:
	docker buildx build -f Dockerfile.py38 --platform linux/arm/v7 -t ml38na:arm32 --load .

buildx-38-a:
	docker buildx build -f Dockerfile.py38-alpine --platform linux/arm/v7 -t m38la:arm32 --load .
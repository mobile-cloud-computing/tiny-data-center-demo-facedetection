VERSION := $(if $(VERSION),$(VERSION),)

deploy:
	docker stack deploy --detach=false -c stack.yml demo-3

publish_ml:
	cd ml \
	&& docker buildx build --platform linux/arm/v7 --no-cache -t juangonzalout/arm7_cloudlet_ml:${VERSION} . \
	&& docker push juangonzalout/arm7_cloudlet_ml:${VERSION} \
	&& cd ..

publish_orchestator:
	cd orchestator \
	&& docker build --no-cache -t juangonzalout/arm7_cloudlet_orchestator:${VERSION} . \
	&& docker push juangonzalout/arm7_cloudlet_orchestator:${VERSION} \
	&& cd ..
run:
	docker rm minio && docker run -d -p 9100:9000 -p 9001:9001 \
		-e "MINIO_ACCESS_KEY=minioadmin" \
		-e "MINIO_SECRET_KEY=minioadmin" \
		-v ~/minio/data:/data \
		--name minio \
		minio:arm32 minio server /data --console-address ":9001"

build:
	docker build -t minio .

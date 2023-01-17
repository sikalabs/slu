IMAGE = sikalabs/install-slu.sikalabs.io
IMAGE_GITHUB = ghcr.io/sikalabs/install-slu.sikalabs.io

build-and-push:
	docker build --platform linux/amd64 -t ${IMAGE} -t ${IMAGE_GITHUB} .
	docker push ${IMAGE}
	docker push ${IMAGE_GITHUB}

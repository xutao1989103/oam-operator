TAG?=latest
VERSION?=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

build-pkg:
	go build  -o bin/oamOperator main.go

build-image:
	go build  -o bin/oamOperator main.go
	docker build -t xtaodocker/oam-operator:$(TAG) . -f Dockerfile

push:
	docker tag xtaodocker/oam-operator:$(TAG) xtaodocker/oam-operator:$(VERSION)
	docker push xtaodocker/oam-operator:$(VERSION)

generate-yaml:
	kustomize build artifacts/oam > artifacts/oam/oam.yaml

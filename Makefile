TAG?=latest
VERSION?=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

build-pkg:
	go build  -o bin/oamOperator main.go

build-image:
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  \
    		-ldflags "-s -w -X github.com/xutao1989103/oam-operator/pkg/version.REVISION=$${GIT_COMMIT}" \
    		-a -installsuffix cgo -o bin/oamOperator main.go
	docker build -t xtaodocker/oam-operator:$(TAG) . -f Dockerfile

push:
	docker tag xtaodocker/oam-operator:$(TAG) xtaodocker/oam-operator:$(VERSION)
	docker push xtaodocker/oam-operator:$(VERSION)

generate-yaml:
	kustomize build artifacts/oam > artifacts/oam/oam.yaml

apply:
	kubectl apply -f artifacts/oam/oam.yaml
module github.com/xutao1989103/oam-operator

go 1.13

require (
	github.com/MEDIGO/go-healthz v0.0.0-20160923151312-9b0725fef657
	github.com/imdario/mergo v0.3.9 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/sample-controller v0.18.1
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7
	k8s.io/api => k8s.io/api v0.0.0-20191109101512-6d4d1612ba53
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191109100837-dffb012825f2
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191109102209-3c0d1af94be5
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191109100332-a9a0d9c0b3aa
)

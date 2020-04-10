package versioned

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

type Clientset struct {
	*discovery.DiscoveryClient
}

func NewForConfig(config *rest.Config) (*Clientset, error) {
	return &Clientset{nil}, nil
}
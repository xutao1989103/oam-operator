package main

import (
	"errors"
	"flag"
	"github.com/MEDIGO/go-healthz"
	"github.com/xutao1989103/oam-operator/pkg/controller"
	"github.com/xutao1989103/oam-operator/pkg/signals"
	"github.com/xutao1989103/oam-operator/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"net/http"
	"time"
)

var (
	masterUrl  string
	kubeconfig string
)

func main() {

	klog.InitFlags(nil)

	flag.Parse()

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	if err = controller.NewController(kubeClient, stopCh).Run(stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}

	go startHealthCheckServer()

	<-stopCh
	klog.Info("Shutting down workers")
}

func startHealthCheckServer()  {
	healthz.Set("version", version.VERSION)

	healthz.Register("important_check", time.Second*5, func() error {
		return errors.New("fail fail fail")
	})

	http.Handle("/healthz", healthz.Handler())
	http.ListenAndServe(":8000", nil)
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterUrl, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}


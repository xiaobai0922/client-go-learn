package client_go_informer

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func ShareInformerDemo() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/bailu/.kube/config")
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//factory := informers.NewSharedInformerFactory(clientSet, 0)
	factory := informers.NewSharedInformerFactoryWithOptions(clientSet, 0, informers.WithNamespace("default"))
	informer := factory.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			println("add event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			println("update event")
		},
		DeleteFunc: func(obj interface{}) {
			println("delete eventls")
		},
	})

	stopChan := make(chan struct{})
	factory.Start(stopChan)
	<-stopChan

}

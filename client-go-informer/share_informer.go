package client_go_informer

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
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

	// 添加一个队列
	rateLimitingQueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			println("add event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				fmt.Printf("get key failed: %v\n", err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			println("update event")
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				fmt.Printf("get key failed: %v\n", err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		DeleteFunc: func(obj interface{}) {
			println("delete eventls")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				fmt.Printf("get key failed: %v\n", err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
	})

	stopChan := make(chan struct{})
	factory.Start(stopChan)
	<-stopChan

}

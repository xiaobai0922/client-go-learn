package client_go_client

import (
	"fmt"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func Clients() {
	// RESTClient 使用
	/*
		// config
		config, err := clientcmd.BuildConfigFromFlags("", "/Users/bailu/.kube/config")
		if err != nil {
			panic(err)
		}
		config.GroupVersion = &v1.SchemeGroupVersion
		config.NegotiatedSerializer = scheme.Codecs
		config.APIPath = "/api"

		// RESTClient
		restClient, err := rest.RESTClientFor(config)
		if err != nil {
			panic(err)
		}

		// get data
		pod := v1.Pod{}
		err = restClient.Get().Namespace("default").Resource("pods").Name("test").Do(context.TODO()).Into(&pod)
		if err != nil {
			panic(err)
		} else {
			print(pod.Name)
		}
	*/

	// Clientset 的使用
	// config
	/*
		config, err := clientcmd.BuildConfigFromFlags("", "/Users/bailu/.kube/config")
		if err != nil {
			panic(err)
		}

		// Clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		// data
		deployment, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), "test-deployment", v1.GetOptions{})
		if err != nil {
			panic(err)
		} else {
			print(deployment.Name)
		}
	*/

	// dynamicClient 的使用
	/*
		config, err := clientcmd.BuildConfigFromFlags("", "/Users/bailu/.kube/config")
		if err != nil {
			panic(err)
		}

		dynamicClient, err := dynamic.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		resource := dynamicClient.Resource(schema.GroupVersionResource{
			//Group:    "core", // pods类型的资源默认是core，但是这里是要省略的
			Version:  "v1",
			Resource: "pods",
		})

		pod, err := resource.Namespace("default").Get(context.TODO(), "test", v1.GetOptions{})
		if err != nil {
			panic(err)
		} else {
			print(pod.GetName())
		}
	*/

	// discoveryClient 的使用
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/bailu/.kube/config")
	if err != nil {
		panic(err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	apiResources, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		panic(err)
	}

	for _, apiResourceList := range apiResources {
		gv := apiResourceList.GroupVersion
		for _, resource := range apiResourceList.APIResources {
			fmt.Printf("Resource: %s, GroupVersion: %s\n", resource.Name, gv)
		}
	}
}

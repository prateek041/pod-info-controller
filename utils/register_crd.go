package main

import (
	"context"
	"fmt"

	controllertest "github.com/prateek041/controller-test.git"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/tools/clientcmd"
)

func RegisterCRDs() {
	// read the CRD from yaml file
	yamlFile, err := controllertest.CRDv1TetragonPod.ReadFile("config/crd/bases/test.prateeksingh.tech_podreplicators.yaml")
	if err != nil {
		panic(err)
	}

	// Create a Kubernetes client using the default kubeconfig file
	kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := apiextclientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// Create an empty CRD object
	crd := &apiextv1.CustomResourceDefinition{}

	// Unmarshal the YAML into the CRD object
	if err := yaml.Unmarshal(yamlFile, crd); err != nil {
		panic(err)
	}

	// Create the CRD using the Kubernetes client
	_, err = clientset.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("CRD created successfully")
}

func main() {
	RegisterCRDs()
}

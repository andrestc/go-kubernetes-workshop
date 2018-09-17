package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func runClient() error {
	// Create client config from ~/.kube/config file
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		return err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// START1 OMIT
	// List all pods on all namespaces
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("%d pods in the cluster\n", len(pods.Items))
	// END1 OMIT

	// START2 OMIT
	// Get one pod
	pod, err := clientset.CoreV1().Pods("default").Get("mypod", metav1.GetOptions{})
	if err != nil {
		return err
	}
	// END2 OMIT
	data, _ := json.MarshalIndent(pod, "", "  ")
	fmt.Printf("%s\n", data)

	return nil
}

func main() {
	err := runClient()
	if err != nil {
		panic(err)
	}
}

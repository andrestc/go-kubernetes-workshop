package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/golang/glog"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func runInformer() error {
	kubeConfigPath := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "kube config path")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfigPath)
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// POD INFORMER OMIT
	resyncInterval := 30 * time.Second
	informerFactory := informers.NewSharedInformerFactory(clientset, resyncInterval) // HL1
	podInformer := informerFactory.Core().V1().Pods()                                // HL2
	// POD INFORMER OMIT

	// POD HOOKS OMIT
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) { // HL
			fmt.Printf("add pod: %v\n", new.(*v1.Pod).Name)
		},
		UpdateFunc: func(old, new interface{}) { // HL
			fmt.Printf("update pod: %v\n", new.(*v1.Pod).Name)
		},
		DeleteFunc: func(obj interface{}) { // HL
			fmt.Printf("delete pod: %v\n", obj.(*v1.Pod).Name)
		},
	})
	go informerFactory.Start(wait.NeverStop) // HL
	// POD HOOKS OMIT

	for {
		time.Sleep(time.Second * 5)
		// POD LISTER OMIT
		pod, err := podInformer.Lister().Pods("default").Get("shell")
		// POD LISTER OMIT
		if err != nil {
			fmt.Println(fmt.Errorf("Error getting pod: %v", err))
			continue
		}
		fmt.Printf("Labels of pod `shell`: %v\n", pod.GetLabels())
	}

}

func main() {
	err := runInformer()
	if err != nil {
		panic(err)
	}
}

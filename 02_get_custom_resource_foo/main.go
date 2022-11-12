package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var defaultKubeConfigPath string
	if home := homedir.HomeDir(); home != "" {
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")
	flag.Parse()

	//k8s-apiのconfigを作成
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	//k8s-clientを作成
	clientset, _ := kubernetes.NewForConfig(config)
	pods, _ := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})

	fmt.Println("INDEX\tNAMESPACE\tNAME")
	for i, pod := range pods.Items {
		fmt.Printf("%d\t%s\t%s\n", i, pod.GetNamespace(), pod.GetName())
	}
}

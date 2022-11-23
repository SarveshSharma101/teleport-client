package k8client

import (
	"context"
	"flag"
	"log"
	"teleport-client/utility"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	edgev1 "github.com/difinative/Edge-MonitoringOperator/api/v1"
)

func GetClient() dynamic.Interface {

	log.Println("Loading kubeconfig file...")
	kubeconfig := flag.String(utility.KUBECONFIG, utility.KUBECONFIG_PATH, "path to load kubeconfig")
	log.Println("Loading the configs from kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Println("!!!!!!!!! Error while trying to load config file using kube config file !!!!!!!!!!!!")
		log.Print("!!! Error >>", err)
		log.Println("Loading config using service account")
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Println("!!!!!!!!! Error while trying to load config from service account !!!!!!!!!!!!")
			log.Print("!!! Error >>", err)
		}
	}

	log.Println("Initiating a dynamic k8 client")
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Println("!!!!!!!!! Error while trying to crete dynamic client !!!!!!!!!!!!")
		log.Println("!!! Error >> ", err)
	}

	return dynClient
}

func GetEdgeList(dynClient *dynamic.Interface) edgev1.EdgeList {

	resource := schema.GroupVersionResource{
		Group:    utility.EDGE_CR_GROUP,
		Version:  utility.EDGE_CR_VERSION,
		Resource: utility.EDGE_CR_RESOURCE,
	}
	log.Printf("Edge resource >> Group: %s, Version: %s, Kind: %s\n", resource.Group, resource.Version, resource.Resource)

	edgeList := edgev1.EdgeList{}

	namespace := utility.EDGE_CR_NAMESPACE
	log.Println("Edge resource namespace: ", namespace)

	unsList, err := (*dynClient).Resource(resource).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("!!!! Error while tring to get list of edge !!!!!")
		log.Printf("Edge resource >> Group: %s, Version: %s, Kind: %s\n", resource.Group, resource.Version, resource.Resource)
		log.Println("Edge resource namespace: ", namespace)
		log.Println("!!!! Error >>", err)
	}

	for _, ui := range unsList.Items {
		edge := edgev1.Edge{}
		runtime.DefaultUnstructuredConverter.FromUnstructured(ui.UnstructuredContent(), &edge)
		edgeList.Items = append(edgeList.Items, edge)
	}

	log.Println("Number of expected edges: ", len(edgeList.Items))
	return edgeList
}

func UpdateEdges(dynClient *dynamic.Interface, edgeList *[]edgev1.Edge) {
	namespace := utility.EDGE_CR_NAMESPACE
	log.Println("Edge resource namespace: ", namespace)

	resource := schema.GroupVersionResource{
		Group:    utility.EDGE_CR_GROUP,
		Version:  utility.EDGE_CR_VERSION,
		Resource: utility.EDGE_CR_RESOURCE,
	}
	log.Printf("Edge resource >> Group: %s, Version: %s, Kind: %s\n", resource.Group, resource.Version, resource.Resource)
	for _, edgeToUpdate := range *edgeList {
		edgeObject, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&edgeToUpdate)

		if err != nil {
			log.Printf("!!!! Error while trying to type cast the edge: %s to it's resource type !!!!\n", edgeToUpdate.Spec.Edgename)
			log.Println(err)
		}
		unStructEdge := unstructured.Unstructured{Object: edgeObject}
		e, err := (*dynClient).Resource(resource).Namespace(namespace).UpdateStatus(context.TODO(), &unStructEdge, metav1.UpdateOptions{})
		if err != nil {
			log.Printf("!!!! Error while trying to update the edge: %s !!!!\n", edgeToUpdate.Spec.Edgename)
			log.Println(err)
		} else {
			log.Printf("Edge: %s updated \n", e.GetName())
		}

	}
}

package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	datamodels "teleport-client/datamodel"
	k8client "teleport-client/k8-Client"
	"teleport-client/teleport"
	"teleport-client/utility"

	edgev1 "github.com/difinative/Edge-MonitoringOperator/api/v1"
	"github.com/gravitational/teleport/api/types"
)

func main() {

	log.Println("Setting background context...")
	ctx := context.Background()

	log.Println("Initiating teleport client....")
	teleportClt := teleport.GetTeleportClient(&ctx)
	log.Println("Teleport client created >> ", teleportClt)

	log.Println("Initiating K8 client....")
	dynClient := k8client.GetClient()
	log.Println("K8 client created >> ", dynClient)

	for {
		log.Println("-------------------------------------------------------------------------------------------------------------------")
		log.Println("Getting squirrel_tts value from env... ")
		timeToSleep, err := strconv.Atoi(os.Getenv("squirrel_tts"))
		if err != nil {
			log.Println("!!!!!!!!!!!! 'squirrel_tts' value is not a number. should represent the number of minutes !!!!!!!!!!!!!!")
			timeToSleep = 1
		}
		log.Println("tts value is >> ", timeToSleep)
		edgeDurations := teleport.TestTeleportSSH()

		log.Println("Getting list of edges from teleport ")
		actualEdgeList := teleport.GetAvailableNodeList(teleportClt, &ctx)

		log.Println("Getting list of edges K8 cluster")
		expectedEdgeList := k8client.GetEdgeList(&dynClient)

		log.Println("Validating the edge lists")
		items := &expectedEdgeList.Items
		validateTheNodes(items, actualEdgeList)
		updateTheInferenceServerUpdateDuration(edgeDurations, items)

		k8client.UpdateEdges(&dynClient, items)
	}

}

func validateTheNodes(expectedEdgeList *[]edgev1.Edge, actualEdgeList []types.Server) {

	isPresent := false

	for i := 0; i < len(*expectedEdgeList); i++ {
		isPresent = false
		edge := &(*expectedEdgeList)[i]
		for _, node := range actualEdgeList {
			if edge.Spec.Edgename == node.GetHostname() {
				isPresent = true
				if strings.EqualFold(strings.ToLower(edge.Status.Workingstatus), strings.ToLower(utility.EDGE_STATUS_DOWN)) || strings.EqualFold(strings.ToLower(edge.Status.Workingstatus), "") {
					edge.Status.Workingstatus = utility.EDGE_STATUS_UP
				}
				if node.GetAllLabels()["memoryAvailable"] != "" {
					edge.Status.AvailableMemory = node.GetAllLabels()["memoryAvailable"]
				}
			}
		}
		if !isPresent && strings.EqualFold(strings.ToLower(edge.Status.Workingstatus), strings.ToLower(utility.EDGE_STATUS_UP)) {
			edge.Status.Workingstatus = utility.EDGE_STATUS_DOWN
		}
	}
}

func updateTheInferenceServerUpdateDuration(edgeDurations datamodels.Stats, items *[]edgev1.Edge) {
	isPresent := false
	for i := 0; i < len(*items); i++ {
		edge := &(*items)[i]
		isPresent = false
		for _, edgeStatData := range edgeDurations.Stats {
			for k, es := range edgeStatData {
				if strings.EqualFold(strings.ToLower(k), strings.ToLower(edge.Spec.Edgename)) {
					isPresent = true
					duration := utility.GetLastUpdateDuration(es.FolderUpdateTime)
					edge.Status.InferenceServerLastUpdate = duration
					cmap := make(map[string]edgev1.Camera)
					// isCamerasPresent := false
					egc := edgev1.Camera{}
					for _, esc := range es.Camera {
						egc.Resolution = esc.Resolution
						cmap[esc.Name] = egc
					}
					edge.Status.Cameras = cmap
				}
			}
		}
		if !isPresent {
			log.Println("Edge: ", edge.Spec.Edgename, " image folder is not present in the inference server")
			edge.Status.InferenceServerLastUpdate = -100
		}
	}
}

package teleport

import (
	"context"
	"log"
	"teleport-client/utility"

	"github.com/gravitational/teleport/api/breaker"
	"github.com/gravitational/teleport/api/client"
	"github.com/gravitational/teleport/api/types"
	"google.golang.org/grpc"
)

func GetTeleportClient(ctx *context.Context, teleportConfig utility.Teleport) *client.Client {
	log.Println("Reading the config file")

	log.Println("Creating teleport client using adress and indentity key")
	clt, err := client.New((*ctx), client.Config{
		Addrs:                      []string{teleportConfig.Address},
		Credentials:                []client.Credentials{client.LoadIdentityFile(utility.IDENTITY_KEY_FILE)},
		Dialer:                     nil,
		DialOpts:                   []grpc.DialOption{},
		DialInBackground:           false,
		DialTimeout:                0,
		KeepAlivePeriod:            0,
		KeepAliveCount:             0,
		InsecureAddressDiscovery:   false,
		ALPNSNIAuthDialClusterName: "",
		CircuitBreakerConfig:       breaker.Config{},
		Context:                    nil,
	})
	if err != nil {
		log.Println("!!! Error while trying to create the teleport client: ")
		panic(err)
	}
	return clt
}

func GetAvailableNodeList(clt *client.Client, ctx *context.Context) []types.Server {
	log.Printf("Getting the list of edges from %s namespace \n", utility.TELEPORT_NAMESPACE)
	nodes, err := clt.GetNodes(*ctx, utility.TELEPORT_NAMESPACE)
	if err != nil {
		log.Println("!!!!! Error while trying to get the list of edges from teleport !!!!")
		panic(err)
	}
	log.Println("Number of active edges: ", len(nodes))
	return nodes

}

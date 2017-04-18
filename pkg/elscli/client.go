package elscli

import (
	"context"
	"github.com/galo/els-go/pkg/api"
	"github.com/prometheus/common/log"
)

func GetServiceInstanceByKey(client api.ElsClient, routingKey string) (*api.ServiceInstanceReponse, error) {
	req := &api.RoutingKeyRequest{routingKey}
	resp, err := client.GetServiceInstanceByKey(context.Background(), req)
	if err != nil {
		log.Fatalf("Error gettting routing jey: %v", err)
		return nil, err
	}

	log.Info("Rotung key %s and tags %s", resp.GetServiceUri(), resp.GetTags())
	return resp, nil
}

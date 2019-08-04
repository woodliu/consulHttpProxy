package util

import (
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"log"
	"os"
)

var Logger = log.New(os.Stdout,"[Consul]",log.LstdFlags)

func NewConsulAgentMetaData(t, dc string)(*api.Catalog, *api.Agent, *api.QueryOptions,error){
	//We don't use https scheme
	config := &api.Config{
		Address:t,
		Scheme:"http",
		Datacenter:dc,
		Transport:cleanhttp.DefaultPooledTransport(),
	}

	// Get a new client
	client, err := api.NewClient(config)
	if err != nil {
		return nil,nil,nil,err
	}

	return client.Catalog(), client.Agent(), &api.QueryOptions{Datacenter:dc, UseCache:false},nil
}

func GetServices(catalog *api.Catalog, service string, q *api.QueryOptions)([]*api.CatalogService,error){
	//Actually，here just get one service
	catalogServices,_,err := catalog.ServiceMultipleTags(service,[]string{service},q)
	if err != nil{
		return nil,err
	}

	return catalogServices,nil
}

func ServiceExist(service string, catalog *api.Catalog, queryOptions *api.QueryOptions)(bool,error){
	catalogServices,err := GetServices(catalog, service, queryOptions)
	if err != nil{
		return false,err
	}
	return len(catalogServices) != 0,nil
}

func DeleteService(service string,catalog *api.Catalog, queryOptions *api.QueryOptions,agent *api.Agent)error{
    return agent.ServiceDeregister(service)
}
package util

import (
	pb "consulHttpProxy/proto"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	Logger = log.New(os.Stdout,"[Consul]",log.LstdFlags)
	ConsulServerPort = "8500"
	DataCenter = "dc1"
)

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
	catalogServices,_,err := catalog.Service(service,"",q)
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

func genSysTags(tags []string)([]string,int){
	/* tags format [<scheme metrics_path team>,custom_tags...] */
	for k,tag := range tags{
		if strings.HasPrefix(tag,"<"){
			tag = strings.TrimLeft(tag,"<")
			tag = strings.TrimRight(tag,">")
			return strings.Split(tag," "),k
		}
	}

	return nil,0
}

func svcTagToAppInfo(team string, services map[string][]string)([]*pb.AppExporterInfo,error){
	var appInfos []*pb.AppExporterInfo
	for id,tags := range services{
		if len(tags) < 2{
			continue
		}

		/* tags format [<scheme metrics_path team>,custom_tags...] */
		if sysTags,k := genSysTags(tags);nil != sysTags{
			/* if this service not contain the right team, continue to next service */
			if team != sysTags[2]{
				continue
			}

			coustomTags := append(tags[:k], tags[k+1:]...)
			appInfos = append(appInfos,&pb.AppExporterInfo{Id: id,Scheme:sysTags[0],MetricPath: sysTags[1],Team: sysTags[2],Tags: coustomTags})
		}
	}

	return appInfos,nil
}

func ListServices(team string, catalog *api.Catalog, q *api.QueryOptions)(*pb.ListRespMsg,error){
	services,_,err := catalog.Services(q)
	if nil != err{
		log.Println(err)
		return nil,err
	}

	listResp,err := svcTagToAppInfo(team,services)
	if nil != err{
		return nil,err
	}

	return &pb.ListRespMsg{AppInfos: listResp},nil
}

func AddServices(agent *api.Agent, catalog *api.Catalog, queryOptions *api.QueryOptions, svcs []api.AgentServiceRegistration)error{
	if 0 == len(svcs) {
		Logger.Println("add 0 servives")
		return errors.New("err:add 0 services")
	}

	for _,service := range svcs{
		exsit,err := ServiceExist(service.Name,catalog,queryOptions)
		if nil != err{
			Logger.Println(err)
			return err
		}

		if exsit{
			Logger.Println("Service:",service.ID,"has existed")
		}else{
			//Register a service
			err = agent.ServiceRegister(&service)
			if err != nil {
				Logger.Println(err)
				return err
			}
		}
	}
	return nil
}

func RemoveServices(team string,service string,catalog *api.Catalog, queryOptions *api.QueryOptions,agent *api.Agent)error{
	//Get service info, one servicename just get one service instance
	catalogServices,err := GetServices(catalog,service,queryOptions)
	if nil != err{
		Logger.Println(err)
		return err
	}

	if 0 == len(catalogServices){
		errStr := fmt.Sprintf("No sunc service %s",service)
		Logger.Println(errStr)
		return errors.New(errStr)
	}else if 1 < len(catalogServices){
		errStr := fmt.Sprintf("Multi service %s",service)
		Logger.Println(errStr)
		return errors.New(errStr)
	}

	sysTags,_ := genSysTags(catalogServices[0].ServiceTags)
	if nil == sysTags || team != sysTags[2]{
		errStr := fmt.Sprintf("service %s team invalid",service)
		Logger.Println(errStr)
		return errors.New(errStr)
	}

	serverNode := catalogServices[0].Address + ":" + ConsulServerPort
	catalog,agent,_,err = NewConsulAgentMetaData(serverNode, DataCenter)
	if nil != err{
		Logger.Println(err)
		return err
	}

    return agent.ServiceDeregister(service)
}

func ExportServices(appInfos []*pb.AppExporterInfo, fileName string)error{
	var fileContent string = "[\n"
	for _,appInfo := range appInfos{
		// there can be only one service by a given name
		tags := "["
		for _, tag := range appInfo.Tags{
			tags = tags + "\"" + tag + "\","
		}
		tags = strings.TrimRight(tags,",")
		tags = tags+"]"
		fileContent = fileContent + "{"+"\"ID\":\"" + appInfo.Id +"\",\"Team\":\"" + appInfo.Team + "\",\"Tags\":" + tags + ",\"Scheme\":\"" + appInfo.Scheme + "\",\"MetricPath\":\"" + appInfo.MetricPath + "\"},\n"

	}
	fileContent = strings.TrimRight(fileContent,",\n")
	fileContent = fileContent+"\n]"

	return ioutil.WriteFile(fileName, []byte(fileContent),0666)
}
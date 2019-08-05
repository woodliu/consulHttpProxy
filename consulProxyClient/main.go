package main

import (
	"consulHttpProxy"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	h bool
	l bool
	a string
	d string
	dc string
	t string
	sn string
)


func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: consulProxyClient [options...] 

#List all services
consulProxyClient -l -t ${host}:${port}
#List service by name
consulProxyClient -l --sn=${serviceName} -t ${host}:${port}
#Add service
consulProxyClient -a ${regFile} -t ${host}:${port}
#Delete service
consulProxyClient -d ${serviceName} -t ${host}:${port}

Options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&l, "l", false, "list services")
	flag.StringVar(&d, "d", "", "delete a service by `servicename`")
	flag.StringVar(&dc, "dc", "dc1", "`datacenter`")
	flag.StringVar(&a, "a", "", "add a service by `filename`")
	flag.StringVar(&t, "t", "", "`target`,format: [${ip}:${port}] / [${dns_name}:${port}]")
	flag.StringVar(&sn, "sn", "", "consul `servicename`,used for getting service by name")
	flag.Usage =usage
	flag.Parse()

	//Help info
	if h{
		usage()
		return
	}

	if t == ""{
		util.Logger.Println("Must specify a target,format like: [${ip}:${port}] / [${dns_name}:${port}]")
		return
	}

	catalog,agent,queryOptons,err := util.NewConsulAgentMetaData(t,dc)
	if nil != err{
		util.Logger.Println(err)
		return
	}
	//List services
	if l {
		if "" == sn{
			listServices(catalog,queryOptons)
		}else{
			listServicesByTags(catalog, sn, queryOptons)
			return
		}
	}

	//Add service
	if a !="" && !strings.HasPrefix(a,"-"){
		addService(agent,catalog,queryOptons,a)
		return
	}

	//Delete service
	if d !="" && !strings.HasPrefix(d,"-"){
		err = util.DeleteService(d,catalog,queryOptons,agent)
		if nil != err{
			util.Logger.Println(err)
		}
		return
	}
}

func listServices(catalog *api.Catalog, q *api.QueryOptions){
	services,_,err := catalog.Services(q)
	if nil != err{
		log.Println(err)
		return
	}

	//Delete service created by consul
	delete(services,"consul")

	title := fmt.Sprintf("%-25s%s\n","ServiceName","ServiceTags")
	fmt.Println(title,"----------------------------------------------------------------------------------------------------")
	for k,v := range services{
		fmt.Printf("%-25s%s\n",k,v)
	}
}

func listServicesByTags(catalog *api.Catalog, service string, q *api.QueryOptions){
	catalogServices,err := util.GetServices(catalog, service ,q)
	if nil != err || 0 == len(catalogServices){
		util.Logger.Println(err)
		return
	}

	title := fmt.Sprintf("%-15s%-20s%-20s%-23s%s\n","Node","ServiceName","ServiceID","ServiceAddress:Port","ServiceTags")
	fmt.Println(title,"----------------------------------------------------------------------------------------------------")
	for _,v := range catalogServices{
		tagsStr := fmt.Sprintf("%s\n",v.ServiceTags)
		fmt.Printf("%-15s%-20s%-20s%s:%-10d%s\n",v.Node,v.ServiceName,v.ServiceID,v.ServiceAddress,v.ServicePort,tagsStr)
	}
}

func addService(agent *api.Agent, catalog *api.Catalog, queryOptions *api.QueryOptions, f string){
	body,err := ioutil.ReadFile(f)
	if err != nil {
		util.Logger.Println(err)
		return
	}

	service := &api.AgentServiceRegistration{}
	err = json.Unmarshal(body,service)
	if err != nil {
		util.Logger.Println(err)
		return
	}

	//Set service.id equal to service.name, we add and delete service by service name. Service register file like below
	service.ID = service.Name
	/*
		{
		  "Name": "test-exporter",
		  "Tags": [
			"primary",
			"v1"
		  ],
		  "Address": "10.10.10.1",
		  "Port": 3333,
		  "EnableTagOverride": false,
		  "Tags": ["aa,bb"]
		}
	 */

	 //Append a tags named service.Name to help get service detail info
	 service.Tags = append(service.Tags,service.Name)

	//Preventing adding service with same name
	exsit,err1 := util.ServiceExist(service.Name,catalog,queryOptions)
	if nil != err1{
		util.Logger.Println(err1)
		return
	}
	if exsit{
		util.Logger.Println("Service:",service.Name,"has existed")
		return
	}

	 //Register a service
	err = agent.ServiceRegister(service)
	if err != nil {
		util.Logger.Println(err)
	}
}

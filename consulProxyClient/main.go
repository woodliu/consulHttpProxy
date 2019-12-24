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
	"strconv"
	"strings"
)

var (
	h bool
	l bool
	a string
	d string
	e string
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
#Exporter the services
consulProxyClient -e ${fileName} -t ${host}:${port}

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
	flag.StringVar(&e, "e", "", "export a service by `filename`")
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

	//Exporter service, default export services in dc1
	if e !="" && !strings.HasPrefix(e,"-"){
		err = exportServices(catalog ,queryOptons,e)
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


func exportServices(catalog *api.Catalog, q *api.QueryOptions, fileName string)error{
	services,_,err := catalog.Services(q)
	if nil != err{
		return err
	}

	//Delete service created by consul
	delete(services,"consul")

	var fileContent string = "[\n"
	for serviceName,_ := range services{
		// there can be only one service by a given name
		catalogServices,_ := util.GetServices(catalog, serviceName ,q)
		tags := "["
		for _, tag := range catalogServices[0].ServiceTags{
			tags = tags + "\"" + tag + "\","
		}
		tags = strings.TrimRight(tags,",")
		tags = tags+"]"
		fileContent = fileContent + "{"+"\"ID\":\"" + catalogServices[0].ServiceID +"\",\"Name\":\"" + catalogServices[0].ServiceName + "\",\"Address\":\"" + catalogServices[0].Address + "\",\"Port\":" + strconv.Itoa(catalogServices[0].ServicePort) + ",\"EnableTagOverride\":false" + ",\"Tags\":" + tags + "},\n"

	}
	fileContent = strings.TrimRight(fileContent,",\n")
	fileContent = fileContent+"\n]"
	err = ioutil.WriteFile(fileName, []byte(fileContent),0666)
	return err
}


func addService(agent *api.Agent, catalog *api.Catalog, queryOptions *api.QueryOptions, f string){
	body,err := ioutil.ReadFile(f)
	if err != nil {
		util.Logger.Println(err)
		return
	}

	var svcs []api.AgentServiceRegistration
	err = json.Unmarshal(body,&svcs)
	if err != nil {
		util.Logger.Println(err)
		return
	}

	for _,service := range svcs{
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
		err = agent.ServiceRegister(&service)
		if err != nil {
			util.Logger.Println(err)
		}
	}
}

package main

import (
	"consulHttpProxy"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var logger = log.New(os.Stdout,"[ConsulProxy]",log.LstdFlags)

var (
	consulServiceAddr = "127.0.0.1:8500"
	agentProxyAddr = "0.0.0.0:"+util.GetEnv("CONSUL_SERVER_PORT","8700")
)
const (
	deleteHostHeader = true
	keepHostHeader   = false
)

func copyHeaders(dst, src http.Header, keepDestHeaders bool) {
	if !keepDestHeaders {
		for k := range dst {
			dst.Del(k)
		}
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

func deleteConsulService(url *url.URL){
	target := consulServiceAddr
	serviceName := url.Path[strings.LastIndex(url.Path,"/")+1:]
	dataCenter := url.RawQuery[strings.LastIndex(url.RawQuery,"=")+1:]
	//Get meta data from local consul server
	catalog,agent,queryOptons,err := util.NewConsulAgentMetaData(target,dataCenter)
	if nil != err{
		util.Logger.Println(err)
		return
	}

	//Get service info, one servicename just get one service instance
	catalogServices,err := util.GetServices(catalog,serviceName,queryOptons)
	if nil != err{
		util.Logger.Println(err)
		return
	}

	if 0 == len(catalogServices){
		util.Logger.Println("Not such service:",serviceName)
		return
	}

	serviceNode := catalogServices[0].Address
	catalog,agent,queryOptons,err = util.NewConsulAgentMetaData(serviceNode,dataCenter)
	if nil != err{
		util.Logger.Println(err)
		return
	}
	err = util.DeleteService(serviceName,catalog,queryOptons,agent)
	if nil != err{
		util.Logger.Println(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Host = consulServiceAddr
		req.URL.Scheme = "http"
		//Create a new Request to consul server
		req,err := http.NewRequest(req.Method,req.URL.String(),req.Body)
		if err != nil {
			logger.Println(err)
			return
		}

		//Proecss the delete event
		if strings.Contains(req.URL.String(),"v1/agent/service/deregister"){
			deleteConsulService(req.URL)
		}

		//Do request to consul server
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.Println(err)
			return
		}

		//Copy consul server Response header
		copyHeaders(w.Header(), res.Header, keepHostHeader)
		w.WriteHeader(res.StatusCode)
		_, err = io.Copy(w, res.Body)
		if err != nil {
			logger.Printf("Error while sending a response for the '/' path: %v", err)
		}
		defer res.Body.Close()
	})

	//Listen and server on agen---tProxyAddr
	util.Logger.Println("Start Listen on:",agentProxyAddr)
	http.ListenAndServe(agentProxyAddr, nil)
}




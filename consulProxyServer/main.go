package main

import (
	"consulHttpProxy"
	pb "consulHttpProxy/proto"
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var logger = log.New(os.Stdout,"[ConsulProxy]",log.LstdFlags)

var (
	Address = kingpin.Flag("address","consul proxy address").Default("0.0.0.0").String()
	Port = kingpin.Flag("port","consul proxy port").Default("8700").String()

	ConsulServerAddr = "127.0.0.1:"+ util.ConsulServerPort
	//ConsulServerAddr = "consul-reg-prometheus.uatapps.opsocp.csvw.com"
)

type server struct {}

func (s *server)ListRequest(ctx context.Context, in *pb.ListReqMsg) (*pb.ListRespMsg, error){
	team := in.GetTeam()
	catalog,_,queryOptons,_ := util.NewConsulAgentMetaData(ConsulServerAddr,util.DataCenter)
	listResp,err := util.ListServices(team, catalog, queryOptons)

	return listResp,err
}

func (s *server)AddRequest(ctx context.Context,in *pb.AddReqMsg) (*pb.RespResult, error){
	var scheme,metrics_path,team string
	var svcReg []api.AgentServiceRegistration

	catalog,agent,queryOptons,_ := util.NewConsulAgentMetaData(ConsulServerAddr,util.DataCenter)
	svcs := in.GetAppInfos()
	for _,svc := range svcs{
		/* save metadata to tags [<scheme metrics_path team>,custom_tags...] */
		var tags []string

		team = svc.Team
		if "" == svc.Scheme{
			scheme = "http"
		}else{
			scheme = svc.Scheme
		}

		if "" == svc.MetricPath{
			metrics_path = "/metrics"
		}else {
			metrics_path = svc.MetricPath
		}

		sysStr := "<" + strings.Join([]string{scheme,metrics_path,team}," ") + ">"

		tags = append(tags,sysStr)
		tags = append(tags,svc.Tags...)

		svcID := strings.Split(svc.Id,":")
		if 2 != len(svcID){
			errStr := fmt.Sprintf("Target:[%s] format error!",svc.Id)
			return nil,errors.New(errStr)
		}

		address,p := svcID[0],svcID[1]
		port,_ := strconv.Atoi(p)
		svcReg = append(svcReg,api.AgentServiceRegistration{ID: svc.Id,Name: svc.Id,Address: address,Port: port,Tags: tags})
	}

	err := util.AddServices(agent, catalog, queryOptons, svcReg)
	if nil == err{
		return &pb.RespResult{Ret: 0},nil
	}
	return nil,err
}

func (s *server)UpdateRequest(ctx context.Context,in *pb.UpdateReqMsg)(*pb.RespResult, error){
	return nil,nil
}

func (s *server)RemoveRequest(ctx context.Context,in *pb.RemoveReqMsg)(*pb.RespResult, error){
	team := in.GetTeam()
	removeIds := in.GetServiceId()

	catalog,agent,queryOptons,_ := util.NewConsulAgentMetaData(ConsulServerAddr,util.DataCenter)
	for _,id := range removeIds{
		if err := util.RemoveServices(team,id,catalog, queryOptons,agent);nil != err{
			util.Logger.Println(err)
			return nil,err
		}
	}

	return nil,nil
}
func main() {
	kingpin.Parse()
	AgentProxyAddr := *Address + ":" + *Port

	util.Logger.Println("Agent listen on:",AgentProxyAddr)
	lis, err := net.Listen("tcp", AgentProxyAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRequestServer(s, &server{})

	s.Serve(lis)
}



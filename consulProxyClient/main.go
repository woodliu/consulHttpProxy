package main

import (
	"consulHttpProxy"
	pb "consulHttpProxy/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var(
	team = kingpin.Flag("team","specify a team").Required().String()
	target = kingpin.Flag("target","consul proxy target, format: 127.0.0.1:8700 ").Required().String()

	list = kingpin.Command("list", "list services for specify team")

	update = kingpin.Command("update", "update services for specify team")
	updateFile = update.Flag("file","file contain services configuration").Required().ExistingFile()

	add = kingpin.Command("add", "add new services for specify team")
	addFile = add.Flag("file","file contain services configuration").Required().ExistingFile()

	remove = kingpin.Command("remove", "remove service by service id")
	removeServiceIds = remove.Flag("serviceid","specify service id to remove").Required().Strings()

	export = kingpin.Command("export", "export service by servicename")
	exportFile = export.Flag("file", "directory to export services").Required().String()
)

type registerInfo struct {
	ID string          `json:",omitempty"`
	Team string      `json:",omitempty"`
	Tags []string      `json:",omitempty"`
	Scheme string      `json:",omitempty"`
	MetricPath string  `json:",omitempty"`
}

func isTeamUnique(regInfo []registerInfo)bool{
	if 0 == len(regInfo){
		return true
	}
	team := regInfo[0].Team
	teamRegInfo := 0
	for _,info := range regInfo{
		if info.Team == team{
			teamRegInfo ++
		}
	}

	return teamRegInfo == len(regInfo)
}

func main() {
	switch kingpin.Parse() {
	case"list":
		conn, err := grpc.Dial(*target, grpc.WithInsecure())
		if err != nil {
			util.Logger.Fatalln(err)
		}

		defer conn.Close()
		c := pb.NewRequestClient(conn)

		resp,err := c.ListRequest(context.Background(), &pb.ListReqMsg{Team:*team})
		if nil != err {
			util.Logger.Fatalln(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Scheme","Id","MetricPath","Tags"})
		for _, v := range resp.AppInfos {
			table.Append([]string{v.Scheme,v.Id,v.MetricPath,fmt.Sprint(v.Tags)})
		}
		table.Render()
	case "add":
		body,err := ioutil.ReadFile(*addFile)
		if err != nil {
			util.Logger.Fatalln(err)
			return
		}

		var svcs []registerInfo
		err = json.Unmarshal(body,&svcs)
		if err != nil {
			util.Logger.Fatalln(err)
			return
		}

		if !isTeamUnique(svcs){
			util.Logger.Fatalln("Team Must be unique!")
			return
		}

		var appExportInfo []*pb.AppExporterInfo
		for _,svc := range svcs{
			if "" == svc.Team{
				util.Logger.Fatalln("Must specify a team")
				return
			}

			if 0 == len(svc.Tags){
				util.Logger.Fatalln("Must specify some tags")
				return
			}

			appExportInfo = append(appExportInfo,&pb.AppExporterInfo{Id:svc.ID,Team: svc.Team,Tags:svc.Tags,Scheme: svc.Scheme,MetricPath: svc.MetricPath})
		}

		conn, err := grpc.Dial(*target, grpc.WithInsecure())
		if err != nil {
			util.Logger.Fatalln(err)
		}

		defer conn.Close()
		c := pb.NewRequestClient(conn)

		_,err = c.AddRequest(context.Background(), &pb.AddReqMsg{Team:svcs[0].Team, AppInfos: appExportInfo})
		if nil != err {
			util.Logger.Fatalln(err)
			return
		}

		util.Logger.Println("Done!")
	case "remove":
		conn, err := grpc.Dial(*target, grpc.WithInsecure())
		if err != nil {
			util.Logger.Fatalln(err)
		}

		defer conn.Close()
		c := pb.NewRequestClient(conn)

		_,err = c.RemoveRequest(context.Background(), &pb.RemoveReqMsg{Team: *team,ServiceId: *removeServiceIds})
		if nil != err{
			util.Logger.Fatalln(err)
			return
		}
		util.Logger.Println("Done!")
	case "update":
		//TODO
	case "export":
		conn, err := grpc.Dial(*target, grpc.WithInsecure())
		if err != nil {
			util.Logger.Fatalln(err)
		}

		defer conn.Close()
		c := pb.NewRequestClient(conn)

		resp,err := c.ListRequest(context.Background(), &pb.ListReqMsg{Team:*team})
		if nil != err {
			util.Logger.Fatalln(err)
			return
		}
		appInfos := resp.AppInfos
		util.ExportServices(appInfos, *exportFile)

	default:
		util.Logger.Fatalln("WRONG COMMAND!")
	}
}
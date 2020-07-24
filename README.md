# ConsulHttpProxy

A proxy for handle consul services behind a load balance or gateway,mainly for register prometheus exporters.

#### Build

Gen `.proto.go`

```
protoc.exe -I proto\ proto\register.proto --go_out=plugins=grpc:proto
```

`consulProxyClient` contains the proxy client，and `consulProxyServer` contains the proxy server.

If you use alpine image to build the docker image, use the command below:

```shell
# GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o consulProxyServer main.go
```

#### Deploy

Deploy sequence：

1. [consul-pvc.yaml](https://github.com/woodliu/consulHttpProxy/blob/master/deploy/consul-pvc.yaml)
2. [consul-service.yaml](https://github.com/woodliu/consulHttpProxy/blob/master/deploy/consul-service.yaml)
3. [consul-statefulset.yaml](https://github.com/woodliu/consulHttpProxy/blob/master/deploy/consul-statefulset.yaml)

#### What is this tool work for?

When use load balance/gateway to register a service, the load balance will choice one consul server to store the service. When delete a service, the load balance/gateway may choice the wrong consul server according to the balance strategy. This tool will make the right decision to do all the operations, include add/list/remove/export consul services.

#### WorkFlow

The work flow like below, consul proxy server works as a sidecar container with consul server. It will deliver the client event to the right consul server. Here, it use the `service ID` as the unique Key for each service, when register a service, there will be only one server store it.

You can build a `consulProxyClient` to add/list/remove/export service

![](./image/workflow.png)

### Client usage

```shell
>consulProxyClient.exe --help
usage: consulProxyClient.exe --team=TEAM --target=TARGET [<flags>] <command> [<args> ...]

Flags:
  --help           Show context-sensitive help (also try --help-long and
                   --help-man).
  --team=TEAM      specify a team
  --target=TARGET  consul proxy target, format: 127.0.0.1:8700

Commands:
  help [<command>...]
    Show help.

  list
    list services for specify team

  update --file=FILE
    update services for specify team

  add --file=FILE
    add new services for specify team

  remove --serviceid=SERVICEID
    remove service by service id

  export --file=FILE
    export service by servicename
```

```shell
# add exporters
consulProxyClient.exe --team=test --target=10.160.243.237:8700 add --file=reg.json
# find exporters of specify team
consulProxyClient.exe --team=test --target=10.160.243.237:8700 list
# remove one exporter of specify team
consulProxyClient.exe --team=test --target=10.160.243.237:8700 remove --serviceid=3.3.3.3:9100
# remove exporters of specify team
consulProxyClient.exe --team=test --target=10.160.243.237:8700 remove --serviceid=1.1.1.1:9100 --serviceid=3.3.3.3:9100 --serviceid=4.4.4.4:6666
```

You can use a file `reg.json` to add service to consul

```json
[
{"ID":"11.11.11.11:9100","Team":"private","Tags":["test1,test11"]},
{"ID":"22.22.22.22:8888","Team":"private","Tags":["test2"],"Scheme":"https","MetricPath":"/path/metrics"},
{"ID":"33.33.33.33:9100","Team":"private","Tags":["test3,test33,test333"]}
]
```

If you want change the exporter metricpath, you need to relabel the prometheus metric `__metrics_path__`.

```yaml
      relabel_configs:
        - source_labels: ['__meta_consul_tags']
          regex: ',<.*>,(.*),'
          replacement: '${1}'
          target_label: service_tags
        - source_labels: ['__meta_consul_service']
          regex: '^consul$'
          action: drop
        - source_labels:  ['__meta_consul_service_id']
          target_label: service_id
        - source_labels:  ['__meta_consul_service_address']
          target_label: service_address
        - source_labels: [__meta_consul_tags]
          regex:  ',<(.*) (.*) (.*)>.*'
          replacement: '${1}'
          target_label: __scheme__
        - source_labels: [__meta_consul_tags]
          regex:  ',<(.*) (.*) (.*)>.*'
          replacement: '${2}'
          target_label: __metrics_path__
        - source_labels: [__meta_consul_tags]
          regex:  ',<(.*) (.*) (.*)>.*'
          replacement: '${3}'
          target_label: team
```

### Server usage

```shell
>consulProxyServer.exe --help
usage: consulProxyServer.exe [<flags>]

Flags:
  --help               Show context-sensitive help (also try --help-long and
                       --help-man).
  --address="0.0.0.0"  consul proxy address
  --port="8700"        consul proxy port
```

#### constraint

- This tool used for `HTTP` scheme not `HTTPs`
- You can specify the `https` `Scheme` in reg.json, but it won't work, because this tool hasn't support TLS or other authentication.

Because this tool base on grpc, if using openshift, the openshift route may not support grpc, should use nodeport to redirect the flow.

```shell
# cat consul-proxy-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: consul-proxy
  labels:
    name: consul-proxy
spec:
  type: NodePort
  ports:
    - name: consul-proxy
      port: 8700
      nodePort: 30001
  selector:
    app: consul
```
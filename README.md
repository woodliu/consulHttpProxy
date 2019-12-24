# ConsulHttpProxy

A proxy for handle consul services behind a load balance or gateway.

#### Build

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

When use load balance/gateway to register a service, the load balance will choice one consul server to store the service. When delete a service, the load balance/gateway may choice the wrong consul server according to the balance strategy.

#### WorkFlow

The work flow like below, consul proxy server works as a sidecar container with consul server. It will deliver the service `deregister` event to the right consul server. Here, it use the serviceName as the unique Key for each service, when register a service, there will be only one server store it.

You can build a `consulProxyClient` to add/list/delete service

![](./image/workflow.png)

### usage

```shell
# ./consulClient --help

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
  -a filename
        add a service by filename
  -d servicename
        delete a service by servicename
  -dc datacenter
        datacenter (default "dc1")
  -e filename
        export a service by filename
  -h    help
  -l    list services
  -sn servicename
        consul servicename,used for getting service by name
  -t target
        target,format: [${ip}:${port}] / [${dns_name}:${port}]
```

#### constraint

- This tool used for `HTTP` scheme not `HTTPs`
- Pods should communicate with each other(Eg. overlay network)

# go-prometheus-exporter

This repository implements a dummy api server as prometheus exporter 

See more on [Medium](https://medium.com/@ejhsu/build-a-monitoring-dashboard-by-prometheus-grafana-741a7d949ec2)

## Prerequisite

- [direnv](https://direnv.net/)

## Before running

- run `go get -x -v github.com/golang/dep/cmd/dep` at root folder

- set GOPATH to root folder `export GOPATH=$(pwd)`

- set PATH to include GOPATH `export PATH=${PATH}:${GOPATH}/bin`

- cd to `src/api` and run `dep ensure`


## Run Server

- `go run server.go exporter.go`


src/api内是生产数据的service
*start.sh是启动prometheus服务端和grafana
config内是prometheus和grafana对应的配置文件
test内为发送给生产者的调用

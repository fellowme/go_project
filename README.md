# go_project

## 执行vendor

- go mod tidy
- go mod vendor

## protobuf_file

    protobuf 文件 执行 sh grpc.sh  生成 pb.go文件

## 安装docker es

    docker run -d --name es -p 9200:9200 -p 9300:9300 -e ES_JAVA_OPTS="-Xms512m -Xmx512m" -e "discovery.type=single-node" 2bd69c322e98

## 安装docker  kibana

    docker run --name kibana -e ELASTICSEARCH_HOSTS=http://172.17.0.3:9200 -p 5601:5601 -d kibana:7.4.2

## 安装docker pulsar

    docker run -it -p 6650:6650  -p 8080:8080 --mount source=pulsardata,target=/pulsar/data --mount source=pulsarconf,target=/pulsar/conf apachepulsar/pulsar:2.9.1 bin/pulsar standalone
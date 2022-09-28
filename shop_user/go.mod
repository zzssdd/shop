module shop_user

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/micro/go-log v0.1.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	github.com/micro/micro/v3 v3.12.3
	github.com/patrickmn/go-cache v2.1.0+incompatible
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.8
)

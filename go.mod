module github.com/smile-im/auth-service

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/micro-kit/micro-common => ../../micro-kit/micro-common
	github.com/micro-kit/microkit => ../../micro-kit/microkit
	github.com/smile-im/microkit-client/client/auth => ../microkit-client/client/auth
	github.com/smile-im/microkit-client/proto/authpb => ../microkit-client/proto/authpb
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0 // grpc对etcd依赖问题
)

require (
	code.aliyun.com/eduzhixin-dev/trade-service v1.60.5
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/micro-kit/micro-common v0.0.0-00010101000000-000000000000
	github.com/micro-kit/microkit v0.0.0-20191109115319-3481ce33bba7
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pingplusplus/pingpp-go v0.0.0-20190620100124-1b3eb9f381ba // indirect
	github.com/prometheus/client_golang v1.9.0 // indirect
	github.com/smile-im/microkit-client/client/auth v0.0.0-00010101000000-000000000000
	github.com/smile-im/microkit-client/proto/authpb v0.0.0-00010101000000-000000000000
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/grpc v1.34.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/gorm v1.20.9
)

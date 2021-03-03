module hellocase

go 1.15

require (
	github.com/micro/micro/v3 v3.0.0
	github.com/asim/nitro/v3 v3.4.2
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/asim/nitro/v3 => github.com/micro/go-micro/v3 v3.4.2

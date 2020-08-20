module go-rpc-happysooner-file

go 1.14

require (
	github.com/facebook/ent v0.4.0
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v3 v3.0.0-alpha.0.20200812115214-1fa3ac5599eb
	github.com/unliar/happysooner-proto v1.6.5
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

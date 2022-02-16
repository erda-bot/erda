// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: mysql.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/orchestrator/addon/mysql/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.orchestrator.addon.mysql",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                 = reflect.TypeOf((*Client)(nil)).Elem()
	addonMySQLServiceClientType = reflect.TypeOf((*pb.AddonMySQLServiceClient)(nil)).Elem()
	addonMySQLServiceServerType = reflect.TypeOf((*pb.AddonMySQLServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.orchestrator.addon.mysql-client":
		return p.client
	case "erda.orchestrator.addon.mysql.AddonMySQLService":
		return &addonMySQLServiceWrapper{client: p.client.AddonMySQLService(), opts: opts}
	case "erda.orchestrator.addon.mysql.AddonMySQLService.client":
		return p.client.AddonMySQLService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case addonMySQLServiceClientType:
		return p.client.AddonMySQLService()
	case addonMySQLServiceServerType:
		return &addonMySQLServiceWrapper{client: p.client.AddonMySQLService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.orchestrator.addon.mysql-client", &servicehub.Spec{
		Services: []string{
			"erda.orchestrator.addon.mysql.AddonMySQLService",
			"erda.orchestrator.addon.mysql.AddonMySQLService.client",
			"erda.orchestrator.addon.mysql-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			addonMySQLServiceClientType,
			// server types
			addonMySQLServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
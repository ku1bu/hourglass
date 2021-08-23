package grpcv1

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/ku1bu/hourglass/api/hourglass/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func TestRegister(t *testing.T) {
	conn, err := grpc.DialInsecure(context.TODO(), grpc.WithEndpoint("localhost:9000"))
	if err != nil {
		t.Fatal(err)
	}

	cli := v1.NewProviderClient(conn)
	resp, err := cli.Register(context.Background(), &v1.ProviderRegisterRequest{
		Capacity:   1000,
		Name:       "111112",
		MinClients: 1,
	})

	eee, ok := err.(*v1.Errorgg)
	if ok {
		fmt.Println(eee.Code)
		fmt.Println(eee.Message)
	} else {
		fmt.Println(">>>>>>")
	}

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

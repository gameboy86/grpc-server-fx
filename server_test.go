package grpcserverfx

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	mocks "github.com/gameboy86/grpc-server-fx/mocks/github.com/gameboy86/grpc-server-fx"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/test/bufconn"
)

type TestLog struct{}

func (t *TestLog) Logf(string, ...interface{})   {}
func (t *TestLog) Errorf(string, ...interface{}) {}
func (t *TestLog) FailNow()                      {}

func NewConfigurerMock() *mocks.GRPCServerConfigurer {
	return new(mocks.GRPCServerConfigurer)
}

func TestGrpcServerFxModule(t *testing.T) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	configurer := func() *mocks.GRPCServerConfigurer {
		c := new(mocks.GRPCServerConfigurer)
		c.On("GRPCServerReflection").Return(true)
		c.On("GRPCServerPort").Return(8000).Times(1)
		return c
	}
	startCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	app := fx.New(
		Module,
		fx.Decorate(
			func() net.Listener {
				return lis
			},
		),
		fx.Provide(
			fx.Annotate(
				configurer,
				fx.As(new(GRPCServerConfigurer)),
			),
		),
	)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("123 Call client...")
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return lis.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}
	fmt.Println("Call client...")
	in := grpc_health_v1.HealthCheckRequest{}
	client := grpc_health_v1.NewHealthClient(conn)
	response, err := client.Check(startCtx, &in)
	fmt.Println(response, err)
	stopCtx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

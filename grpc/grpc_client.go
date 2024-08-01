package grpc

import (
	"fmt"

	"github.com/anhhuy1010/cms-menu/config"
	"google.golang.org/grpc"
)

type GrpcService struct {
	UserConnect *grpc.ClientConn
}

var grpcService *GrpcService

func (sv *GrpcService) NewService() (*GrpcService, error) {
	if grpcService == nil {
		cfg := config.GetConfig()
		grpcArr := cfg.GetStringMap("grpc")
		for k, _ := range grpcArr {
			hostName := fmt.Sprintf(`grpc.%s.host`, k)
			portName := fmt.Sprintf(`grpc.%s.port`, k)
			host := cfg.GetString(hostName)
			port := fmt.Sprintf("%v", cfg.GetString(portName))

			conn, err := Connect(host, port)
			if err != nil {
				return nil, err
			}

			if k == "user" {
				sv.UserConnect = conn
			}
		}
		grpcService = sv
	}
	return grpcService, nil
}

func Connect(host string, port string) (*grpc.ClientConn, error) {
	var clientConn *grpc.ClientConn
	address := fmt.Sprintf(`%s:%s`, host, port)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	clientConn = conn

	return clientConn, nil
}

func GetInstance() *GrpcService {
	return grpcService
}

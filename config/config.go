package config

import rpc "buf.build/gen/go/bytebase/bytebase/grpc/go/v1/bytebasev1grpc"

type Config struct {
	ProjectServiceClient rpc.ProjectServiceClient
}

func New() *Config {
	return &Config{}
}

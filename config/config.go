package config

import v1pb "buf.build/gen/go/bytebase/bytebase/grpc/go/v1/bytebasev1grpc"

type Config struct {
	URL   string
	Token string

	projectServiceClient v1pb.ProjectServiceClient
}

func New() (*Config, error) {
	return &Config{
		URL:   "",
		Token: "",
	}, nil
}

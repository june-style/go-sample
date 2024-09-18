package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/derrors"
)

func NewConfig(ctx context.Context, cfg *dconfig.Config) (aws.Config, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.Aws.Region),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, opts ...any) (aws.Endpoint, error) {
			if region == cfg.Aws.Region && len(cfg.Aws.Endpoint) > 0 {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           cfg.Aws.Endpoint,
					SigningRegion: cfg.Aws.Region,
				}, nil
			}
			// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})),
	)
	if err != nil {
		return aws.Config{}, derrors.Wrap(err)
	}
	return awsCfg, nil
}

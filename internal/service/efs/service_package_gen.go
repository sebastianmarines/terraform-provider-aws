// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package efs

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	efs_sdkv2 "github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAccessPoint,
			TypeName: "aws_efs_access_point",
			Name:     "Access Point",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  DataSourceAccessPoints,
			TypeName: "aws_efs_access_points",
		},
		{
			Factory:  dataSourceFileSystem,
			TypeName: "aws_efs_file_system",
			Name:     "File System",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  DataSourceMountTarget,
			TypeName: "aws_efs_mount_target",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccessPoint,
			TypeName: "aws_efs_access_point",
			Name:     "Access Point",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceBackupPolicy,
			TypeName: "aws_efs_backup_policy",
		},
		{
			Factory:  resourceFileSystem,
			TypeName: "aws_efs_file_system",
			Name:     "File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceFileSystemPolicy,
			TypeName: "aws_efs_file_system_policy",
			Name:     "File System Policy",
		},
		{
			Factory:  ResourceMountTarget,
			TypeName: "aws_efs_mount_target",
			Name:     "Mount Target",
		},
		{
			Factory:  ResourceReplicationConfiguration,
			TypeName: "aws_efs_replication_configuration",
			Name:     "Replication Configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EFS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*efs_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return efs_sdkv2.NewFromConfig(cfg,
		efs_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

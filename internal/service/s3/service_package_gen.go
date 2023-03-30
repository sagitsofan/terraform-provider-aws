// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package s3

import (
	"context"

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
			Factory:  DataSourceCanonicalUserID,
			TypeName: "aws_canonical_user_id",
		},
		{
			Factory:  DataSourceBucket,
			TypeName: "aws_s3_bucket",
		},
		{
			Factory:  DataSourceBucketObject,
			TypeName: "aws_s3_bucket_object",
		},
		{
			Factory:  DataSourceBucketObjects,
			TypeName: "aws_s3_bucket_objects",
		},
		{
			Factory:  DataSourceBucketPolicy,
			TypeName: "aws_s3_bucket_policy",
		},
		{
			Factory:  DataSourceObject,
			TypeName: "aws_s3_object",
		},
		{
			Factory:  DataSourceObjects,
			TypeName: "aws_s3_objects",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBucket,
			TypeName: "aws_s3_bucket",
			Name:     "Bucket",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceBucketAccelerateConfiguration,
			TypeName: "aws_s3_bucket_accelerate_configuration",
		},
		{
			Factory:  ResourceBucketACL,
			TypeName: "aws_s3_bucket_acl",
		},
		{
			Factory:  ResourceBucketAnalyticsConfiguration,
			TypeName: "aws_s3_bucket_analytics_configuration",
		},
		{
			Factory:  ResourceBucketCorsConfiguration,
			TypeName: "aws_s3_bucket_cors_configuration",
		},
		{
			Factory:  ResourceBucketIntelligentTieringConfiguration,
			TypeName: "aws_s3_bucket_intelligent_tiering_configuration",
		},
		{
			Factory:  ResourceBucketInventory,
			TypeName: "aws_s3_bucket_inventory",
		},
		{
			Factory:  ResourceBucketLifecycleConfiguration,
			TypeName: "aws_s3_bucket_lifecycle_configuration",
		},
		{
			Factory:  ResourceBucketLogging,
			TypeName: "aws_s3_bucket_logging",
		},
		{
			Factory:  ResourceBucketMetric,
			TypeName: "aws_s3_bucket_metric",
		},
		{
			Factory:  ResourceBucketNotification,
			TypeName: "aws_s3_bucket_notification",
		},
		{
			Factory:  ResourceBucketObject,
			TypeName: "aws_s3_bucket_object",
			Name:     "Object",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceBucketObjectLockConfiguration,
			TypeName: "aws_s3_bucket_object_lock_configuration",
		},
		{
			Factory:  ResourceBucketOwnershipControls,
			TypeName: "aws_s3_bucket_ownership_controls",
		},
		{
			Factory:  ResourceBucketPolicy,
			TypeName: "aws_s3_bucket_policy",
		},
		{
			Factory:  ResourceBucketPublicAccessBlock,
			TypeName: "aws_s3_bucket_public_access_block",
		},
		{
			Factory:  ResourceBucketReplicationConfiguration,
			TypeName: "aws_s3_bucket_replication_configuration",
		},
		{
			Factory:  ResourceBucketRequestPaymentConfiguration,
			TypeName: "aws_s3_bucket_request_payment_configuration",
		},
		{
			Factory:  ResourceBucketServerSideEncryptionConfiguration,
			TypeName: "aws_s3_bucket_server_side_encryption_configuration",
		},
		{
			Factory:  ResourceBucketVersioning,
			TypeName: "aws_s3_bucket_versioning",
		},
		{
			Factory:  ResourceBucketWebsiteConfiguration,
			TypeName: "aws_s3_bucket_website_configuration",
		},
		{
			Factory:  ResourceObject,
			TypeName: "aws_s3_object",
			Name:     "Object",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceObjectCopy,
			TypeName: "aws_s3_object_copy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.S3
}

var ServicePackage = &servicePackage{}

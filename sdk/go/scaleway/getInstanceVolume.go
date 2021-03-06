// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupInstanceVolume(ctx *pulumi.Context, args *LookupInstanceVolumeArgs, opts ...pulumi.InvokeOption) (*LookupInstanceVolumeResult, error) {
	var rv LookupInstanceVolumeResult
	err := ctx.Invoke("scaleway:index/getInstanceVolume:getInstanceVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceVolume.
type LookupInstanceVolumeArgs struct {
	Name     *string `pulumi:"name"`
	VolumeId *string `pulumi:"volumeId"`
	Zone     *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceVolume.
type LookupInstanceVolumeResult struct {
	FromSnapshotId string `pulumi:"fromSnapshotId"`
	FromVolumeId   string `pulumi:"fromVolumeId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ServerId       string  `pulumi:"serverId"`
	SizeInGb       int     `pulumi:"sizeInGb"`
	Type           string  `pulumi:"type"`
	VolumeId       *string `pulumi:"volumeId"`
	Zone           *string `pulumi:"zone"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("scaleway:index/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	Server   string `pulumi:"server"`
	SizeInGb int    `pulumi:"sizeInGb"`
	Type     string `pulumi:"type"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	// the name of the volume
	Name pulumi.StringOutput `pulumi:"name"`
	// the server the volume is attached to
	Server pulumi.StringOutput `pulumi:"server"`
	// the size of the volume in GB
	SizeInGb pulumi.IntOutput `pulumi:"sizeInGb"`
	// the type of backing storage
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil || args.SizeInGb == nil {
		return nil, errors.New("missing required argument 'SizeInGb'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &VolumeArgs{}
	}
	var resource Volume
	err := ctx.RegisterResource("scaleway:index/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("scaleway:index/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// the name of the volume
	Name *string `pulumi:"name"`
	// the server the volume is attached to
	Server *string `pulumi:"server"`
	// the size of the volume in GB
	SizeInGb *int `pulumi:"sizeInGb"`
	// the type of backing storage
	Type *string `pulumi:"type"`
}

type VolumeState struct {
	// the name of the volume
	Name pulumi.StringPtrInput
	// the server the volume is attached to
	Server pulumi.StringPtrInput
	// the size of the volume in GB
	SizeInGb pulumi.IntPtrInput
	// the type of backing storage
	Type pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// the name of the volume
	Name *string `pulumi:"name"`
	// the size of the volume in GB
	SizeInGb int `pulumi:"sizeInGb"`
	// the type of backing storage
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// the name of the volume
	Name pulumi.StringPtrInput
	// the size of the volume in GB
	SizeInGb pulumi.IntInput
	// the type of backing storage
	Type pulumi.StringInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

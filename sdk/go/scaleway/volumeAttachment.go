// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VolumeAttachment struct {
	pulumi.CustomResourceState

	// the server a volume should be attached to
	Server pulumi.StringOutput `pulumi:"server"`
	// the volume to attach
	Volume pulumi.StringOutput `pulumi:"volume"`
}

// NewVolumeAttachment registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachment(ctx *pulumi.Context,
	name string, args *VolumeAttachmentArgs, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	if args == nil || args.Server == nil {
		return nil, errors.New("missing required argument 'Server'")
	}
	if args == nil || args.Volume == nil {
		return nil, errors.New("missing required argument 'Volume'")
	}
	if args == nil {
		args = &VolumeAttachmentArgs{}
	}
	var resource VolumeAttachment
	err := ctx.RegisterResource("scaleway:index/volumeAttachment:VolumeAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachment gets an existing VolumeAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachmentState, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	var resource VolumeAttachment
	err := ctx.ReadResource("scaleway:index/volumeAttachment:VolumeAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachment resources.
type volumeAttachmentState struct {
	// the server a volume should be attached to
	Server *string `pulumi:"server"`
	// the volume to attach
	Volume *string `pulumi:"volume"`
}

type VolumeAttachmentState struct {
	// the server a volume should be attached to
	Server pulumi.StringPtrInput
	// the volume to attach
	Volume pulumi.StringPtrInput
}

func (VolumeAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentState)(nil)).Elem()
}

type volumeAttachmentArgs struct {
	// the server a volume should be attached to
	Server string `pulumi:"server"`
	// the volume to attach
	Volume string `pulumi:"volume"`
}

// The set of arguments for constructing a VolumeAttachment resource.
type VolumeAttachmentArgs struct {
	// the server a volume should be attached to
	Server pulumi.StringInput
	// the volume to attach
	Volume pulumi.StringInput
}

func (VolumeAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentArgs)(nil)).Elem()
}

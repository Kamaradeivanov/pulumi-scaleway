// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type InstanceIPReverseDns struct {
	pulumi.CustomResourceState

	// The IP ID or IP address
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The reverse DNS for this IP
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceIPReverseDns registers a new resource with the given unique name, arguments, and options.
func NewInstanceIPReverseDns(ctx *pulumi.Context,
	name string, args *InstanceIPReverseDnsArgs, opts ...pulumi.ResourceOption) (*InstanceIPReverseDns, error) {
	if args == nil || args.IpId == nil {
		return nil, errors.New("missing required argument 'IpId'")
	}
	if args == nil || args.Reverse == nil {
		return nil, errors.New("missing required argument 'Reverse'")
	}
	if args == nil {
		args = &InstanceIPReverseDnsArgs{}
	}
	var resource InstanceIPReverseDns
	err := ctx.RegisterResource("scaleway:index/instanceIPReverseDns:InstanceIPReverseDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIPReverseDns gets an existing InstanceIPReverseDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIPReverseDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIPReverseDnsState, opts ...pulumi.ResourceOption) (*InstanceIPReverseDns, error) {
	var resource InstanceIPReverseDns
	err := ctx.ReadResource("scaleway:index/instanceIPReverseDns:InstanceIPReverseDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIPReverseDns resources.
type instanceIPReverseDnsState struct {
	// The IP ID or IP address
	IpId *string `pulumi:"ipId"`
	// The reverse DNS for this IP
	Reverse *string `pulumi:"reverse"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type InstanceIPReverseDnsState struct {
	// The IP ID or IP address
	IpId pulumi.StringPtrInput
	// The reverse DNS for this IP
	Reverse pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceIPReverseDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIPReverseDnsState)(nil)).Elem()
}

type instanceIPReverseDnsArgs struct {
	// The IP ID or IP address
	IpId string `pulumi:"ipId"`
	// The reverse DNS for this IP
	Reverse string `pulumi:"reverse"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceIPReverseDns resource.
type InstanceIPReverseDnsArgs struct {
	// The IP ID or IP address
	IpId pulumi.StringInput
	// The reverse DNS for this IP
	Reverse pulumi.StringInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceIPReverseDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIPReverseDnsArgs)(nil)).Elem()
}
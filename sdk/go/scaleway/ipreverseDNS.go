// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IPReverseDNS struct {
	pulumi.CustomResourceState

	// The ipv4 address of the ip, or IP ID
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The ipv4 reverse dns
	Reverse pulumi.StringOutput `pulumi:"reverse"`
}

// NewIPReverseDNS registers a new resource with the given unique name, arguments, and options.
func NewIPReverseDNS(ctx *pulumi.Context,
	name string, args *IPReverseDNSArgs, opts ...pulumi.ResourceOption) (*IPReverseDNS, error) {
	if args == nil || args.Ip == nil {
		return nil, errors.New("missing required argument 'Ip'")
	}
	if args == nil || args.Reverse == nil {
		return nil, errors.New("missing required argument 'Reverse'")
	}
	if args == nil {
		args = &IPReverseDNSArgs{}
	}
	var resource IPReverseDNS
	err := ctx.RegisterResource("scaleway:index/iPReverseDNS:IPReverseDNS", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPReverseDNS gets an existing IPReverseDNS resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPReverseDNS(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPReverseDNSState, opts ...pulumi.ResourceOption) (*IPReverseDNS, error) {
	var resource IPReverseDNS
	err := ctx.ReadResource("scaleway:index/iPReverseDNS:IPReverseDNS", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPReverseDNS resources.
type ipreverseDNSState struct {
	// The ipv4 address of the ip, or IP ID
	Ip *string `pulumi:"ip"`
	// The ipv4 reverse dns
	Reverse *string `pulumi:"reverse"`
}

type IPReverseDNSState struct {
	// The ipv4 address of the ip, or IP ID
	Ip pulumi.StringPtrInput
	// The ipv4 reverse dns
	Reverse pulumi.StringPtrInput
}

func (IPReverseDNSState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipreverseDNSState)(nil)).Elem()
}

type ipreverseDNSArgs struct {
	// The ipv4 address of the ip, or IP ID
	Ip string `pulumi:"ip"`
	// The ipv4 reverse dns
	Reverse string `pulumi:"reverse"`
}

// The set of arguments for constructing a IPReverseDNS resource.
type IPReverseDNSArgs struct {
	// The ipv4 address of the ip, or IP ID
	Ip pulumi.StringInput
	// The ipv4 reverse dns
	Reverse pulumi.StringInput
}

func (IPReverseDNSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipreverseDNSArgs)(nil)).Elem()
}

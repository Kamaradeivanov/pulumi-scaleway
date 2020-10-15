// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecurityGroup struct {
	pulumi.CustomResourceState

	// The description of the security group
	Description pulumi.StringOutput `pulumi:"description"`
	// Add default security group rules
	EnableDefaultSecurity pulumi.BoolPtrOutput `pulumi:"enableDefaultSecurity"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"inboundDefaultPolicy"`
	// The name of the security group
	Name pulumi.StringOutput `pulumi:"name"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"outboundDefaultPolicy"`
	// Mark security group as stateful
	Stateful pulumi.BoolPtrOutput `pulumi:"stateful"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil {
		args = &SecurityGroupArgs{}
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("scaleway:index/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("scaleway:index/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The description of the security group
	Description *string `pulumi:"description"`
	// Add default security group rules
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// Mark security group as stateful
	Stateful *bool `pulumi:"stateful"`
}

type SecurityGroupState struct {
	// The description of the security group
	Description pulumi.StringPtrInput
	// Add default security group rules
	EnableDefaultSecurity pulumi.BoolPtrInput
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrInput
	// Mark security group as stateful
	Stateful pulumi.BoolPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The description of the security group
	Description string `pulumi:"description"`
	// Add default security group rules
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// Mark security group as stateful
	Stateful *bool `pulumi:"stateful"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The description of the security group
	Description pulumi.StringInput
	// Add default security group rules
	EnableDefaultSecurity pulumi.BoolPtrInput
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrInput
	// Mark security group as stateful
	Stateful pulumi.BoolPtrInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

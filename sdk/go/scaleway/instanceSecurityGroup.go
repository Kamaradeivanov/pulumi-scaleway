// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type InstanceSecurityGroup struct {
	pulumi.CustomResourceState

	// The description of the security group
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	ExternalRules pulumi.BoolPtrOutput   `pulumi:"externalRules"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"inboundDefaultPolicy"`
	// Inbound rules for this security group
	InboundRules InstanceSecurityGroupInboundRuleArrayOutput `pulumi:"inboundRules"`
	// The name of the security group
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"outboundDefaultPolicy"`
	// Outbound rules for this security group
	OutboundRules InstanceSecurityGroupOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// The stateful value of the security group
	Stateful pulumi.BoolPtrOutput `pulumi:"stateful"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceSecurityGroup(ctx *pulumi.Context,
	name string, args *InstanceSecurityGroupArgs, opts ...pulumi.ResourceOption) (*InstanceSecurityGroup, error) {
	if args == nil {
		args = &InstanceSecurityGroupArgs{}
	}
	var resource InstanceSecurityGroup
	err := ctx.RegisterResource("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceSecurityGroup gets an existing InstanceSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceSecurityGroupState, opts ...pulumi.ResourceOption) (*InstanceSecurityGroup, error) {
	var resource InstanceSecurityGroup
	err := ctx.ReadResource("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceSecurityGroup resources.
type instanceSecurityGroupState struct {
	// The description of the security group
	Description   *string `pulumi:"description"`
	ExternalRules *bool   `pulumi:"externalRules"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// Inbound rules for this security group
	InboundRules []InstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// Outbound rules for this security group
	OutboundRules []InstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// The stateful value of the security group
	Stateful *bool `pulumi:"stateful"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type InstanceSecurityGroupState struct {
	// The description of the security group
	Description   pulumi.StringPtrInput
	ExternalRules pulumi.BoolPtrInput
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrInput
	// Inbound rules for this security group
	InboundRules InstanceSecurityGroupInboundRuleArrayInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrInput
	// Outbound rules for this security group
	OutboundRules InstanceSecurityGroupOutboundRuleArrayInput
	// The stateful value of the security group
	Stateful pulumi.BoolPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupState)(nil)).Elem()
}

type instanceSecurityGroupArgs struct {
	// The description of the security group
	Description   *string `pulumi:"description"`
	ExternalRules *bool   `pulumi:"externalRules"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// Inbound rules for this security group
	InboundRules []InstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// Outbound rules for this security group
	OutboundRules []InstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// The stateful value of the security group
	Stateful *bool `pulumi:"stateful"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceSecurityGroup resource.
type InstanceSecurityGroupArgs struct {
	// The description of the security group
	Description   pulumi.StringPtrInput
	ExternalRules pulumi.BoolPtrInput
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrInput
	// Inbound rules for this security group
	InboundRules InstanceSecurityGroupInboundRuleArrayInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrInput
	// Outbound rules for this security group
	OutboundRules InstanceSecurityGroupOutboundRuleArrayInput
	// The stateful value of the security group
	Stateful pulumi.BoolPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupArgs)(nil)).Elem()
}

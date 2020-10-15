// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type K8SPoolBeta struct {
	pulumi.CustomResourceState

	// Enable the autohealing on the pool
	Autohealing pulumi.BoolPtrOutput `pulumi:"autohealing"`
	// Enable the autoscaling on the pool
	Autoscaling pulumi.BoolPtrOutput `pulumi:"autoscaling"`
	// The ID of the cluster on which this pool will be created
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Container runtime for the pool
	ContainerRuntime pulumi.StringPtrOutput `pulumi:"containerRuntime"`
	// The date and time of the creation of the pool
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The actual size of the pool
	CurrentSize pulumi.IntOutput `pulumi:"currentSize"`
	// Maximum size of the pool
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// Minimun size of the pool
	MinSize pulumi.IntPtrOutput `pulumi:"minSize"`
	// The name of the cluster
	Name pulumi.StringOutput `pulumi:"name"`
	// Server type of the pool servers
	NodeType pulumi.StringOutput        `pulumi:"nodeType"`
	Nodes    K8SPoolBetaNodeArrayOutput `pulumi:"nodes"`
	// ID of the placement group
	PlacementGroupId pulumi.StringPtrOutput `pulumi:"placementGroupId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// Size of the pool
	Size pulumi.IntOutput `pulumi:"size"`
	// The status of the pool
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags associated with the pool
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the pool
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The Kubernetes version of the pool
	Version pulumi.StringOutput `pulumi:"version"`
	// Whether to wait for the pool to be ready
	WaitForPoolReady pulumi.BoolPtrOutput `pulumi:"waitForPoolReady"`
}

// NewK8SPoolBeta registers a new resource with the given unique name, arguments, and options.
func NewK8SPoolBeta(ctx *pulumi.Context,
	name string, args *K8SPoolBetaArgs, opts ...pulumi.ResourceOption) (*K8SPoolBeta, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.NodeType == nil {
		return nil, errors.New("missing required argument 'NodeType'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &K8SPoolBetaArgs{}
	}
	var resource K8SPoolBeta
	err := ctx.RegisterResource("scaleway:index/k8SPoolBeta:K8SPoolBeta", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetK8SPoolBeta gets an existing K8SPoolBeta resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetK8SPoolBeta(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *K8SPoolBetaState, opts ...pulumi.ResourceOption) (*K8SPoolBeta, error) {
	var resource K8SPoolBeta
	err := ctx.ReadResource("scaleway:index/k8SPoolBeta:K8SPoolBeta", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering K8SPoolBeta resources.
type k8spoolBetaState struct {
	// Enable the autohealing on the pool
	Autohealing *bool `pulumi:"autohealing"`
	// Enable the autoscaling on the pool
	Autoscaling *bool `pulumi:"autoscaling"`
	// The ID of the cluster on which this pool will be created
	ClusterId *string `pulumi:"clusterId"`
	// Container runtime for the pool
	ContainerRuntime *string `pulumi:"containerRuntime"`
	// The date and time of the creation of the pool
	CreatedAt *string `pulumi:"createdAt"`
	// The actual size of the pool
	CurrentSize *int `pulumi:"currentSize"`
	// Maximum size of the pool
	MaxSize *int `pulumi:"maxSize"`
	// Minimun size of the pool
	MinSize *int `pulumi:"minSize"`
	// The name of the cluster
	Name *string `pulumi:"name"`
	// Server type of the pool servers
	NodeType *string           `pulumi:"nodeType"`
	Nodes    []K8SPoolBetaNode `pulumi:"nodes"`
	// ID of the placement group
	PlacementGroupId *string `pulumi:"placementGroupId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Size of the pool
	Size *int `pulumi:"size"`
	// The status of the pool
	Status *string `pulumi:"status"`
	// The tags associated with the pool
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the pool
	UpdatedAt *string `pulumi:"updatedAt"`
	// The Kubernetes version of the pool
	Version *string `pulumi:"version"`
	// Whether to wait for the pool to be ready
	WaitForPoolReady *bool `pulumi:"waitForPoolReady"`
}

type K8SPoolBetaState struct {
	// Enable the autohealing on the pool
	Autohealing pulumi.BoolPtrInput
	// Enable the autoscaling on the pool
	Autoscaling pulumi.BoolPtrInput
	// The ID of the cluster on which this pool will be created
	ClusterId pulumi.StringPtrInput
	// Container runtime for the pool
	ContainerRuntime pulumi.StringPtrInput
	// The date and time of the creation of the pool
	CreatedAt pulumi.StringPtrInput
	// The actual size of the pool
	CurrentSize pulumi.IntPtrInput
	// Maximum size of the pool
	MaxSize pulumi.IntPtrInput
	// Minimun size of the pool
	MinSize pulumi.IntPtrInput
	// The name of the cluster
	Name pulumi.StringPtrInput
	// Server type of the pool servers
	NodeType pulumi.StringPtrInput
	Nodes    K8SPoolBetaNodeArrayInput
	// ID of the placement group
	PlacementGroupId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Size of the pool
	Size pulumi.IntPtrInput
	// The status of the pool
	Status pulumi.StringPtrInput
	// The tags associated with the pool
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the pool
	UpdatedAt pulumi.StringPtrInput
	// The Kubernetes version of the pool
	Version pulumi.StringPtrInput
	// Whether to wait for the pool to be ready
	WaitForPoolReady pulumi.BoolPtrInput
}

func (K8SPoolBetaState) ElementType() reflect.Type {
	return reflect.TypeOf((*k8spoolBetaState)(nil)).Elem()
}

type k8spoolBetaArgs struct {
	// Enable the autohealing on the pool
	Autohealing *bool `pulumi:"autohealing"`
	// Enable the autoscaling on the pool
	Autoscaling *bool `pulumi:"autoscaling"`
	// The ID of the cluster on which this pool will be created
	ClusterId string `pulumi:"clusterId"`
	// Container runtime for the pool
	ContainerRuntime *string `pulumi:"containerRuntime"`
	// Maximum size of the pool
	MaxSize *int `pulumi:"maxSize"`
	// Minimun size of the pool
	MinSize *int `pulumi:"minSize"`
	// The name of the cluster
	Name *string `pulumi:"name"`
	// Server type of the pool servers
	NodeType string `pulumi:"nodeType"`
	// ID of the placement group
	PlacementGroupId *string `pulumi:"placementGroupId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Size of the pool
	Size int `pulumi:"size"`
	// The tags associated with the pool
	Tags []string `pulumi:"tags"`
	// Whether to wait for the pool to be ready
	WaitForPoolReady *bool `pulumi:"waitForPoolReady"`
}

// The set of arguments for constructing a K8SPoolBeta resource.
type K8SPoolBetaArgs struct {
	// Enable the autohealing on the pool
	Autohealing pulumi.BoolPtrInput
	// Enable the autoscaling on the pool
	Autoscaling pulumi.BoolPtrInput
	// The ID of the cluster on which this pool will be created
	ClusterId pulumi.StringInput
	// Container runtime for the pool
	ContainerRuntime pulumi.StringPtrInput
	// Maximum size of the pool
	MaxSize pulumi.IntPtrInput
	// Minimun size of the pool
	MinSize pulumi.IntPtrInput
	// The name of the cluster
	Name pulumi.StringPtrInput
	// Server type of the pool servers
	NodeType pulumi.StringInput
	// ID of the placement group
	PlacementGroupId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Size of the pool
	Size pulumi.IntInput
	// The tags associated with the pool
	Tags pulumi.StringArrayInput
	// Whether to wait for the pool to be ready
	WaitForPoolReady pulumi.BoolPtrInput
}

func (K8SPoolBetaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*k8spoolBetaArgs)(nil)).Elem()
}
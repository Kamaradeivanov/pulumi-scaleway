// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RdbInstanceBeta struct {
	pulumi.CustomResourceState

	// Certificate of the database instance
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Disable automated backup for the database instance
	DisableBackup pulumi.BoolPtrOutput `pulumi:"disableBackup"`
	// Endpoint IP of the database instance
	EndpointIp pulumi.StringOutput `pulumi:"endpointIp"`
	// Endpoint port of the database instance
	EndpointPort pulumi.IntOutput `pulumi:"endpointPort"`
	// Database's engine version id
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Enable or disable high availability for the database instance
	IsHaCluster pulumi.BoolPtrOutput `pulumi:"isHaCluster"`
	// Name of the database instance
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of database instance you want to create
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Password for the first user of the database instance
	Password pulumi.StringOutput `pulumi:"password"`
	// Read replicas of the database instance
	ReadReplicas RdbInstanceBetaReadReplicaArrayOutput `pulumi:"readReplicas"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Identifier for the first user of the database instance
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewRdbInstanceBeta registers a new resource with the given unique name, arguments, and options.
func NewRdbInstanceBeta(ctx *pulumi.Context,
	name string, args *RdbInstanceBetaArgs, opts ...pulumi.ResourceOption) (*RdbInstanceBeta, error) {
	if args == nil || args.Engine == nil {
		return nil, errors.New("missing required argument 'Engine'")
	}
	if args == nil || args.NodeType == nil {
		return nil, errors.New("missing required argument 'NodeType'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	if args == nil {
		args = &RdbInstanceBetaArgs{}
	}
	var resource RdbInstanceBeta
	err := ctx.RegisterResource("scaleway:index/rdbInstanceBeta:RdbInstanceBeta", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbInstanceBeta gets an existing RdbInstanceBeta resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbInstanceBeta(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbInstanceBetaState, opts ...pulumi.ResourceOption) (*RdbInstanceBeta, error) {
	var resource RdbInstanceBeta
	err := ctx.ReadResource("scaleway:index/rdbInstanceBeta:RdbInstanceBeta", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbInstanceBeta resources.
type rdbInstanceBetaState struct {
	// Certificate of the database instance
	Certificate *string `pulumi:"certificate"`
	// Disable automated backup for the database instance
	DisableBackup *bool `pulumi:"disableBackup"`
	// Endpoint IP of the database instance
	EndpointIp *string `pulumi:"endpointIp"`
	// Endpoint port of the database instance
	EndpointPort *int `pulumi:"endpointPort"`
	// Database's engine version id
	Engine *string `pulumi:"engine"`
	// Enable or disable high availability for the database instance
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// Name of the database instance
	Name *string `pulumi:"name"`
	// The type of database instance you want to create
	NodeType *string `pulumi:"nodeType"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Password for the first user of the database instance
	Password *string `pulumi:"password"`
	// Read replicas of the database instance
	ReadReplicas []RdbInstanceBetaReadReplica `pulumi:"readReplicas"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance
	UserName *string `pulumi:"userName"`
}

type RdbInstanceBetaState struct {
	// Certificate of the database instance
	Certificate pulumi.StringPtrInput
	// Disable automated backup for the database instance
	DisableBackup pulumi.BoolPtrInput
	// Endpoint IP of the database instance
	EndpointIp pulumi.StringPtrInput
	// Endpoint port of the database instance
	EndpointPort pulumi.IntPtrInput
	// Database's engine version id
	Engine pulumi.StringPtrInput
	// Enable or disable high availability for the database instance
	IsHaCluster pulumi.BoolPtrInput
	// Name of the database instance
	Name pulumi.StringPtrInput
	// The type of database instance you want to create
	NodeType pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Password for the first user of the database instance
	Password pulumi.StringPtrInput
	// Read replicas of the database instance
	ReadReplicas RdbInstanceBetaReadReplicaArrayInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance
	UserName pulumi.StringPtrInput
}

func (RdbInstanceBetaState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbInstanceBetaState)(nil)).Elem()
}

type rdbInstanceBetaArgs struct {
	// Disable automated backup for the database instance
	DisableBackup *bool `pulumi:"disableBackup"`
	// Database's engine version id
	Engine string `pulumi:"engine"`
	// Enable or disable high availability for the database instance
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// Name of the database instance
	Name *string `pulumi:"name"`
	// The type of database instance you want to create
	NodeType string `pulumi:"nodeType"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Password for the first user of the database instance
	Password string `pulumi:"password"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a RdbInstanceBeta resource.
type RdbInstanceBetaArgs struct {
	// Disable automated backup for the database instance
	DisableBackup pulumi.BoolPtrInput
	// Database's engine version id
	Engine pulumi.StringInput
	// Enable or disable high availability for the database instance
	IsHaCluster pulumi.BoolPtrInput
	// Name of the database instance
	Name pulumi.StringPtrInput
	// The type of database instance you want to create
	NodeType pulumi.StringInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Password for the first user of the database instance
	Password pulumi.StringInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance
	UserName pulumi.StringInput
}

func (RdbInstanceBetaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbInstanceBetaArgs)(nil)).Elem()
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type UserData struct {
	pulumi.CustomResourceState

	// The key of the user data to manage
	Key pulumi.StringOutput `pulumi:"key"`
	// The server the meta data is associated with
	Server pulumi.StringOutput `pulumi:"server"`
	// The value of the user
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewUserData registers a new resource with the given unique name, arguments, and options.
func NewUserData(ctx *pulumi.Context,
	name string, args *UserDataArgs, opts ...pulumi.ResourceOption) (*UserData, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Server == nil {
		return nil, errors.New("missing required argument 'Server'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &UserDataArgs{}
	}
	var resource UserData
	err := ctx.RegisterResource("scaleway:index/userData:UserData", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserData gets an existing UserData resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserData(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDataState, opts ...pulumi.ResourceOption) (*UserData, error) {
	var resource UserData
	err := ctx.ReadResource("scaleway:index/userData:UserData", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserData resources.
type userDataState struct {
	// The key of the user data to manage
	Key *string `pulumi:"key"`
	// The server the meta data is associated with
	Server *string `pulumi:"server"`
	// The value of the user
	Value *string `pulumi:"value"`
}

type UserDataState struct {
	// The key of the user data to manage
	Key pulumi.StringPtrInput
	// The server the meta data is associated with
	Server pulumi.StringPtrInput
	// The value of the user
	Value pulumi.StringPtrInput
}

func (UserDataState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDataState)(nil)).Elem()
}

type userDataArgs struct {
	// The key of the user data to manage
	Key string `pulumi:"key"`
	// The server the meta data is associated with
	Server string `pulumi:"server"`
	// The value of the user
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a UserData resource.
type UserDataArgs struct {
	// The key of the user data to manage
	Key pulumi.StringInput
	// The server the meta data is associated with
	Server pulumi.StringInput
	// The value of the user
	Value pulumi.StringInput
}

func (UserDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDataArgs)(nil)).Elem()
}

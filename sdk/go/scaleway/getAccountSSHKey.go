// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAccountSSHKey(ctx *pulumi.Context, args *LookupAccountSSHKeyArgs, opts ...pulumi.InvokeOption) (*LookupAccountSSHKeyResult, error) {
	var rv LookupAccountSSHKeyResult
	err := ctx.Invoke("scaleway:index/getAccountSSHKey:getAccountSSHKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountSSHKey.
type LookupAccountSSHKeyArgs struct {
	Name           *string `pulumi:"name"`
	OrganizationId *string `pulumi:"organizationId"`
	SshKeyId       *string `pulumi:"sshKeyId"`
}

// A collection of values returned by getAccountSSHKey.
type LookupAccountSSHKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	OrganizationId string `pulumi:"organizationId"`
	PublicKey      string `pulumi:"publicKey"`
	SshKeyId       string `pulumi:"sshKeyId"`
}
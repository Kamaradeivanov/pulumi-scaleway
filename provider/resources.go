// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scaleway

import (
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-scaleway/scaleway"
)

// all of the token components used below.
const (
	// packages:
	scalewayPkg = "scaleway"
	// modules:
	scalewayMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(scalewayPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := scaleway.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "scaleway",
		Description: "A Pulumi package for creating and managing scaleway cloud resources.",
		Keywords:    []string{"pulumi", "scaleway"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/Kamaradeivanov/pulumi-scaleway",
		GitHubOrg:   "scaleway",
		Config:      map[string]*tfbridge.SchemaInfo{
			"access_key": {
				// Type: makeType("access_key", "AccessKey"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_ACCESS_KEY"},
				},
			},
			"secret_key": {
				// Type: makeType("secret_key", "SecretKey"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_SECRET_KEY"},
				},
			},
			"organization_id": {
				// Type: makeType("organization_id", "OrganizationId"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ORGANIZATION_ID"},
				},
			},
			"region": {
				// Type: makeType("region", "Region"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_REGION"},
				},
			},
			"zone": {
				// Type: makeType("zone", "Zone"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ZONE"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources:            map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: makeResource(scalewayMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: makeResource(scalewayMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: makeType(scalewayPkg, "Tags")},
			// 	},
			// },
			"scaleway_account_ssh_key":               {Tok: makeResource(scalewayMod, "AccountSSHKey")},
			"scaleway_baremetal_server":         			{Tok: makeResource(scalewayMod, "BaremetalServer"),},
			// "scaleway_bucket":                        {Tok: makeResource(scalewayMod, "Bucket")},
			"scaleway_instance_ip":                   {Tok: makeResource(scalewayMod, "InstanceIP")},
			"scaleway_instance_ip_reverse_dns":       {Tok: makeResource(scalewayMod, "InstanceIPReverseDns")},
			"scaleway_instance_placement_group":      {Tok: makeResource(scalewayMod, "InstancePlacementGroup")},
			"scaleway_instance_security_group":       {Tok: makeResource(scalewayMod, "InstanceSecurityGroup")},
			"scaleway_instance_security_group_rules": {Tok: makeResource(scalewayMod, "InstanceSecurityGroupRules")},
			"scaleway_instance_server":               {Tok: makeResource(scalewayMod, "InstanceServer")},
			"scaleway_instance_volume":               {Tok: makeResource(scalewayMod, "InstanceVolume")},
			"scaleway_ip":                            {Tok: makeResource(scalewayMod, "IP")},
			"scaleway_ip_reverse_dns":                {Tok: makeResource(scalewayMod, "IPReverseDNS")},
			"scaleway_k8s_cluster_beta":              {Tok: makeResource(scalewayMod, "K8SClusterBeta")},
			"scaleway_k8s_pool_beta":                 {Tok: makeResource(scalewayMod, "K8SPoolBeta")},
			"scaleway_lb_backend_beta":               {Tok: makeResource(scalewayMod, "LbBackendBeta")},
			"scaleway_lb_beta":                       {Tok: makeResource(scalewayMod, "LbBeta")},
			"scaleway_lb_certificate_beta":           {Tok: makeResource(scalewayMod, "LbCertificateBeta")},
			"scaleway_lb_frontend_beta":              {Tok: makeResource(scalewayMod, "LbFrontendBeta")},
			"scaleway_lb_ip_beta":              			{Tok: makeResource(scalewayMod, "LbIPBeta")},
			"scaleway_object_bucket":                 {Tok: makeResource(scalewayMod, "ObjectBucket")},
			"scaleway_rdb_instance_beta":             {Tok: makeResource(scalewayMod, "RdbInstanceBeta")},
			"scaleway_registry_namespace_beta":       {Tok: makeResource(scalewayMod, "RegistryNamespaceBeta")},
			"scaleway_security_group":                {Tok: makeResource(scalewayMod, "SecurityGroup")},
			"scaleway_security_group_rule":           {Tok: makeResource(scalewayMod, "SecurityGroupRule")},
			"scaleway_server":                        {Tok: makeResource(scalewayMod, "Server")},
			"scaleway_ssh_key":                       {Tok: makeResource(scalewayMod, "SSHKey")},
			"scaleway_token":                         {Tok: makeResource(scalewayMod, "Token")},
			"scaleway_user_data":                     {Tok: makeResource(scalewayMod, "UserData")},
			"scaleway_volume":                        {Tok: makeResource(scalewayMod, "Volume")},
			"scaleway_volume_attachment":             {Tok: makeResource(scalewayMod, "VolumeAttachment")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(scalewayMod, "getAmi")},
			"scaleway_account_ssh_key":         {Tok: makeDataSource(scalewayMod, "getAccountSSHKey")},
			"scaleway_baremetal_offer":    			{
				Tok: makeDataSource(scalewayMod, "getBaremetalOffer"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"cpu": {
						Name: "cpus",
					},
				},
			},
			"scaleway_bootscript":              {Tok: makeDataSource(scalewayMod, "getBootscript")},
			"scaleway_image":                   {Tok: makeDataSource(scalewayMod, "getImage")},
			"scaleway_instance_image":          {Tok: makeDataSource(scalewayMod, "getInstanceImage")},
			"scaleway_instance_security_group": {Tok: makeDataSource(scalewayMod, "getInstanceSecurityGroup")},
			"scaleway_instance_server":         {Tok: makeDataSource(scalewayMod, "getInstanceServer")},
			"scaleway_instance_volume":         {Tok: makeDataSource(scalewayMod, "getInstanceVolume")},
			"scaleway_lb_ip_beta": 		          {Tok: makeDataSource(scalewayMod, "getLbIPBeta")},
			"scaleway_marketplace_image_beta":  {Tok: makeDataSource(scalewayMod, "getMarketplaceImageBeta")},
			"scaleway_registry_image_beta":     {Tok: makeDataSource(scalewayMod, "getRegistryImageBeta")},
			"scaleway_registry_namespace_beta": {Tok: makeDataSource(scalewayMod, "getRegistryNamespaceBeta")},
			"scaleway_security_group":          {Tok: makeDataSource(scalewayMod, "getSecurityGroup")},
			"scaleway_volume":                  {Tok: makeDataSource(scalewayMod, "getVolume")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// AsyncDataSources: true,
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}

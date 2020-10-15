// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the scaleway package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'scaleway';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        inputs["accessKey"] = (args ? args.accessKey : undefined) || utilities.getEnv("SCW_ACCESS_KEY");
        inputs["organization"] = args ? args.organization : undefined;
        inputs["organizationId"] = (args ? args.organizationId : undefined) || utilities.getEnv("SCW_DEFAULT_ORGANIZATION_ID");
        inputs["region"] = (args ? args.region : undefined) || utilities.getEnv("SCW_DEFAULT_REGION");
        inputs["secretKey"] = (args ? args.secretKey : undefined) || utilities.getEnv("SCW_SECRET_KEY");
        inputs["token"] = args ? args.token : undefined;
        inputs["zone"] = (args ? args.zone : undefined) || utilities.getEnv("SCW_DEFAULT_ZONE");
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The Scaleway access key.
     */
    readonly accessKey?: pulumi.Input<string>;
    /**
     * @deprecated Use `organization_id` instead.
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * The Scaleway organization ID.
     */
    readonly organizationId?: pulumi.Input<string>;
    /**
     * The Scaleway default region to use for your resources.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The Scaleway secret Key.
     */
    readonly secretKey?: pulumi.Input<string>;
    /**
     * @deprecated Use `secret_key` instead.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * The Scaleway default zone to use for your resources.
     */
    readonly zone?: pulumi.Input<string>;
}

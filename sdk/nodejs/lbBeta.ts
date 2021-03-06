// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LbBeta extends pulumi.CustomResource {
    /**
     * Get an existing LbBeta resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbBetaState, opts?: pulumi.CustomResourceOptions): LbBeta {
        return new LbBeta(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/lbBeta:LbBeta';

    /**
     * Returns true if the given object is an instance of LbBeta.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbBeta {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbBeta.__pulumiType;
    }

    /**
     * The load-balance public IP address
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * The load-balance public IP ID
     */
    public readonly ipId!: pulumi.Output<string>;
    /**
     * Name of the lb
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization_id you want to attach the resource to
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Array of tags to associate with the load-balancer
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of load-balancer you want to create
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a LbBeta resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbBetaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbBetaArgs | LbBetaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LbBetaState | undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipId"] = state ? state.ipId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["organizationId"] = state ? state.organizationId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as LbBetaArgs | undefined;
            if (!args || args.ipId === undefined) {
                throw new Error("Missing required property 'ipId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["ipId"] = args ? args.ipId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["ipAddress"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LbBeta.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbBeta resources.
 */
export interface LbBetaState {
    /**
     * The load-balance public IP address
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The load-balance public IP ID
     */
    readonly ipId?: pulumi.Input<string>;
    /**
     * Name of the lb
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The organization_id you want to attach the resource to
     */
    readonly organizationId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Array of tags to associate with the load-balancer
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of load-balancer you want to create
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbBeta resource.
 */
export interface LbBetaArgs {
    /**
     * The load-balance public IP ID
     */
    readonly ipId: pulumi.Input<string>;
    /**
     * Name of the lb
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The organization_id you want to attach the resource to
     */
    readonly organizationId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Array of tags to associate with the load-balancer
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of load-balancer you want to create
     */
    readonly type: pulumi.Input<string>;
}

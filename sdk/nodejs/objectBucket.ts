// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ObjectBucket extends pulumi.CustomResource {
    /**
     * Get an existing ObjectBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectBucketState, opts?: pulumi.CustomResourceOptions): ObjectBucket {
        return new ObjectBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectBucket:ObjectBucket';

    /**
     * Returns true if the given object is an instance of ObjectBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectBucket.__pulumiType;
    }

    /**
     * ACL of the bucket: either 'public-read' or 'private'.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The name of the bucket
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a ObjectBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ObjectBucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectBucketArgs | ObjectBucketState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ObjectBucketState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ObjectBucketArgs | undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ObjectBucket.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectBucket resources.
 */
export interface ObjectBucketState {
    /**
     * ACL of the bucket: either 'public-read' or 'private'.
     */
    readonly acl?: pulumi.Input<string>;
    /**
     * The name of the bucket
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectBucket resource.
 */
export interface ObjectBucketArgs {
    /**
     * ACL of the bucket: either 'public-read' or 'private'.
     */
    readonly acl?: pulumi.Input<string>;
    /**
     * The name of the bucket
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    readonly region?: pulumi.Input<string>;
}
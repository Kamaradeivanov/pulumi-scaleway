// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class InstanceIPReverseDns extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIPReverseDns resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIPReverseDnsState, opts?: pulumi.CustomResourceOptions): InstanceIPReverseDns {
        return new InstanceIPReverseDns(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceIPReverseDns:InstanceIPReverseDns';

    /**
     * Returns true if the given object is an instance of InstanceIPReverseDns.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIPReverseDns {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIPReverseDns.__pulumiType;
    }

    /**
     * The IP ID or IP address
     */
    public readonly ipId!: pulumi.Output<string>;
    /**
     * The reverse DNS for this IP
     */
    public readonly reverse!: pulumi.Output<string>;
    /**
     * The zone you want to attach the resource to
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceIPReverseDns resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIPReverseDnsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIPReverseDnsArgs | InstanceIPReverseDnsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceIPReverseDnsState | undefined;
            inputs["ipId"] = state ? state.ipId : undefined;
            inputs["reverse"] = state ? state.reverse : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceIPReverseDnsArgs | undefined;
            if (!args || args.ipId === undefined) {
                throw new Error("Missing required property 'ipId'");
            }
            if (!args || args.reverse === undefined) {
                throw new Error("Missing required property 'reverse'");
            }
            inputs["ipId"] = args ? args.ipId : undefined;
            inputs["reverse"] = args ? args.reverse : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceIPReverseDns.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIPReverseDns resources.
 */
export interface InstanceIPReverseDnsState {
    /**
     * The IP ID or IP address
     */
    readonly ipId?: pulumi.Input<string>;
    /**
     * The reverse DNS for this IP
     */
    readonly reverse?: pulumi.Input<string>;
    /**
     * The zone you want to attach the resource to
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIPReverseDns resource.
 */
export interface InstanceIPReverseDnsArgs {
    /**
     * The IP ID or IP address
     */
    readonly ipId: pulumi.Input<string>;
    /**
     * The reverse DNS for this IP
     */
    readonly reverse: pulumi.Input<string>;
    /**
     * The zone you want to attach the resource to
     */
    readonly zone?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class BaremetalServer extends pulumi.CustomResource {
    /**
     * Get an existing BaremetalServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BaremetalServerState, opts?: pulumi.CustomResourceOptions): BaremetalServer {
        return new BaremetalServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/baremetalServer:BaremetalServer';

    /**
     * Returns true if the given object is an instance of BaremetalServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BaremetalServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BaremetalServer.__pulumiType;
    }

    /**
     * Some description to associate to the server, max 255 characters
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * Hostname of the server
     */
    public readonly hostname!: pulumi.Output<string | undefined>;
    public /*out*/ readonly ips!: pulumi.Output<outputs.BaremetalServerIp[]>;
    /**
     * Name of the server
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the server offer
     */
    public readonly offer!: pulumi.Output<string>;
    /**
     * ID of the server offer
     */
    public /*out*/ readonly offerId!: pulumi.Output<string>;
    /**
     * The organization_id you want to attach the resource to
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The base image of the server
     */
    public readonly os!: pulumi.Output<string>;
    /**
     * The base image id of the server
     */
    public /*out*/ readonly osId!: pulumi.Output<string>;
    /**
     * Array of SSH key IDs allowed to SSH to the server
     */
    public readonly sshKeyIds!: pulumi.Output<string[]>;
    /**
     * Array of tags to associate with the server
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The zone you want to attach the resource to
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a BaremetalServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BaremetalServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BaremetalServerArgs | BaremetalServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BaremetalServerState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["ips"] = state ? state.ips : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["offer"] = state ? state.offer : undefined;
            inputs["offerId"] = state ? state.offerId : undefined;
            inputs["organizationId"] = state ? state.organizationId : undefined;
            inputs["os"] = state ? state.os : undefined;
            inputs["osId"] = state ? state.osId : undefined;
            inputs["sshKeyIds"] = state ? state.sshKeyIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as BaremetalServerArgs | undefined;
            if (!args || args.offer === undefined) {
                throw new Error("Missing required property 'offer'");
            }
            if (!args || args.os === undefined) {
                throw new Error("Missing required property 'os'");
            }
            if (!args || args.sshKeyIds === undefined) {
                throw new Error("Missing required property 'sshKeyIds'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["offer"] = args ? args.offer : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["os"] = args ? args.os : undefined;
            inputs["sshKeyIds"] = args ? args.sshKeyIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["domain"] = undefined /*out*/;
            inputs["ips"] = undefined /*out*/;
            inputs["offerId"] = undefined /*out*/;
            inputs["osId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BaremetalServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BaremetalServer resources.
 */
export interface BaremetalServerState {
    /**
     * Some description to associate to the server, max 255 characters
     */
    readonly description?: pulumi.Input<string>;
    readonly domain?: pulumi.Input<string>;
    /**
     * Hostname of the server
     */
    readonly hostname?: pulumi.Input<string>;
    readonly ips?: pulumi.Input<pulumi.Input<inputs.BaremetalServerIp>[]>;
    /**
     * Name of the server
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the server offer
     */
    readonly offer?: pulumi.Input<string>;
    /**
     * ID of the server offer
     */
    readonly offerId?: pulumi.Input<string>;
    /**
     * The organization_id you want to attach the resource to
     */
    readonly organizationId?: pulumi.Input<string>;
    /**
     * The base image of the server
     */
    readonly os?: pulumi.Input<string>;
    /**
     * The base image id of the server
     */
    readonly osId?: pulumi.Input<string>;
    /**
     * Array of SSH key IDs allowed to SSH to the server
     */
    readonly sshKeyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of tags to associate with the server
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone you want to attach the resource to
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BaremetalServer resource.
 */
export interface BaremetalServerArgs {
    /**
     * Some description to associate to the server, max 255 characters
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Hostname of the server
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * Name of the server
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the server offer
     */
    readonly offer: pulumi.Input<string>;
    /**
     * The organization_id you want to attach the resource to
     */
    readonly organizationId?: pulumi.Input<string>;
    /**
     * The base image of the server
     */
    readonly os: pulumi.Input<string>;
    /**
     * Array of SSH key IDs allowed to SSH to the server
     */
    readonly sshKeyIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of tags to associate with the server
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone you want to attach the resource to
     */
    readonly zone?: pulumi.Input<string>;
}

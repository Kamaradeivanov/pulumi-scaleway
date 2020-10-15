// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LbCertificateBeta extends pulumi.CustomResource {
    /**
     * Get an existing LbCertificateBeta resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbCertificateBetaState, opts?: pulumi.CustomResourceOptions): LbCertificateBeta {
        return new LbCertificateBeta(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/lbCertificateBeta:LbCertificateBeta';

    /**
     * Returns true if the given object is an instance of LbCertificateBeta.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbCertificateBeta {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbCertificateBeta.__pulumiType;
    }

    /**
     * The main domain name of the certificate
     */
    public /*out*/ readonly commonName!: pulumi.Output<string>;
    /**
     * The custom type certificate type configuration
     */
    public readonly customCertificate!: pulumi.Output<outputs.LbCertificateBetaCustomCertificate | undefined>;
    /**
     * The identifier (SHA-1) of the certificate
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The load-balancer ID
     */
    public readonly lbId!: pulumi.Output<string>;
    /**
     * The Let's Encrypt type certificate configuration
     */
    public readonly letsencrypt!: pulumi.Output<outputs.LbCertificateBetaLetsencrypt | undefined>;
    /**
     * The name of the load-balancer certificate
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The not valid after validity bound timestamp
     */
    public /*out*/ readonly notValidAfter!: pulumi.Output<string>;
    /**
     * The not valid before validity bound timestamp
     */
    public /*out*/ readonly notValidBefore!: pulumi.Output<string>;
    /**
     * The status of certificate
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The alternative domain names of the certificate
     */
    public /*out*/ readonly subjectAlternativeName!: pulumi.Output<string>;

    /**
     * Create a LbCertificateBeta resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbCertificateBetaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbCertificateBetaArgs | LbCertificateBetaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LbCertificateBetaState | undefined;
            inputs["commonName"] = state ? state.commonName : undefined;
            inputs["customCertificate"] = state ? state.customCertificate : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["lbId"] = state ? state.lbId : undefined;
            inputs["letsencrypt"] = state ? state.letsencrypt : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notValidAfter"] = state ? state.notValidAfter : undefined;
            inputs["notValidBefore"] = state ? state.notValidBefore : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subjectAlternativeName"] = state ? state.subjectAlternativeName : undefined;
        } else {
            const args = argsOrState as LbCertificateBetaArgs | undefined;
            if (!args || args.lbId === undefined) {
                throw new Error("Missing required property 'lbId'");
            }
            inputs["customCertificate"] = args ? args.customCertificate : undefined;
            inputs["lbId"] = args ? args.lbId : undefined;
            inputs["letsencrypt"] = args ? args.letsencrypt : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["commonName"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["notValidAfter"] = undefined /*out*/;
            inputs["notValidBefore"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["subjectAlternativeName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LbCertificateBeta.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbCertificateBeta resources.
 */
export interface LbCertificateBetaState {
    /**
     * The main domain name of the certificate
     */
    readonly commonName?: pulumi.Input<string>;
    /**
     * The custom type certificate type configuration
     */
    readonly customCertificate?: pulumi.Input<inputs.LbCertificateBetaCustomCertificate>;
    /**
     * The identifier (SHA-1) of the certificate
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * The load-balancer ID
     */
    readonly lbId?: pulumi.Input<string>;
    /**
     * The Let's Encrypt type certificate configuration
     */
    readonly letsencrypt?: pulumi.Input<inputs.LbCertificateBetaLetsencrypt>;
    /**
     * The name of the load-balancer certificate
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The not valid after validity bound timestamp
     */
    readonly notValidAfter?: pulumi.Input<string>;
    /**
     * The not valid before validity bound timestamp
     */
    readonly notValidBefore?: pulumi.Input<string>;
    /**
     * The status of certificate
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The alternative domain names of the certificate
     */
    readonly subjectAlternativeName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbCertificateBeta resource.
 */
export interface LbCertificateBetaArgs {
    /**
     * The custom type certificate type configuration
     */
    readonly customCertificate?: pulumi.Input<inputs.LbCertificateBetaCustomCertificate>;
    /**
     * The load-balancer ID
     */
    readonly lbId: pulumi.Input<string>;
    /**
     * The Let's Encrypt type certificate configuration
     */
    readonly letsencrypt?: pulumi.Input<inputs.LbCertificateBetaLetsencrypt>;
    /**
     * The name of the load-balancer certificate
     */
    readonly name?: pulumi.Input<string>;
}

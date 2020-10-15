// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getMarketplaceImageBeta(args: GetMarketplaceImageBetaArgs, opts?: pulumi.InvokeOptions): Promise<GetMarketplaceImageBetaResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getMarketplaceImageBeta:getMarketplaceImageBeta", {
        "instanceType": args.instanceType,
        "label": args.label,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getMarketplaceImageBeta.
 */
export interface GetMarketplaceImageBetaArgs {
    readonly instanceType?: string;
    readonly label: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getMarketplaceImageBeta.
 */
export interface GetMarketplaceImageBetaResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceType?: string;
    readonly label: string;
    readonly zone: string;
}

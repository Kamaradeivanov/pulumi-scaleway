// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSecurityGroup(args: GetSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getSecurityGroup:getSecurityGroup", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupArgs {
    readonly name: string;
}

/**
 * A collection of values returned by getSecurityGroup.
 */
export interface GetSecurityGroupResult {
    readonly description: string;
    readonly enableDefaultSecurity: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}

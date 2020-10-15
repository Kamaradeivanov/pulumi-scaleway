// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getInstanceSecurityGroup(args?: GetInstanceSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceSecurityGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getInstanceSecurityGroup:getInstanceSecurityGroup", {
        "name": args.name,
        "securityGroupId": args.securityGroupId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceSecurityGroup.
 */
export interface GetInstanceSecurityGroupArgs {
    readonly name?: string;
    readonly securityGroupId?: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getInstanceSecurityGroup.
 */
export interface GetInstanceSecurityGroupResult {
    readonly description: string;
    readonly externalRules: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly inboundDefaultPolicy: string;
    readonly inboundRules: outputs.GetInstanceSecurityGroupInboundRule[];
    readonly name?: string;
    readonly organizationId: string;
    readonly outboundDefaultPolicy: string;
    readonly outboundRules: outputs.GetInstanceSecurityGroupOutboundRule[];
    readonly securityGroupId?: string;
    readonly stateful: boolean;
    readonly zone?: string;
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getInstanceServer(args?: GetInstanceServerArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceServerResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getInstanceServer:getInstanceServer", {
        "name": args.name,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceServer.
 */
export interface GetInstanceServerArgs {
    readonly name?: string;
    readonly serverId?: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getInstanceServer.
 */
export interface GetInstanceServerResult {
    readonly additionalVolumeIds: string[];
    readonly bootType: string;
    readonly cloudInit: string;
    readonly disableDynamicIp: boolean;
    readonly disablePublicIp: boolean;
    readonly enableDynamicIp: boolean;
    readonly enableIpv6: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly image: string;
    readonly ipId: string;
    readonly ipv6Address: string;
    readonly ipv6Gateway: string;
    readonly ipv6PrefixLength: number;
    readonly name?: string;
    readonly organizationId: string;
    readonly placementGroupId: string;
    readonly placementGroupPolicyRespected: boolean;
    readonly privateIp: string;
    readonly publicIp: string;
    readonly rootVolumes: outputs.GetInstanceServerRootVolume[];
    readonly securityGroupId: string;
    readonly serverId?: string;
    readonly state: string;
    readonly tags: string[];
    readonly type: string;
    readonly userDatas: outputs.GetInstanceServerUserData[];
    readonly zone?: string;
}

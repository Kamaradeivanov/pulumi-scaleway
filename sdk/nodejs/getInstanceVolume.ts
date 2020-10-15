// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getInstanceVolume(args?: GetInstanceVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceVolumeResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getInstanceVolume:getInstanceVolume", {
        "name": args.name,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceVolume.
 */
export interface GetInstanceVolumeArgs {
    readonly name?: string;
    readonly volumeId?: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getInstanceVolume.
 */
export interface GetInstanceVolumeResult {
    readonly fromSnapshotId: string;
    readonly fromVolumeId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly organizationId: string;
    readonly serverId: string;
    readonly sizeInGb: number;
    readonly type: string;
    readonly volumeId?: string;
    readonly zone?: string;
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getVolume(args: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getVolume:getVolume", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeArgs {
    readonly name: string;
}

/**
 * A collection of values returned by getVolume.
 */
export interface GetVolumeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly server: string;
    readonly sizeInGb: number;
    readonly type: string;
}

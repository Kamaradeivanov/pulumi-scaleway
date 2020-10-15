// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("scaleway:index/getImage:getImage", {
        "architecture": args.architecture,
        "mostRecent": args.mostRecent,
        "name": args.name,
        "nameFilter": args.nameFilter,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    readonly architecture: string;
    readonly mostRecent?: boolean;
    readonly name?: string;
    readonly nameFilter?: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    readonly architecture: string;
    readonly creationDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mostRecent?: boolean;
    readonly name: string;
    readonly nameFilter?: string;
    readonly organization: string;
    readonly public: boolean;
}

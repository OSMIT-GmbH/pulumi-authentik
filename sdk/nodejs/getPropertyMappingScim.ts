// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getPropertyMappingScim(args?: GetPropertyMappingScimArgs, opts?: pulumi.InvokeOptions): Promise<GetPropertyMappingScimResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("authentik:index/getPropertyMappingScim:getPropertyMappingScim", {
        "ids": args.ids,
        "managed": args.managed,
        "managedLists": args.managedLists,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPropertyMappingScim.
 */
export interface GetPropertyMappingScimArgs {
    ids?: string[];
    managed?: string;
    managedLists?: string[];
    name?: string;
}

/**
 * A collection of values returned by getPropertyMappingScim.
 */
export interface GetPropertyMappingScimResult {
    readonly expression: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly managed?: string;
    readonly managedLists?: string[];
    readonly name?: string;
}
export function getPropertyMappingScimOutput(args?: GetPropertyMappingScimOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPropertyMappingScimResult> {
    return pulumi.output(args).apply((a: any) => getPropertyMappingScim(a, opts))
}

/**
 * A collection of arguments for invoking getPropertyMappingScim.
 */
export interface GetPropertyMappingScimOutputArgs {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    managed?: pulumi.Input<string>;
    managedLists?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
}

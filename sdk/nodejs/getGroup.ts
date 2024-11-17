// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get groups by pk or name
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * const admins = authentik.getGroup({
 *     name: "authentik Admins",
 * });
 * ```
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("authentik:index/getGroup:getGroup", {
        "includeUsers": args.includeUsers,
        "name": args.name,
        "pk": args.pk,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    includeUsers?: boolean;
    name?: string;
    pk?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * Generated.
     */
    readonly attributes: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeUsers?: boolean;
    /**
     * Generated.
     */
    readonly isSuperuser: boolean;
    readonly name?: string;
    /**
     * Generated.
     */
    readonly numPk: number;
    /**
     * Generated.
     */
    readonly parent: string;
    /**
     * Generated.
     */
    readonly parentName: string;
    readonly pk?: string;
    /**
     * Generated.
     */
    readonly users: number[];
    /**
     * Generated.
     */
    readonly usersObjs: outputs.GetGroupUsersObj[];
}
/**
 * Get groups by pk or name
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * const admins = authentik.getGroup({
 *     name: "authentik Admins",
 * });
 * ```
 */
export function getGroupOutput(args?: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("authentik:index/getGroup:getGroup", {
        "includeUsers": args.includeUsers,
        "name": args.name,
        "pk": args.pk,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    includeUsers?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    pk?: pulumi.Input<string>;
}

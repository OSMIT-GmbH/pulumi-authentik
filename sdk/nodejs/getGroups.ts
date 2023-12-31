// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get groups list
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * const all = authentik.getGroups({});
 * const admins = authentik.getGroups({
 *     isSuperuser: true,
 * });
 * ```
 */
export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("authentik:index/getGroups:getGroups", {
        "attributes": args.attributes,
        "isSuperuser": args.isSuperuser,
        "membersByPks": args.membersByPks,
        "membersByUsernames": args.membersByUsernames,
        "name": args.name,
        "ordering": args.ordering,
        "search": args.search,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    attributes?: string;
    isSuperuser?: boolean;
    membersByPks?: number[];
    membersByUsernames?: string[];
    name?: string;
    ordering?: string;
    search?: string;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    readonly attributes?: string;
    /**
     * Generated.
     */
    readonly groups: outputs.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isSuperuser?: boolean;
    readonly membersByPks?: number[];
    readonly membersByUsernames?: string[];
    readonly name?: string;
    readonly ordering?: string;
    readonly search?: string;
}
/**
 * Get groups list
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * const all = authentik.getGroups({});
 * const admins = authentik.getGroups({
 *     isSuperuser: true,
 * });
 * ```
 */
export function getGroupsOutput(args?: GetGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupsResult> {
    return pulumi.output(args).apply((a: any) => getGroups(a, opts))
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsOutputArgs {
    attributes?: pulumi.Input<string>;
    isSuperuser?: pulumi.Input<boolean>;
    membersByPks?: pulumi.Input<pulumi.Input<number>[]>;
    membersByUsernames?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    ordering?: pulumi.Input<string>;
    search?: pulumi.Input<string>;
}

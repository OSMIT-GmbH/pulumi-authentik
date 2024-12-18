// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get brands by domain
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * // To get the details of a brand by domain
 * const authentik-default = authentik.getBrand({
 *     domain: "authentik-default",
 * });
 * ```
 */
export function getBrand(args?: GetBrandArgs, opts?: pulumi.InvokeOptions): Promise<GetBrandResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("authentik:index/getBrand:getBrand", {
        "brandingFavicon": args.brandingFavicon,
        "brandingLogo": args.brandingLogo,
        "brandingTitle": args.brandingTitle,
        "default": args.default,
        "defaultApplication": args.defaultApplication,
        "domain": args.domain,
        "flowAuthentication": args.flowAuthentication,
        "flowDeviceCode": args.flowDeviceCode,
        "flowInvalidation": args.flowInvalidation,
        "flowRecovery": args.flowRecovery,
        "flowUnenrollment": args.flowUnenrollment,
        "flowUserSettings": args.flowUserSettings,
        "webCertificate": args.webCertificate,
    }, opts);
}

/**
 * A collection of arguments for invoking getBrand.
 */
export interface GetBrandArgs {
    /**
     * Generated.
     */
    brandingFavicon?: string;
    /**
     * Generated.
     */
    brandingLogo?: string;
    /**
     * Generated.
     */
    brandingTitle?: string;
    /**
     * Generated.
     */
    default?: boolean;
    /**
     * Generated.
     */
    defaultApplication?: string;
    /**
     * Generated.
     */
    domain?: string;
    /**
     * Generated.
     */
    flowAuthentication?: string;
    /**
     * Generated.
     */
    flowDeviceCode?: string;
    /**
     * Generated.
     */
    flowInvalidation?: string;
    /**
     * Generated.
     */
    flowRecovery?: string;
    /**
     * Generated.
     */
    flowUnenrollment?: string;
    /**
     * Generated.
     */
    flowUserSettings?: string;
    /**
     * Generated.
     */
    webCertificate?: string;
}

/**
 * A collection of values returned by getBrand.
 */
export interface GetBrandResult {
    /**
     * Generated.
     */
    readonly brandingFavicon: string;
    /**
     * Generated.
     */
    readonly brandingLogo: string;
    /**
     * Generated.
     */
    readonly brandingTitle: string;
    /**
     * Generated.
     */
    readonly default: boolean;
    /**
     * Generated.
     */
    readonly defaultApplication: string;
    /**
     * Generated.
     */
    readonly domain: string;
    /**
     * Generated.
     */
    readonly flowAuthentication: string;
    /**
     * Generated.
     */
    readonly flowDeviceCode: string;
    /**
     * Generated.
     */
    readonly flowInvalidation: string;
    /**
     * Generated.
     */
    readonly flowRecovery: string;
    /**
     * Generated.
     */
    readonly flowUnenrollment: string;
    /**
     * Generated.
     */
    readonly flowUserSettings: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Generated.
     */
    readonly webCertificate: string;
}
/**
 * Get brands by domain
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * // To get the details of a brand by domain
 * const authentik-default = authentik.getBrand({
 *     domain: "authentik-default",
 * });
 * ```
 */
export function getBrandOutput(args?: GetBrandOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBrandResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("authentik:index/getBrand:getBrand", {
        "brandingFavicon": args.brandingFavicon,
        "brandingLogo": args.brandingLogo,
        "brandingTitle": args.brandingTitle,
        "default": args.default,
        "defaultApplication": args.defaultApplication,
        "domain": args.domain,
        "flowAuthentication": args.flowAuthentication,
        "flowDeviceCode": args.flowDeviceCode,
        "flowInvalidation": args.flowInvalidation,
        "flowRecovery": args.flowRecovery,
        "flowUnenrollment": args.flowUnenrollment,
        "flowUserSettings": args.flowUserSettings,
        "webCertificate": args.webCertificate,
    }, opts);
}

/**
 * A collection of arguments for invoking getBrand.
 */
export interface GetBrandOutputArgs {
    /**
     * Generated.
     */
    brandingFavicon?: pulumi.Input<string>;
    /**
     * Generated.
     */
    brandingLogo?: pulumi.Input<string>;
    /**
     * Generated.
     */
    brandingTitle?: pulumi.Input<string>;
    /**
     * Generated.
     */
    default?: pulumi.Input<boolean>;
    /**
     * Generated.
     */
    defaultApplication?: pulumi.Input<string>;
    /**
     * Generated.
     */
    domain?: pulumi.Input<string>;
    /**
     * Generated.
     */
    flowAuthentication?: pulumi.Input<string>;
    /**
     * Generated.
     */
    flowDeviceCode?: pulumi.Input<string>;
    /**
     * Generated.
     */
    flowInvalidation?: pulumi.Input<string>;
    /**
     * Generated.
     */
    flowRecovery?: pulumi.Input<string>;
    /**
     * Generated.
     */
    flowUnenrollment?: pulumi.Input<string>;
    /**
     * Generated.
     */
    flowUserSettings?: pulumi.Input<string>;
    /**
     * Generated.
     */
    webCertificate?: pulumi.Input<string>;
}

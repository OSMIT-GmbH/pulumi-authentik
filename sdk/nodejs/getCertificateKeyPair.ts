// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get certificate-key pairs by name
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * const generated = authentik.getCertificateKeyPair({
 *     name: "authentik Self-signed Certificate",
 * });
 * ```
 */
export function getCertificateKeyPair(args: GetCertificateKeyPairArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateKeyPairResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("authentik:index/getCertificateKeyPair:getCertificateKeyPair", {
        "fetchCertificate": args.fetchCertificate,
        "fetchKey": args.fetchKey,
        "keyData": args.keyData,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificateKeyPair.
 */
export interface GetCertificateKeyPairArgs {
    /**
     * If set to true, certificate data will be fetched. Defaults to `true`.
     */
    fetchCertificate?: boolean;
    /**
     * If set to true, private key data will be fetched. Defaults to `true`.
     */
    fetchKey?: boolean;
    /**
     * Generated.
     */
    keyData?: string;
    name: string;
}

/**
 * A collection of values returned by getCertificateKeyPair.
 */
export interface GetCertificateKeyPairResult {
    /**
     * Generated.
     */
    readonly certificateData: string;
    /**
     * Generated.
     */
    readonly expiry: string;
    /**
     * If set to true, certificate data will be fetched. Defaults to `true`.
     */
    readonly fetchCertificate?: boolean;
    /**
     * If set to true, private key data will be fetched. Defaults to `true`.
     */
    readonly fetchKey?: boolean;
    /**
     * SHA1-hashed certificate fingerprint Generated.
     */
    readonly fingerprint1: string;
    /**
     * SHA256-hashed certificate fingerprint Generated.
     */
    readonly fingerprint256: string;
    /**
     * Generated.
     */
    readonly id: string;
    /**
     * Generated.
     */
    readonly keyData: string;
    readonly name: string;
    /**
     * Generated.
     */
    readonly subject: string;
}
/**
 * Get certificate-key pairs by name
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@pulumi/authentik";
 *
 * const generated = authentik.getCertificateKeyPair({
 *     name: "authentik Self-signed Certificate",
 * });
 * ```
 */
export function getCertificateKeyPairOutput(args: GetCertificateKeyPairOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateKeyPairResult> {
    return pulumi.output(args).apply((a: any) => getCertificateKeyPair(a, opts))
}

/**
 * A collection of arguments for invoking getCertificateKeyPair.
 */
export interface GetCertificateKeyPairOutputArgs {
    /**
     * If set to true, certificate data will be fetched. Defaults to `true`.
     */
    fetchCertificate?: pulumi.Input<boolean>;
    /**
     * If set to true, private key data will be fetched. Defaults to `true`.
     */
    fetchKey?: pulumi.Input<boolean>;
    /**
     * Generated.
     */
    keyData?: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

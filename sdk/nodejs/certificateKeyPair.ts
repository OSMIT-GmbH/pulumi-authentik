// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as authentik from "@osmit-gmbh/pulumi-authentik";
 * import * as tls from "@pulumi/tls";
 *
 * // Generate a certificate-key pair
 * const example = new tls.index.PrivateKey("example", {
 *     algorithm: "ECDSA",
 *     ecdsaCurve: "P384",
 * });
 * const exampleSelfSignedCert = new tls.index.SelfSignedCert("example", {
 *     keyAlgorithm: "ECDSA",
 *     privateKeyPem: example.privateKeyPem,
 *     subject: [{
 *         commonName: "example.com",
 *         organization: "ACME Examples, Inc",
 *     }],
 *     validityPeriodHours: 12,
 *     allowedUses: [
 *         "key_encipherment",
 *         "digital_signature",
 *         "server_auth",
 *     ],
 * });
 * const name = new authentik.CertificateKeyPair("name", {
 *     name: "keypair",
 *     certificateData: exampleSelfSignedCert.certPem,
 *     keyData: example.privateKeyPem,
 * });
 * ```
 */
export class CertificateKeyPair extends pulumi.CustomResource {
    /**
     * Get an existing CertificateKeyPair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateKeyPairState, opts?: pulumi.CustomResourceOptions): CertificateKeyPair {
        return new CertificateKeyPair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/certificateKeyPair:CertificateKeyPair';

    /**
     * Returns true if the given object is an instance of CertificateKeyPair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateKeyPair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateKeyPair.__pulumiType;
    }

    public readonly certificateData!: pulumi.Output<string>;
    public readonly keyData!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a CertificateKeyPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateKeyPairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateKeyPairArgs | CertificateKeyPairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateKeyPairState | undefined;
            resourceInputs["certificateData"] = state ? state.certificateData : undefined;
            resourceInputs["keyData"] = state ? state.keyData : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as CertificateKeyPairArgs | undefined;
            if ((!args || args.certificateData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateData'");
            }
            resourceInputs["certificateData"] = args ? args.certificateData : undefined;
            resourceInputs["keyData"] = args?.keyData ? pulumi.secret(args.keyData) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["keyData"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CertificateKeyPair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateKeyPair resources.
 */
export interface CertificateKeyPairState {
    certificateData?: pulumi.Input<string>;
    keyData?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertificateKeyPair resource.
 */
export interface CertificateKeyPairArgs {
    certificateData: pulumi.Input<string>;
    keyData?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
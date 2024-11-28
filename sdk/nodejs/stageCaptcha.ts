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
 *
 * // Create a captcha stage
 * const name = new authentik.StageCaptcha("name", {
 *     name: "captcha",
 *     privateKey: "foo",
 *     publicKey: "bar",
 * });
 * ```
 */
export class StageCaptcha extends pulumi.CustomResource {
    /**
     * Get an existing StageCaptcha resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageCaptchaState, opts?: pulumi.CustomResourceOptions): StageCaptcha {
        return new StageCaptcha(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageCaptcha:StageCaptcha';

    /**
     * Returns true if the given object is an instance of StageCaptcha.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageCaptcha {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageCaptcha.__pulumiType;
    }

    /**
     * Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
     */
    public readonly apiUrl!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `true`.
     */
    public readonly errorOnInvalidScore!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to `false`.
     */
    public readonly interactive!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
     */
    public readonly jsUrl!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly privateKey!: pulumi.Output<string>;
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * Defaults to `0.5`.
     */
    public readonly scoreMaxThreshold!: pulumi.Output<number | undefined>;
    /**
     * Defaults to `1`.
     */
    public readonly scoreMinThreshold!: pulumi.Output<number | undefined>;

    /**
     * Create a StageCaptcha resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StageCaptchaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageCaptchaArgs | StageCaptchaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageCaptchaState | undefined;
            resourceInputs["apiUrl"] = state ? state.apiUrl : undefined;
            resourceInputs["errorOnInvalidScore"] = state ? state.errorOnInvalidScore : undefined;
            resourceInputs["interactive"] = state ? state.interactive : undefined;
            resourceInputs["jsUrl"] = state ? state.jsUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["scoreMaxThreshold"] = state ? state.scoreMaxThreshold : undefined;
            resourceInputs["scoreMinThreshold"] = state ? state.scoreMinThreshold : undefined;
        } else {
            const args = argsOrState as StageCaptchaArgs | undefined;
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["apiUrl"] = args ? args.apiUrl : undefined;
            resourceInputs["errorOnInvalidScore"] = args ? args.errorOnInvalidScore : undefined;
            resourceInputs["interactive"] = args ? args.interactive : undefined;
            resourceInputs["jsUrl"] = args ? args.jsUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["scoreMaxThreshold"] = args ? args.scoreMaxThreshold : undefined;
            resourceInputs["scoreMinThreshold"] = args ? args.scoreMinThreshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(StageCaptcha.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageCaptcha resources.
 */
export interface StageCaptchaState {
    /**
     * Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
     */
    apiUrl?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    errorOnInvalidScore?: pulumi.Input<boolean>;
    /**
     * Defaults to `false`.
     */
    interactive?: pulumi.Input<boolean>;
    /**
     * Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
     */
    jsUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    privateKey?: pulumi.Input<string>;
    publicKey?: pulumi.Input<string>;
    /**
     * Defaults to `0.5`.
     */
    scoreMaxThreshold?: pulumi.Input<number>;
    /**
     * Defaults to `1`.
     */
    scoreMinThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a StageCaptcha resource.
 */
export interface StageCaptchaArgs {
    /**
     * Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
     */
    apiUrl?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    errorOnInvalidScore?: pulumi.Input<boolean>;
    /**
     * Defaults to `false`.
     */
    interactive?: pulumi.Input<boolean>;
    /**
     * Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
     */
    jsUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    privateKey: pulumi.Input<string>;
    publicKey: pulumi.Input<string>;
    /**
     * Defaults to `0.5`.
     */
    scoreMaxThreshold?: pulumi.Input<number>;
    /**
     * Defaults to `1`.
     */
    scoreMinThreshold?: pulumi.Input<number>;
}

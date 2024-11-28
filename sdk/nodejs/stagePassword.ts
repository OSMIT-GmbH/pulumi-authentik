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
 * // Create a password stage that tests against the interla database
 * const test = new authentik.StagePassword("test", {
 *     name: "test-stage",
 *     backends: ["authentik.core.auth.InbuiltBackend"],
 * });
 * ```
 */
export class StagePassword extends pulumi.CustomResource {
    /**
     * Get an existing StagePassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StagePasswordState, opts?: pulumi.CustomResourceOptions): StagePassword {
        return new StagePassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stagePassword:StagePassword';

    /**
     * Returns true if the given object is an instance of StagePassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StagePassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StagePassword.__pulumiType;
    }

    /**
     * Defaults to `false`.
     */
    public readonly allowShowPassword!: pulumi.Output<boolean | undefined>;
    public readonly backends!: pulumi.Output<string[]>;
    public readonly configureFlow!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `5`.
     */
    public readonly failedAttemptsBeforeCancel!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a StagePassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StagePasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StagePasswordArgs | StagePasswordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StagePasswordState | undefined;
            resourceInputs["allowShowPassword"] = state ? state.allowShowPassword : undefined;
            resourceInputs["backends"] = state ? state.backends : undefined;
            resourceInputs["configureFlow"] = state ? state.configureFlow : undefined;
            resourceInputs["failedAttemptsBeforeCancel"] = state ? state.failedAttemptsBeforeCancel : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as StagePasswordArgs | undefined;
            if ((!args || args.backends === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backends'");
            }
            resourceInputs["allowShowPassword"] = args ? args.allowShowPassword : undefined;
            resourceInputs["backends"] = args ? args.backends : undefined;
            resourceInputs["configureFlow"] = args ? args.configureFlow : undefined;
            resourceInputs["failedAttemptsBeforeCancel"] = args ? args.failedAttemptsBeforeCancel : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StagePassword.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StagePassword resources.
 */
export interface StagePasswordState {
    /**
     * Defaults to `false`.
     */
    allowShowPassword?: pulumi.Input<boolean>;
    backends?: pulumi.Input<pulumi.Input<string>[]>;
    configureFlow?: pulumi.Input<string>;
    /**
     * Defaults to `5`.
     */
    failedAttemptsBeforeCancel?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StagePassword resource.
 */
export interface StagePasswordArgs {
    /**
     * Defaults to `false`.
     */
    allowShowPassword?: pulumi.Input<boolean>;
    backends: pulumi.Input<pulumi.Input<string>[]>;
    configureFlow?: pulumi.Input<string>;
    /**
     * Defaults to `5`.
     */
    failedAttemptsBeforeCancel?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
}
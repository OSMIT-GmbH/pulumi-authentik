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
 * const name = new authentik.StageAuthenticatorValidate("name", {
 *     deviceClasses: ["static"],
 *     notConfiguredAction: "skip",
 * });
 * ```
 */
export class StageAuthenticatorValidate extends pulumi.CustomResource {
    /**
     * Get an existing StageAuthenticatorValidate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageAuthenticatorValidateState, opts?: pulumi.CustomResourceOptions): StageAuthenticatorValidate {
        return new StageAuthenticatorValidate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageAuthenticatorValidate:StageAuthenticatorValidate';

    /**
     * Returns true if the given object is an instance of StageAuthenticatorValidate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageAuthenticatorValidate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageAuthenticatorValidate.__pulumiType;
    }

    public readonly configurationStages!: pulumi.Output<string[] | undefined>;
    public readonly deviceClasses!: pulumi.Output<string[] | undefined>;
    public readonly lastAuthThreshold!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Allowed values: - `skip` - `deny` - `configure`
     */
    public readonly notConfiguredAction!: pulumi.Output<string>;
    public readonly webauthnAllowedDeviceTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Allowed values: - `required` - `preferred` - `discouraged`
     */
    public readonly webauthnUserVerification!: pulumi.Output<string | undefined>;

    /**
     * Create a StageAuthenticatorValidate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StageAuthenticatorValidateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageAuthenticatorValidateArgs | StageAuthenticatorValidateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageAuthenticatorValidateState | undefined;
            resourceInputs["configurationStages"] = state ? state.configurationStages : undefined;
            resourceInputs["deviceClasses"] = state ? state.deviceClasses : undefined;
            resourceInputs["lastAuthThreshold"] = state ? state.lastAuthThreshold : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notConfiguredAction"] = state ? state.notConfiguredAction : undefined;
            resourceInputs["webauthnAllowedDeviceTypes"] = state ? state.webauthnAllowedDeviceTypes : undefined;
            resourceInputs["webauthnUserVerification"] = state ? state.webauthnUserVerification : undefined;
        } else {
            const args = argsOrState as StageAuthenticatorValidateArgs | undefined;
            if ((!args || args.notConfiguredAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notConfiguredAction'");
            }
            resourceInputs["configurationStages"] = args ? args.configurationStages : undefined;
            resourceInputs["deviceClasses"] = args ? args.deviceClasses : undefined;
            resourceInputs["lastAuthThreshold"] = args ? args.lastAuthThreshold : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notConfiguredAction"] = args ? args.notConfiguredAction : undefined;
            resourceInputs["webauthnAllowedDeviceTypes"] = args ? args.webauthnAllowedDeviceTypes : undefined;
            resourceInputs["webauthnUserVerification"] = args ? args.webauthnUserVerification : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StageAuthenticatorValidate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageAuthenticatorValidate resources.
 */
export interface StageAuthenticatorValidateState {
    configurationStages?: pulumi.Input<pulumi.Input<string>[]>;
    deviceClasses?: pulumi.Input<pulumi.Input<string>[]>;
    lastAuthThreshold?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Allowed values: - `skip` - `deny` - `configure`
     */
    notConfiguredAction?: pulumi.Input<string>;
    webauthnAllowedDeviceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed values: - `required` - `preferred` - `discouraged`
     */
    webauthnUserVerification?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StageAuthenticatorValidate resource.
 */
export interface StageAuthenticatorValidateArgs {
    configurationStages?: pulumi.Input<pulumi.Input<string>[]>;
    deviceClasses?: pulumi.Input<pulumi.Input<string>[]>;
    lastAuthThreshold?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Allowed values: - `skip` - `deny` - `configure`
     */
    notConfiguredAction: pulumi.Input<string>;
    webauthnAllowedDeviceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed values: - `required` - `preferred` - `discouraged`
     */
    webauthnUserVerification?: pulumi.Input<string>;
}

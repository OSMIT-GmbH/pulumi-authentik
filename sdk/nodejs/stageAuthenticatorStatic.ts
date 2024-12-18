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
 * // Create a static TOTP Setup stage
 * const name = new authentik.StageAuthenticatorStatic("name", {name: "static-totp-setup"});
 * ```
 */
export class StageAuthenticatorStatic extends pulumi.CustomResource {
    /**
     * Get an existing StageAuthenticatorStatic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageAuthenticatorStaticState, opts?: pulumi.CustomResourceOptions): StageAuthenticatorStatic {
        return new StageAuthenticatorStatic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageAuthenticatorStatic:StageAuthenticatorStatic';

    /**
     * Returns true if the given object is an instance of StageAuthenticatorStatic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageAuthenticatorStatic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageAuthenticatorStatic.__pulumiType;
    }

    public readonly configureFlow!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Defaults to `6`.
     */
    public readonly tokenCount!: pulumi.Output<number | undefined>;
    /**
     * Defaults to `12`.
     */
    public readonly tokenLength!: pulumi.Output<number | undefined>;

    /**
     * Create a StageAuthenticatorStatic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StageAuthenticatorStaticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageAuthenticatorStaticArgs | StageAuthenticatorStaticState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageAuthenticatorStaticState | undefined;
            resourceInputs["configureFlow"] = state ? state.configureFlow : undefined;
            resourceInputs["friendlyName"] = state ? state.friendlyName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tokenCount"] = state ? state.tokenCount : undefined;
            resourceInputs["tokenLength"] = state ? state.tokenLength : undefined;
        } else {
            const args = argsOrState as StageAuthenticatorStaticArgs | undefined;
            resourceInputs["configureFlow"] = args ? args.configureFlow : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tokenCount"] = args ? args.tokenCount : undefined;
            resourceInputs["tokenLength"] = args ? args.tokenLength : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StageAuthenticatorStatic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageAuthenticatorStatic resources.
 */
export interface StageAuthenticatorStaticState {
    configureFlow?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Defaults to `6`.
     */
    tokenCount?: pulumi.Input<number>;
    /**
     * Defaults to `12`.
     */
    tokenLength?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a StageAuthenticatorStatic resource.
 */
export interface StageAuthenticatorStaticArgs {
    configureFlow?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Defaults to `6`.
     */
    tokenCount?: pulumi.Input<number>;
    /**
     * Defaults to `12`.
     */
    tokenLength?: pulumi.Input<number>;
}

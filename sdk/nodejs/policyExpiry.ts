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
 * // Create expiry policy
 * const name = new authentik.PolicyExpiry("name", {
 *     name: "expiry",
 *     days: 3,
 * });
 * ```
 */
export class PolicyExpiry extends pulumi.CustomResource {
    /**
     * Get an existing PolicyExpiry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyExpiryState, opts?: pulumi.CustomResourceOptions): PolicyExpiry {
        return new PolicyExpiry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/policyExpiry:PolicyExpiry';

    /**
     * Returns true if the given object is an instance of PolicyExpiry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyExpiry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyExpiry.__pulumiType;
    }

    public readonly days!: pulumi.Output<number>;
    /**
     * Defaults to `false`.
     */
    public readonly denyOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to `false`.
     */
    public readonly executionLogging!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a PolicyExpiry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyExpiryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyExpiryArgs | PolicyExpiryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyExpiryState | undefined;
            resourceInputs["days"] = state ? state.days : undefined;
            resourceInputs["denyOnly"] = state ? state.denyOnly : undefined;
            resourceInputs["executionLogging"] = state ? state.executionLogging : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PolicyExpiryArgs | undefined;
            if ((!args || args.days === undefined) && !opts.urn) {
                throw new Error("Missing required property 'days'");
            }
            resourceInputs["days"] = args ? args.days : undefined;
            resourceInputs["denyOnly"] = args ? args.denyOnly : undefined;
            resourceInputs["executionLogging"] = args ? args.executionLogging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyExpiry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyExpiry resources.
 */
export interface PolicyExpiryState {
    days?: pulumi.Input<number>;
    /**
     * Defaults to `false`.
     */
    denyOnly?: pulumi.Input<boolean>;
    /**
     * Defaults to `false`.
     */
    executionLogging?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyExpiry resource.
 */
export interface PolicyExpiryArgs {
    days: pulumi.Input<number>;
    /**
     * Defaults to `false`.
     */
    denyOnly?: pulumi.Input<boolean>;
    /**
     * Defaults to `false`.
     */
    executionLogging?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
}
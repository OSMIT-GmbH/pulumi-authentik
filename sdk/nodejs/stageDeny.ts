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
 * const name = new authentik.StageDeny("name", {});
 * ```
 */
export class StageDeny extends pulumi.CustomResource {
    /**
     * Get an existing StageDeny resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageDenyState, opts?: pulumi.CustomResourceOptions): StageDeny {
        return new StageDeny(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageDeny:StageDeny';

    /**
     * Returns true if the given object is an instance of StageDeny.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageDeny {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageDeny.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;

    /**
     * Create a StageDeny resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StageDenyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageDenyArgs | StageDenyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageDenyState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as StageDenyArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StageDeny.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageDeny resources.
 */
export interface StageDenyState {
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StageDeny resource.
 */
export interface StageDenyArgs {
    name?: pulumi.Input<string>;
}

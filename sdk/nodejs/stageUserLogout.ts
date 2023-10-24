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
 * const name = new authentik.StageUserLogout("name", {});
 * ```
 */
export class StageUserLogout extends pulumi.CustomResource {
    /**
     * Get an existing StageUserLogout resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageUserLogoutState, opts?: pulumi.CustomResourceOptions): StageUserLogout {
        return new StageUserLogout(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageUserLogout:StageUserLogout';

    /**
     * Returns true if the given object is an instance of StageUserLogout.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageUserLogout {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageUserLogout.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;

    /**
     * Create a StageUserLogout resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StageUserLogoutArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageUserLogoutArgs | StageUserLogoutState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageUserLogoutState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as StageUserLogoutArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StageUserLogout.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageUserLogout resources.
 */
export interface StageUserLogoutState {
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StageUserLogout resource.
 */
export interface StageUserLogoutArgs {
    name?: pulumi.Input<string>;
}

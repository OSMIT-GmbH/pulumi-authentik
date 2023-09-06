// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class StageUserWrite extends pulumi.CustomResource {
    /**
     * Get an existing StageUserWrite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageUserWriteState, opts?: pulumi.CustomResourceOptions): StageUserWrite {
        return new StageUserWrite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageUserWrite:StageUserWrite';

    /**
     * Returns true if the given object is an instance of StageUserWrite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageUserWrite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageUserWrite.__pulumiType;
    }

    public readonly createUsersAsInactive!: pulumi.Output<boolean | undefined>;
    public readonly createUsersGroup!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly userCreationMode!: pulumi.Output<string | undefined>;
    public readonly userPathTemplate!: pulumi.Output<string | undefined>;

    /**
     * Create a StageUserWrite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StageUserWriteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageUserWriteArgs | StageUserWriteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageUserWriteState | undefined;
            resourceInputs["createUsersAsInactive"] = state ? state.createUsersAsInactive : undefined;
            resourceInputs["createUsersGroup"] = state ? state.createUsersGroup : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["userCreationMode"] = state ? state.userCreationMode : undefined;
            resourceInputs["userPathTemplate"] = state ? state.userPathTemplate : undefined;
        } else {
            const args = argsOrState as StageUserWriteArgs | undefined;
            resourceInputs["createUsersAsInactive"] = args ? args.createUsersAsInactive : undefined;
            resourceInputs["createUsersGroup"] = args ? args.createUsersGroup : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["userCreationMode"] = args ? args.userCreationMode : undefined;
            resourceInputs["userPathTemplate"] = args ? args.userPathTemplate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StageUserWrite.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageUserWrite resources.
 */
export interface StageUserWriteState {
    createUsersAsInactive?: pulumi.Input<boolean>;
    createUsersGroup?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    userCreationMode?: pulumi.Input<string>;
    userPathTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StageUserWrite resource.
 */
export interface StageUserWriteArgs {
    createUsersAsInactive?: pulumi.Input<boolean>;
    createUsersGroup?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    userCreationMode?: pulumi.Input<string>;
    userPathTemplate?: pulumi.Input<string>;
}

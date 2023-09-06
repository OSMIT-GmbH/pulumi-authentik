// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class StageAuthenticatorWebauthn extends pulumi.CustomResource {
    /**
     * Get an existing StageAuthenticatorWebauthn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageAuthenticatorWebauthnState, opts?: pulumi.CustomResourceOptions): StageAuthenticatorWebauthn {
        return new StageAuthenticatorWebauthn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageAuthenticatorWebauthn:StageAuthenticatorWebauthn';

    /**
     * Returns true if the given object is an instance of StageAuthenticatorWebauthn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageAuthenticatorWebauthn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageAuthenticatorWebauthn.__pulumiType;
    }

    public readonly authenticatorAttachment!: pulumi.Output<string | undefined>;
    public readonly configureFlow!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly residentKeyRequirement!: pulumi.Output<string | undefined>;
    public readonly userVerification!: pulumi.Output<string | undefined>;

    /**
     * Create a StageAuthenticatorWebauthn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StageAuthenticatorWebauthnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageAuthenticatorWebauthnArgs | StageAuthenticatorWebauthnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageAuthenticatorWebauthnState | undefined;
            resourceInputs["authenticatorAttachment"] = state ? state.authenticatorAttachment : undefined;
            resourceInputs["configureFlow"] = state ? state.configureFlow : undefined;
            resourceInputs["friendlyName"] = state ? state.friendlyName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["residentKeyRequirement"] = state ? state.residentKeyRequirement : undefined;
            resourceInputs["userVerification"] = state ? state.userVerification : undefined;
        } else {
            const args = argsOrState as StageAuthenticatorWebauthnArgs | undefined;
            resourceInputs["authenticatorAttachment"] = args ? args.authenticatorAttachment : undefined;
            resourceInputs["configureFlow"] = args ? args.configureFlow : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["residentKeyRequirement"] = args ? args.residentKeyRequirement : undefined;
            resourceInputs["userVerification"] = args ? args.userVerification : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StageAuthenticatorWebauthn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageAuthenticatorWebauthn resources.
 */
export interface StageAuthenticatorWebauthnState {
    authenticatorAttachment?: pulumi.Input<string>;
    configureFlow?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    residentKeyRequirement?: pulumi.Input<string>;
    userVerification?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StageAuthenticatorWebauthn resource.
 */
export interface StageAuthenticatorWebauthnArgs {
    authenticatorAttachment?: pulumi.Input<string>;
    configureFlow?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    residentKeyRequirement?: pulumi.Input<string>;
    userVerification?: pulumi.Input<string>;
}

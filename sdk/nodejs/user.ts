// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * JSON format expected. Use jsonencode() to pass objects.
     */
    public readonly attributes!: pulumi.Output<string | undefined>;
    public readonly email!: pulumi.Output<string | undefined>;
    public readonly groups!: pulumi.Output<string[]>;
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly path!: pulumi.Output<string | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * JSON format expected. Use jsonencode() to pass objects.
     */
    attributes?: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    isActive?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    path?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * JSON format expected. Use jsonencode() to pass objects.
     */
    attributes?: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    isActive?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    path?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

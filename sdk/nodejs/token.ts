// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Token extends pulumi.CustomResource {
    /**
     * Get an existing Token resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TokenState, opts?: pulumi.CustomResourceOptions): Token {
        return new Token(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/token:Token';

    /**
     * Returns true if the given object is an instance of Token.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Token {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Token.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly expires!: pulumi.Output<string | undefined>;
    /**
     * Generated.
     */
    public /*out*/ readonly expiresIn!: pulumi.Output<number>;
    /**
     * Defaults to `true`.
     */
    public readonly expiring!: pulumi.Output<boolean | undefined>;
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Defaults to `api`.
     */
    public readonly intent!: pulumi.Output<string | undefined>;
    /**
     * Generated.
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * Defaults to `false`.
     */
    public readonly retrieveKey!: pulumi.Output<boolean | undefined>;
    public readonly user!: pulumi.Output<number>;

    /**
     * Create a Token resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TokenArgs | TokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TokenState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expires"] = state ? state.expires : undefined;
            resourceInputs["expiresIn"] = state ? state.expiresIn : undefined;
            resourceInputs["expiring"] = state ? state.expiring : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["intent"] = state ? state.intent : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["retrieveKey"] = state ? state.retrieveKey : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as TokenArgs | undefined;
            if ((!args || args.identifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifier'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expires"] = args ? args.expires : undefined;
            resourceInputs["expiring"] = args ? args.expiring : undefined;
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["intent"] = args ? args.intent : undefined;
            resourceInputs["retrieveKey"] = args ? args.retrieveKey : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["expiresIn"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Token.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Token resources.
 */
export interface TokenState {
    description?: pulumi.Input<string>;
    expires?: pulumi.Input<string>;
    /**
     * Generated.
     */
    expiresIn?: pulumi.Input<number>;
    /**
     * Defaults to `true`.
     */
    expiring?: pulumi.Input<boolean>;
    identifier?: pulumi.Input<string>;
    /**
     * Defaults to `api`.
     */
    intent?: pulumi.Input<string>;
    /**
     * Generated.
     */
    key?: pulumi.Input<string>;
    /**
     * Defaults to `false`.
     */
    retrieveKey?: pulumi.Input<boolean>;
    user?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Token resource.
 */
export interface TokenArgs {
    description?: pulumi.Input<string>;
    expires?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    expiring?: pulumi.Input<boolean>;
    identifier: pulumi.Input<string>;
    /**
     * Defaults to `api`.
     */
    intent?: pulumi.Input<string>;
    /**
     * Defaults to `false`.
     */
    retrieveKey?: pulumi.Input<boolean>;
    user: pulumi.Input<number>;
}

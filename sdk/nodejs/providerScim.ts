// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class ProviderScim extends pulumi.CustomResource {
    /**
     * Get an existing ProviderScim resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProviderScimState, opts?: pulumi.CustomResourceOptions): ProviderScim {
        return new ProviderScim(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/providerScim:ProviderScim';

    /**
     * Returns true if the given object is an instance of ProviderScim.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderScim {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderScim.__pulumiType;
    }

    public readonly excludeUsersServiceAccount!: pulumi.Output<boolean | undefined>;
    public readonly filterGroup!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly propertyMappings!: pulumi.Output<string[] | undefined>;
    public readonly propertyMappingsGroups!: pulumi.Output<string[] | undefined>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ProviderScim resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderScimArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProviderScimArgs | ProviderScimState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProviderScimState | undefined;
            resourceInputs["excludeUsersServiceAccount"] = state ? state.excludeUsersServiceAccount : undefined;
            resourceInputs["filterGroup"] = state ? state.filterGroup : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["propertyMappings"] = state ? state.propertyMappings : undefined;
            resourceInputs["propertyMappingsGroups"] = state ? state.propertyMappingsGroups : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ProviderScimArgs | undefined;
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["excludeUsersServiceAccount"] = args ? args.excludeUsersServiceAccount : undefined;
            resourceInputs["filterGroup"] = args ? args.filterGroup : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["propertyMappings"] = args ? args.propertyMappings : undefined;
            resourceInputs["propertyMappingsGroups"] = args ? args.propertyMappingsGroups : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProviderScim.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProviderScim resources.
 */
export interface ProviderScimState {
    excludeUsersServiceAccount?: pulumi.Input<boolean>;
    filterGroup?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    propertyMappingsGroups?: pulumi.Input<pulumi.Input<string>[]>;
    token?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderScim resource.
 */
export interface ProviderScimArgs {
    excludeUsersServiceAccount?: pulumi.Input<boolean>;
    filterGroup?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    propertyMappingsGroups?: pulumi.Input<pulumi.Input<string>[]>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
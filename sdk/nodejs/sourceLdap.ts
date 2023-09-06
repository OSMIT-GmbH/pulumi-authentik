// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SourceLdap extends pulumi.CustomResource {
    /**
     * Get an existing SourceLdap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceLdapState, opts?: pulumi.CustomResourceOptions): SourceLdap {
        return new SourceLdap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/sourceLdap:SourceLdap';

    /**
     * Returns true if the given object is an instance of SourceLdap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceLdap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceLdap.__pulumiType;
    }

    public readonly additionalGroupDn!: pulumi.Output<string | undefined>;
    public readonly additionalUserDn!: pulumi.Output<string | undefined>;
    public readonly baseDn!: pulumi.Output<string>;
    public readonly bindCn!: pulumi.Output<string>;
    public readonly bindPassword!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly groupMembershipField!: pulumi.Output<string | undefined>;
    public readonly groupObjectFilter!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly objectUniquenessField!: pulumi.Output<string | undefined>;
    public readonly propertyMappings!: pulumi.Output<string[] | undefined>;
    public readonly propertyMappingsGroups!: pulumi.Output<string[] | undefined>;
    public readonly serverUri!: pulumi.Output<string>;
    public readonly slug!: pulumi.Output<string>;
    public readonly startTls!: pulumi.Output<boolean | undefined>;
    public readonly syncGroups!: pulumi.Output<boolean | undefined>;
    public readonly syncParentGroup!: pulumi.Output<string | undefined>;
    public readonly syncUsers!: pulumi.Output<boolean | undefined>;
    public readonly syncUsersPassword!: pulumi.Output<boolean | undefined>;
    public readonly userObjectFilter!: pulumi.Output<string | undefined>;
    public readonly userPathTemplate!: pulumi.Output<string | undefined>;
    public readonly uuid!: pulumi.Output<string>;

    /**
     * Create a SourceLdap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceLdapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceLdapArgs | SourceLdapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SourceLdapState | undefined;
            resourceInputs["additionalGroupDn"] = state ? state.additionalGroupDn : undefined;
            resourceInputs["additionalUserDn"] = state ? state.additionalUserDn : undefined;
            resourceInputs["baseDn"] = state ? state.baseDn : undefined;
            resourceInputs["bindCn"] = state ? state.bindCn : undefined;
            resourceInputs["bindPassword"] = state ? state.bindPassword : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["groupMembershipField"] = state ? state.groupMembershipField : undefined;
            resourceInputs["groupObjectFilter"] = state ? state.groupObjectFilter : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["objectUniquenessField"] = state ? state.objectUniquenessField : undefined;
            resourceInputs["propertyMappings"] = state ? state.propertyMappings : undefined;
            resourceInputs["propertyMappingsGroups"] = state ? state.propertyMappingsGroups : undefined;
            resourceInputs["serverUri"] = state ? state.serverUri : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["startTls"] = state ? state.startTls : undefined;
            resourceInputs["syncGroups"] = state ? state.syncGroups : undefined;
            resourceInputs["syncParentGroup"] = state ? state.syncParentGroup : undefined;
            resourceInputs["syncUsers"] = state ? state.syncUsers : undefined;
            resourceInputs["syncUsersPassword"] = state ? state.syncUsersPassword : undefined;
            resourceInputs["userObjectFilter"] = state ? state.userObjectFilter : undefined;
            resourceInputs["userPathTemplate"] = state ? state.userPathTemplate : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as SourceLdapArgs | undefined;
            if ((!args || args.baseDn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseDn'");
            }
            if ((!args || args.bindCn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindCn'");
            }
            if ((!args || args.bindPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindPassword'");
            }
            if ((!args || args.serverUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverUri'");
            }
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            resourceInputs["additionalGroupDn"] = args ? args.additionalGroupDn : undefined;
            resourceInputs["additionalUserDn"] = args ? args.additionalUserDn : undefined;
            resourceInputs["baseDn"] = args ? args.baseDn : undefined;
            resourceInputs["bindCn"] = args ? args.bindCn : undefined;
            resourceInputs["bindPassword"] = args?.bindPassword ? pulumi.secret(args.bindPassword) : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["groupMembershipField"] = args ? args.groupMembershipField : undefined;
            resourceInputs["groupObjectFilter"] = args ? args.groupObjectFilter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["objectUniquenessField"] = args ? args.objectUniquenessField : undefined;
            resourceInputs["propertyMappings"] = args ? args.propertyMappings : undefined;
            resourceInputs["propertyMappingsGroups"] = args ? args.propertyMappingsGroups : undefined;
            resourceInputs["serverUri"] = args ? args.serverUri : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["startTls"] = args ? args.startTls : undefined;
            resourceInputs["syncGroups"] = args ? args.syncGroups : undefined;
            resourceInputs["syncParentGroup"] = args ? args.syncParentGroup : undefined;
            resourceInputs["syncUsers"] = args ? args.syncUsers : undefined;
            resourceInputs["syncUsersPassword"] = args ? args.syncUsersPassword : undefined;
            resourceInputs["userObjectFilter"] = args ? args.userObjectFilter : undefined;
            resourceInputs["userPathTemplate"] = args ? args.userPathTemplate : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["bindPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SourceLdap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceLdap resources.
 */
export interface SourceLdapState {
    additionalGroupDn?: pulumi.Input<string>;
    additionalUserDn?: pulumi.Input<string>;
    baseDn?: pulumi.Input<string>;
    bindCn?: pulumi.Input<string>;
    bindPassword?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    groupMembershipField?: pulumi.Input<string>;
    groupObjectFilter?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    objectUniquenessField?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    propertyMappingsGroups?: pulumi.Input<pulumi.Input<string>[]>;
    serverUri?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
    startTls?: pulumi.Input<boolean>;
    syncGroups?: pulumi.Input<boolean>;
    syncParentGroup?: pulumi.Input<string>;
    syncUsers?: pulumi.Input<boolean>;
    syncUsersPassword?: pulumi.Input<boolean>;
    userObjectFilter?: pulumi.Input<string>;
    userPathTemplate?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SourceLdap resource.
 */
export interface SourceLdapArgs {
    additionalGroupDn?: pulumi.Input<string>;
    additionalUserDn?: pulumi.Input<string>;
    baseDn: pulumi.Input<string>;
    bindCn: pulumi.Input<string>;
    bindPassword: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    groupMembershipField?: pulumi.Input<string>;
    groupObjectFilter?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    objectUniquenessField?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    propertyMappingsGroups?: pulumi.Input<pulumi.Input<string>[]>;
    serverUri: pulumi.Input<string>;
    slug: pulumi.Input<string>;
    startTls?: pulumi.Input<boolean>;
    syncGroups?: pulumi.Input<boolean>;
    syncParentGroup?: pulumi.Input<string>;
    syncUsers?: pulumi.Input<boolean>;
    syncUsersPassword?: pulumi.Input<boolean>;
    userObjectFilter?: pulumi.Input<string>;
    userPathTemplate?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
}

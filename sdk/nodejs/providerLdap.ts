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
 * import * as authentik from "@pulumi/authentik";
 *
 * // Create an LDAP Provider
 * const default-authentication-flow = authentik.getFlow({
 *     slug: "default-authentication-flow",
 * });
 * const name = new authentik.ProviderLdap("name", {
 *     name: "ldap-app",
 *     baseDn: "dc=ldap,dc=goauthentik,dc=io",
 *     bindFlow: default_authentication_flow.then(default_authentication_flow => default_authentication_flow.id),
 * });
 * const nameApplication = new authentik.Application("name", {
 *     name: "ldap-app",
 *     slug: "ldap-app",
 *     protocolProvider: name.id,
 * });
 * ```
 */
export class ProviderLdap extends pulumi.CustomResource {
    /**
     * Get an existing ProviderLdap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProviderLdapState, opts?: pulumi.CustomResourceOptions): ProviderLdap {
        return new ProviderLdap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/providerLdap:ProviderLdap';

    /**
     * Returns true if the given object is an instance of ProviderLdap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderLdap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderLdap.__pulumiType;
    }

    public readonly baseDn!: pulumi.Output<string>;
    public readonly bindFlow!: pulumi.Output<string>;
    /**
     * Defaults to `direct`.
     */
    public readonly bindMode!: pulumi.Output<string | undefined>;
    public readonly certificate!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `4000`.
     */
    public readonly gidStartNumber!: pulumi.Output<number | undefined>;
    /**
     * Defaults to `true`.
     */
    public readonly mfaSupport!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Defaults to `direct`.
     */
    public readonly searchMode!: pulumi.Output<string | undefined>;
    public readonly tlsServerName!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `2000`.
     */
    public readonly uidStartNumber!: pulumi.Output<number | undefined>;
    public readonly unbindFlow!: pulumi.Output<string>;

    /**
     * Create a ProviderLdap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderLdapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProviderLdapArgs | ProviderLdapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProviderLdapState | undefined;
            resourceInputs["baseDn"] = state ? state.baseDn : undefined;
            resourceInputs["bindFlow"] = state ? state.bindFlow : undefined;
            resourceInputs["bindMode"] = state ? state.bindMode : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["gidStartNumber"] = state ? state.gidStartNumber : undefined;
            resourceInputs["mfaSupport"] = state ? state.mfaSupport : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["searchMode"] = state ? state.searchMode : undefined;
            resourceInputs["tlsServerName"] = state ? state.tlsServerName : undefined;
            resourceInputs["uidStartNumber"] = state ? state.uidStartNumber : undefined;
            resourceInputs["unbindFlow"] = state ? state.unbindFlow : undefined;
        } else {
            const args = argsOrState as ProviderLdapArgs | undefined;
            if ((!args || args.baseDn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseDn'");
            }
            if ((!args || args.bindFlow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindFlow'");
            }
            if ((!args || args.unbindFlow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'unbindFlow'");
            }
            resourceInputs["baseDn"] = args ? args.baseDn : undefined;
            resourceInputs["bindFlow"] = args ? args.bindFlow : undefined;
            resourceInputs["bindMode"] = args ? args.bindMode : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["gidStartNumber"] = args ? args.gidStartNumber : undefined;
            resourceInputs["mfaSupport"] = args ? args.mfaSupport : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["searchMode"] = args ? args.searchMode : undefined;
            resourceInputs["tlsServerName"] = args ? args.tlsServerName : undefined;
            resourceInputs["uidStartNumber"] = args ? args.uidStartNumber : undefined;
            resourceInputs["unbindFlow"] = args ? args.unbindFlow : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProviderLdap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProviderLdap resources.
 */
export interface ProviderLdapState {
    baseDn?: pulumi.Input<string>;
    bindFlow?: pulumi.Input<string>;
    /**
     * Defaults to `direct`.
     */
    bindMode?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    /**
     * Defaults to `4000`.
     */
    gidStartNumber?: pulumi.Input<number>;
    /**
     * Defaults to `true`.
     */
    mfaSupport?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * Defaults to `direct`.
     */
    searchMode?: pulumi.Input<string>;
    tlsServerName?: pulumi.Input<string>;
    /**
     * Defaults to `2000`.
     */
    uidStartNumber?: pulumi.Input<number>;
    unbindFlow?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderLdap resource.
 */
export interface ProviderLdapArgs {
    baseDn: pulumi.Input<string>;
    bindFlow: pulumi.Input<string>;
    /**
     * Defaults to `direct`.
     */
    bindMode?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    /**
     * Defaults to `4000`.
     */
    gidStartNumber?: pulumi.Input<number>;
    /**
     * Defaults to `true`.
     */
    mfaSupport?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * Defaults to `direct`.
     */
    searchMode?: pulumi.Input<string>;
    tlsServerName?: pulumi.Input<string>;
    /**
     * Defaults to `2000`.
     */
    uidStartNumber?: pulumi.Input<number>;
    unbindFlow: pulumi.Input<string>;
}
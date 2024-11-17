// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SourceKerberos extends pulumi.CustomResource {
    /**
     * Get an existing SourceKerberos resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceKerberosState, opts?: pulumi.CustomResourceOptions): SourceKerberos {
        return new SourceKerberos(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/sourceKerberos:SourceKerberos';

    /**
     * Returns true if the given object is an instance of SourceKerberos.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceKerberos {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceKerberos.__pulumiType;
    }

    public readonly authenticationFlow!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly enrollmentFlow!: pulumi.Output<string | undefined>;
    /**
     * Allowed values: - `identifier` - `nameLink` - `nameDeny`
     */
    public readonly groupMatchingMode!: pulumi.Output<string | undefined>;
    /**
     * Custom krb5.conf to use. Uses the system one by default
     */
    public readonly krb5Conf!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend
     */
    public readonly passwordLoginUpdateInternalPassword!: pulumi.Output<boolean | undefined>;
    /**
     * Allowed values: - `all` - `any`
     */
    public readonly policyEngineMode!: pulumi.Output<string | undefined>;
    /**
     * Kerberos realm
     */
    public readonly realm!: pulumi.Output<string>;
    public readonly slug!: pulumi.Output<string>;
    /**
     * Credential cache to use for SPNEGO in form type:residual
     */
    public readonly spnegoCcache!: pulumi.Output<string | undefined>;
    /**
     * SPNEGO keytab base64-encoded or path to keytab in the form FILE:path
     */
    public readonly spnegoKeytab!: pulumi.Output<string | undefined>;
    /**
     * Force the use of a specific server name for SPNEGO
     */
    public readonly spnegoServerName!: pulumi.Output<string | undefined>;
    /**
     * Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual
     */
    public readonly syncCcache!: pulumi.Output<string | undefined>;
    /**
     * Keytab to authenticate to kadmin for sync. Must be base64-encoded or in the form TYPE:residual
     */
    public readonly syncKeytab!: pulumi.Output<string | undefined>;
    /**
     * Password to authenticate to kadmin for sync
     */
    public readonly syncPassword!: pulumi.Output<string | undefined>;
    /**
     * Principal to authenticate to kadmin for sync.
     */
    public readonly syncPrincipal!: pulumi.Output<string | undefined>;
    /**
     * Sync users from Kerberos into authentik
     */
    public readonly syncUsers!: pulumi.Output<boolean | undefined>;
    /**
     * When a user changes their password, sync it back to Kerberos
     */
    public readonly syncUsersPassword!: pulumi.Output<boolean | undefined>;
    /**
     * Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
     */
    public readonly userMatchingMode!: pulumi.Output<string | undefined>;
    public readonly userPathTemplate!: pulumi.Output<string | undefined>;
    public readonly uuid!: pulumi.Output<string>;

    /**
     * Create a SourceKerberos resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceKerberosArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceKerberosArgs | SourceKerberosState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SourceKerberosState | undefined;
            resourceInputs["authenticationFlow"] = state ? state.authenticationFlow : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["enrollmentFlow"] = state ? state.enrollmentFlow : undefined;
            resourceInputs["groupMatchingMode"] = state ? state.groupMatchingMode : undefined;
            resourceInputs["krb5Conf"] = state ? state.krb5Conf : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["passwordLoginUpdateInternalPassword"] = state ? state.passwordLoginUpdateInternalPassword : undefined;
            resourceInputs["policyEngineMode"] = state ? state.policyEngineMode : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["spnegoCcache"] = state ? state.spnegoCcache : undefined;
            resourceInputs["spnegoKeytab"] = state ? state.spnegoKeytab : undefined;
            resourceInputs["spnegoServerName"] = state ? state.spnegoServerName : undefined;
            resourceInputs["syncCcache"] = state ? state.syncCcache : undefined;
            resourceInputs["syncKeytab"] = state ? state.syncKeytab : undefined;
            resourceInputs["syncPassword"] = state ? state.syncPassword : undefined;
            resourceInputs["syncPrincipal"] = state ? state.syncPrincipal : undefined;
            resourceInputs["syncUsers"] = state ? state.syncUsers : undefined;
            resourceInputs["syncUsersPassword"] = state ? state.syncUsersPassword : undefined;
            resourceInputs["userMatchingMode"] = state ? state.userMatchingMode : undefined;
            resourceInputs["userPathTemplate"] = state ? state.userPathTemplate : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as SourceKerberosArgs | undefined;
            if ((!args || args.realm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realm'");
            }
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            resourceInputs["authenticationFlow"] = args ? args.authenticationFlow : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["enrollmentFlow"] = args ? args.enrollmentFlow : undefined;
            resourceInputs["groupMatchingMode"] = args ? args.groupMatchingMode : undefined;
            resourceInputs["krb5Conf"] = args ? args.krb5Conf : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passwordLoginUpdateInternalPassword"] = args ? args.passwordLoginUpdateInternalPassword : undefined;
            resourceInputs["policyEngineMode"] = args ? args.policyEngineMode : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["spnegoCcache"] = args ? args.spnegoCcache : undefined;
            resourceInputs["spnegoKeytab"] = args?.spnegoKeytab ? pulumi.secret(args.spnegoKeytab) : undefined;
            resourceInputs["spnegoServerName"] = args ? args.spnegoServerName : undefined;
            resourceInputs["syncCcache"] = args ? args.syncCcache : undefined;
            resourceInputs["syncKeytab"] = args?.syncKeytab ? pulumi.secret(args.syncKeytab) : undefined;
            resourceInputs["syncPassword"] = args?.syncPassword ? pulumi.secret(args.syncPassword) : undefined;
            resourceInputs["syncPrincipal"] = args ? args.syncPrincipal : undefined;
            resourceInputs["syncUsers"] = args ? args.syncUsers : undefined;
            resourceInputs["syncUsersPassword"] = args ? args.syncUsersPassword : undefined;
            resourceInputs["userMatchingMode"] = args ? args.userMatchingMode : undefined;
            resourceInputs["userPathTemplate"] = args ? args.userPathTemplate : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["spnegoKeytab", "syncKeytab", "syncPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SourceKerberos.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceKerberos resources.
 */
export interface SourceKerberosState {
    authenticationFlow?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    enrollmentFlow?: pulumi.Input<string>;
    /**
     * Allowed values: - `identifier` - `nameLink` - `nameDeny`
     */
    groupMatchingMode?: pulumi.Input<string>;
    /**
     * Custom krb5.conf to use. Uses the system one by default
     */
    krb5Conf?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend
     */
    passwordLoginUpdateInternalPassword?: pulumi.Input<boolean>;
    /**
     * Allowed values: - `all` - `any`
     */
    policyEngineMode?: pulumi.Input<string>;
    /**
     * Kerberos realm
     */
    realm?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
    /**
     * Credential cache to use for SPNEGO in form type:residual
     */
    spnegoCcache?: pulumi.Input<string>;
    /**
     * SPNEGO keytab base64-encoded or path to keytab in the form FILE:path
     */
    spnegoKeytab?: pulumi.Input<string>;
    /**
     * Force the use of a specific server name for SPNEGO
     */
    spnegoServerName?: pulumi.Input<string>;
    /**
     * Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual
     */
    syncCcache?: pulumi.Input<string>;
    /**
     * Keytab to authenticate to kadmin for sync. Must be base64-encoded or in the form TYPE:residual
     */
    syncKeytab?: pulumi.Input<string>;
    /**
     * Password to authenticate to kadmin for sync
     */
    syncPassword?: pulumi.Input<string>;
    /**
     * Principal to authenticate to kadmin for sync.
     */
    syncPrincipal?: pulumi.Input<string>;
    /**
     * Sync users from Kerberos into authentik
     */
    syncUsers?: pulumi.Input<boolean>;
    /**
     * When a user changes their password, sync it back to Kerberos
     */
    syncUsersPassword?: pulumi.Input<boolean>;
    /**
     * Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
     */
    userMatchingMode?: pulumi.Input<string>;
    userPathTemplate?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SourceKerberos resource.
 */
export interface SourceKerberosArgs {
    authenticationFlow?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    enrollmentFlow?: pulumi.Input<string>;
    /**
     * Allowed values: - `identifier` - `nameLink` - `nameDeny`
     */
    groupMatchingMode?: pulumi.Input<string>;
    /**
     * Custom krb5.conf to use. Uses the system one by default
     */
    krb5Conf?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend
     */
    passwordLoginUpdateInternalPassword?: pulumi.Input<boolean>;
    /**
     * Allowed values: - `all` - `any`
     */
    policyEngineMode?: pulumi.Input<string>;
    /**
     * Kerberos realm
     */
    realm: pulumi.Input<string>;
    slug: pulumi.Input<string>;
    /**
     * Credential cache to use for SPNEGO in form type:residual
     */
    spnegoCcache?: pulumi.Input<string>;
    /**
     * SPNEGO keytab base64-encoded or path to keytab in the form FILE:path
     */
    spnegoKeytab?: pulumi.Input<string>;
    /**
     * Force the use of a specific server name for SPNEGO
     */
    spnegoServerName?: pulumi.Input<string>;
    /**
     * Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual
     */
    syncCcache?: pulumi.Input<string>;
    /**
     * Keytab to authenticate to kadmin for sync. Must be base64-encoded or in the form TYPE:residual
     */
    syncKeytab?: pulumi.Input<string>;
    /**
     * Password to authenticate to kadmin for sync
     */
    syncPassword?: pulumi.Input<string>;
    /**
     * Principal to authenticate to kadmin for sync.
     */
    syncPrincipal?: pulumi.Input<string>;
    /**
     * Sync users from Kerberos into authentik
     */
    syncUsers?: pulumi.Input<boolean>;
    /**
     * When a user changes their password, sync it back to Kerberos
     */
    syncUsersPassword?: pulumi.Input<boolean>;
    /**
     * Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
     */
    userMatchingMode?: pulumi.Input<string>;
    userPathTemplate?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
}

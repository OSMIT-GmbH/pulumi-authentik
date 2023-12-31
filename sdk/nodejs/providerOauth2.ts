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
 * // Create an OAuth2 Provider
 * const nameProviderOauth2 = new authentik.ProviderOauth2("nameProviderOauth2", {clientId: "grafana"});
 * const nameApplication = new authentik.Application("nameApplication", {
 *     slug: "test-app",
 *     protocolProvider: nameProviderOauth2.id,
 * });
 * ```
 */
export class ProviderOauth2 extends pulumi.CustomResource {
    /**
     * Get an existing ProviderOauth2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProviderOauth2State, opts?: pulumi.CustomResourceOptions): ProviderOauth2 {
        return new ProviderOauth2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/providerOauth2:ProviderOauth2';

    /**
     * Returns true if the given object is an instance of ProviderOauth2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderOauth2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderOauth2.__pulumiType;
    }

    /**
     * Defaults to `minutes=1`.
     */
    public readonly accessCodeValidity!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `minutes=10`.
     */
    public readonly accessTokenValidity!: pulumi.Output<string | undefined>;
    public readonly authenticationFlow!: pulumi.Output<string | undefined>;
    public readonly authorizationFlow!: pulumi.Output<string>;
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Generated.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * Defaults to `confidential`.
     */
    public readonly clientType!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `true`.
     */
    public readonly includeClaimsInIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to `perProvider`.
     */
    public readonly issuerMode!: pulumi.Output<string | undefined>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    public readonly jwksSources!: pulumi.Output<string[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly propertyMappings!: pulumi.Output<string[] | undefined>;
    public readonly redirectUris!: pulumi.Output<string[] | undefined>;
    /**
     * Defaults to `days=30`.
     */
    public readonly refreshTokenValidity!: pulumi.Output<string | undefined>;
    public readonly signingKey!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `hashedUserId`.
     */
    public readonly subMode!: pulumi.Output<string | undefined>;

    /**
     * Create a ProviderOauth2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderOauth2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProviderOauth2Args | ProviderOauth2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProviderOauth2State | undefined;
            resourceInputs["accessCodeValidity"] = state ? state.accessCodeValidity : undefined;
            resourceInputs["accessTokenValidity"] = state ? state.accessTokenValidity : undefined;
            resourceInputs["authenticationFlow"] = state ? state.authenticationFlow : undefined;
            resourceInputs["authorizationFlow"] = state ? state.authorizationFlow : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["clientType"] = state ? state.clientType : undefined;
            resourceInputs["includeClaimsInIdToken"] = state ? state.includeClaimsInIdToken : undefined;
            resourceInputs["issuerMode"] = state ? state.issuerMode : undefined;
            resourceInputs["jwksSources"] = state ? state.jwksSources : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["propertyMappings"] = state ? state.propertyMappings : undefined;
            resourceInputs["redirectUris"] = state ? state.redirectUris : undefined;
            resourceInputs["refreshTokenValidity"] = state ? state.refreshTokenValidity : undefined;
            resourceInputs["signingKey"] = state ? state.signingKey : undefined;
            resourceInputs["subMode"] = state ? state.subMode : undefined;
        } else {
            const args = argsOrState as ProviderOauth2Args | undefined;
            if ((!args || args.authorizationFlow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationFlow'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            resourceInputs["accessCodeValidity"] = args ? args.accessCodeValidity : undefined;
            resourceInputs["accessTokenValidity"] = args ? args.accessTokenValidity : undefined;
            resourceInputs["authenticationFlow"] = args ? args.authenticationFlow : undefined;
            resourceInputs["authorizationFlow"] = args ? args.authorizationFlow : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["clientType"] = args ? args.clientType : undefined;
            resourceInputs["includeClaimsInIdToken"] = args ? args.includeClaimsInIdToken : undefined;
            resourceInputs["issuerMode"] = args ? args.issuerMode : undefined;
            resourceInputs["jwksSources"] = args ? args.jwksSources : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["propertyMappings"] = args ? args.propertyMappings : undefined;
            resourceInputs["redirectUris"] = args ? args.redirectUris : undefined;
            resourceInputs["refreshTokenValidity"] = args ? args.refreshTokenValidity : undefined;
            resourceInputs["signingKey"] = args ? args.signingKey : undefined;
            resourceInputs["subMode"] = args ? args.subMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProviderOauth2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProviderOauth2 resources.
 */
export interface ProviderOauth2State {
    /**
     * Defaults to `minutes=1`.
     */
    accessCodeValidity?: pulumi.Input<string>;
    /**
     * Defaults to `minutes=10`.
     */
    accessTokenValidity?: pulumi.Input<string>;
    authenticationFlow?: pulumi.Input<string>;
    authorizationFlow?: pulumi.Input<string>;
    clientId?: pulumi.Input<string>;
    /**
     * Generated.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Defaults to `confidential`.
     */
    clientType?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    includeClaimsInIdToken?: pulumi.Input<boolean>;
    /**
     * Defaults to `perProvider`.
     */
    issuerMode?: pulumi.Input<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    jwksSources?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defaults to `days=30`.
     */
    refreshTokenValidity?: pulumi.Input<string>;
    signingKey?: pulumi.Input<string>;
    /**
     * Defaults to `hashedUserId`.
     */
    subMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderOauth2 resource.
 */
export interface ProviderOauth2Args {
    /**
     * Defaults to `minutes=1`.
     */
    accessCodeValidity?: pulumi.Input<string>;
    /**
     * Defaults to `minutes=10`.
     */
    accessTokenValidity?: pulumi.Input<string>;
    authenticationFlow?: pulumi.Input<string>;
    authorizationFlow: pulumi.Input<string>;
    clientId: pulumi.Input<string>;
    /**
     * Generated.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Defaults to `confidential`.
     */
    clientType?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    includeClaimsInIdToken?: pulumi.Input<boolean>;
    /**
     * Defaults to `perProvider`.
     */
    issuerMode?: pulumi.Input<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    jwksSources?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defaults to `days=30`.
     */
    refreshTokenValidity?: pulumi.Input<string>;
    signingKey?: pulumi.Input<string>;
    /**
     * Defaults to `hashedUserId`.
     */
    subMode?: pulumi.Input<string>;
}

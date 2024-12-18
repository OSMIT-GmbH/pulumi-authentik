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
 * const name = new authentik.ProviderOauth2("name", {
 *     name: "grafana",
 *     clientId: "grafana",
 *     allowedRedirectUris: [{
 *         matching_mode: "strict",
 *         url: "http://localhost",
 *     }],
 * });
 * const nameApplication = new authentik.Application("name", {
 *     name: "test app",
 *     slug: "test-app",
 *     protocolProvider: name.id,
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

    public readonly accessCodeValidity!: pulumi.Output<string | undefined>;
    public readonly accessTokenValidity!: pulumi.Output<string | undefined>;
    public readonly allowedRedirectUris!: pulumi.Output<{[key: string]: string}[] | undefined>;
    public readonly authenticationFlow!: pulumi.Output<string | undefined>;
    public readonly authorizationFlow!: pulumi.Output<string>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * Allowed values: - `confidential` - `public`
     */
    public readonly clientType!: pulumi.Output<string | undefined>;
    public readonly encryptionKey!: pulumi.Output<string | undefined>;
    public readonly includeClaimsInIdToken!: pulumi.Output<boolean | undefined>;
    public readonly invalidationFlow!: pulumi.Output<string>;
    /**
     * Allowed values: - `global` - `perProvider`
     */
    public readonly issuerMode!: pulumi.Output<string | undefined>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    public readonly jwksSources!: pulumi.Output<string[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly propertyMappings!: pulumi.Output<string[] | undefined>;
    public readonly refreshTokenValidity!: pulumi.Output<string | undefined>;
    public readonly signingKey!: pulumi.Output<string | undefined>;
    /**
     * Allowed values: - `hashedUserId` - `userId` - `userUuid` - `userUsername` - `userEmail` - `userUpn`
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
            resourceInputs["allowedRedirectUris"] = state ? state.allowedRedirectUris : undefined;
            resourceInputs["authenticationFlow"] = state ? state.authenticationFlow : undefined;
            resourceInputs["authorizationFlow"] = state ? state.authorizationFlow : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["clientType"] = state ? state.clientType : undefined;
            resourceInputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            resourceInputs["includeClaimsInIdToken"] = state ? state.includeClaimsInIdToken : undefined;
            resourceInputs["invalidationFlow"] = state ? state.invalidationFlow : undefined;
            resourceInputs["issuerMode"] = state ? state.issuerMode : undefined;
            resourceInputs["jwksSources"] = state ? state.jwksSources : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["propertyMappings"] = state ? state.propertyMappings : undefined;
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
            if ((!args || args.invalidationFlow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'invalidationFlow'");
            }
            resourceInputs["accessCodeValidity"] = args ? args.accessCodeValidity : undefined;
            resourceInputs["accessTokenValidity"] = args ? args.accessTokenValidity : undefined;
            resourceInputs["allowedRedirectUris"] = args ? args.allowedRedirectUris : undefined;
            resourceInputs["authenticationFlow"] = args ? args.authenticationFlow : undefined;
            resourceInputs["authorizationFlow"] = args ? args.authorizationFlow : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["clientType"] = args ? args.clientType : undefined;
            resourceInputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            resourceInputs["includeClaimsInIdToken"] = args ? args.includeClaimsInIdToken : undefined;
            resourceInputs["invalidationFlow"] = args ? args.invalidationFlow : undefined;
            resourceInputs["issuerMode"] = args ? args.issuerMode : undefined;
            resourceInputs["jwksSources"] = args ? args.jwksSources : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["propertyMappings"] = args ? args.propertyMappings : undefined;
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
    accessCodeValidity?: pulumi.Input<string>;
    accessTokenValidity?: pulumi.Input<string>;
    allowedRedirectUris?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    authenticationFlow?: pulumi.Input<string>;
    authorizationFlow?: pulumi.Input<string>;
    clientId?: pulumi.Input<string>;
    clientSecret?: pulumi.Input<string>;
    /**
     * Allowed values: - `confidential` - `public`
     */
    clientType?: pulumi.Input<string>;
    encryptionKey?: pulumi.Input<string>;
    includeClaimsInIdToken?: pulumi.Input<boolean>;
    invalidationFlow?: pulumi.Input<string>;
    /**
     * Allowed values: - `global` - `perProvider`
     */
    issuerMode?: pulumi.Input<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    jwksSources?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    refreshTokenValidity?: pulumi.Input<string>;
    signingKey?: pulumi.Input<string>;
    /**
     * Allowed values: - `hashedUserId` - `userId` - `userUuid` - `userUsername` - `userEmail` - `userUpn`
     */
    subMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderOauth2 resource.
 */
export interface ProviderOauth2Args {
    accessCodeValidity?: pulumi.Input<string>;
    accessTokenValidity?: pulumi.Input<string>;
    allowedRedirectUris?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    authenticationFlow?: pulumi.Input<string>;
    authorizationFlow: pulumi.Input<string>;
    clientId: pulumi.Input<string>;
    clientSecret?: pulumi.Input<string>;
    /**
     * Allowed values: - `confidential` - `public`
     */
    clientType?: pulumi.Input<string>;
    encryptionKey?: pulumi.Input<string>;
    includeClaimsInIdToken?: pulumi.Input<boolean>;
    invalidationFlow: pulumi.Input<string>;
    /**
     * Allowed values: - `global` - `perProvider`
     */
    issuerMode?: pulumi.Input<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    jwksSources?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    refreshTokenValidity?: pulumi.Input<string>;
    signingKey?: pulumi.Input<string>;
    /**
     * Allowed values: - `hashedUserId` - `userId` - `userUuid` - `userUsername` - `userEmail` - `userUpn`
     */
    subMode?: pulumi.Input<string>;
}

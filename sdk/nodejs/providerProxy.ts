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
 * // Create a proxy provider
 * const default-authorization-flow = authentik.getFlow({
 *     slug: "default-provider-authorization-implicit-consent",
 * });
 * const name = new authentik.ProviderProxy("name", {
 *     name: "test-app",
 *     internalHost: "http://foo.bar.baz",
 *     externalHost: "http://internal.service",
 *     authorizationFlow: default_authorization_flow.then(default_authorization_flow => default_authorization_flow.id),
 * });
 * const nameApplication = new authentik.Application("name", {
 *     name: "test-app",
 *     slug: "test-app",
 *     protocolProvider: name.id,
 * });
 * ```
 */
export class ProviderProxy extends pulumi.CustomResource {
    /**
     * Get an existing ProviderProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProviderProxyState, opts?: pulumi.CustomResourceOptions): ProviderProxy {
        return new ProviderProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/providerProxy:ProviderProxy';

    /**
     * Returns true if the given object is an instance of ProviderProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderProxy.__pulumiType;
    }

    public readonly accessTokenValidity!: pulumi.Output<string | undefined>;
    public readonly authenticationFlow!: pulumi.Output<string | undefined>;
    public readonly authorizationFlow!: pulumi.Output<string>;
    public readonly basicAuthEnabled!: pulumi.Output<boolean | undefined>;
    public readonly basicAuthPasswordAttribute!: pulumi.Output<string | undefined>;
    public readonly basicAuthUsernameAttribute!: pulumi.Output<string | undefined>;
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    public readonly cookieDomain!: pulumi.Output<string | undefined>;
    public readonly externalHost!: pulumi.Output<string>;
    public readonly interceptHeaderAuth!: pulumi.Output<boolean | undefined>;
    public readonly internalHost!: pulumi.Output<string | undefined>;
    public readonly internalHostSslValidation!: pulumi.Output<boolean | undefined>;
    public readonly invalidationFlow!: pulumi.Output<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    public readonly jwksSources!: pulumi.Output<string[] | undefined>;
    /**
     * Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly propertyMappings!: pulumi.Output<string[] | undefined>;
    public readonly refreshTokenValidity!: pulumi.Output<string | undefined>;
    public readonly skipPathRegex!: pulumi.Output<string | undefined>;

    /**
     * Create a ProviderProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProviderProxyArgs | ProviderProxyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProviderProxyState | undefined;
            resourceInputs["accessTokenValidity"] = state ? state.accessTokenValidity : undefined;
            resourceInputs["authenticationFlow"] = state ? state.authenticationFlow : undefined;
            resourceInputs["authorizationFlow"] = state ? state.authorizationFlow : undefined;
            resourceInputs["basicAuthEnabled"] = state ? state.basicAuthEnabled : undefined;
            resourceInputs["basicAuthPasswordAttribute"] = state ? state.basicAuthPasswordAttribute : undefined;
            resourceInputs["basicAuthUsernameAttribute"] = state ? state.basicAuthUsernameAttribute : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["cookieDomain"] = state ? state.cookieDomain : undefined;
            resourceInputs["externalHost"] = state ? state.externalHost : undefined;
            resourceInputs["interceptHeaderAuth"] = state ? state.interceptHeaderAuth : undefined;
            resourceInputs["internalHost"] = state ? state.internalHost : undefined;
            resourceInputs["internalHostSslValidation"] = state ? state.internalHostSslValidation : undefined;
            resourceInputs["invalidationFlow"] = state ? state.invalidationFlow : undefined;
            resourceInputs["jwksSources"] = state ? state.jwksSources : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["propertyMappings"] = state ? state.propertyMappings : undefined;
            resourceInputs["refreshTokenValidity"] = state ? state.refreshTokenValidity : undefined;
            resourceInputs["skipPathRegex"] = state ? state.skipPathRegex : undefined;
        } else {
            const args = argsOrState as ProviderProxyArgs | undefined;
            if ((!args || args.authorizationFlow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationFlow'");
            }
            if ((!args || args.externalHost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalHost'");
            }
            if ((!args || args.invalidationFlow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'invalidationFlow'");
            }
            resourceInputs["accessTokenValidity"] = args ? args.accessTokenValidity : undefined;
            resourceInputs["authenticationFlow"] = args ? args.authenticationFlow : undefined;
            resourceInputs["authorizationFlow"] = args ? args.authorizationFlow : undefined;
            resourceInputs["basicAuthEnabled"] = args ? args.basicAuthEnabled : undefined;
            resourceInputs["basicAuthPasswordAttribute"] = args ? args.basicAuthPasswordAttribute : undefined;
            resourceInputs["basicAuthUsernameAttribute"] = args ? args.basicAuthUsernameAttribute : undefined;
            resourceInputs["cookieDomain"] = args ? args.cookieDomain : undefined;
            resourceInputs["externalHost"] = args ? args.externalHost : undefined;
            resourceInputs["interceptHeaderAuth"] = args ? args.interceptHeaderAuth : undefined;
            resourceInputs["internalHost"] = args ? args.internalHost : undefined;
            resourceInputs["internalHostSslValidation"] = args ? args.internalHostSslValidation : undefined;
            resourceInputs["invalidationFlow"] = args ? args.invalidationFlow : undefined;
            resourceInputs["jwksSources"] = args ? args.jwksSources : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["propertyMappings"] = args ? args.propertyMappings : undefined;
            resourceInputs["refreshTokenValidity"] = args ? args.refreshTokenValidity : undefined;
            resourceInputs["skipPathRegex"] = args ? args.skipPathRegex : undefined;
            resourceInputs["clientId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProviderProxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProviderProxy resources.
 */
export interface ProviderProxyState {
    accessTokenValidity?: pulumi.Input<string>;
    authenticationFlow?: pulumi.Input<string>;
    authorizationFlow?: pulumi.Input<string>;
    basicAuthEnabled?: pulumi.Input<boolean>;
    basicAuthPasswordAttribute?: pulumi.Input<string>;
    basicAuthUsernameAttribute?: pulumi.Input<string>;
    clientId?: pulumi.Input<string>;
    cookieDomain?: pulumi.Input<string>;
    externalHost?: pulumi.Input<string>;
    interceptHeaderAuth?: pulumi.Input<boolean>;
    internalHost?: pulumi.Input<string>;
    internalHostSslValidation?: pulumi.Input<boolean>;
    invalidationFlow?: pulumi.Input<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    jwksSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
     */
    mode?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    refreshTokenValidity?: pulumi.Input<string>;
    skipPathRegex?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProviderProxy resource.
 */
export interface ProviderProxyArgs {
    accessTokenValidity?: pulumi.Input<string>;
    authenticationFlow?: pulumi.Input<string>;
    authorizationFlow: pulumi.Input<string>;
    basicAuthEnabled?: pulumi.Input<boolean>;
    basicAuthPasswordAttribute?: pulumi.Input<string>;
    basicAuthUsernameAttribute?: pulumi.Input<string>;
    cookieDomain?: pulumi.Input<string>;
    externalHost: pulumi.Input<string>;
    interceptHeaderAuth?: pulumi.Input<boolean>;
    internalHost?: pulumi.Input<string>;
    internalHostSslValidation?: pulumi.Input<boolean>;
    invalidationFlow: pulumi.Input<string>;
    /**
     * JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
     */
    jwksSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
     */
    mode?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    propertyMappings?: pulumi.Input<pulumi.Input<string>[]>;
    refreshTokenValidity?: pulumi.Input<string>;
    skipPathRegex?: pulumi.Input<string>;
}
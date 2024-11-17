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
 * // Create a local docker connection
 * const local = new authentik.ServiceConnectionDocker("local", {
 *     name: "local",
 *     local: true,
 * });
 * // Create a remote docker connection
 * const tls_auth = new authentik.CertificateKeyPair("tls-auth", {
 *     name: "docker-tls-auth",
 *     certificateData: "...",
 *     keyData: "...",
 * });
 * const tls_verification = new authentik.CertificateKeyPair("tls-verification", {
 *     name: "docker-tls-verification",
 *     certificateData: "...",
 * });
 * const remote_host = new authentik.ServiceConnectionDocker("remote-host", {
 *     name: "remote-host",
 *     url: "http://1.2.3.4:2368",
 *     tlsVerification: tls_auth.id,
 *     tlsAuthentication: tls_verification.id,
 * });
 * ```
 */
export class ServiceConnectionDocker extends pulumi.CustomResource {
    /**
     * Get an existing ServiceConnectionDocker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceConnectionDockerState, opts?: pulumi.CustomResourceOptions): ServiceConnectionDocker {
        return new ServiceConnectionDocker(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/serviceConnectionDocker:ServiceConnectionDocker';

    /**
     * Returns true if the given object is an instance of ServiceConnectionDocker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceConnectionDocker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceConnectionDocker.__pulumiType;
    }

    /**
     * Defaults to `false`.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly tlsAuthentication!: pulumi.Output<string | undefined>;
    public readonly tlsVerification!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `http+unix:///var/run/docker.sock`.
     */
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceConnectionDocker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceConnectionDockerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceConnectionDockerArgs | ServiceConnectionDockerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceConnectionDockerState | undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tlsAuthentication"] = state ? state.tlsAuthentication : undefined;
            resourceInputs["tlsVerification"] = state ? state.tlsVerification : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceConnectionDockerArgs | undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tlsAuthentication"] = args ? args.tlsAuthentication : undefined;
            resourceInputs["tlsVerification"] = args ? args.tlsVerification : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceConnectionDocker.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceConnectionDocker resources.
 */
export interface ServiceConnectionDockerState {
    /**
     * Defaults to `false`.
     */
    local?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    tlsAuthentication?: pulumi.Input<string>;
    tlsVerification?: pulumi.Input<string>;
    /**
     * Defaults to `http+unix:///var/run/docker.sock`.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceConnectionDocker resource.
 */
export interface ServiceConnectionDockerArgs {
    /**
     * Defaults to `false`.
     */
    local?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    tlsAuthentication?: pulumi.Input<string>;
    tlsVerification?: pulumi.Input<string>;
    /**
     * Defaults to `http+unix:///var/run/docker.sock`.
     */
    url?: pulumi.Input<string>;
}

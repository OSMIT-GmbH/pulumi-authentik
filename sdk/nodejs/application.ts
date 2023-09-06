// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    public readonly backchannelProviders!: pulumi.Output<number[] | undefined>;
    public readonly group!: pulumi.Output<string | undefined>;
    public readonly metaDescription!: pulumi.Output<string | undefined>;
    public readonly metaIcon!: pulumi.Output<string | undefined>;
    public readonly metaLaunchUrl!: pulumi.Output<string | undefined>;
    public readonly metaPublisher!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly openInNewTab!: pulumi.Output<boolean | undefined>;
    public readonly policyEngineMode!: pulumi.Output<string | undefined>;
    public readonly protocolProvider!: pulumi.Output<number | undefined>;
    public readonly slug!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["backchannelProviders"] = state ? state.backchannelProviders : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["metaDescription"] = state ? state.metaDescription : undefined;
            resourceInputs["metaIcon"] = state ? state.metaIcon : undefined;
            resourceInputs["metaLaunchUrl"] = state ? state.metaLaunchUrl : undefined;
            resourceInputs["metaPublisher"] = state ? state.metaPublisher : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["openInNewTab"] = state ? state.openInNewTab : undefined;
            resourceInputs["policyEngineMode"] = state ? state.policyEngineMode : undefined;
            resourceInputs["protocolProvider"] = state ? state.protocolProvider : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            resourceInputs["backchannelProviders"] = args ? args.backchannelProviders : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["metaDescription"] = args ? args.metaDescription : undefined;
            resourceInputs["metaIcon"] = args ? args.metaIcon : undefined;
            resourceInputs["metaLaunchUrl"] = args ? args.metaLaunchUrl : undefined;
            resourceInputs["metaPublisher"] = args ? args.metaPublisher : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openInNewTab"] = args ? args.openInNewTab : undefined;
            resourceInputs["policyEngineMode"] = args ? args.policyEngineMode : undefined;
            resourceInputs["protocolProvider"] = args ? args.protocolProvider : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    backchannelProviders?: pulumi.Input<pulumi.Input<number>[]>;
    group?: pulumi.Input<string>;
    metaDescription?: pulumi.Input<string>;
    metaIcon?: pulumi.Input<string>;
    metaLaunchUrl?: pulumi.Input<string>;
    metaPublisher?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    openInNewTab?: pulumi.Input<boolean>;
    policyEngineMode?: pulumi.Input<string>;
    protocolProvider?: pulumi.Input<number>;
    slug?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    backchannelProviders?: pulumi.Input<pulumi.Input<number>[]>;
    group?: pulumi.Input<string>;
    metaDescription?: pulumi.Input<string>;
    metaIcon?: pulumi.Input<string>;
    metaLaunchUrl?: pulumi.Input<string>;
    metaPublisher?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    openInNewTab?: pulumi.Input<boolean>;
    policyEngineMode?: pulumi.Input<string>;
    protocolProvider?: pulumi.Input<number>;
    slug: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
}

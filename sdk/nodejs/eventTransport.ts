// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class EventTransport extends pulumi.CustomResource {
    /**
     * Get an existing EventTransport resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTransportState, opts?: pulumi.CustomResourceOptions): EventTransport {
        return new EventTransport(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/eventTransport:EventTransport';

    /**
     * Returns true if the given object is an instance of EventTransport.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTransport {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTransport.__pulumiType;
    }

    /**
     * Allowed values: - `local` - `webhook` - `webhookSlack` - `email`
     */
    public readonly mode!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly sendOnce!: pulumi.Output<boolean | undefined>;
    public readonly webhookMapping!: pulumi.Output<string | undefined>;
    public readonly webhookUrl!: pulumi.Output<string | undefined>;

    /**
     * Create a EventTransport resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTransportArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTransportArgs | EventTransportState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventTransportState | undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sendOnce"] = state ? state.sendOnce : undefined;
            resourceInputs["webhookMapping"] = state ? state.webhookMapping : undefined;
            resourceInputs["webhookUrl"] = state ? state.webhookUrl : undefined;
        } else {
            const args = argsOrState as EventTransportArgs | undefined;
            if ((!args || args.mode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mode'");
            }
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sendOnce"] = args ? args.sendOnce : undefined;
            resourceInputs["webhookMapping"] = args ? args.webhookMapping : undefined;
            resourceInputs["webhookUrl"] = args ? args.webhookUrl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventTransport.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTransport resources.
 */
export interface EventTransportState {
    /**
     * Allowed values: - `local` - `webhook` - `webhookSlack` - `email`
     */
    mode?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    sendOnce?: pulumi.Input<boolean>;
    webhookMapping?: pulumi.Input<string>;
    webhookUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTransport resource.
 */
export interface EventTransportArgs {
    /**
     * Allowed values: - `local` - `webhook` - `webhookSlack` - `email`
     */
    mode: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    sendOnce?: pulumi.Input<boolean>;
    webhookMapping?: pulumi.Input<string>;
    webhookUrl?: pulumi.Input<string>;
}

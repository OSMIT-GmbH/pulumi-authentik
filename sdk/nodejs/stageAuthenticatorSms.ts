// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class StageAuthenticatorSms extends pulumi.CustomResource {
    /**
     * Get an existing StageAuthenticatorSms resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageAuthenticatorSmsState, opts?: pulumi.CustomResourceOptions): StageAuthenticatorSms {
        return new StageAuthenticatorSms(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/stageAuthenticatorSms:StageAuthenticatorSms';

    /**
     * Returns true if the given object is an instance of StageAuthenticatorSms.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StageAuthenticatorSms {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StageAuthenticatorSms.__pulumiType;
    }

    public readonly accountSid!: pulumi.Output<string>;
    public readonly auth!: pulumi.Output<string>;
    public readonly authPassword!: pulumi.Output<string | undefined>;
    /**
     * Allowed values: - `basic` - `bearer`
     */
    public readonly authType!: pulumi.Output<string | undefined>;
    public readonly configureFlow!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    public readonly fromNumber!: pulumi.Output<string>;
    public readonly mapping!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Allowed values: - `twilio` - `generic`
     */
    public readonly smsProvider!: pulumi.Output<string | undefined>;
    public readonly verifyOnly!: pulumi.Output<boolean | undefined>;

    /**
     * Create a StageAuthenticatorSms resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StageAuthenticatorSmsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageAuthenticatorSmsArgs | StageAuthenticatorSmsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageAuthenticatorSmsState | undefined;
            resourceInputs["accountSid"] = state ? state.accountSid : undefined;
            resourceInputs["auth"] = state ? state.auth : undefined;
            resourceInputs["authPassword"] = state ? state.authPassword : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["configureFlow"] = state ? state.configureFlow : undefined;
            resourceInputs["friendlyName"] = state ? state.friendlyName : undefined;
            resourceInputs["fromNumber"] = state ? state.fromNumber : undefined;
            resourceInputs["mapping"] = state ? state.mapping : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["smsProvider"] = state ? state.smsProvider : undefined;
            resourceInputs["verifyOnly"] = state ? state.verifyOnly : undefined;
        } else {
            const args = argsOrState as StageAuthenticatorSmsArgs | undefined;
            if ((!args || args.accountSid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountSid'");
            }
            if ((!args || args.auth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'auth'");
            }
            if ((!args || args.fromNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fromNumber'");
            }
            resourceInputs["accountSid"] = args?.accountSid ? pulumi.secret(args.accountSid) : undefined;
            resourceInputs["auth"] = args?.auth ? pulumi.secret(args.auth) : undefined;
            resourceInputs["authPassword"] = args?.authPassword ? pulumi.secret(args.authPassword) : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["configureFlow"] = args ? args.configureFlow : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["fromNumber"] = args ? args.fromNumber : undefined;
            resourceInputs["mapping"] = args ? args.mapping : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["smsProvider"] = args ? args.smsProvider : undefined;
            resourceInputs["verifyOnly"] = args ? args.verifyOnly : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountSid", "auth", "authPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(StageAuthenticatorSms.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StageAuthenticatorSms resources.
 */
export interface StageAuthenticatorSmsState {
    accountSid?: pulumi.Input<string>;
    auth?: pulumi.Input<string>;
    authPassword?: pulumi.Input<string>;
    /**
     * Allowed values: - `basic` - `bearer`
     */
    authType?: pulumi.Input<string>;
    configureFlow?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    fromNumber?: pulumi.Input<string>;
    mapping?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Allowed values: - `twilio` - `generic`
     */
    smsProvider?: pulumi.Input<string>;
    verifyOnly?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a StageAuthenticatorSms resource.
 */
export interface StageAuthenticatorSmsArgs {
    accountSid: pulumi.Input<string>;
    auth: pulumi.Input<string>;
    authPassword?: pulumi.Input<string>;
    /**
     * Allowed values: - `basic` - `bearer`
     */
    authType?: pulumi.Input<string>;
    configureFlow?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    fromNumber: pulumi.Input<string>;
    mapping?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Allowed values: - `twilio` - `generic`
     */
    smsProvider?: pulumi.Input<string>;
    verifyOnly?: pulumi.Input<boolean>;
}

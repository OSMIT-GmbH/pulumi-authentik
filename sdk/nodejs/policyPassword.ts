// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class PolicyPassword extends pulumi.CustomResource {
    /**
     * Get an existing PolicyPassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyPasswordState, opts?: pulumi.CustomResourceOptions): PolicyPassword {
        return new PolicyPassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/policyPassword:PolicyPassword';

    /**
     * Returns true if the given object is an instance of PolicyPassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyPassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyPassword.__pulumiType;
    }

    public readonly amountDigits!: pulumi.Output<number | undefined>;
    public readonly amountLowercase!: pulumi.Output<number | undefined>;
    public readonly amountSymbols!: pulumi.Output<number | undefined>;
    public readonly amountUppercase!: pulumi.Output<number | undefined>;
    public readonly checkHaveIBeenPwned!: pulumi.Output<boolean | undefined>;
    public readonly checkStaticRules!: pulumi.Output<boolean | undefined>;
    public readonly checkZxcvbn!: pulumi.Output<boolean | undefined>;
    public readonly errorMessage!: pulumi.Output<string>;
    public readonly executionLogging!: pulumi.Output<boolean | undefined>;
    public readonly hibpAllowedCount!: pulumi.Output<number | undefined>;
    public readonly lengthMin!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly passwordField!: pulumi.Output<string | undefined>;
    public readonly symbolCharset!: pulumi.Output<string | undefined>;
    public readonly zxcvbnScoreThreshold!: pulumi.Output<number | undefined>;

    /**
     * Create a PolicyPassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyPasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyPasswordArgs | PolicyPasswordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyPasswordState | undefined;
            resourceInputs["amountDigits"] = state ? state.amountDigits : undefined;
            resourceInputs["amountLowercase"] = state ? state.amountLowercase : undefined;
            resourceInputs["amountSymbols"] = state ? state.amountSymbols : undefined;
            resourceInputs["amountUppercase"] = state ? state.amountUppercase : undefined;
            resourceInputs["checkHaveIBeenPwned"] = state ? state.checkHaveIBeenPwned : undefined;
            resourceInputs["checkStaticRules"] = state ? state.checkStaticRules : undefined;
            resourceInputs["checkZxcvbn"] = state ? state.checkZxcvbn : undefined;
            resourceInputs["errorMessage"] = state ? state.errorMessage : undefined;
            resourceInputs["executionLogging"] = state ? state.executionLogging : undefined;
            resourceInputs["hibpAllowedCount"] = state ? state.hibpAllowedCount : undefined;
            resourceInputs["lengthMin"] = state ? state.lengthMin : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["passwordField"] = state ? state.passwordField : undefined;
            resourceInputs["symbolCharset"] = state ? state.symbolCharset : undefined;
            resourceInputs["zxcvbnScoreThreshold"] = state ? state.zxcvbnScoreThreshold : undefined;
        } else {
            const args = argsOrState as PolicyPasswordArgs | undefined;
            if ((!args || args.errorMessage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'errorMessage'");
            }
            resourceInputs["amountDigits"] = args ? args.amountDigits : undefined;
            resourceInputs["amountLowercase"] = args ? args.amountLowercase : undefined;
            resourceInputs["amountSymbols"] = args ? args.amountSymbols : undefined;
            resourceInputs["amountUppercase"] = args ? args.amountUppercase : undefined;
            resourceInputs["checkHaveIBeenPwned"] = args ? args.checkHaveIBeenPwned : undefined;
            resourceInputs["checkStaticRules"] = args ? args.checkStaticRules : undefined;
            resourceInputs["checkZxcvbn"] = args ? args.checkZxcvbn : undefined;
            resourceInputs["errorMessage"] = args ? args.errorMessage : undefined;
            resourceInputs["executionLogging"] = args ? args.executionLogging : undefined;
            resourceInputs["hibpAllowedCount"] = args ? args.hibpAllowedCount : undefined;
            resourceInputs["lengthMin"] = args ? args.lengthMin : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passwordField"] = args ? args.passwordField : undefined;
            resourceInputs["symbolCharset"] = args ? args.symbolCharset : undefined;
            resourceInputs["zxcvbnScoreThreshold"] = args ? args.zxcvbnScoreThreshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyPassword.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyPassword resources.
 */
export interface PolicyPasswordState {
    amountDigits?: pulumi.Input<number>;
    amountLowercase?: pulumi.Input<number>;
    amountSymbols?: pulumi.Input<number>;
    amountUppercase?: pulumi.Input<number>;
    checkHaveIBeenPwned?: pulumi.Input<boolean>;
    checkStaticRules?: pulumi.Input<boolean>;
    checkZxcvbn?: pulumi.Input<boolean>;
    errorMessage?: pulumi.Input<string>;
    executionLogging?: pulumi.Input<boolean>;
    hibpAllowedCount?: pulumi.Input<number>;
    lengthMin?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    passwordField?: pulumi.Input<string>;
    symbolCharset?: pulumi.Input<string>;
    zxcvbnScoreThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PolicyPassword resource.
 */
export interface PolicyPasswordArgs {
    amountDigits?: pulumi.Input<number>;
    amountLowercase?: pulumi.Input<number>;
    amountSymbols?: pulumi.Input<number>;
    amountUppercase?: pulumi.Input<number>;
    checkHaveIBeenPwned?: pulumi.Input<boolean>;
    checkStaticRules?: pulumi.Input<boolean>;
    checkZxcvbn?: pulumi.Input<boolean>;
    errorMessage: pulumi.Input<string>;
    executionLogging?: pulumi.Input<boolean>;
    hibpAllowedCount?: pulumi.Input<number>;
    lengthMin?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    passwordField?: pulumi.Input<string>;
    symbolCharset?: pulumi.Input<string>;
    zxcvbnScoreThreshold?: pulumi.Input<number>;
}

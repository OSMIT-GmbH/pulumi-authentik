// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FlowStageBinding extends pulumi.CustomResource {
    /**
     * Get an existing FlowStageBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowStageBindingState, opts?: pulumi.CustomResourceOptions): FlowStageBinding {
        return new FlowStageBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/flowStageBinding:FlowStageBinding';

    /**
     * Returns true if the given object is an instance of FlowStageBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowStageBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowStageBinding.__pulumiType;
    }

    public readonly evaluateOnPlan!: pulumi.Output<boolean | undefined>;
    public readonly invalidResponseAction!: pulumi.Output<string | undefined>;
    public readonly order!: pulumi.Output<number>;
    public readonly policyEngineMode!: pulumi.Output<string | undefined>;
    public readonly reEvaluatePolicies!: pulumi.Output<boolean | undefined>;
    public readonly stage!: pulumi.Output<string>;
    public readonly target!: pulumi.Output<string>;

    /**
     * Create a FlowStageBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowStageBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlowStageBindingArgs | FlowStageBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlowStageBindingState | undefined;
            resourceInputs["evaluateOnPlan"] = state ? state.evaluateOnPlan : undefined;
            resourceInputs["invalidResponseAction"] = state ? state.invalidResponseAction : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["policyEngineMode"] = state ? state.policyEngineMode : undefined;
            resourceInputs["reEvaluatePolicies"] = state ? state.reEvaluatePolicies : undefined;
            resourceInputs["stage"] = state ? state.stage : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as FlowStageBindingArgs | undefined;
            if ((!args || args.order === undefined) && !opts.urn) {
                throw new Error("Missing required property 'order'");
            }
            if ((!args || args.stage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stage'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["evaluateOnPlan"] = args ? args.evaluateOnPlan : undefined;
            resourceInputs["invalidResponseAction"] = args ? args.invalidResponseAction : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["policyEngineMode"] = args ? args.policyEngineMode : undefined;
            resourceInputs["reEvaluatePolicies"] = args ? args.reEvaluatePolicies : undefined;
            resourceInputs["stage"] = args ? args.stage : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FlowStageBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlowStageBinding resources.
 */
export interface FlowStageBindingState {
    evaluateOnPlan?: pulumi.Input<boolean>;
    invalidResponseAction?: pulumi.Input<string>;
    order?: pulumi.Input<number>;
    policyEngineMode?: pulumi.Input<string>;
    reEvaluatePolicies?: pulumi.Input<boolean>;
    stage?: pulumi.Input<string>;
    target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlowStageBinding resource.
 */
export interface FlowStageBindingArgs {
    evaluateOnPlan?: pulumi.Input<boolean>;
    invalidResponseAction?: pulumi.Input<string>;
    order: pulumi.Input<number>;
    policyEngineMode?: pulumi.Input<string>;
    reEvaluatePolicies?: pulumi.Input<boolean>;
    stage: pulumi.Input<string>;
    target: pulumi.Input<string>;
}

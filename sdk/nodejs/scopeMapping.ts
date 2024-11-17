// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Scope Provider Property mappings
 *
 * > This resource is deprecated. Migrate to `authentik.PropertyMappingProviderScope`.
 */
export class ScopeMapping extends pulumi.CustomResource {
    /**
     * Get an existing ScopeMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScopeMappingState, opts?: pulumi.CustomResourceOptions): ScopeMapping {
        return new ScopeMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/scopeMapping:ScopeMapping';

    /**
     * Returns true if the given object is an instance of ScopeMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScopeMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScopeMapping.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly expression!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly scopeName!: pulumi.Output<string>;

    /**
     * Create a ScopeMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScopeMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScopeMappingArgs | ScopeMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScopeMappingState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expression"] = state ? state.expression : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeName"] = state ? state.scopeName : undefined;
        } else {
            const args = argsOrState as ScopeMappingArgs | undefined;
            if ((!args || args.expression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expression'");
            }
            if ((!args || args.scopeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expression"] = args ? args.expression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeName"] = args ? args.scopeName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScopeMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScopeMapping resources.
 */
export interface ScopeMappingState {
    description?: pulumi.Input<string>;
    expression?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    scopeName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScopeMapping resource.
 */
export interface ScopeMappingArgs {
    description?: pulumi.Input<string>;
    expression: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    scopeName: pulumi.Input<string>;
}

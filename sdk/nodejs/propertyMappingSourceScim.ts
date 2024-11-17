// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage SCIM Source Property mappings
 */
export class PropertyMappingSourceScim extends pulumi.CustomResource {
    /**
     * Get an existing PropertyMappingSourceScim resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyMappingSourceScimState, opts?: pulumi.CustomResourceOptions): PropertyMappingSourceScim {
        return new PropertyMappingSourceScim(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/propertyMappingSourceScim:PropertyMappingSourceScim';

    /**
     * Returns true if the given object is an instance of PropertyMappingSourceScim.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertyMappingSourceScim {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertyMappingSourceScim.__pulumiType;
    }

    public readonly expression!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a PropertyMappingSourceScim resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PropertyMappingSourceScimArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertyMappingSourceScimArgs | PropertyMappingSourceScimState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyMappingSourceScimState | undefined;
            resourceInputs["expression"] = state ? state.expression : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PropertyMappingSourceScimArgs | undefined;
            if ((!args || args.expression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expression'");
            }
            resourceInputs["expression"] = args ? args.expression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PropertyMappingSourceScim.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyMappingSourceScim resources.
 */
export interface PropertyMappingSourceScimState {
    expression?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PropertyMappingSourceScim resource.
 */
export interface PropertyMappingSourceScimArgs {
    expression: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
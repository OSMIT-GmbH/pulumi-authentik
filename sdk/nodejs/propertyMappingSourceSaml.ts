// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage SAML Source Property mappings
 */
export class PropertyMappingSourceSaml extends pulumi.CustomResource {
    /**
     * Get an existing PropertyMappingSourceSaml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyMappingSourceSamlState, opts?: pulumi.CustomResourceOptions): PropertyMappingSourceSaml {
        return new PropertyMappingSourceSaml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/propertyMappingSourceSaml:PropertyMappingSourceSaml';

    /**
     * Returns true if the given object is an instance of PropertyMappingSourceSaml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertyMappingSourceSaml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertyMappingSourceSaml.__pulumiType;
    }

    public readonly expression!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a PropertyMappingSourceSaml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PropertyMappingSourceSamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertyMappingSourceSamlArgs | PropertyMappingSourceSamlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyMappingSourceSamlState | undefined;
            resourceInputs["expression"] = state ? state.expression : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PropertyMappingSourceSamlArgs | undefined;
            if ((!args || args.expression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expression'");
            }
            resourceInputs["expression"] = args ? args.expression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PropertyMappingSourceSaml.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyMappingSourceSaml resources.
 */
export interface PropertyMappingSourceSamlState {
    expression?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PropertyMappingSourceSaml resource.
 */
export interface PropertyMappingSourceSamlArgs {
    expression: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
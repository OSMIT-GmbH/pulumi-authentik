// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Radius Provider Property mappings
 *
 * > This resource is deprecated. Migrate to `authentik.PropertyMappingProviderRadius`.
 */
export class PropertyMappingRadius extends pulumi.CustomResource {
    /**
     * Get an existing PropertyMappingRadius resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyMappingRadiusState, opts?: pulumi.CustomResourceOptions): PropertyMappingRadius {
        return new PropertyMappingRadius(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/propertyMappingRadius:PropertyMappingRadius';

    /**
     * Returns true if the given object is an instance of PropertyMappingRadius.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertyMappingRadius {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertyMappingRadius.__pulumiType;
    }

    public readonly expression!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a PropertyMappingRadius resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PropertyMappingRadiusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertyMappingRadiusArgs | PropertyMappingRadiusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyMappingRadiusState | undefined;
            resourceInputs["expression"] = state ? state.expression : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PropertyMappingRadiusArgs | undefined;
            if ((!args || args.expression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expression'");
            }
            resourceInputs["expression"] = args ? args.expression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PropertyMappingRadius.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyMappingRadius resources.
 */
export interface PropertyMappingRadiusState {
    expression?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PropertyMappingRadius resource.
 */
export interface PropertyMappingRadiusArgs {
    expression: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Kerberos Source Property mappings
 */
export class PropertyMappingSourceKerberos extends pulumi.CustomResource {
    /**
     * Get an existing PropertyMappingSourceKerberos resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyMappingSourceKerberosState, opts?: pulumi.CustomResourceOptions): PropertyMappingSourceKerberos {
        return new PropertyMappingSourceKerberos(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'authentik:index/propertyMappingSourceKerberos:PropertyMappingSourceKerberos';

    /**
     * Returns true if the given object is an instance of PropertyMappingSourceKerberos.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertyMappingSourceKerberos {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertyMappingSourceKerberos.__pulumiType;
    }

    public readonly expression!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a PropertyMappingSourceKerberos resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PropertyMappingSourceKerberosArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertyMappingSourceKerberosArgs | PropertyMappingSourceKerberosState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertyMappingSourceKerberosState | undefined;
            resourceInputs["expression"] = state ? state.expression : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PropertyMappingSourceKerberosArgs | undefined;
            if ((!args || args.expression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expression'");
            }
            resourceInputs["expression"] = args ? args.expression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PropertyMappingSourceKerberos.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertyMappingSourceKerberos resources.
 */
export interface PropertyMappingSourceKerberosState {
    expression?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PropertyMappingSourceKerberos resource.
 */
export interface PropertyMappingSourceKerberosArgs {
    expression: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

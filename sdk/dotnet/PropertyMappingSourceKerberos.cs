// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace OSMIT_GmbH.Authentik
{
    /// <summary>
    /// Manage Kerberos Source Property mappings
    /// </summary>
    [AuthentikResourceType("authentik:index/propertyMappingSourceKerberos:PropertyMappingSourceKerberos")]
    public partial class PropertyMappingSourceKerberos : global::Pulumi.CustomResource
    {
        [Output("expression")]
        public Output<string> Expression { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyMappingSourceKerberos resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyMappingSourceKerberos(string name, PropertyMappingSourceKerberosArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingSourceKerberos:PropertyMappingSourceKerberos", name, args ?? new PropertyMappingSourceKerberosArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyMappingSourceKerberos(string name, Input<string> id, PropertyMappingSourceKerberosState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingSourceKerberos:PropertyMappingSourceKerberos", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OSMIT-GmbH/pulumi-authentik/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PropertyMappingSourceKerberos resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertyMappingSourceKerberos Get(string name, Input<string> id, PropertyMappingSourceKerberosState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertyMappingSourceKerberos(name, id, state, options);
        }
    }

    public sealed class PropertyMappingSourceKerberosArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingSourceKerberosArgs()
        {
        }
        public static new PropertyMappingSourceKerberosArgs Empty => new PropertyMappingSourceKerberosArgs();
    }

    public sealed class PropertyMappingSourceKerberosState : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingSourceKerberosState()
        {
        }
        public static new PropertyMappingSourceKerberosState Empty => new PropertyMappingSourceKerberosState();
    }
}
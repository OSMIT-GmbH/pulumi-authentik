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
    /// Manage Microsoft Entra Provider Property mappings
    /// 
    /// &gt; This resource is deprecated. Migrate to `authentik.PropertyMappingProviderMicrosoftEntra`.
    /// </summary>
    [AuthentikResourceType("authentik:index/propertyMappingMicrosoftEntra:PropertyMappingMicrosoftEntra")]
    public partial class PropertyMappingMicrosoftEntra : global::Pulumi.CustomResource
    {
        [Output("expression")]
        public Output<string> Expression { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyMappingMicrosoftEntra resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyMappingMicrosoftEntra(string name, PropertyMappingMicrosoftEntraArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingMicrosoftEntra:PropertyMappingMicrosoftEntra", name, args ?? new PropertyMappingMicrosoftEntraArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyMappingMicrosoftEntra(string name, Input<string> id, PropertyMappingMicrosoftEntraState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingMicrosoftEntra:PropertyMappingMicrosoftEntra", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PropertyMappingMicrosoftEntra resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertyMappingMicrosoftEntra Get(string name, Input<string> id, PropertyMappingMicrosoftEntraState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertyMappingMicrosoftEntra(name, id, state, options);
        }
    }

    public sealed class PropertyMappingMicrosoftEntraArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingMicrosoftEntraArgs()
        {
        }
        public static new PropertyMappingMicrosoftEntraArgs Empty => new PropertyMappingMicrosoftEntraArgs();
    }

    public sealed class PropertyMappingMicrosoftEntraState : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingMicrosoftEntraState()
        {
        }
        public static new PropertyMappingMicrosoftEntraState Empty => new PropertyMappingMicrosoftEntraState();
    }
}

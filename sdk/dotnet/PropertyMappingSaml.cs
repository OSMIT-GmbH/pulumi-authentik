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
    /// Manage SAML Provider Property mappings
    /// 
    /// &gt; This resource is deprecated. Migrate to `authentik.PropertyMappingProviderSaml`.
    /// </summary>
    [AuthentikResourceType("authentik:index/propertyMappingSaml:PropertyMappingSaml")]
    public partial class PropertyMappingSaml : global::Pulumi.CustomResource
    {
        [Output("expression")]
        public Output<string> Expression { get; private set; } = null!;

        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("samlName")]
        public Output<string> SamlName { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyMappingSaml resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyMappingSaml(string name, PropertyMappingSamlArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingSaml:PropertyMappingSaml", name, args ?? new PropertyMappingSamlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyMappingSaml(string name, Input<string> id, PropertyMappingSamlState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingSaml:PropertyMappingSaml", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PropertyMappingSaml resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertyMappingSaml Get(string name, Input<string> id, PropertyMappingSamlState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertyMappingSaml(name, id, state, options);
        }
    }

    public sealed class PropertyMappingSamlArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("samlName", required: true)]
        public Input<string> SamlName { get; set; } = null!;

        public PropertyMappingSamlArgs()
        {
        }
        public static new PropertyMappingSamlArgs Empty => new PropertyMappingSamlArgs();
    }

    public sealed class PropertyMappingSamlState : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("samlName")]
        public Input<string>? SamlName { get; set; }

        public PropertyMappingSamlState()
        {
        }
        public static new PropertyMappingSamlState Empty => new PropertyMappingSamlState();
    }
}

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
    /// Manage SCIM Source Property mappings
    /// </summary>
    [AuthentikResourceType("authentik:index/propertyMappingSourceScim:PropertyMappingSourceScim")]
    public partial class PropertyMappingSourceScim : global::Pulumi.CustomResource
    {
        [Output("expression")]
        public Output<string> Expression { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyMappingSourceScim resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyMappingSourceScim(string name, PropertyMappingSourceScimArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingSourceScim:PropertyMappingSourceScim", name, args ?? new PropertyMappingSourceScimArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyMappingSourceScim(string name, Input<string> id, PropertyMappingSourceScimState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingSourceScim:PropertyMappingSourceScim", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PropertyMappingSourceScim resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertyMappingSourceScim Get(string name, Input<string> id, PropertyMappingSourceScimState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertyMappingSourceScim(name, id, state, options);
        }
    }

    public sealed class PropertyMappingSourceScimArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingSourceScimArgs()
        {
        }
        public static new PropertyMappingSourceScimArgs Empty => new PropertyMappingSourceScimArgs();
    }

    public sealed class PropertyMappingSourceScimState : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingSourceScimState()
        {
        }
        public static new PropertyMappingSourceScimState Empty => new PropertyMappingSourceScimState();
    }
}

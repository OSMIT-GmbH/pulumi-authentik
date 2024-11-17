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
    /// Manage Notification Property mappings
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Authentik = OSMIT_GmbH.Authentik;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var name = new Authentik.PropertyMappingNotification("name", new()
    ///     {
    ///         Expression = "return {\"foo\": context['foo']}",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/propertyMappingNotification:PropertyMappingNotification")]
    public partial class PropertyMappingNotification : global::Pulumi.CustomResource
    {
        [Output("expression")]
        public Output<string> Expression { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyMappingNotification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyMappingNotification(string name, PropertyMappingNotificationArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingNotification:PropertyMappingNotification", name, args ?? new PropertyMappingNotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyMappingNotification(string name, Input<string> id, PropertyMappingNotificationState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/propertyMappingNotification:PropertyMappingNotification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PropertyMappingNotification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertyMappingNotification Get(string name, Input<string> id, PropertyMappingNotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertyMappingNotification(name, id, state, options);
        }
    }

    public sealed class PropertyMappingNotificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingNotificationArgs()
        {
        }
        public static new PropertyMappingNotificationArgs Empty => new PropertyMappingNotificationArgs();
    }

    public sealed class PropertyMappingNotificationState : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PropertyMappingNotificationState()
        {
        }
        public static new PropertyMappingNotificationState Empty => new PropertyMappingNotificationState();
    }
}

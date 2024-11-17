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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Authentik = OSMIT_GmbH.Authentik;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Authentik.Blueprint("instance", new()
    ///     {
    ///         Path = "default/flow-default-authentication-flow.yaml",
    ///         Context = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["foo"] = "bar",
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/blueprint:Blueprint")]
    public partial class Blueprint : global::Pulumi.CustomResource
    {
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        /// </summary>
        [Output("context")]
        public Output<string?> Context { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;


        /// <summary>
        /// Create a Blueprint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Blueprint(string name, BlueprintArgs? args = null, CustomResourceOptions? options = null)
            : base("authentik:index/blueprint:Blueprint", name, args ?? new BlueprintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Blueprint(string name, Input<string> id, BlueprintState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/blueprint:Blueprint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Blueprint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Blueprint Get(string name, Input<string> id, BlueprintState? state = null, CustomResourceOptions? options = null)
        {
            return new Blueprint(name, id, state, options);
        }
    }

    public sealed class BlueprintArgs : global::Pulumi.ResourceArgs
    {
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        public BlueprintArgs()
        {
        }
        public static new BlueprintArgs Empty => new BlueprintArgs();
    }

    public sealed class BlueprintState : global::Pulumi.ResourceArgs
    {
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        public BlueprintState()
        {
        }
        public static new BlueprintState Empty => new BlueprintState();
    }
}

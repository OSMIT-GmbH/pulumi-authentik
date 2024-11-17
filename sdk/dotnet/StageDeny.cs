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
    /// using Pulumi;
    /// using Authentik = OSMIT_GmbH.Authentik;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create deny stage, can be used with policies
    ///     var name = new Authentik.StageDeny("name", new()
    ///     {
    ///         Name = "deny",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/stageDeny:StageDeny")]
    public partial class StageDeny : global::Pulumi.CustomResource
    {
        [Output("denyMessage")]
        public Output<string?> DenyMessage { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a StageDeny resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StageDeny(string name, StageDenyArgs? args = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageDeny:StageDeny", name, args ?? new StageDenyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StageDeny(string name, Input<string> id, StageDenyState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageDeny:StageDeny", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StageDeny resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StageDeny Get(string name, Input<string> id, StageDenyState? state = null, CustomResourceOptions? options = null)
        {
            return new StageDeny(name, id, state, options);
        }
    }

    public sealed class StageDenyArgs : global::Pulumi.ResourceArgs
    {
        [Input("denyMessage")]
        public Input<string>? DenyMessage { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StageDenyArgs()
        {
        }
        public static new StageDenyArgs Empty => new StageDenyArgs();
    }

    public sealed class StageDenyState : global::Pulumi.ResourceArgs
    {
        [Input("denyMessage")]
        public Input<string>? DenyMessage { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StageDenyState()
        {
        }
        public static new StageDenyState Empty => new StageDenyState();
    }
}

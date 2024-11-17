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
    ///     // Create a flow with a stage attached
    ///     var name = new Authentik.StageDummy("name", new()
    ///     {
    ///         Name = "test-stage",
    ///     });
    /// 
    ///     var flow = new Authentik.Flow("flow", new()
    ///     {
    ///         Name = "test-flow",
    ///         Title = "Test flow",
    ///         Slug = "test-flow",
    ///         Designation = "authorization",
    ///     });
    /// 
    ///     var dummy_flow = new Authentik.FlowStageBinding("dummy-flow", new()
    ///     {
    ///         Target = flow.Uuid,
    ///         Stage = name.Id,
    ///         Order = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/flow:Flow")]
    public partial class Flow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allowed values: - `none` - `require_authenticated` - `require_unauthenticated` - `require_superuser` - `require_outpost`
        /// </summary>
        [Output("authentication")]
        public Output<string?> Authentication { get; private set; } = null!;

        /// <summary>
        /// Optional URL to an image which will be used as the background during the flow.
        /// </summary>
        [Output("background")]
        public Output<string?> Background { get; private set; } = null!;

        [Output("compatibilityMode")]
        public Output<bool?> CompatibilityMode { get; private set; } = null!;

        [Output("deniedAction")]
        public Output<string?> DeniedAction { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
        /// `stage_configuration`
        /// </summary>
        [Output("designation")]
        public Output<string> Designation { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `stacked` - `content_left` - `content_right` - `sidebar_left` - `sidebar_right`
        /// </summary>
        [Output("layout")]
        public Output<string?> Layout { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `all` - `any`
        /// </summary>
        [Output("policyEngineMode")]
        public Output<string?> PolicyEngineMode { get; private set; } = null!;

        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a Flow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Flow(string name, FlowArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/flow:Flow", name, args ?? new FlowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Flow(string name, Input<string> id, FlowState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/flow:Flow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Flow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Flow Get(string name, Input<string> id, FlowState? state = null, CustomResourceOptions? options = null)
        {
            return new Flow(name, id, state, options);
        }
    }

    public sealed class FlowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allowed values: - `none` - `require_authenticated` - `require_unauthenticated` - `require_superuser` - `require_outpost`
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Optional URL to an image which will be used as the background during the flow.
        /// </summary>
        [Input("background")]
        public Input<string>? Background { get; set; }

        [Input("compatibilityMode")]
        public Input<bool>? CompatibilityMode { get; set; }

        [Input("deniedAction")]
        public Input<string>? DeniedAction { get; set; }

        /// <summary>
        /// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
        /// `stage_configuration`
        /// </summary>
        [Input("designation", required: true)]
        public Input<string> Designation { get; set; } = null!;

        /// <summary>
        /// Allowed values: - `stacked` - `content_left` - `content_right` - `sidebar_left` - `sidebar_right`
        /// </summary>
        [Input("layout")]
        public Input<string>? Layout { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `all` - `any`
        /// </summary>
        [Input("policyEngineMode")]
        public Input<string>? PolicyEngineMode { get; set; }

        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public FlowArgs()
        {
        }
        public static new FlowArgs Empty => new FlowArgs();
    }

    public sealed class FlowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allowed values: - `none` - `require_authenticated` - `require_unauthenticated` - `require_superuser` - `require_outpost`
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Optional URL to an image which will be used as the background during the flow.
        /// </summary>
        [Input("background")]
        public Input<string>? Background { get; set; }

        [Input("compatibilityMode")]
        public Input<bool>? CompatibilityMode { get; set; }

        [Input("deniedAction")]
        public Input<string>? DeniedAction { get; set; }

        /// <summary>
        /// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
        /// `stage_configuration`
        /// </summary>
        [Input("designation")]
        public Input<string>? Designation { get; set; }

        /// <summary>
        /// Allowed values: - `stacked` - `content_left` - `content_right` - `sidebar_left` - `sidebar_right`
        /// </summary>
        [Input("layout")]
        public Input<string>? Layout { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `all` - `any`
        /// </summary>
        [Input("policyEngineMode")]
        public Input<string>? PolicyEngineMode { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public FlowState()
        {
        }
        public static new FlowState Empty => new FlowState();
    }
}

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
    ///     // Create a prompt stage with 1 field
    ///     var field = new Authentik.StagePromptField("field", new()
    ///     {
    ///         FieldKey = "username",
    ///         Label = "Username",
    ///         Type = "username",
    ///     });
    /// 
    ///     var name = new Authentik.StagePrompt("name", new()
    ///     {
    ///         Name = "test",
    ///         Fields = new[]
    ///         {
    ///             authentikStagePromptField.Field.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/stagePrompt:StagePrompt")]
    public partial class StagePrompt : global::Pulumi.CustomResource
    {
        [Output("fields")]
        public Output<ImmutableArray<string>> Fields { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("validationPolicies")]
        public Output<ImmutableArray<string>> ValidationPolicies { get; private set; } = null!;


        /// <summary>
        /// Create a StagePrompt resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StagePrompt(string name, StagePromptArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/stagePrompt:StagePrompt", name, args ?? new StagePromptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StagePrompt(string name, Input<string> id, StagePromptState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/stagePrompt:StagePrompt", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StagePrompt resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StagePrompt Get(string name, Input<string> id, StagePromptState? state = null, CustomResourceOptions? options = null)
        {
            return new StagePrompt(name, id, state, options);
        }
    }

    public sealed class StagePromptArgs : global::Pulumi.ResourceArgs
    {
        [Input("fields", required: true)]
        private InputList<string>? _fields;
        public InputList<string> Fields
        {
            get => _fields ?? (_fields = new InputList<string>());
            set => _fields = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("validationPolicies")]
        private InputList<string>? _validationPolicies;
        public InputList<string> ValidationPolicies
        {
            get => _validationPolicies ?? (_validationPolicies = new InputList<string>());
            set => _validationPolicies = value;
        }

        public StagePromptArgs()
        {
        }
        public static new StagePromptArgs Empty => new StagePromptArgs();
    }

    public sealed class StagePromptState : global::Pulumi.ResourceArgs
    {
        [Input("fields")]
        private InputList<string>? _fields;
        public InputList<string> Fields
        {
            get => _fields ?? (_fields = new InputList<string>());
            set => _fields = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("validationPolicies")]
        private InputList<string>? _validationPolicies;
        public InputList<string> ValidationPolicies
        {
            get => _validationPolicies ?? (_validationPolicies = new InputList<string>());
            set => _validationPolicies = value;
        }

        public StagePromptState()
        {
        }
        public static new StagePromptState Empty => new StagePromptState();
    }
}
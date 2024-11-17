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
    /// using Authentik = Pulumi.Authentik;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create identification stage with sources and showing a password field
    ///     var default_authorization_flow = Authentik.GetFlow.Invoke(new()
    ///     {
    ///         Slug = "default-provider-authorization-implicit-consent",
    ///     });
    /// 
    ///     var name = new Authentik.SourceOauth("name", new()
    ///     {
    ///         Name = "test",
    ///         Slug = "test",
    ///         AuthenticationFlow = default_authorization_flow.Apply(default_authorization_flow =&gt; default_authorization_flow.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         EnrollmentFlow = default_authorization_flow.Apply(default_authorization_flow =&gt; default_authorization_flow.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         ProviderType = "discord",
    ///         ConsumerKey = "foo",
    ///         ConsumerSecret = "bar",
    ///     });
    /// 
    ///     var nameStagePassword = new Authentik.StagePassword("name", new()
    ///     {
    ///         Name = "test-pass",
    ///         Backends = new[]
    ///         {
    ///             "authentik.core.auth.InbuiltBackend",
    ///         },
    ///     });
    /// 
    ///     var nameStageIdentification = new Authentik.StageIdentification("name", new()
    ///     {
    ///         Name = "test-ident",
    ///         UserFields = new[]
    ///         {
    ///             "username",
    ///         },
    ///         Sources = new[]
    ///         {
    ///             name.Uuid,
    ///         },
    ///         PasswordStage = nameStagePassword.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/stageIdentification:StageIdentification")]
    public partial class StageIdentification : global::Pulumi.CustomResource
    {
        [Output("captchaStage")]
        public Output<string?> CaptchaStage { get; private set; } = null!;

        [Output("caseInsensitiveMatching")]
        public Output<bool?> CaseInsensitiveMatching { get; private set; } = null!;

        [Output("enrollmentFlow")]
        public Output<string?> EnrollmentFlow { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("passwordStage")]
        public Output<string?> PasswordStage { get; private set; } = null!;

        [Output("passwordlessFlow")]
        public Output<string?> PasswordlessFlow { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("pretendUserExists")]
        public Output<bool?> PretendUserExists { get; private set; } = null!;

        [Output("recoveryFlow")]
        public Output<string?> RecoveryFlow { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("showMatchedUser")]
        public Output<bool?> ShowMatchedUser { get; private set; } = null!;

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Output("showSourceLabels")]
        public Output<bool?> ShowSourceLabels { get; private set; } = null!;

        [Output("sources")]
        public Output<ImmutableArray<string>> Sources { get; private set; } = null!;

        [Output("userFields")]
        public Output<ImmutableArray<string>> UserFields { get; private set; } = null!;


        /// <summary>
        /// Create a StageIdentification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StageIdentification(string name, StageIdentificationArgs? args = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageIdentification:StageIdentification", name, args ?? new StageIdentificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StageIdentification(string name, Input<string> id, StageIdentificationState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageIdentification:StageIdentification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StageIdentification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StageIdentification Get(string name, Input<string> id, StageIdentificationState? state = null, CustomResourceOptions? options = null)
        {
            return new StageIdentification(name, id, state, options);
        }
    }

    public sealed class StageIdentificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("captchaStage")]
        public Input<string>? CaptchaStage { get; set; }

        [Input("caseInsensitiveMatching")]
        public Input<bool>? CaseInsensitiveMatching { get; set; }

        [Input("enrollmentFlow")]
        public Input<string>? EnrollmentFlow { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passwordStage")]
        public Input<string>? PasswordStage { get; set; }

        [Input("passwordlessFlow")]
        public Input<string>? PasswordlessFlow { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("pretendUserExists")]
        public Input<bool>? PretendUserExists { get; set; }

        [Input("recoveryFlow")]
        public Input<string>? RecoveryFlow { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("showMatchedUser")]
        public Input<bool>? ShowMatchedUser { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("showSourceLabels")]
        public Input<bool>? ShowSourceLabels { get; set; }

        [Input("sources")]
        private InputList<string>? _sources;
        public InputList<string> Sources
        {
            get => _sources ?? (_sources = new InputList<string>());
            set => _sources = value;
        }

        [Input("userFields")]
        private InputList<string>? _userFields;
        public InputList<string> UserFields
        {
            get => _userFields ?? (_userFields = new InputList<string>());
            set => _userFields = value;
        }

        public StageIdentificationArgs()
        {
        }
        public static new StageIdentificationArgs Empty => new StageIdentificationArgs();
    }

    public sealed class StageIdentificationState : global::Pulumi.ResourceArgs
    {
        [Input("captchaStage")]
        public Input<string>? CaptchaStage { get; set; }

        [Input("caseInsensitiveMatching")]
        public Input<bool>? CaseInsensitiveMatching { get; set; }

        [Input("enrollmentFlow")]
        public Input<string>? EnrollmentFlow { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passwordStage")]
        public Input<string>? PasswordStage { get; set; }

        [Input("passwordlessFlow")]
        public Input<string>? PasswordlessFlow { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("pretendUserExists")]
        public Input<bool>? PretendUserExists { get; set; }

        [Input("recoveryFlow")]
        public Input<string>? RecoveryFlow { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("showMatchedUser")]
        public Input<bool>? ShowMatchedUser { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("showSourceLabels")]
        public Input<bool>? ShowSourceLabels { get; set; }

        [Input("sources")]
        private InputList<string>? _sources;
        public InputList<string> Sources
        {
            get => _sources ?? (_sources = new InputList<string>());
            set => _sources = value;
        }

        [Input("userFields")]
        private InputList<string>? _userFields;
        public InputList<string> UserFields
        {
            get => _userFields ?? (_userFields = new InputList<string>());
            set => _userFields = value;
        }

        public StageIdentificationState()
        {
        }
        public static new StageIdentificationState Empty => new StageIdentificationState();
    }
}

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
    ///     // Create consent stage
    ///     var name = new Authentik.StageConsent("name", new()
    ///     {
    ///         Name = "consent",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/stageConsent:StageConsent")]
    public partial class StageConsent : global::Pulumi.CustomResource
    {
        [Output("consentExpireIn")]
        public Output<string?> ConsentExpireIn { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `always_require` - `permanent` - `expiring`
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a StageConsent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StageConsent(string name, StageConsentArgs? args = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageConsent:StageConsent", name, args ?? new StageConsentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StageConsent(string name, Input<string> id, StageConsentState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageConsent:StageConsent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StageConsent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StageConsent Get(string name, Input<string> id, StageConsentState? state = null, CustomResourceOptions? options = null)
        {
            return new StageConsent(name, id, state, options);
        }
    }

    public sealed class StageConsentArgs : global::Pulumi.ResourceArgs
    {
        [Input("consentExpireIn")]
        public Input<string>? ConsentExpireIn { get; set; }

        /// <summary>
        /// Allowed values: - `always_require` - `permanent` - `expiring`
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StageConsentArgs()
        {
        }
        public static new StageConsentArgs Empty => new StageConsentArgs();
    }

    public sealed class StageConsentState : global::Pulumi.ResourceArgs
    {
        [Input("consentExpireIn")]
        public Input<string>? ConsentExpireIn { get; set; }

        /// <summary>
        /// Allowed values: - `always_require` - `permanent` - `expiring`
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StageConsentState()
        {
        }
        public static new StageConsentState Empty => new StageConsentState();
    }
}

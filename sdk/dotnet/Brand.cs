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
    ///     var @default = new Authentik.Brand("default", new()
    ///     {
    ///         BrandingTitle = "test",
    ///         Default = true,
    ///         Domain = ".",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/brand:Brand")]
    public partial class Brand : global::Pulumi.CustomResource
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        /// </summary>
        [Output("attributes")]
        public Output<string?> Attributes { get; private set; } = null!;

        [Output("brandingFavicon")]
        public Output<string?> BrandingFavicon { get; private set; } = null!;

        [Output("brandingLogo")]
        public Output<string?> BrandingLogo { get; private set; } = null!;

        /// <summary>
        /// Defaults to `authentik`.
        /// </summary>
        [Output("brandingTitle")]
        public Output<string?> BrandingTitle { get; private set; } = null!;

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Output("default")]
        public Output<bool?> Default { get; private set; } = null!;

        [Output("defaultApplication")]
        public Output<string?> DefaultApplication { get; private set; } = null!;

        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("flowAuthentication")]
        public Output<string?> FlowAuthentication { get; private set; } = null!;

        [Output("flowDeviceCode")]
        public Output<string?> FlowDeviceCode { get; private set; } = null!;

        [Output("flowInvalidation")]
        public Output<string?> FlowInvalidation { get; private set; } = null!;

        [Output("flowRecovery")]
        public Output<string?> FlowRecovery { get; private set; } = null!;

        [Output("flowUnenrollment")]
        public Output<string?> FlowUnenrollment { get; private set; } = null!;

        [Output("flowUserSettings")]
        public Output<string?> FlowUserSettings { get; private set; } = null!;

        [Output("webCertificate")]
        public Output<string?> WebCertificate { get; private set; } = null!;


        /// <summary>
        /// Create a Brand resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Brand(string name, BrandArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/brand:Brand", name, args ?? new BrandArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Brand(string name, Input<string> id, BrandState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/brand:Brand", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Brand resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Brand Get(string name, Input<string> id, BrandState? state = null, CustomResourceOptions? options = null)
        {
            return new Brand(name, id, state, options);
        }
    }

    public sealed class BrandArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        /// </summary>
        [Input("attributes")]
        public Input<string>? Attributes { get; set; }

        [Input("brandingFavicon")]
        public Input<string>? BrandingFavicon { get; set; }

        [Input("brandingLogo")]
        public Input<string>? BrandingLogo { get; set; }

        /// <summary>
        /// Defaults to `authentik`.
        /// </summary>
        [Input("brandingTitle")]
        public Input<string>? BrandingTitle { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        [Input("defaultApplication")]
        public Input<string>? DefaultApplication { get; set; }

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("flowAuthentication")]
        public Input<string>? FlowAuthentication { get; set; }

        [Input("flowDeviceCode")]
        public Input<string>? FlowDeviceCode { get; set; }

        [Input("flowInvalidation")]
        public Input<string>? FlowInvalidation { get; set; }

        [Input("flowRecovery")]
        public Input<string>? FlowRecovery { get; set; }

        [Input("flowUnenrollment")]
        public Input<string>? FlowUnenrollment { get; set; }

        [Input("flowUserSettings")]
        public Input<string>? FlowUserSettings { get; set; }

        [Input("webCertificate")]
        public Input<string>? WebCertificate { get; set; }

        public BrandArgs()
        {
        }
        public static new BrandArgs Empty => new BrandArgs();
    }

    public sealed class BrandState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        /// </summary>
        [Input("attributes")]
        public Input<string>? Attributes { get; set; }

        [Input("brandingFavicon")]
        public Input<string>? BrandingFavicon { get; set; }

        [Input("brandingLogo")]
        public Input<string>? BrandingLogo { get; set; }

        /// <summary>
        /// Defaults to `authentik`.
        /// </summary>
        [Input("brandingTitle")]
        public Input<string>? BrandingTitle { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        [Input("defaultApplication")]
        public Input<string>? DefaultApplication { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("flowAuthentication")]
        public Input<string>? FlowAuthentication { get; set; }

        [Input("flowDeviceCode")]
        public Input<string>? FlowDeviceCode { get; set; }

        [Input("flowInvalidation")]
        public Input<string>? FlowInvalidation { get; set; }

        [Input("flowRecovery")]
        public Input<string>? FlowRecovery { get; set; }

        [Input("flowUnenrollment")]
        public Input<string>? FlowUnenrollment { get; set; }

        [Input("flowUserSettings")]
        public Input<string>? FlowUserSettings { get; set; }

        [Input("webCertificate")]
        public Input<string>? WebCertificate { get; set; }

        public BrandState()
        {
        }
        public static new BrandState Empty => new BrandState();
    }
}

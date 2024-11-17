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
    ///     var default_authorization_flow = Authentik.GetFlow.Invoke(new()
    ///     {
    ///         Slug = "default-provider-authorization-implicit-consent",
    ///     });
    /// 
    ///     var proxy = new Authentik.ProviderProxy("proxy", new()
    ///     {
    ///         AuthorizationFlow = default_authorization_flow.Apply(default_authorization_flow =&gt; default_authorization_flow.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         ExternalHost = "http://foo.bar.baz",
    ///         InternalHost = "http://internal.local",
    ///     });
    /// 
    ///     var outpost = new Authentik.Outpost("outpost", new()
    ///     {
    ///         ProtocolProviders = new[]
    ///         {
    ///             proxy.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/outpost:Outpost")]
    public partial class Outpost : global::Pulumi.CustomResource
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("protocolProviders")]
        public Output<ImmutableArray<int>> ProtocolProviders { get; private set; } = null!;

        [Output("serviceConnection")]
        public Output<string?> ServiceConnection { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `proxy` - `ldap` - `radius` - `rac`
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Outpost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Outpost(string name, OutpostArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/outpost:Outpost", name, args ?? new OutpostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Outpost(string name, Input<string> id, OutpostState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/outpost:Outpost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Outpost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Outpost Get(string name, Input<string> id, OutpostState? state = null, CustomResourceOptions? options = null)
        {
            return new Outpost(name, id, state, options);
        }
    }

    public sealed class OutpostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocolProviders", required: true)]
        private InputList<int>? _protocolProviders;
        public InputList<int> ProtocolProviders
        {
            get => _protocolProviders ?? (_protocolProviders = new InputList<int>());
            set => _protocolProviders = value;
        }

        [Input("serviceConnection")]
        public Input<string>? ServiceConnection { get; set; }

        /// <summary>
        /// Allowed values: - `proxy` - `ldap` - `radius` - `rac`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OutpostArgs()
        {
        }
        public static new OutpostArgs Empty => new OutpostArgs();
    }

    public sealed class OutpostState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocolProviders")]
        private InputList<int>? _protocolProviders;
        public InputList<int> ProtocolProviders
        {
            get => _protocolProviders ?? (_protocolProviders = new InputList<int>());
            set => _protocolProviders = value;
        }

        [Input("serviceConnection")]
        public Input<string>? ServiceConnection { get; set; }

        /// <summary>
        /// Allowed values: - `proxy` - `ldap` - `radius` - `rac`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OutpostState()
        {
        }
        public static new OutpostState Empty => new OutpostState();
    }
}

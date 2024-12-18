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
    ///     // Create a Radius Provider
    ///     var default_authentication_flow = Authentik.GetFlow.Invoke(new()
    ///     {
    ///         Slug = "default-authentication-flow",
    ///     });
    /// 
    ///     var name = new Authentik.ProviderRadius("name", new()
    ///     {
    ///         Name = "radius-app",
    ///         AuthorizationFlow = default_authentication_flow.Apply(default_authentication_flow =&gt; default_authentication_flow.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         ClientNetworks = "10.10.0.0/24",
    ///         SharedSecret = "my-shared-secret",
    ///     });
    /// 
    ///     var nameApplication = new Authentik.Application("name", new()
    ///     {
    ///         Name = "radius-app",
    ///         Slug = "radius-app",
    ///         ProtocolProvider = name.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/providerRadius:ProviderRadius")]
    public partial class ProviderRadius : global::Pulumi.CustomResource
    {
        [Output("authorizationFlow")]
        public Output<string> AuthorizationFlow { get; private set; } = null!;

        /// <summary>
        /// Defaults to `0.0.0.0/0, ::/0`.
        /// </summary>
        [Output("clientNetworks")]
        public Output<string?> ClientNetworks { get; private set; } = null!;

        [Output("invalidationFlow")]
        public Output<string> InvalidationFlow { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("mfaSupport")]
        public Output<bool?> MfaSupport { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("propertyMappings")]
        public Output<ImmutableArray<string>> PropertyMappings { get; private set; } = null!;

        [Output("sharedSecret")]
        public Output<string> SharedSecret { get; private set; } = null!;


        /// <summary>
        /// Create a ProviderRadius resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProviderRadius(string name, ProviderRadiusArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/providerRadius:ProviderRadius", name, args ?? new ProviderRadiusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProviderRadius(string name, Input<string> id, ProviderRadiusState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/providerRadius:ProviderRadius", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OSMIT-GmbH/pulumi-authentik/releases/download/v${VERSION}",
                AdditionalSecretOutputs =
                {
                    "sharedSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProviderRadius resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProviderRadius Get(string name, Input<string> id, ProviderRadiusState? state = null, CustomResourceOptions? options = null)
        {
            return new ProviderRadius(name, id, state, options);
        }
    }

    public sealed class ProviderRadiusArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizationFlow", required: true)]
        public Input<string> AuthorizationFlow { get; set; } = null!;

        /// <summary>
        /// Defaults to `0.0.0.0/0, ::/0`.
        /// </summary>
        [Input("clientNetworks")]
        public Input<string>? ClientNetworks { get; set; }

        [Input("invalidationFlow", required: true)]
        public Input<string> InvalidationFlow { get; set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("mfaSupport")]
        public Input<bool>? MfaSupport { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("propertyMappings")]
        private InputList<string>? _propertyMappings;
        public InputList<string> PropertyMappings
        {
            get => _propertyMappings ?? (_propertyMappings = new InputList<string>());
            set => _propertyMappings = value;
        }

        [Input("sharedSecret", required: true)]
        private Input<string>? _sharedSecret;
        public Input<string>? SharedSecret
        {
            get => _sharedSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sharedSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProviderRadiusArgs()
        {
        }
        public static new ProviderRadiusArgs Empty => new ProviderRadiusArgs();
    }

    public sealed class ProviderRadiusState : global::Pulumi.ResourceArgs
    {
        [Input("authorizationFlow")]
        public Input<string>? AuthorizationFlow { get; set; }

        /// <summary>
        /// Defaults to `0.0.0.0/0, ::/0`.
        /// </summary>
        [Input("clientNetworks")]
        public Input<string>? ClientNetworks { get; set; }

        [Input("invalidationFlow")]
        public Input<string>? InvalidationFlow { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("mfaSupport")]
        public Input<bool>? MfaSupport { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("propertyMappings")]
        private InputList<string>? _propertyMappings;
        public InputList<string> PropertyMappings
        {
            get => _propertyMappings ?? (_propertyMappings = new InputList<string>());
            set => _propertyMappings = value;
        }

        [Input("sharedSecret")]
        private Input<string>? _sharedSecret;
        public Input<string>? SharedSecret
        {
            get => _sharedSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sharedSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProviderRadiusState()
        {
        }
        public static new ProviderRadiusState Empty => new ProviderRadiusState();
    }
}

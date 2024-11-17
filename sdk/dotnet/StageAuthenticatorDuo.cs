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
    ///     // Create a duo setup stage
    ///     var name = new Authentik.StageAuthenticatorDuo("name", new()
    ///     {
    ///         Name = "duo-setup",
    ///         ClientId = "foo",
    ///         ClientSecret = "bar",
    ///         ApiHostname = "http://foo.bar.baz",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/stageAuthenticatorDuo:StageAuthenticatorDuo")]
    public partial class StageAuthenticatorDuo : global::Pulumi.CustomResource
    {
        [Output("adminIntegrationKey")]
        public Output<string?> AdminIntegrationKey { get; private set; } = null!;

        [Output("adminSecretKey")]
        public Output<string?> AdminSecretKey { get; private set; } = null!;

        [Output("apiHostname")]
        public Output<string> ApiHostname { get; private set; } = null!;

        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        [Output("configureFlow")]
        public Output<string?> ConfigureFlow { get; private set; } = null!;

        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a StageAuthenticatorDuo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StageAuthenticatorDuo(string name, StageAuthenticatorDuoArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/stageAuthenticatorDuo:StageAuthenticatorDuo", name, args ?? new StageAuthenticatorDuoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StageAuthenticatorDuo(string name, Input<string> id, StageAuthenticatorDuoState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageAuthenticatorDuo:StageAuthenticatorDuo", name, state, MakeResourceOptions(options, id))
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
                    "adminSecretKey",
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StageAuthenticatorDuo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StageAuthenticatorDuo Get(string name, Input<string> id, StageAuthenticatorDuoState? state = null, CustomResourceOptions? options = null)
        {
            return new StageAuthenticatorDuo(name, id, state, options);
        }
    }

    public sealed class StageAuthenticatorDuoArgs : global::Pulumi.ResourceArgs
    {
        [Input("adminIntegrationKey")]
        public Input<string>? AdminIntegrationKey { get; set; }

        [Input("adminSecretKey")]
        private Input<string>? _adminSecretKey;
        public Input<string>? AdminSecretKey
        {
            get => _adminSecretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _adminSecretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("apiHostname", required: true)]
        public Input<string> ApiHostname { get; set; } = null!;

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        private Input<string>? _clientSecret;
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("configureFlow")]
        public Input<string>? ConfigureFlow { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StageAuthenticatorDuoArgs()
        {
        }
        public static new StageAuthenticatorDuoArgs Empty => new StageAuthenticatorDuoArgs();
    }

    public sealed class StageAuthenticatorDuoState : global::Pulumi.ResourceArgs
    {
        [Input("adminIntegrationKey")]
        public Input<string>? AdminIntegrationKey { get; set; }

        [Input("adminSecretKey")]
        private Input<string>? _adminSecretKey;
        public Input<string>? AdminSecretKey
        {
            get => _adminSecretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _adminSecretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("apiHostname")]
        public Input<string>? ApiHostname { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("configureFlow")]
        public Input<string>? ConfigureFlow { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public StageAuthenticatorDuoState()
        {
        }
        public static new StageAuthenticatorDuoState Empty => new StageAuthenticatorDuoState();
    }
}

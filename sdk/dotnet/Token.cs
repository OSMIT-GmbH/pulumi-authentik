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
    ///     // Create a token for a user
    ///     var name = new Authentik.User("name", new()
    ///     {
    ///         Username = "user",
    ///         Name = "User",
    ///     });
    /// 
    ///     var @default = new Authentik.Token("default", new()
    ///     {
    ///         Identifier = "my-token",
    ///         User = name.Id,
    ///         Description = "My secret token",
    ///         Expires = "2025-01-01T15:04:05Z",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/token:Token")]
    public partial class Token : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("expires")]
        public Output<string?> Expires { get; private set; } = null!;

        [Output("expiresIn")]
        public Output<int> ExpiresIn { get; private set; } = null!;

        [Output("expiring")]
        public Output<bool?> Expiring { get; private set; } = null!;

        [Output("identifier")]
        public Output<string> Identifier { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `verification` - `api` - `recovery` - `app_password`
        /// </summary>
        [Output("intent")]
        public Output<string?> Intent { get; private set; } = null!;

        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        [Output("retrieveKey")]
        public Output<bool?> RetrieveKey { get; private set; } = null!;

        [Output("user")]
        public Output<int> User { get; private set; } = null!;


        /// <summary>
        /// Create a Token resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Token(string name, TokenArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/token:Token", name, args ?? new TokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Token(string name, Input<string> id, TokenState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/token:Token", name, state, MakeResourceOptions(options, id))
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
                    "key",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Token resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Token Get(string name, Input<string> id, TokenState? state = null, CustomResourceOptions? options = null)
        {
            return new Token(name, id, state, options);
        }
    }

    public sealed class TokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expires")]
        public Input<string>? Expires { get; set; }

        [Input("expiring")]
        public Input<bool>? Expiring { get; set; }

        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        /// <summary>
        /// Allowed values: - `verification` - `api` - `recovery` - `app_password`
        /// </summary>
        [Input("intent")]
        public Input<string>? Intent { get; set; }

        [Input("retrieveKey")]
        public Input<bool>? RetrieveKey { get; set; }

        [Input("user", required: true)]
        public Input<int> User { get; set; } = null!;

        public TokenArgs()
        {
        }
        public static new TokenArgs Empty => new TokenArgs();
    }

    public sealed class TokenState : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expires")]
        public Input<string>? Expires { get; set; }

        [Input("expiresIn")]
        public Input<int>? ExpiresIn { get; set; }

        [Input("expiring")]
        public Input<bool>? Expiring { get; set; }

        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// Allowed values: - `verification` - `api` - `recovery` - `app_password`
        /// </summary>
        [Input("intent")]
        public Input<string>? Intent { get; set; }

        [Input("key")]
        private Input<string>? _key;
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("retrieveKey")]
        public Input<bool>? RetrieveKey { get; set; }

        [Input("user")]
        public Input<int>? User { get; set; }

        public TokenState()
        {
        }
        public static new TokenState Empty => new TokenState();
    }
}
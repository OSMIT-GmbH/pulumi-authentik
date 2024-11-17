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
    /// </summary>
    [AuthentikResourceType("authentik:index/providerScim:ProviderScim")]
    public partial class ProviderScim : global::Pulumi.CustomResource
    {
        [Output("excludeUsersServiceAccount")]
        public Output<bool?> ExcludeUsersServiceAccount { get; private set; } = null!;

        [Output("filterGroup")]
        public Output<string?> FilterGroup { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("propertyMappings")]
        public Output<ImmutableArray<string>> PropertyMappings { get; private set; } = null!;

        [Output("propertyMappingsGroups")]
        public Output<ImmutableArray<string>> PropertyMappingsGroups { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ProviderScim resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProviderScim(string name, ProviderScimArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/providerScim:ProviderScim", name, args ?? new ProviderScimArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProviderScim(string name, Input<string> id, ProviderScimState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/providerScim:ProviderScim", name, state, MakeResourceOptions(options, id))
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
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProviderScim resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProviderScim Get(string name, Input<string> id, ProviderScimState? state = null, CustomResourceOptions? options = null)
        {
            return new ProviderScim(name, id, state, options);
        }
    }

    public sealed class ProviderScimArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludeUsersServiceAccount")]
        public Input<bool>? ExcludeUsersServiceAccount { get; set; }

        [Input("filterGroup")]
        public Input<string>? FilterGroup { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("propertyMappings")]
        private InputList<string>? _propertyMappings;
        public InputList<string> PropertyMappings
        {
            get => _propertyMappings ?? (_propertyMappings = new InputList<string>());
            set => _propertyMappings = value;
        }

        [Input("propertyMappingsGroups")]
        private InputList<string>? _propertyMappingsGroups;
        public InputList<string> PropertyMappingsGroups
        {
            get => _propertyMappingsGroups ?? (_propertyMappingsGroups = new InputList<string>());
            set => _propertyMappingsGroups = value;
        }

        [Input("token", required: true)]
        private Input<string>? _token;
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ProviderScimArgs()
        {
        }
        public static new ProviderScimArgs Empty => new ProviderScimArgs();
    }

    public sealed class ProviderScimState : global::Pulumi.ResourceArgs
    {
        [Input("excludeUsersServiceAccount")]
        public Input<bool>? ExcludeUsersServiceAccount { get; set; }

        [Input("filterGroup")]
        public Input<string>? FilterGroup { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("propertyMappings")]
        private InputList<string>? _propertyMappings;
        public InputList<string> PropertyMappings
        {
            get => _propertyMappings ?? (_propertyMappings = new InputList<string>());
            set => _propertyMappings = value;
        }

        [Input("propertyMappingsGroups")]
        private InputList<string>? _propertyMappingsGroups;
        public InputList<string> PropertyMappingsGroups
        {
            get => _propertyMappingsGroups ?? (_propertyMappingsGroups = new InputList<string>());
            set => _propertyMappingsGroups = value;
        }

        [Input("token")]
        private Input<string>? _token;
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProviderScimState()
        {
        }
        public static new ProviderScimState Empty => new ProviderScimState();
    }
}

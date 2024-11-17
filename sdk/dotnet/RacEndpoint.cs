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
    [AuthentikResourceType("authentik:index/racEndpoint:RacEndpoint")]
    public partial class RacEndpoint : global::Pulumi.CustomResource
    {
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        [Output("maximumConnections")]
        public Output<int?> MaximumConnections { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("propertyMappings")]
        public Output<ImmutableArray<string>> PropertyMappings { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `rdp` - `vnc` - `ssh`
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        [Output("protocolProvider")]
        public Output<int> ProtocolProvider { get; private set; } = null!;

        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Output("settings")]
        public Output<string?> Settings { get; private set; } = null!;


        /// <summary>
        /// Create a RacEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RacEndpoint(string name, RacEndpointArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/racEndpoint:RacEndpoint", name, args ?? new RacEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RacEndpoint(string name, Input<string> id, RacEndpointState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/racEndpoint:RacEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RacEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RacEndpoint Get(string name, Input<string> id, RacEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new RacEndpoint(name, id, state, options);
        }
    }

    public sealed class RacEndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("maximumConnections")]
        public Input<int>? MaximumConnections { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("propertyMappings")]
        private InputList<string>? _propertyMappings;
        public InputList<string> PropertyMappings
        {
            get => _propertyMappings ?? (_propertyMappings = new InputList<string>());
            set => _propertyMappings = value;
        }

        /// <summary>
        /// Allowed values: - `rdp` - `vnc` - `ssh`
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("protocolProvider", required: true)]
        public Input<int> ProtocolProvider { get; set; } = null!;

        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        public RacEndpointArgs()
        {
        }
        public static new RacEndpointArgs Empty => new RacEndpointArgs();
    }

    public sealed class RacEndpointState : global::Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("maximumConnections")]
        public Input<int>? MaximumConnections { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("propertyMappings")]
        private InputList<string>? _propertyMappings;
        public InputList<string> PropertyMappings
        {
            get => _propertyMappings ?? (_propertyMappings = new InputList<string>());
            set => _propertyMappings = value;
        }

        /// <summary>
        /// Allowed values: - `rdp` - `vnc` - `ssh`
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("protocolProvider")]
        public Input<int>? ProtocolProvider { get; set; }

        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        public RacEndpointState()
        {
        }
        public static new RacEndpointState Empty => new RacEndpointState();
    }
}

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
    [AuthentikResourceType("authentik:index/providerGoogleWorkspace:ProviderGoogleWorkspace")]
    public partial class ProviderGoogleWorkspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Output("credentials")]
        public Output<string?> Credentials { get; private set; } = null!;

        [Output("defaultGroupEmailDomain")]
        public Output<string> DefaultGroupEmailDomain { get; private set; } = null!;

        [Output("delegatedSubject")]
        public Output<string?> DelegatedSubject { get; private set; } = null!;

        [Output("excludeUsersServiceAccount")]
        public Output<bool?> ExcludeUsersServiceAccount { get; private set; } = null!;

        [Output("filterGroup")]
        public Output<string?> FilterGroup { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `delete` - `do_nothing`
        /// </summary>
        [Output("groupDeleteAction")]
        public Output<string?> GroupDeleteAction { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("propertyMappings")]
        public Output<ImmutableArray<string>> PropertyMappings { get; private set; } = null!;

        [Output("propertyMappingsGroups")]
        public Output<ImmutableArray<string>> PropertyMappingsGroups { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `do_nothing` - `delete` - `suspend`
        /// </summary>
        [Output("userDeleteAction")]
        public Output<string?> UserDeleteAction { get; private set; } = null!;


        /// <summary>
        /// Create a ProviderGoogleWorkspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProviderGoogleWorkspace(string name, ProviderGoogleWorkspaceArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/providerGoogleWorkspace:ProviderGoogleWorkspace", name, args ?? new ProviderGoogleWorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProviderGoogleWorkspace(string name, Input<string> id, ProviderGoogleWorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/providerGoogleWorkspace:ProviderGoogleWorkspace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProviderGoogleWorkspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProviderGoogleWorkspace Get(string name, Input<string> id, ProviderGoogleWorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new ProviderGoogleWorkspace(name, id, state, options);
        }
    }

    public sealed class ProviderGoogleWorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        [Input("defaultGroupEmailDomain", required: true)]
        public Input<string> DefaultGroupEmailDomain { get; set; } = null!;

        [Input("delegatedSubject")]
        public Input<string>? DelegatedSubject { get; set; }

        [Input("excludeUsersServiceAccount")]
        public Input<bool>? ExcludeUsersServiceAccount { get; set; }

        [Input("filterGroup")]
        public Input<string>? FilterGroup { get; set; }

        /// <summary>
        /// Allowed values: - `delete` - `do_nothing`
        /// </summary>
        [Input("groupDeleteAction")]
        public Input<string>? GroupDeleteAction { get; set; }

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

        /// <summary>
        /// Allowed values: - `do_nothing` - `delete` - `suspend`
        /// </summary>
        [Input("userDeleteAction")]
        public Input<string>? UserDeleteAction { get; set; }

        public ProviderGoogleWorkspaceArgs()
        {
        }
        public static new ProviderGoogleWorkspaceArgs Empty => new ProviderGoogleWorkspaceArgs();
    }

    public sealed class ProviderGoogleWorkspaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON format expected. Use jsonencode() to pass objects.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        [Input("defaultGroupEmailDomain")]
        public Input<string>? DefaultGroupEmailDomain { get; set; }

        [Input("delegatedSubject")]
        public Input<string>? DelegatedSubject { get; set; }

        [Input("excludeUsersServiceAccount")]
        public Input<bool>? ExcludeUsersServiceAccount { get; set; }

        [Input("filterGroup")]
        public Input<string>? FilterGroup { get; set; }

        /// <summary>
        /// Allowed values: - `delete` - `do_nothing`
        /// </summary>
        [Input("groupDeleteAction")]
        public Input<string>? GroupDeleteAction { get; set; }

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

        /// <summary>
        /// Allowed values: - `do_nothing` - `delete` - `suspend`
        /// </summary>
        [Input("userDeleteAction")]
        public Input<string>? UserDeleteAction { get; set; }

        public ProviderGoogleWorkspaceState()
        {
        }
        public static new ProviderGoogleWorkspaceState Empty => new ProviderGoogleWorkspaceState();
    }
}

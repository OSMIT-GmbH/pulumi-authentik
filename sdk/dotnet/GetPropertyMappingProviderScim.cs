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
    public static class GetPropertyMappingProviderScim
    {
        /// <summary>
        /// Get SCIM Provider Property mappings
        /// </summary>
        public static Task<GetPropertyMappingProviderScimResult> InvokeAsync(GetPropertyMappingProviderScimArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPropertyMappingProviderScimResult>("authentik:index/getPropertyMappingProviderScim:getPropertyMappingProviderScim", args ?? new GetPropertyMappingProviderScimArgs(), options.WithDefaults());

        /// <summary>
        /// Get SCIM Provider Property mappings
        /// </summary>
        public static Output<GetPropertyMappingProviderScimResult> Invoke(GetPropertyMappingProviderScimInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPropertyMappingProviderScimResult>("authentik:index/getPropertyMappingProviderScim:getPropertyMappingProviderScim", args ?? new GetPropertyMappingProviderScimInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPropertyMappingProviderScimArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// List of ids when `managed_list` is set. Generated.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("managed")]
        public string? Managed { get; set; }

        [Input("managedLists")]
        private List<string>? _managedLists;

        /// <summary>
        /// Retrieve multiple property mappings
        /// </summary>
        public List<string> ManagedLists
        {
            get => _managedLists ?? (_managedLists = new List<string>());
            set => _managedLists = value;
        }

        [Input("name")]
        public string? Name { get; set; }

        public GetPropertyMappingProviderScimArgs()
        {
        }
        public static new GetPropertyMappingProviderScimArgs Empty => new GetPropertyMappingProviderScimArgs();
    }

    public sealed class GetPropertyMappingProviderScimInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// List of ids when `managed_list` is set. Generated.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("managed")]
        public Input<string>? Managed { get; set; }

        [Input("managedLists")]
        private InputList<string>? _managedLists;

        /// <summary>
        /// Retrieve multiple property mappings
        /// </summary>
        public InputList<string> ManagedLists
        {
            get => _managedLists ?? (_managedLists = new InputList<string>());
            set => _managedLists = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetPropertyMappingProviderScimInvokeArgs()
        {
        }
        public static new GetPropertyMappingProviderScimInvokeArgs Empty => new GetPropertyMappingProviderScimInvokeArgs();
    }


    [OutputType]
    public sealed class GetPropertyMappingProviderScimResult
    {
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of ids when `managed_list` is set. Generated.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? Managed;
        /// <summary>
        /// Retrieve multiple property mappings
        /// </summary>
        public readonly ImmutableArray<string> ManagedLists;
        public readonly string? Name;

        [OutputConstructor]
        private GetPropertyMappingProviderScimResult(
            string expression,

            string id,

            ImmutableArray<string> ids,

            string? managed,

            ImmutableArray<string> managedLists,

            string? name)
        {
            Expression = expression;
            Id = id;
            Ids = ids;
            Managed = managed;
            ManagedLists = managedLists;
            Name = name;
        }
    }
}

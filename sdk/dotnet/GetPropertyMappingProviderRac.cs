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
    public static class GetPropertyMappingProviderRac
    {
        /// <summary>
        /// Get RAC Provider Property mappings
        /// </summary>
        public static Task<GetPropertyMappingProviderRacResult> InvokeAsync(GetPropertyMappingProviderRacArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPropertyMappingProviderRacResult>("authentik:index/getPropertyMappingProviderRac:getPropertyMappingProviderRac", args ?? new GetPropertyMappingProviderRacArgs(), options.WithDefaults());

        /// <summary>
        /// Get RAC Provider Property mappings
        /// </summary>
        public static Output<GetPropertyMappingProviderRacResult> Invoke(GetPropertyMappingProviderRacInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPropertyMappingProviderRacResult>("authentik:index/getPropertyMappingProviderRac:getPropertyMappingProviderRac", args ?? new GetPropertyMappingProviderRacInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPropertyMappingProviderRacArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// Generated.
        /// </summary>
        [Input("settings")]
        public string? Settings { get; set; }

        public GetPropertyMappingProviderRacArgs()
        {
        }
        public static new GetPropertyMappingProviderRacArgs Empty => new GetPropertyMappingProviderRacArgs();
    }

    public sealed class GetPropertyMappingProviderRacInvokeArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// Generated.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        public GetPropertyMappingProviderRacInvokeArgs()
        {
        }
        public static new GetPropertyMappingProviderRacInvokeArgs Empty => new GetPropertyMappingProviderRacInvokeArgs();
    }


    [OutputType]
    public sealed class GetPropertyMappingProviderRacResult
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
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Settings;

        [OutputConstructor]
        private GetPropertyMappingProviderRacResult(
            string expression,

            string id,

            ImmutableArray<string> ids,

            string? managed,

            ImmutableArray<string> managedLists,

            string? name,

            string settings)
        {
            Expression = expression;
            Id = id;
            Ids = ids;
            Managed = managed;
            ManagedLists = managedLists;
            Name = name;
            Settings = settings;
        }
    }
}
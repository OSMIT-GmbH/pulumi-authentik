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
    public static class GetGroup
    {
        /// <summary>
        /// Get groups by pk or name
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Authentik = Pulumi.Authentik;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // To get the ID of a group by name
        ///     var admins = Authentik.GetGroup.Invoke(new()
        ///     {
        ///         Name = "authentik Admins",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("authentik:index/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Get groups by pk or name
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Authentik = Pulumi.Authentik;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // To get the ID of a group by name
        ///     var admins = Authentik.GetGroup.Invoke(new()
        ///     {
        ///         Name = "authentik Admins",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("authentik:index/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("includeUsers")]
        public bool? IncludeUsers { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("pk")]
        public string? Pk { get; set; }

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("includeUsers")]
        public Input<bool>? IncludeUsers { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pk")]
        public Input<string>? Pk { get; set; }

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Attributes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeUsers;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly bool IsSuperuser;
        public readonly string? Name;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly int NumPk;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Parent;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string ParentName;
        public readonly string? Pk;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly ImmutableArray<int> Users;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupUsersObjResult> UsersObjs;

        [OutputConstructor]
        private GetGroupResult(
            string attributes,

            string id,

            bool? includeUsers,

            bool isSuperuser,

            string? name,

            int numPk,

            string parent,

            string parentName,

            string? pk,

            ImmutableArray<int> users,

            ImmutableArray<Outputs.GetGroupUsersObjResult> usersObjs)
        {
            Attributes = attributes;
            Id = id;
            IncludeUsers = includeUsers;
            IsSuperuser = isSuperuser;
            Name = name;
            NumPk = numPk;
            Parent = parent;
            ParentName = parentName;
            Pk = pk;
            Users = users;
            UsersObjs = usersObjs;
        }
    }
}
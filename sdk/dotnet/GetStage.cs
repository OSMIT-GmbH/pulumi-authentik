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
    public static class GetStage
    {
        /// <summary>
        /// Get stages by name
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
        ///     var default_authentication_identification = Authentik.GetStage.Invoke(new()
        ///     {
        ///         Name = "default-authentication-identification",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetStageResult> InvokeAsync(GetStageArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStageResult>("authentik:index/getStage:getStage", args ?? new GetStageArgs(), options.WithDefaults());

        /// <summary>
        /// Get stages by name
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
        ///     var default_authentication_identification = Authentik.GetStage.Invoke(new()
        ///     {
        ///         Name = "default-authentication-identification",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetStageResult> Invoke(GetStageInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStageResult>("authentik:index/getStage:getStage", args ?? new GetStageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Generated.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetStageArgs()
        {
        }
        public static new GetStageArgs Empty => new GetStageArgs();
    }

    public sealed class GetStageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetStageInvokeArgs()
        {
        }
        public static new GetStageInvokeArgs Empty => new GetStageInvokeArgs();
    }


    [OutputType]
    public sealed class GetStageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetStageResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}

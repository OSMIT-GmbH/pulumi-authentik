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
    ///     // Create expiry policy
    ///     var name = new Authentik.PolicyExpiry("name", new()
    ///     {
    ///         Name = "expiry",
    ///         Days = 3,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/policyExpiry:PolicyExpiry")]
    public partial class PolicyExpiry : global::Pulumi.CustomResource
    {
        [Output("days")]
        public Output<int> Days { get; private set; } = null!;

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Output("denyOnly")]
        public Output<bool?> DenyOnly { get; private set; } = null!;

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Output("executionLogging")]
        public Output<bool?> ExecutionLogging { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyExpiry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyExpiry(string name, PolicyExpiryArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/policyExpiry:PolicyExpiry", name, args ?? new PolicyExpiryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyExpiry(string name, Input<string> id, PolicyExpiryState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/policyExpiry:PolicyExpiry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyExpiry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyExpiry Get(string name, Input<string> id, PolicyExpiryState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyExpiry(name, id, state, options);
        }
    }

    public sealed class PolicyExpiryArgs : global::Pulumi.ResourceArgs
    {
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("denyOnly")]
        public Input<bool>? DenyOnly { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("executionLogging")]
        public Input<bool>? ExecutionLogging { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PolicyExpiryArgs()
        {
        }
        public static new PolicyExpiryArgs Empty => new PolicyExpiryArgs();
    }

    public sealed class PolicyExpiryState : global::Pulumi.ResourceArgs
    {
        [Input("days")]
        public Input<int>? Days { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("denyOnly")]
        public Input<bool>? DenyOnly { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("executionLogging")]
        public Input<bool>? ExecutionLogging { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PolicyExpiryState()
        {
        }
        public static new PolicyExpiryState Empty => new PolicyExpiryState();
    }
}

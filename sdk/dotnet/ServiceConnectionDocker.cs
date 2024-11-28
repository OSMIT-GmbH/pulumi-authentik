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
    ///     // Create a local docker connection
    ///     var local = new Authentik.ServiceConnectionDocker("local", new()
    ///     {
    ///         Name = "local",
    ///         Local = true,
    ///     });
    /// 
    ///     // Create a remote docker connection
    ///     var tls_auth = new Authentik.CertificateKeyPair("tls-auth", new()
    ///     {
    ///         Name = "docker-tls-auth",
    ///         CertificateData = "...",
    ///         KeyData = "...",
    ///     });
    /// 
    ///     var tls_verification = new Authentik.CertificateKeyPair("tls-verification", new()
    ///     {
    ///         Name = "docker-tls-verification",
    ///         CertificateData = "...",
    ///     });
    /// 
    ///     var remote_host = new Authentik.ServiceConnectionDocker("remote-host", new()
    ///     {
    ///         Name = "remote-host",
    ///         Url = "http://1.2.3.4:2368",
    ///         TlsVerification = tls_auth.Id,
    ///         TlsAuthentication = tls_verification.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/serviceConnectionDocker:ServiceConnectionDocker")]
    public partial class ServiceConnectionDocker : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tlsAuthentication")]
        public Output<string?> TlsAuthentication { get; private set; } = null!;

        [Output("tlsVerification")]
        public Output<string?> TlsVerification { get; private set; } = null!;

        /// <summary>
        /// Defaults to `http+unix:///var/run/docker.sock`.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceConnectionDocker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceConnectionDocker(string name, ServiceConnectionDockerArgs? args = null, CustomResourceOptions? options = null)
            : base("authentik:index/serviceConnectionDocker:ServiceConnectionDocker", name, args ?? new ServiceConnectionDockerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceConnectionDocker(string name, Input<string> id, ServiceConnectionDockerState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/serviceConnectionDocker:ServiceConnectionDocker", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceConnectionDocker resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceConnectionDocker Get(string name, Input<string> id, ServiceConnectionDockerState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceConnectionDocker(name, id, state, options);
        }
    }

    public sealed class ServiceConnectionDockerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tlsAuthentication")]
        public Input<string>? TlsAuthentication { get; set; }

        [Input("tlsVerification")]
        public Input<string>? TlsVerification { get; set; }

        /// <summary>
        /// Defaults to `http+unix:///var/run/docker.sock`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceConnectionDockerArgs()
        {
        }
        public static new ServiceConnectionDockerArgs Empty => new ServiceConnectionDockerArgs();
    }

    public sealed class ServiceConnectionDockerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tlsAuthentication")]
        public Input<string>? TlsAuthentication { get; set; }

        [Input("tlsVerification")]
        public Input<string>? TlsVerification { get; set; }

        /// <summary>
        /// Defaults to `http+unix:///var/run/docker.sock`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceConnectionDockerState()
        {
        }
        public static new ServiceConnectionDockerState Empty => new ServiceConnectionDockerState();
    }
}
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
    public static class GetCertificateKeyPair
    {
        /// <summary>
        /// Get certificate-key pairs by name
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
        ///     // To get the the ID and other info about a certificate
        ///     var generated = Authentik.GetCertificateKeyPair.Invoke(new()
        ///     {
        ///         Name = "authentik Self-signed Certificate",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCertificateKeyPairResult> InvokeAsync(GetCertificateKeyPairArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateKeyPairResult>("authentik:index/getCertificateKeyPair:getCertificateKeyPair", args ?? new GetCertificateKeyPairArgs(), options.WithDefaults());

        /// <summary>
        /// Get certificate-key pairs by name
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
        ///     // To get the the ID and other info about a certificate
        ///     var generated = Authentik.GetCertificateKeyPair.Invoke(new()
        ///     {
        ///         Name = "authentik Self-signed Certificate",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCertificateKeyPairResult> Invoke(GetCertificateKeyPairInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateKeyPairResult>("authentik:index/getCertificateKeyPair:getCertificateKeyPair", args ?? new GetCertificateKeyPairInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateKeyPairArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If set to true, certificate data will be fetched. Defaults to `true`.
        /// </summary>
        [Input("fetchCertificate")]
        public bool? FetchCertificate { get; set; }

        /// <summary>
        /// If set to true, private key data will be fetched. Defaults to `true`.
        /// </summary>
        [Input("fetchKey")]
        public bool? FetchKey { get; set; }

        [Input("keyData")]
        private string? _keyData;

        /// <summary>
        /// Generated.
        /// </summary>
        public string? KeyData
        {
            get => _keyData;
            set => _keyData = value;
        }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCertificateKeyPairArgs()
        {
        }
        public static new GetCertificateKeyPairArgs Empty => new GetCertificateKeyPairArgs();
    }

    public sealed class GetCertificateKeyPairInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If set to true, certificate data will be fetched. Defaults to `true`.
        /// </summary>
        [Input("fetchCertificate")]
        public Input<bool>? FetchCertificate { get; set; }

        /// <summary>
        /// If set to true, private key data will be fetched. Defaults to `true`.
        /// </summary>
        [Input("fetchKey")]
        public Input<bool>? FetchKey { get; set; }

        [Input("keyData")]
        private Input<string>? _keyData;

        /// <summary>
        /// Generated.
        /// </summary>
        public Input<string>? KeyData
        {
            get => _keyData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetCertificateKeyPairInvokeArgs()
        {
        }
        public static new GetCertificateKeyPairInvokeArgs Empty => new GetCertificateKeyPairInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateKeyPairResult
    {
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string CertificateData;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Expiry;
        /// <summary>
        /// If set to true, certificate data will be fetched. Defaults to `true`.
        /// </summary>
        public readonly bool? FetchCertificate;
        /// <summary>
        /// If set to true, private key data will be fetched. Defaults to `true`.
        /// </summary>
        public readonly bool? FetchKey;
        /// <summary>
        /// SHA1-hashed certificate fingerprint Generated.
        /// </summary>
        public readonly string Fingerprint1;
        /// <summary>
        /// SHA256-hashed certificate fingerprint Generated.
        /// </summary>
        public readonly string Fingerprint256;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string KeyData;
        public readonly string Name;
        /// <summary>
        /// Generated.
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private GetCertificateKeyPairResult(
            string certificateData,

            string expiry,

            bool? fetchCertificate,

            bool? fetchKey,

            string fingerprint1,

            string fingerprint256,

            string id,

            string keyData,

            string name,

            string subject)
        {
            CertificateData = certificateData;
            Expiry = expiry;
            FetchCertificate = fetchCertificate;
            FetchKey = fetchKey;
            Fingerprint1 = fingerprint1;
            Fingerprint256 = fingerprint256;
            Id = id;
            KeyData = keyData;
            Name = name;
            Subject = subject;
        }
    }
}
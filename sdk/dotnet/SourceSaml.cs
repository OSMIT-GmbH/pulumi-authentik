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
    /// using Authentik = Pulumi.Authentik;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a SAML Source
    ///     var default_source_pre_authentication = Authentik.GetFlow.Invoke(new()
    ///     {
    ///         Slug = "default-source-pre-authentication",
    ///     });
    /// 
    ///     var default_source_authentication = Authentik.GetFlow.Invoke(new()
    ///     {
    ///         Slug = "default-source-authentication",
    ///     });
    /// 
    ///     var default_source_enrollment = Authentik.GetFlow.Invoke(new()
    ///     {
    ///         Slug = "default-source-enrollment",
    ///     });
    /// 
    ///     var name = new Authentik.SourceSaml("name", new()
    ///     {
    ///         Name = "test-source",
    ///         Slug = "test-source",
    ///         AuthenticationFlow = default_source_authentication.Apply(default_source_authentication =&gt; default_source_authentication.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         EnrollmentFlow = default_source_enrollment.Apply(default_source_enrollment =&gt; default_source_enrollment.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         PreAuthenticationFlow = default_source_pre_authentication.Apply(default_source_pre_authentication =&gt; default_source_pre_authentication.Apply(getFlowResult =&gt; getFlowResult.Id)),
    ///         SsoUrl = "http://localhost",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AuthentikResourceType("authentik:index/sourceSaml:SourceSaml")]
    public partial class SourceSaml : global::Pulumi.CustomResource
    {
        [Output("allowIdpInitiated")]
        public Output<bool?> AllowIdpInitiated { get; private set; } = null!;

        [Output("authenticationFlow")]
        public Output<string?> AuthenticationFlow { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `REDIRECT` - `POST` - `POST_AUTO`
        /// </summary>
        [Output("bindingType")]
        public Output<string?> BindingType { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `http://www.w3.org/2000/09/xmldsig#sha1` - `http://www.w3.org/2001/04/xmlenc#sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#sha384` - `http://www.w3.org/2001/04/xmlenc#sha512`
        /// </summary>
        [Output("digestAlgorithm")]
        public Output<string?> DigestAlgorithm { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("encryptionKp")]
        public Output<string?> EncryptionKp { get; private set; } = null!;

        [Output("enrollmentFlow")]
        public Output<string?> EnrollmentFlow { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `identifier` - `name_link` - `name_deny`
        /// </summary>
        [Output("groupMatchingMode")]
        public Output<string?> GroupMatchingMode { get; private set; } = null!;

        [Output("issuer")]
        public Output<string?> Issuer { get; private set; } = null!;

        /// <summary>
        /// SAML Metadata
        /// </summary>
        [Output("metadata")]
        public Output<string> Metadata { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent` - `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`
        /// </summary>
        [Output("nameIdPolicy")]
        public Output<string?> NameIdPolicy { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `all` - `any`
        /// </summary>
        [Output("policyEngineMode")]
        public Output<string?> PolicyEngineMode { get; private set; } = null!;

        [Output("preAuthenticationFlow")]
        public Output<string> PreAuthenticationFlow { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `http://www.w3.org/2000/09/xmldsig#rsa-sha1` - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384` - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512` -
        /// `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1` - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384` - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512` -
        /// `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
        /// </summary>
        [Output("signatureAlgorithm")]
        public Output<string?> SignatureAlgorithm { get; private set; } = null!;

        [Output("signingKp")]
        public Output<string?> SigningKp { get; private set; } = null!;

        [Output("sloUrl")]
        public Output<string?> SloUrl { get; private set; } = null!;

        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("ssoUrl")]
        public Output<string> SsoUrl { get; private set; } = null!;

        [Output("temporaryUserDeleteAfter")]
        public Output<string?> TemporaryUserDeleteAfter { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        /// </summary>
        [Output("userMatchingMode")]
        public Output<string?> UserMatchingMode { get; private set; } = null!;

        [Output("userPathTemplate")]
        public Output<string?> UserPathTemplate { get; private set; } = null!;

        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a SourceSaml resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceSaml(string name, SourceSamlArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/sourceSaml:SourceSaml", name, args ?? new SourceSamlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceSaml(string name, Input<string> id, SourceSamlState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/sourceSaml:SourceSaml", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SourceSaml resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceSaml Get(string name, Input<string> id, SourceSamlState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceSaml(name, id, state, options);
        }
    }

    public sealed class SourceSamlArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowIdpInitiated")]
        public Input<bool>? AllowIdpInitiated { get; set; }

        [Input("authenticationFlow")]
        public Input<string>? AuthenticationFlow { get; set; }

        /// <summary>
        /// Allowed values: - `REDIRECT` - `POST` - `POST_AUTO`
        /// </summary>
        [Input("bindingType")]
        public Input<string>? BindingType { get; set; }

        /// <summary>
        /// Allowed values: - `http://www.w3.org/2000/09/xmldsig#sha1` - `http://www.w3.org/2001/04/xmlenc#sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#sha384` - `http://www.w3.org/2001/04/xmlenc#sha512`
        /// </summary>
        [Input("digestAlgorithm")]
        public Input<string>? DigestAlgorithm { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("encryptionKp")]
        public Input<string>? EncryptionKp { get; set; }

        [Input("enrollmentFlow")]
        public Input<string>? EnrollmentFlow { get; set; }

        /// <summary>
        /// Allowed values: - `identifier` - `name_link` - `name_deny`
        /// </summary>
        [Input("groupMatchingMode")]
        public Input<string>? GroupMatchingMode { get; set; }

        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent` - `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`
        /// </summary>
        [Input("nameIdPolicy")]
        public Input<string>? NameIdPolicy { get; set; }

        /// <summary>
        /// Allowed values: - `all` - `any`
        /// </summary>
        [Input("policyEngineMode")]
        public Input<string>? PolicyEngineMode { get; set; }

        [Input("preAuthenticationFlow", required: true)]
        public Input<string> PreAuthenticationFlow { get; set; } = null!;

        /// <summary>
        /// Allowed values: - `http://www.w3.org/2000/09/xmldsig#rsa-sha1` - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384` - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512` -
        /// `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1` - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384` - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512` -
        /// `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        [Input("signingKp")]
        public Input<string>? SigningKp { get; set; }

        [Input("sloUrl")]
        public Input<string>? SloUrl { get; set; }

        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        [Input("ssoUrl", required: true)]
        public Input<string> SsoUrl { get; set; } = null!;

        [Input("temporaryUserDeleteAfter")]
        public Input<string>? TemporaryUserDeleteAfter { get; set; }

        /// <summary>
        /// Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        /// </summary>
        [Input("userMatchingMode")]
        public Input<string>? UserMatchingMode { get; set; }

        [Input("userPathTemplate")]
        public Input<string>? UserPathTemplate { get; set; }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public SourceSamlArgs()
        {
        }
        public static new SourceSamlArgs Empty => new SourceSamlArgs();
    }

    public sealed class SourceSamlState : global::Pulumi.ResourceArgs
    {
        [Input("allowIdpInitiated")]
        public Input<bool>? AllowIdpInitiated { get; set; }

        [Input("authenticationFlow")]
        public Input<string>? AuthenticationFlow { get; set; }

        /// <summary>
        /// Allowed values: - `REDIRECT` - `POST` - `POST_AUTO`
        /// </summary>
        [Input("bindingType")]
        public Input<string>? BindingType { get; set; }

        /// <summary>
        /// Allowed values: - `http://www.w3.org/2000/09/xmldsig#sha1` - `http://www.w3.org/2001/04/xmlenc#sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#sha384` - `http://www.w3.org/2001/04/xmlenc#sha512`
        /// </summary>
        [Input("digestAlgorithm")]
        public Input<string>? DigestAlgorithm { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("encryptionKp")]
        public Input<string>? EncryptionKp { get; set; }

        [Input("enrollmentFlow")]
        public Input<string>? EnrollmentFlow { get; set; }

        /// <summary>
        /// Allowed values: - `identifier` - `name_link` - `name_deny`
        /// </summary>
        [Input("groupMatchingMode")]
        public Input<string>? GroupMatchingMode { get; set; }

        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// SAML Metadata
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent` - `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName` -
        /// `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`
        /// </summary>
        [Input("nameIdPolicy")]
        public Input<string>? NameIdPolicy { get; set; }

        /// <summary>
        /// Allowed values: - `all` - `any`
        /// </summary>
        [Input("policyEngineMode")]
        public Input<string>? PolicyEngineMode { get; set; }

        [Input("preAuthenticationFlow")]
        public Input<string>? PreAuthenticationFlow { get; set; }

        /// <summary>
        /// Allowed values: - `http://www.w3.org/2000/09/xmldsig#rsa-sha1` - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384` - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512` -
        /// `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1` - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256` -
        /// `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384` - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512` -
        /// `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        [Input("signingKp")]
        public Input<string>? SigningKp { get; set; }

        [Input("sloUrl")]
        public Input<string>? SloUrl { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("ssoUrl")]
        public Input<string>? SsoUrl { get; set; }

        [Input("temporaryUserDeleteAfter")]
        public Input<string>? TemporaryUserDeleteAfter { get; set; }

        /// <summary>
        /// Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        /// </summary>
        [Input("userMatchingMode")]
        public Input<string>? UserMatchingMode { get; set; }

        [Input("userPathTemplate")]
        public Input<string>? UserPathTemplate { get; set; }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public SourceSamlState()
        {
        }
        public static new SourceSamlState Empty => new SourceSamlState();
    }
}
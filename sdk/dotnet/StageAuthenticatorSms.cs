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
    [AuthentikResourceType("authentik:index/stageAuthenticatorSms:StageAuthenticatorSms")]
    public partial class StageAuthenticatorSms : global::Pulumi.CustomResource
    {
        [Output("accountSid")]
        public Output<string> AccountSid { get; private set; } = null!;

        [Output("auth")]
        public Output<string> Auth { get; private set; } = null!;

        [Output("authPassword")]
        public Output<string?> AuthPassword { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `basic` - `bearer`
        /// </summary>
        [Output("authType")]
        public Output<string?> AuthType { get; private set; } = null!;

        [Output("configureFlow")]
        public Output<string?> ConfigureFlow { get; private set; } = null!;

        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        [Output("fromNumber")]
        public Output<string> FromNumber { get; private set; } = null!;

        [Output("mapping")]
        public Output<string?> Mapping { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `twilio` - `generic`
        /// </summary>
        [Output("smsProvider")]
        public Output<string?> SmsProvider { get; private set; } = null!;

        [Output("verifyOnly")]
        public Output<bool?> VerifyOnly { get; private set; } = null!;


        /// <summary>
        /// Create a StageAuthenticatorSms resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StageAuthenticatorSms(string name, StageAuthenticatorSmsArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/stageAuthenticatorSms:StageAuthenticatorSms", name, args ?? new StageAuthenticatorSmsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StageAuthenticatorSms(string name, Input<string> id, StageAuthenticatorSmsState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/stageAuthenticatorSms:StageAuthenticatorSms", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OSMIT-GmbH/pulumi-authentik/releases/download/v${VERSION}",
                AdditionalSecretOutputs =
                {
                    "accountSid",
                    "auth",
                    "authPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StageAuthenticatorSms resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StageAuthenticatorSms Get(string name, Input<string> id, StageAuthenticatorSmsState? state = null, CustomResourceOptions? options = null)
        {
            return new StageAuthenticatorSms(name, id, state, options);
        }
    }

    public sealed class StageAuthenticatorSmsArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountSid", required: true)]
        private Input<string>? _accountSid;
        public Input<string>? AccountSid
        {
            get => _accountSid;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountSid = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("auth", required: true)]
        private Input<string>? _auth;
        public Input<string>? Auth
        {
            get => _auth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _auth = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("authPassword")]
        private Input<string>? _authPassword;
        public Input<string>? AuthPassword
        {
            get => _authPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Allowed values: - `basic` - `bearer`
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("configureFlow")]
        public Input<string>? ConfigureFlow { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("fromNumber", required: true)]
        public Input<string> FromNumber { get; set; } = null!;

        [Input("mapping")]
        public Input<string>? Mapping { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `twilio` - `generic`
        /// </summary>
        [Input("smsProvider")]
        public Input<string>? SmsProvider { get; set; }

        [Input("verifyOnly")]
        public Input<bool>? VerifyOnly { get; set; }

        public StageAuthenticatorSmsArgs()
        {
        }
        public static new StageAuthenticatorSmsArgs Empty => new StageAuthenticatorSmsArgs();
    }

    public sealed class StageAuthenticatorSmsState : global::Pulumi.ResourceArgs
    {
        [Input("accountSid")]
        private Input<string>? _accountSid;
        public Input<string>? AccountSid
        {
            get => _accountSid;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountSid = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("auth")]
        private Input<string>? _auth;
        public Input<string>? Auth
        {
            get => _auth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _auth = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("authPassword")]
        private Input<string>? _authPassword;
        public Input<string>? AuthPassword
        {
            get => _authPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Allowed values: - `basic` - `bearer`
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("configureFlow")]
        public Input<string>? ConfigureFlow { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("fromNumber")]
        public Input<string>? FromNumber { get; set; }

        [Input("mapping")]
        public Input<string>? Mapping { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `twilio` - `generic`
        /// </summary>
        [Input("smsProvider")]
        public Input<string>? SmsProvider { get; set; }

        [Input("verifyOnly")]
        public Input<bool>? VerifyOnly { get; set; }

        public StageAuthenticatorSmsState()
        {
        }
        public static new StageAuthenticatorSmsState Empty => new StageAuthenticatorSmsState();
    }
}
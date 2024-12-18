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
    [AuthentikResourceType("authentik:index/eventRule:EventRule")]
    public partial class EventRule : global::Pulumi.CustomResource
    {
        [Output("group")]
        public Output<string?> Group { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Allowed values: - `notice` - `warning` - `alert`
        /// </summary>
        [Output("severity")]
        public Output<string?> Severity { get; private set; } = null!;

        [Output("transports")]
        public Output<ImmutableArray<string>> Transports { get; private set; } = null!;

        [Output("webhookMapping")]
        public Output<string?> WebhookMapping { get; private set; } = null!;


        /// <summary>
        /// Create a EventRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventRule(string name, EventRuleArgs args, CustomResourceOptions? options = null)
            : base("authentik:index/eventRule:EventRule", name, args ?? new EventRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventRule(string name, Input<string> id, EventRuleState? state = null, CustomResourceOptions? options = null)
            : base("authentik:index/eventRule:EventRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventRule Get(string name, Input<string> id, EventRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new EventRule(name, id, state, options);
        }
    }

    public sealed class EventRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `notice` - `warning` - `alert`
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        [Input("transports", required: true)]
        private InputList<string>? _transports;
        public InputList<string> Transports
        {
            get => _transports ?? (_transports = new InputList<string>());
            set => _transports = value;
        }

        [Input("webhookMapping")]
        public Input<string>? WebhookMapping { get; set; }

        public EventRuleArgs()
        {
        }
        public static new EventRuleArgs Empty => new EventRuleArgs();
    }

    public sealed class EventRuleState : global::Pulumi.ResourceArgs
    {
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Allowed values: - `notice` - `warning` - `alert`
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        [Input("transports")]
        private InputList<string>? _transports;
        public InputList<string> Transports
        {
            get => _transports ?? (_transports = new InputList<string>());
            set => _transports = value;
        }

        [Input("webhookMapping")]
        public Input<string>? WebhookMapping { get; set; }

        public EventRuleState()
        {
        }
        public static new EventRuleState Empty => new EventRuleState();
    }
}

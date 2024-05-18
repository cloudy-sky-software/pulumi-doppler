// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.IntegrationsV3
{
    [DopplerNativeResourceType("doppler-native:integrations/v3:Integrations")]
    public partial class Integrations : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The authentication data for the integration
        /// </summary>
        [Output("data")]
        public Output<object?> Data { get; private set; } = null!;

        [Output("integration")]
        public Output<Outputs.IntegrationProperties?> Integration { get; private set; } = null!;

        /// <summary>
        /// The name of the integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The integration type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Integrations resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integrations(string name, IntegrationsArgs args, CustomResourceOptions? options = null)
            : base("doppler-native:integrations/v3:Integrations", name, args ?? new IntegrationsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integrations(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("doppler-native:integrations/v3:Integrations", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-doppler-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integrations resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integrations Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Integrations(name, id, options);
        }
    }

    public sealed class IntegrationsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication data for the integration
        /// </summary>
        [Input("data")]
        public Input<object>? Data { get; set; }

        /// <summary>
        /// The name of the integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The integration type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IntegrationsArgs()
        {
        }
        public static new IntegrationsArgs Empty => new IntegrationsArgs();
    }
}
// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ProjectsV3
{
    [DopplerNativeResourceType("doppler-native:projects/v3:ProjectsUpdate")]
    public partial class ProjectsUpdate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the project.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<Outputs.ProjectProperties> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectsUpdate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectsUpdate(string name, ProjectsUpdateArgs args, CustomResourceOptions? options = null)
            : base("doppler-native:projects/v3:ProjectsUpdate", name, args ?? new ProjectsUpdateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectsUpdate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("doppler-native:projects/v3:ProjectsUpdate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectsUpdate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectsUpdate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProjectsUpdate(name, id, options);
        }
    }

    public sealed class ProjectsUpdateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Unique identifier for the project object.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ProjectsUpdateArgs()
        {
            Description = "PROJECT_DESCRIPTION";
            Name = "PROJECT_NEW_NAME";
            Project = "PROJECT_NAME";
        }
        public static new ProjectsUpdateArgs Empty => new ProjectsUpdateArgs();
    }
}
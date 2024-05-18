// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ProjectsV3
{
    public static class GetProjectMember
    {
        public static Task<GetProjectMemberResult> InvokeAsync(GetProjectMemberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectMemberResult>("doppler-native:projects/v3:getProjectMember", args ?? new GetProjectMemberArgs(), options.WithDefaults());

        public static Output<GetProjectMemberResult> Invoke(GetProjectMemberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectMemberResult>("doppler-native:projects/v3:getProjectMember", args ?? new GetProjectMemberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectMemberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Member's slug
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetProjectMemberArgs()
        {
        }
        public static new GetProjectMemberArgs Empty => new GetProjectMemberArgs();
    }

    public sealed class GetProjectMemberInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Member's slug
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetProjectMemberInvokeArgs()
        {
        }
        public static new GetProjectMemberInvokeArgs Empty => new GetProjectMemberInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectMemberResult
    {
        public readonly Outputs.GetProjectMemberProperties Items;

        [OutputConstructor]
        private GetProjectMemberResult(Outputs.GetProjectMemberProperties items)
        {
            Items = items;
        }
    }
}

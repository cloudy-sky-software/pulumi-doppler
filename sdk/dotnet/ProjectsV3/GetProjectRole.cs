// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ProjectsV3
{
    public static class GetProjectRole
    {
        public static Task<GetProjectRoleResult> InvokeAsync(GetProjectRoleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectRoleResult>("doppler-native:projects/v3:getProjectRole", args ?? new GetProjectRoleArgs(), options.WithDefaults());

        public static Output<GetProjectRoleResult> Invoke(GetProjectRoleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectRoleResult>("doppler-native:projects/v3:getProjectRole", args ?? new GetProjectRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectRoleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The role's unique identifier
        /// </summary>
        [Input("role", required: true)]
        public string Role { get; set; } = null!;

        public GetProjectRoleArgs()
        {
        }
        public static new GetProjectRoleArgs Empty => new GetProjectRoleArgs();
    }

    public sealed class GetProjectRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The role's unique identifier
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public GetProjectRoleInvokeArgs()
        {
        }
        public static new GetProjectRoleInvokeArgs Empty => new GetProjectRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectRoleResult
    {
        public readonly Outputs.GetProjectRoleProperties Items;

        [OutputConstructor]
        private GetProjectRoleResult(Outputs.GetProjectRoleProperties items)
        {
            Items = items;
        }
    }
}

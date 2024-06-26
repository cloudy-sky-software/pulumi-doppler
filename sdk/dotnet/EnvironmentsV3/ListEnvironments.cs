// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.EnvironmentsV3
{
    public static class ListEnvironments
    {
        public static Task<ListEnvironmentsResult> InvokeAsync(ListEnvironmentsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListEnvironmentsResult>("doppler-native:environments/v3:listEnvironments", args ?? new ListEnvironmentsArgs(), options.WithDefaults());

        public static Output<ListEnvironmentsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListEnvironmentsResult>("doppler-native:environments/v3:listEnvironments", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListEnvironmentsArgs : global::Pulumi.InvokeArgs
    {
        public ListEnvironmentsArgs()
        {
        }
        public static new ListEnvironmentsArgs Empty => new ListEnvironmentsArgs();
    }


    [OutputType]
    public sealed class ListEnvironmentsResult
    {
        public readonly Outputs.ListEnvironmentsProperties Items;

        [OutputConstructor]
        private ListEnvironmentsResult(Outputs.ListEnvironmentsProperties items)
        {
            Items = items;
        }
    }
}

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3
{
    public static class ListConfigLogs
    {
        public static Task<ListConfigLogsResult> InvokeAsync(ListConfigLogsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListConfigLogsResult>("doppler-native:configs/v3:listConfigLogs", args ?? new ListConfigLogsArgs(), options.WithDefaults());

        public static Output<ListConfigLogsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListConfigLogsResult>("doppler-native:configs/v3:listConfigLogs", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListConfigLogsArgs : global::Pulumi.InvokeArgs
    {
        public ListConfigLogsArgs()
        {
        }
        public static new ListConfigLogsArgs Empty => new ListConfigLogsArgs();
    }


    [OutputType]
    public sealed class ListConfigLogsResult
    {
        public readonly Outputs.ListConfigLogsProperties Items;

        [OutputConstructor]
        private ListConfigLogsResult(Outputs.ListConfigLogsProperties items)
        {
            Items = items;
        }
    }
}

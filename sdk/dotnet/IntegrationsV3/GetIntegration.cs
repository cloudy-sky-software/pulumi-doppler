// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.IntegrationsV3
{
    public static class GetIntegration
    {
        public static Task<GetIntegrationResult> InvokeAsync(GetIntegrationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationResult>("doppler-native:integrations/v3:getIntegration", args ?? new GetIntegrationArgs(), options.WithDefaults());

        public static Output<GetIntegrationResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationResult>("doppler-native:integrations/v3:getIntegration", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetIntegrationArgs : global::Pulumi.InvokeArgs
    {
        public GetIntegrationArgs()
        {
        }
        public static new GetIntegrationArgs Empty => new GetIntegrationArgs();
    }


    [OutputType]
    public sealed class GetIntegrationResult
    {
        public readonly Outputs.GetIntegrationProperties Items;

        [OutputConstructor]
        private GetIntegrationResult(Outputs.GetIntegrationProperties items)
        {
            Items = items;
        }
    }
}

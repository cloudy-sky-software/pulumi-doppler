// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.IntegrationsV3.Outputs
{

    [OutputType]
    public sealed class GetIntegrationPropertiesIntegrationProperties
    {
        public readonly string? Name;
        public readonly string? Slug;
        public readonly string? Type;

        [OutputConstructor]
        private GetIntegrationPropertiesIntegrationProperties(
            string? name,

            string? slug,

            string? type)
        {
            Name = name;
            Slug = slug;
            Type = type;
        }
    }
}

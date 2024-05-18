// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3.Outputs
{

    [OutputType]
    public sealed class ListConfigsProperties
    {
        public readonly ImmutableArray<Outputs.ListConfigsPropertiesConfigsItemProperties> Configs;
        public readonly int? Page;

        [OutputConstructor]
        private ListConfigsProperties(
            ImmutableArray<Outputs.ListConfigsPropertiesConfigsItemProperties> configs,

            int? page)
        {
            Configs = configs;
            Page = page;
        }
    }
}

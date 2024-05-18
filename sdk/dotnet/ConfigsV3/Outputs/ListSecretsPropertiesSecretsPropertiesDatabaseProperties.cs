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
    public sealed class ListSecretsPropertiesSecretsPropertiesDatabaseProperties
    {
        public readonly string? Computed;
        public readonly string? ComputedVisibility;
        public readonly string? Note;
        public readonly string? Raw;
        public readonly string? RawVisibility;

        [OutputConstructor]
        private ListSecretsPropertiesSecretsPropertiesDatabaseProperties(
            string? computed,

            string? computedVisibility,

            string? note,

            string? raw,

            string? rawVisibility)
        {
            Computed = computed;
            ComputedVisibility = computedVisibility;
            Note = note;
            Raw = raw;
            RawVisibility = rawVisibility;
        }
    }
}

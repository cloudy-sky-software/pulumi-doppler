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
    public sealed class SecretsPropertiesStripeProperties
    {
        public readonly string? Computed;
        public readonly string? Note;
        public readonly string? Raw;

        [OutputConstructor]
        private SecretsPropertiesStripeProperties(
            string? computed,

            string? note,

            string? raw)
        {
            Computed = computed;
            Note = note;
            Raw = raw;
        }
    }
}
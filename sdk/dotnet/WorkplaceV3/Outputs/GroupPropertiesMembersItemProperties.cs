// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.WorkplaceV3.Outputs
{

    [OutputType]
    public sealed class GroupPropertiesMembersItemProperties
    {
        public readonly string? Slug;
        public readonly string? Type;

        [OutputConstructor]
        private GroupPropertiesMembersItemProperties(
            string? slug,

            string? type)
        {
            Slug = slug;
            Type = type;
        }
    }
}

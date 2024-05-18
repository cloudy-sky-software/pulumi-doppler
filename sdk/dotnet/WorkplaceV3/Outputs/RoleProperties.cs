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
    public sealed class RoleProperties
    {
        public readonly string? CreatedAt;
        public readonly string? Identifier;
        public readonly bool? IsCustomRole;
        public readonly bool? IsInlineRole;
        public readonly string? Name;
        public readonly ImmutableArray<string> Permissions;

        [OutputConstructor]
        private RoleProperties(
            string? createdAt,

            string? identifier,

            bool? isCustomRole,

            bool? isInlineRole,

            string? name,

            ImmutableArray<string> permissions)
        {
            CreatedAt = createdAt;
            Identifier = identifier;
            IsCustomRole = isCustomRole;
            IsInlineRole = isInlineRole;
            Name = name;
            Permissions = permissions;
        }
    }
}

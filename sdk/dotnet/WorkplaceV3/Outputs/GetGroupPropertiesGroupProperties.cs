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
    public sealed class GetGroupPropertiesGroupProperties
    {
        public readonly string? CreatedAt;
        public readonly Outputs.GetGroupPropertiesGroupPropertiesDefaultProjectRoleProperties? DefaultProjectRole;
        public readonly ImmutableArray<Outputs.GetGroupPropertiesGroupPropertiesMembersItemProperties> Members;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetGroupPropertiesGroupPropertiesProjectsItemProperties> Projects;
        public readonly string? Slug;

        [OutputConstructor]
        private GetGroupPropertiesGroupProperties(
            string? createdAt,

            Outputs.GetGroupPropertiesGroupPropertiesDefaultProjectRoleProperties? defaultProjectRole,

            ImmutableArray<Outputs.GetGroupPropertiesGroupPropertiesMembersItemProperties> members,

            string? name,

            ImmutableArray<Outputs.GetGroupPropertiesGroupPropertiesProjectsItemProperties> projects,

            string? slug)
        {
            CreatedAt = createdAt;
            DefaultProjectRole = defaultProjectRole;
            Members = members;
            Name = name;
            Projects = projects;
            Slug = slug;
        }
    }
}

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ProjectsV3.Outputs
{

    [OutputType]
    public sealed class GetProjectRoleProperties
    {
        public readonly Outputs.GetProjectRolePropertiesRoleProperties? Role;

        [OutputConstructor]
        private GetProjectRoleProperties(Outputs.GetProjectRolePropertiesRoleProperties? role)
        {
            Role = role;
        }
    }
}

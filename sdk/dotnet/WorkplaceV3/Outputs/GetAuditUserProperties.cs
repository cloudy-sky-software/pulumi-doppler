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
    public sealed class GetAuditUserProperties
    {
        public readonly bool? Success;
        public readonly Outputs.GetAuditUserPropertiesWorkplaceUserProperties? WorkplaceUser;

        [OutputConstructor]
        private GetAuditUserProperties(
            bool? success,

            Outputs.GetAuditUserPropertiesWorkplaceUserProperties? workplaceUser)
        {
            Success = success;
            WorkplaceUser = workplaceUser;
        }
    }
}
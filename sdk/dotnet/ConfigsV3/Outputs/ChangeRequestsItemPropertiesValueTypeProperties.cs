// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3.Outputs
{

    /// <summary>
    /// The default valueType (string) will result in no validations being applied.
    /// </summary>
    [OutputType]
    public sealed class ChangeRequestsItemPropertiesValueTypeProperties
    {
        public readonly Pulumi.DopplerNative.ConfigsV3.SecretsUpdateChangeRequestsItemPropertiesValueTypePropertiesType? Type;

        [OutputConstructor]
        private ChangeRequestsItemPropertiesValueTypeProperties(Pulumi.DopplerNative.ConfigsV3.SecretsUpdateChangeRequestsItemPropertiesValueTypePropertiesType? type)
        {
            Type = type;
        }
    }
}

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3.Inputs
{

    /// <summary>
    /// The default valueType (string) will result in no validations being applied.
    /// </summary>
    public sealed class ChangeRequestsItemPropertiesValueTypePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("type")]
        public Input<Pulumi.DopplerNative.ConfigsV3.SecretsUpdateChangeRequestsItemPropertiesValueTypePropertiesType>? Type { get; set; }

        public ChangeRequestsItemPropertiesValueTypePropertiesArgs()
        {
        }
        public static new ChangeRequestsItemPropertiesValueTypePropertiesArgs Empty => new ChangeRequestsItemPropertiesValueTypePropertiesArgs();
    }
}

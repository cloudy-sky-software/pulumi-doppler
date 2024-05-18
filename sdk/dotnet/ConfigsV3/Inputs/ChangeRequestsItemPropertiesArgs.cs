// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DopplerNative.ConfigsV3.Inputs
{

    public sealed class ChangeRequestsItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the secret.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The original name of the secret. Use `null` (an actual `null`, not the string `null`) or omit this parameter for new secrets. If it differs from `name` then a rename is inferred.
        /// </summary>
        [Input("originalName", required: true)]
        public Input<string> OriginalName { get; set; } = null!;

        /// <summary>
        /// The value you expect the secret to have before `name` is applied. If specified, the request will only be processed if the provided value matches what's found in Doppler.
        /// </summary>
        [Input("originalValue")]
        public Input<string>? OriginalValue { get; set; }

        /// <summary>
        /// The valueType you expect the secret to have before `valueType` is applied. If specified, the request will only be processed if the provided valueType matches what's found in Doppler.
        /// </summary>
        [Input("originalValueType")]
        public Input<Inputs.ChangeRequestsItemPropertiesOriginalValueTypePropertiesArgs>? OriginalValueType { get; set; }

        /// <summary>
        /// Must be set to either `masked`, `unmasked`, or `restricted`. The visibility you expect the secret to have before `visibility` is applied. If specified, the request will only be processed if the provided visibility matches what's found in Doppler.
        /// </summary>
        [Input("originalVisibility")]
        public Input<string>? OriginalVisibility { get; set; }

        /// <summary>
        /// Defaults to `false`. Can only be set to `true` if the config being updated is a branch config and there is a secret with the same name in the root config. In this case, the branch secret will inherit the value and visibility type from the root secret.
        /// </summary>
        [Input("shouldConverge")]
        public Input<bool>? ShouldConverge { get; set; }

        /// <summary>
        /// Defaults to `false`. If set to `true`, will delete the secret matching the `name` field.
        /// </summary>
        [Input("shouldDelete")]
        public Input<bool>? ShouldDelete { get; set; }

        /// <summary>
        /// Defaults to `false`. Can only be set to `true` if the config being updated is a branch config. If set to `true`, the provided secret will be set in both the branch config as well as the root config in that environment.
        /// </summary>
        [Input("shouldPromote")]
        public Input<bool>? ShouldPromote { get; set; }

        /// <summary>
        /// The value the secret should have. Use `null` (an actual `null`, not the string `null`) to leave the existing secret value unchanged.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The default valueType (string) will result in no validations being applied.
        /// </summary>
        [Input("valueType")]
        public Input<Inputs.ChangeRequestsItemPropertiesValueTypePropertiesArgs>? ValueType { get; set; }

        /// <summary>
        /// Must be set to either `masked`, `unmasked`, or `restricted`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ChangeRequestsItemPropertiesArgs()
        {
        }
        public static new ChangeRequestsItemPropertiesArgs Empty => new ChangeRequestsItemPropertiesArgs();
    }
}

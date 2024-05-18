// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getConfig(args?: GetConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:configs/v3:getConfig", {
    }, opts);
}

export interface GetConfigArgs {
}

export interface GetConfigResult {
    readonly items: outputs.configs.v3.GetConfigProperties;
}
export function getConfigOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigResult> {
    return pulumi.output(getConfig(opts))
}

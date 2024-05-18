// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getSecretsDownload(args?: GetSecretsDownloadArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretsDownloadResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:configs/v3:getSecretsDownload", {
    }, opts);
}

export interface GetSecretsDownloadArgs {
}

export interface GetSecretsDownloadResult {
    readonly items: outputs.configs.v3.GetSecretsDownloadProperties;
}
export function getSecretsDownloadOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretsDownloadResult> {
    return pulumi.output(getSecretsDownload(opts))
}

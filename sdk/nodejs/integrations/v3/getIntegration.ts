// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getIntegration(args?: GetIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:integrations/v3:getIntegration", {
    }, opts);
}

export interface GetIntegrationArgs {
}

export interface GetIntegrationResult {
    readonly items: outputs.integrations.v3.GetIntegrationProperties;
}
export function getIntegrationOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetIntegrationResult> {
    return pulumi.output(getIntegration(opts))
}

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listGroups(args?: ListGroupsArgs, opts?: pulumi.InvokeOptions): Promise<ListGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:listGroups", {
    }, opts);
}

export interface ListGroupsArgs {
}

export interface ListGroupsResult {
    readonly items: outputs.workplace.v3.ListGroupsProperties;
}
export function listGroupsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListGroupsResult> {
    return pulumi.output(listGroups(opts))
}

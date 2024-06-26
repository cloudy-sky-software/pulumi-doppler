// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listInvites(args?: ListInvitesArgs, opts?: pulumi.InvokeOptions): Promise<ListInvitesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:listInvites", {
    }, opts);
}

export interface ListInvitesArgs {
}

export interface ListInvitesResult {
    readonly items: outputs.workplace.v3.ListInvitesProperties;
}
export function listInvitesOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListInvitesResult> {
    return pulumi.output(listInvites(opts))
}

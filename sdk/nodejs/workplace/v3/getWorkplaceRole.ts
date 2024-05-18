// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getWorkplaceRole(args: GetWorkplaceRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkplaceRoleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("doppler-native:workplace/v3:getWorkplaceRole", {
        "role": args.role,
    }, opts);
}

export interface GetWorkplaceRoleArgs {
    /**
     * The role's unique identifier
     */
    role: string;
}

export interface GetWorkplaceRoleResult {
    readonly items: outputs.workplace.v3.GetWorkplaceRoleProperties;
}
export function getWorkplaceRoleOutput(args: GetWorkplaceRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkplaceRoleResult> {
    return pulumi.output(args).apply((a: any) => getWorkplaceRole(a, opts))
}

export interface GetWorkplaceRoleOutputArgs {
    /**
     * The role's unique identifier
     */
    role: pulumi.Input<string>;
}
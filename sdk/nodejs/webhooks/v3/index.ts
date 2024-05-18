// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetWebhookArgs, GetWebhookResult, GetWebhookOutputArgs } from "./getWebhook";
export const getWebhook: typeof import("./getWebhook").getWebhook = null as any;
export const getWebhookOutput: typeof import("./getWebhook").getWebhookOutput = null as any;
utilities.lazyLoad(exports, ["getWebhook","getWebhookOutput"], () => require("./getWebhook"));

export { WebhooksArgs } from "./webhooks";
export type Webhooks = import("./webhooks").Webhooks;
export const Webhooks: typeof import("./webhooks").Webhooks = null as any;
utilities.lazyLoad(exports, ["Webhooks"], () => require("./webhooks"));

export { WebhooksDisableArgs } from "./webhooksDisable";
export type WebhooksDisable = import("./webhooksDisable").WebhooksDisable;
export const WebhooksDisable: typeof import("./webhooksDisable").WebhooksDisable = null as any;
utilities.lazyLoad(exports, ["WebhooksDisable"], () => require("./webhooksDisable"));

export { WebhooksEnableArgs } from "./webhooksEnable";
export type WebhooksEnable = import("./webhooksEnable").WebhooksEnable;
export const WebhooksEnable: typeof import("./webhooksEnable").WebhooksEnable = null as any;
utilities.lazyLoad(exports, ["WebhooksEnable"], () => require("./webhooksEnable"));


// Export enums:
export * from "../../types/enums/webhooks/v3";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "doppler-native:webhooks/v3:Webhooks":
                return new Webhooks(name, <any>undefined, { urn })
            case "doppler-native:webhooks/v3:WebhooksDisable":
                return new WebhooksDisable(name, <any>undefined, { urn })
            case "doppler-native:webhooks/v3:WebhooksEnable":
                return new WebhooksEnable(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("doppler-native", "webhooks/v3", _module)

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

export class DynamicSecretsIssueLease extends pulumi.CustomResource {
    /**
     * Get an existing DynamicSecretsIssueLease resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DynamicSecretsIssueLease {
        return new DynamicSecretsIssueLease(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'doppler-native:configs/v3:DynamicSecretsIssueLease';

    /**
     * Returns true if the given object is an instance of DynamicSecretsIssueLease.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DynamicSecretsIssueLease {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DynamicSecretsIssueLease.__pulumiType;
    }

    /**
     * The config where the dynamic secret is located
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The name of the dynamic secret for which to issue this lease
     */
    public readonly dynamicSecret!: pulumi.Output<string>;
    public /*out*/ readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The project where the dynamic secret is located
     */
    public readonly project!: pulumi.Output<string>;
    public /*out*/ readonly success!: pulumi.Output<boolean | undefined>;
    /**
     * The number of seconds until this lease is automatically revoked
     */
    public readonly ttlSec!: pulumi.Output<number>;
    public /*out*/ readonly value!: pulumi.Output<any | undefined>;

    /**
     * Create a DynamicSecretsIssueLease resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DynamicSecretsIssueLeaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.dynamicSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dynamicSecret'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.ttlSec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttlSec'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["dynamicSecret"] = args ? args.dynamicSecret : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["ttlSec"] = args ? args.ttlSec : undefined;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["success"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        } else {
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["dynamicSecret"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["success"] = undefined /*out*/;
            resourceInputs["ttlSec"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DynamicSecretsIssueLease.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DynamicSecretsIssueLease resource.
 */
export interface DynamicSecretsIssueLeaseArgs {
    /**
     * The config where the dynamic secret is located
     */
    config: pulumi.Input<string>;
    /**
     * The name of the dynamic secret for which to issue this lease
     */
    dynamicSecret: pulumi.Input<string>;
    /**
     * The project where the dynamic secret is located
     */
    project: pulumi.Input<string>;
    /**
     * The number of seconds until this lease is automatically revoked
     */
    ttlSec: pulumi.Input<number>;
}

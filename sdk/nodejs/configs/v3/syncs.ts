// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Syncs extends pulumi.CustomResource {
    /**
     * Get an existing Syncs resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Syncs {
        return new Syncs(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'doppler-native:configs/v3:Syncs';

    /**
     * Returns true if the given object is an instance of Syncs.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Syncs {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Syncs.__pulumiType;
    }

    /**
     * Causes sync creation to wait for the initial sync to complete before returning.
     */
    public readonly awaitInitialSync!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration data for the sync
     */
    public readonly data!: pulumi.Output<any>;
    /**
     * An option indicating if and how Doppler should attempt to import secrets from the sync destination
     */
    public readonly importOption!: pulumi.Output<enums.configs.v3.SyncsImportOption | undefined>;
    /**
     * The integration slug which the sync will use
     */
    public readonly integration!: pulumi.Output<string>;
    public /*out*/ readonly sync!: pulumi.Output<outputs.configs.v3.SyncProperties | undefined>;

    /**
     * Create a Syncs resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncsArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.data === undefined) && !opts.urn) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.integration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integration'");
            }
            resourceInputs["awaitInitialSync"] = (args ? args.awaitInitialSync : undefined) ?? true;
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["importOption"] = (args ? args.importOption : undefined) ?? "none";
            resourceInputs["integration"] = args ? args.integration : undefined;
            resourceInputs["sync"] = undefined /*out*/;
        } else {
            resourceInputs["awaitInitialSync"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["importOption"] = undefined /*out*/;
            resourceInputs["integration"] = undefined /*out*/;
            resourceInputs["sync"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Syncs.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Syncs resource.
 */
export interface SyncsArgs {
    /**
     * Causes sync creation to wait for the initial sync to complete before returning.
     */
    awaitInitialSync?: pulumi.Input<boolean>;
    /**
     * Configuration data for the sync
     */
    data: any;
    /**
     * An option indicating if and how Doppler should attempt to import secrets from the sync destination
     */
    importOption?: pulumi.Input<enums.configs.v3.SyncsImportOption>;
    /**
     * The integration slug which the sync will use
     */
    integration: pulumi.Input<string>;
}

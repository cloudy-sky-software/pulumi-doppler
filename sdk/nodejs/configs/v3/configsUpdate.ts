// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class ConfigsUpdate extends pulumi.CustomResource {
    /**
     * Get an existing ConfigsUpdate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigsUpdate {
        return new ConfigsUpdate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'doppler-native:configs/v3:ConfigsUpdate';

    /**
     * Returns true if the given object is an instance of ConfigsUpdate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigsUpdate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigsUpdate.__pulumiType;
    }

    public readonly config!: pulumi.Output<outputs.configs.v3.ConfigProperties>;
    /**
     * The new name of config.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Unique identifier for the project object.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ConfigsUpdate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigsUpdateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["config"] = (args ? args.config : undefined) ?? "CONFIG_NAME";
            resourceInputs["name"] = (args ? args.name : undefined) ?? "CONFIG_NEW_NAME";
            resourceInputs["project"] = (args ? args.project : undefined) ?? "PROJECT_NAME";
        } else {
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigsUpdate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigsUpdate resource.
 */
export interface ConfigsUpdateArgs {
    /**
     * Name of the config object.
     */
    config: pulumi.Input<string>;
    /**
     * The new name of config.
     */
    name?: pulumi.Input<string>;
    /**
     * Unique identifier for the project object.
     */
    project: pulumi.Input<string>;
}

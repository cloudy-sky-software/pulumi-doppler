// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigsTrustedIp struct {
	pulumi.CustomResourceState

	Ip pulumi.StringOutput `pulumi:"ip"`
}

// NewConfigsTrustedIp registers a new resource with the given unique name, arguments, and options.
func NewConfigsTrustedIp(ctx *pulumi.Context,
	name string, args *ConfigsTrustedIpArgs, opts ...pulumi.ResourceOption) (*ConfigsTrustedIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigsTrustedIp
	err := ctx.RegisterResource("doppler-native:configs/v3:ConfigsTrustedIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigsTrustedIp gets an existing ConfigsTrustedIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigsTrustedIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigsTrustedIpState, opts ...pulumi.ResourceOption) (*ConfigsTrustedIp, error) {
	var resource ConfigsTrustedIp
	err := ctx.ReadResource("doppler-native:configs/v3:ConfigsTrustedIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigsTrustedIp resources.
type configsTrustedIpState struct {
}

type ConfigsTrustedIpState struct {
}

func (ConfigsTrustedIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*configsTrustedIpState)(nil)).Elem()
}

type configsTrustedIpArgs struct {
	// An IP address or CIDR range
	Ip string `pulumi:"ip"`
}

// The set of arguments for constructing a ConfigsTrustedIp resource.
type ConfigsTrustedIpArgs struct {
	// An IP address or CIDR range
	Ip pulumi.StringInput
}

func (ConfigsTrustedIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configsTrustedIpArgs)(nil)).Elem()
}

type ConfigsTrustedIpInput interface {
	pulumi.Input

	ToConfigsTrustedIpOutput() ConfigsTrustedIpOutput
	ToConfigsTrustedIpOutputWithContext(ctx context.Context) ConfigsTrustedIpOutput
}

func (*ConfigsTrustedIp) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigsTrustedIp)(nil)).Elem()
}

func (i *ConfigsTrustedIp) ToConfigsTrustedIpOutput() ConfigsTrustedIpOutput {
	return i.ToConfigsTrustedIpOutputWithContext(context.Background())
}

func (i *ConfigsTrustedIp) ToConfigsTrustedIpOutputWithContext(ctx context.Context) ConfigsTrustedIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigsTrustedIpOutput)
}

type ConfigsTrustedIpOutput struct{ *pulumi.OutputState }

func (ConfigsTrustedIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigsTrustedIp)(nil)).Elem()
}

func (o ConfigsTrustedIpOutput) ToConfigsTrustedIpOutput() ConfigsTrustedIpOutput {
	return o
}

func (o ConfigsTrustedIpOutput) ToConfigsTrustedIpOutputWithContext(ctx context.Context) ConfigsTrustedIpOutput {
	return o
}

func (o ConfigsTrustedIpOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigsTrustedIp) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigsTrustedIpInput)(nil)).Elem(), &ConfigsTrustedIp{})
	pulumi.RegisterOutputType(ConfigsTrustedIpOutput{})
}
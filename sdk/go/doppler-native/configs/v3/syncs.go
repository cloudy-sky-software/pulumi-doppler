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

type Syncs struct {
	pulumi.CustomResourceState

	// Causes sync creation to wait for the initial sync to complete before returning.
	AwaitInitialSync pulumi.BoolPtrOutput `pulumi:"awaitInitialSync"`
	// Configuration data for the sync
	Data pulumi.AnyOutput `pulumi:"data"`
	// An option indicating if and how Doppler should attempt to import secrets from the sync destination
	ImportOption SyncsImportOptionPtrOutput `pulumi:"importOption"`
	// The integration slug which the sync will use
	Integration pulumi.StringOutput     `pulumi:"integration"`
	Sync        SyncPropertiesPtrOutput `pulumi:"sync"`
}

// NewSyncs registers a new resource with the given unique name, arguments, and options.
func NewSyncs(ctx *pulumi.Context,
	name string, args *SyncsArgs, opts ...pulumi.ResourceOption) (*Syncs, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.Integration == nil {
		return nil, errors.New("invalid value for required argument 'Integration'")
	}
	if args.AwaitInitialSync == nil {
		args.AwaitInitialSync = pulumi.BoolPtr(true)
	}
	if args.ImportOption == nil {
		args.ImportOption = SyncsImportOption("none")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Syncs
	err := ctx.RegisterResource("doppler-native:configs/v3:Syncs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncs gets an existing Syncs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncsState, opts ...pulumi.ResourceOption) (*Syncs, error) {
	var resource Syncs
	err := ctx.ReadResource("doppler-native:configs/v3:Syncs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Syncs resources.
type syncsState struct {
}

type SyncsState struct {
}

func (SyncsState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncsState)(nil)).Elem()
}

type syncsArgs struct {
	// Causes sync creation to wait for the initial sync to complete before returning.
	AwaitInitialSync *bool `pulumi:"awaitInitialSync"`
	// Configuration data for the sync
	Data interface{} `pulumi:"data"`
	// An option indicating if and how Doppler should attempt to import secrets from the sync destination
	ImportOption *SyncsImportOption `pulumi:"importOption"`
	// The integration slug which the sync will use
	Integration string `pulumi:"integration"`
}

// The set of arguments for constructing a Syncs resource.
type SyncsArgs struct {
	// Causes sync creation to wait for the initial sync to complete before returning.
	AwaitInitialSync pulumi.BoolPtrInput
	// Configuration data for the sync
	Data pulumi.Input
	// An option indicating if and how Doppler should attempt to import secrets from the sync destination
	ImportOption SyncsImportOptionPtrInput
	// The integration slug which the sync will use
	Integration pulumi.StringInput
}

func (SyncsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncsArgs)(nil)).Elem()
}

type SyncsInput interface {
	pulumi.Input

	ToSyncsOutput() SyncsOutput
	ToSyncsOutputWithContext(ctx context.Context) SyncsOutput
}

func (*Syncs) ElementType() reflect.Type {
	return reflect.TypeOf((**Syncs)(nil)).Elem()
}

func (i *Syncs) ToSyncsOutput() SyncsOutput {
	return i.ToSyncsOutputWithContext(context.Background())
}

func (i *Syncs) ToSyncsOutputWithContext(ctx context.Context) SyncsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncsOutput)
}

type SyncsOutput struct{ *pulumi.OutputState }

func (SyncsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Syncs)(nil)).Elem()
}

func (o SyncsOutput) ToSyncsOutput() SyncsOutput {
	return o
}

func (o SyncsOutput) ToSyncsOutputWithContext(ctx context.Context) SyncsOutput {
	return o
}

// Causes sync creation to wait for the initial sync to complete before returning.
func (o SyncsOutput) AwaitInitialSync() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Syncs) pulumi.BoolPtrOutput { return v.AwaitInitialSync }).(pulumi.BoolPtrOutput)
}

// Configuration data for the sync
func (o SyncsOutput) Data() pulumi.AnyOutput {
	return o.ApplyT(func(v *Syncs) pulumi.AnyOutput { return v.Data }).(pulumi.AnyOutput)
}

// An option indicating if and how Doppler should attempt to import secrets from the sync destination
func (o SyncsOutput) ImportOption() SyncsImportOptionPtrOutput {
	return o.ApplyT(func(v *Syncs) SyncsImportOptionPtrOutput { return v.ImportOption }).(SyncsImportOptionPtrOutput)
}

// The integration slug which the sync will use
func (o SyncsOutput) Integration() pulumi.StringOutput {
	return o.ApplyT(func(v *Syncs) pulumi.StringOutput { return v.Integration }).(pulumi.StringOutput)
}

func (o SyncsOutput) Sync() SyncPropertiesPtrOutput {
	return o.ApplyT(func(v *Syncs) SyncPropertiesPtrOutput { return v.Sync }).(SyncPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncsInput)(nil)).Elem(), &Syncs{})
	pulumi.RegisterOutputType(SyncsOutput{})
}
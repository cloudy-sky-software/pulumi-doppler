// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Groups struct {
	pulumi.CustomResourceState

	// Identifier of the project role
	DefaultProjectRole pulumi.StringPtrOutput   `pulumi:"defaultProjectRole"`
	Group              GroupPropertiesPtrOutput `pulumi:"group"`
	Name               pulumi.StringOutput      `pulumi:"name"`
}

// NewGroups registers a new resource with the given unique name, arguments, and options.
func NewGroups(ctx *pulumi.Context,
	name string, args *GroupsArgs, opts ...pulumi.ResourceOption) (*Groups, error) {
	if args == nil {
		args = &GroupsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Groups
	err := ctx.RegisterResource("doppler-native:workplace/v3:Groups", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroups gets an existing Groups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroups(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupsState, opts ...pulumi.ResourceOption) (*Groups, error) {
	var resource Groups
	err := ctx.ReadResource("doppler-native:workplace/v3:Groups", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Groups resources.
type groupsState struct {
}

type GroupsState struct {
}

func (GroupsState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupsState)(nil)).Elem()
}

type groupsArgs struct {
	// Identifier of the project role
	DefaultProjectRole *string `pulumi:"defaultProjectRole"`
	Name               *string `pulumi:"name"`
}

// The set of arguments for constructing a Groups resource.
type GroupsArgs struct {
	// Identifier of the project role
	DefaultProjectRole pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
}

func (GroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupsArgs)(nil)).Elem()
}

type GroupsInput interface {
	pulumi.Input

	ToGroupsOutput() GroupsOutput
	ToGroupsOutputWithContext(ctx context.Context) GroupsOutput
}

func (*Groups) ElementType() reflect.Type {
	return reflect.TypeOf((**Groups)(nil)).Elem()
}

func (i *Groups) ToGroupsOutput() GroupsOutput {
	return i.ToGroupsOutputWithContext(context.Background())
}

func (i *Groups) ToGroupsOutputWithContext(ctx context.Context) GroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupsOutput)
}

type GroupsOutput struct{ *pulumi.OutputState }

func (GroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Groups)(nil)).Elem()
}

func (o GroupsOutput) ToGroupsOutput() GroupsOutput {
	return o
}

func (o GroupsOutput) ToGroupsOutputWithContext(ctx context.Context) GroupsOutput {
	return o
}

// Identifier of the project role
func (o GroupsOutput) DefaultProjectRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringPtrOutput { return v.DefaultProjectRole }).(pulumi.StringPtrOutput)
}

func (o GroupsOutput) Group() GroupPropertiesPtrOutput {
	return o.ApplyT(func(v *Groups) GroupPropertiesPtrOutput { return v.Group }).(GroupPropertiesPtrOutput)
}

func (o GroupsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupsInput)(nil)).Elem(), &Groups{})
	pulumi.RegisterOutputType(GroupsOutput{})
}

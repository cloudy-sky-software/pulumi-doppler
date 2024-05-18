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

type ProjectMembers struct {
	pulumi.CustomResourceState

	// Environment slugs to grant the member access to
	Environments pulumi.StringArrayOutput  `pulumi:"environments"`
	Member       MemberPropertiesPtrOutput `pulumi:"member"`
	// Identifier of the project role
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// Member's slug
	Slug pulumi.StringOutput      `pulumi:"slug"`
	Type ProjectMembersTypeOutput `pulumi:"type"`
}

// NewProjectMembers registers a new resource with the given unique name, arguments, and options.
func NewProjectMembers(ctx *pulumi.Context,
	name string, args *ProjectMembersArgs, opts ...pulumi.ResourceOption) (*ProjectMembers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectMembers
	err := ctx.RegisterResource("doppler-native:projects/v3:ProjectMembers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMembers gets an existing ProjectMembers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMembers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMembersState, opts ...pulumi.ResourceOption) (*ProjectMembers, error) {
	var resource ProjectMembers
	err := ctx.ReadResource("doppler-native:projects/v3:ProjectMembers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMembers resources.
type projectMembersState struct {
}

type ProjectMembersState struct {
}

func (ProjectMembersState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMembersState)(nil)).Elem()
}

type projectMembersArgs struct {
	// Environment slugs to grant the member access to
	Environments []string `pulumi:"environments"`
	// Identifier of the project role
	Role *string `pulumi:"role"`
	// Member's slug
	Slug string             `pulumi:"slug"`
	Type ProjectMembersType `pulumi:"type"`
}

// The set of arguments for constructing a ProjectMembers resource.
type ProjectMembersArgs struct {
	// Environment slugs to grant the member access to
	Environments pulumi.StringArrayInput
	// Identifier of the project role
	Role pulumi.StringPtrInput
	// Member's slug
	Slug pulumi.StringInput
	Type ProjectMembersTypeInput
}

func (ProjectMembersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMembersArgs)(nil)).Elem()
}

type ProjectMembersInput interface {
	pulumi.Input

	ToProjectMembersOutput() ProjectMembersOutput
	ToProjectMembersOutputWithContext(ctx context.Context) ProjectMembersOutput
}

func (*ProjectMembers) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMembers)(nil)).Elem()
}

func (i *ProjectMembers) ToProjectMembersOutput() ProjectMembersOutput {
	return i.ToProjectMembersOutputWithContext(context.Background())
}

func (i *ProjectMembers) ToProjectMembersOutputWithContext(ctx context.Context) ProjectMembersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMembersOutput)
}

type ProjectMembersOutput struct{ *pulumi.OutputState }

func (ProjectMembersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMembers)(nil)).Elem()
}

func (o ProjectMembersOutput) ToProjectMembersOutput() ProjectMembersOutput {
	return o
}

func (o ProjectMembersOutput) ToProjectMembersOutputWithContext(ctx context.Context) ProjectMembersOutput {
	return o
}

// Environment slugs to grant the member access to
func (o ProjectMembersOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectMembers) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o ProjectMembersOutput) Member() MemberPropertiesPtrOutput {
	return o.ApplyT(func(v *ProjectMembers) MemberPropertiesPtrOutput { return v.Member }).(MemberPropertiesPtrOutput)
}

// Identifier of the project role
func (o ProjectMembersOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMembers) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// Member's slug
func (o ProjectMembersOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMembers) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o ProjectMembersOutput) Type() ProjectMembersTypeOutput {
	return o.ApplyT(func(v *ProjectMembers) ProjectMembersTypeOutput { return v.Type }).(ProjectMembersTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMembersInput)(nil)).Elem(), &ProjectMembers{})
	pulumi.RegisterOutputType(ProjectMembersOutput{})
}
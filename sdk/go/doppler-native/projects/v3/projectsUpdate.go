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

type ProjectsUpdate struct {
	pulumi.CustomResourceState

	// Description of the project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the project.
	Name    pulumi.StringOutput     `pulumi:"name"`
	Project ProjectPropertiesOutput `pulumi:"project"`
}

// NewProjectsUpdate registers a new resource with the given unique name, arguments, and options.
func NewProjectsUpdate(ctx *pulumi.Context,
	name string, args *ProjectsUpdateArgs, opts ...pulumi.ResourceOption) (*ProjectsUpdate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		args.Description = pulumi.StringPtr("PROJECT_DESCRIPTION")
	}
	if args.Name == nil {
		args.Name = pulumi.StringPtr("PROJECT_NEW_NAME")
	}
	if args.Project == nil {
		args.Project = pulumi.String("PROJECT_NAME")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectsUpdate
	err := ctx.RegisterResource("doppler-native:projects/v3:ProjectsUpdate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectsUpdate gets an existing ProjectsUpdate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectsUpdate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectsUpdateState, opts ...pulumi.ResourceOption) (*ProjectsUpdate, error) {
	var resource ProjectsUpdate
	err := ctx.ReadResource("doppler-native:projects/v3:ProjectsUpdate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectsUpdate resources.
type projectsUpdateState struct {
}

type ProjectsUpdateState struct {
}

func (ProjectsUpdateState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsUpdateState)(nil)).Elem()
}

type projectsUpdateArgs struct {
	// Description of the project.
	Description *string `pulumi:"description"`
	// Name of the project.
	Name *string `pulumi:"name"`
	// Unique identifier for the project object.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectsUpdate resource.
type ProjectsUpdateArgs struct {
	// Description of the project.
	Description pulumi.StringPtrInput
	// Name of the project.
	Name pulumi.StringPtrInput
	// Unique identifier for the project object.
	Project pulumi.StringInput
}

func (ProjectsUpdateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsUpdateArgs)(nil)).Elem()
}

type ProjectsUpdateInput interface {
	pulumi.Input

	ToProjectsUpdateOutput() ProjectsUpdateOutput
	ToProjectsUpdateOutputWithContext(ctx context.Context) ProjectsUpdateOutput
}

func (*ProjectsUpdate) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectsUpdate)(nil)).Elem()
}

func (i *ProjectsUpdate) ToProjectsUpdateOutput() ProjectsUpdateOutput {
	return i.ToProjectsUpdateOutputWithContext(context.Background())
}

func (i *ProjectsUpdate) ToProjectsUpdateOutputWithContext(ctx context.Context) ProjectsUpdateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectsUpdateOutput)
}

type ProjectsUpdateOutput struct{ *pulumi.OutputState }

func (ProjectsUpdateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectsUpdate)(nil)).Elem()
}

func (o ProjectsUpdateOutput) ToProjectsUpdateOutput() ProjectsUpdateOutput {
	return o
}

func (o ProjectsUpdateOutput) ToProjectsUpdateOutputWithContext(ctx context.Context) ProjectsUpdateOutput {
	return o
}

// Description of the project.
func (o ProjectsUpdateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectsUpdate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o ProjectsUpdateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectsUpdate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProjectsUpdateOutput) Project() ProjectPropertiesOutput {
	return o.ApplyT(func(v *ProjectsUpdate) ProjectPropertiesOutput { return v.Project }).(ProjectPropertiesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectsUpdateInput)(nil)).Elem(), &ProjectsUpdate{})
	pulumi.RegisterOutputType(ProjectsUpdateOutput{})
}

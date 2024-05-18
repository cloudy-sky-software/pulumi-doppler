// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProjectRolesPermissions(ctx *pulumi.Context, args *ListProjectRolesPermissionsArgs, opts ...pulumi.InvokeOption) (*ListProjectRolesPermissionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListProjectRolesPermissionsResult
	err := ctx.Invoke("doppler-native:projects/v3:listProjectRolesPermissions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProjectRolesPermissionsArgs struct {
}

type ListProjectRolesPermissionsResult struct {
	Items ListProjectRolesPermissionsProperties `pulumi:"items"`
}

func ListProjectRolesPermissionsOutput(ctx *pulumi.Context, args ListProjectRolesPermissionsOutputArgs, opts ...pulumi.InvokeOption) ListProjectRolesPermissionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProjectRolesPermissionsResult, error) {
			args := v.(ListProjectRolesPermissionsArgs)
			r, err := ListProjectRolesPermissions(ctx, &args, opts...)
			var s ListProjectRolesPermissionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProjectRolesPermissionsResultOutput)
}

type ListProjectRolesPermissionsOutputArgs struct {
}

func (ListProjectRolesPermissionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectRolesPermissionsArgs)(nil)).Elem()
}

type ListProjectRolesPermissionsResultOutput struct{ *pulumi.OutputState }

func (ListProjectRolesPermissionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectRolesPermissionsResult)(nil)).Elem()
}

func (o ListProjectRolesPermissionsResultOutput) ToListProjectRolesPermissionsResultOutput() ListProjectRolesPermissionsResultOutput {
	return o
}

func (o ListProjectRolesPermissionsResultOutput) ToListProjectRolesPermissionsResultOutputWithContext(ctx context.Context) ListProjectRolesPermissionsResultOutput {
	return o
}

func (o ListProjectRolesPermissionsResultOutput) Items() ListProjectRolesPermissionsPropertiesOutput {
	return o.ApplyT(func(v ListProjectRolesPermissionsResult) ListProjectRolesPermissionsProperties { return v.Items }).(ListProjectRolesPermissionsPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProjectRolesPermissionsResultOutput{})
}
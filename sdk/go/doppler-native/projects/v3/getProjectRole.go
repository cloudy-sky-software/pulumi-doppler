// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProjectRole(ctx *pulumi.Context, args *GetProjectRoleArgs, opts ...pulumi.InvokeOption) (*GetProjectRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectRoleResult
	err := ctx.Invoke("doppler-native:projects/v3:getProjectRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetProjectRoleArgs struct {
	// The role's unique identifier
	Role string `pulumi:"role"`
}

type GetProjectRoleResult struct {
	Items GetProjectRoleProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetProjectRoleResult
func (val *GetProjectRoleResult) Defaults() *GetProjectRoleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetProjectRoleOutput(ctx *pulumi.Context, args GetProjectRoleOutputArgs, opts ...pulumi.InvokeOption) GetProjectRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectRoleResult, error) {
			args := v.(GetProjectRoleArgs)
			r, err := GetProjectRole(ctx, &args, opts...)
			var s GetProjectRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectRoleResultOutput)
}

type GetProjectRoleOutputArgs struct {
	// The role's unique identifier
	Role pulumi.StringInput `pulumi:"role"`
}

func (GetProjectRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectRoleArgs)(nil)).Elem()
}

type GetProjectRoleResultOutput struct{ *pulumi.OutputState }

func (GetProjectRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectRoleResult)(nil)).Elem()
}

func (o GetProjectRoleResultOutput) ToGetProjectRoleResultOutput() GetProjectRoleResultOutput {
	return o
}

func (o GetProjectRoleResultOutput) ToGetProjectRoleResultOutputWithContext(ctx context.Context) GetProjectRoleResultOutput {
	return o
}

func (o GetProjectRoleResultOutput) Items() GetProjectRolePropertiesOutput {
	return o.ApplyT(func(v GetProjectRoleResult) GetProjectRoleProperties { return v.Items }).(GetProjectRolePropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectRoleResultOutput{})
}
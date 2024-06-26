// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-doppler-native/sdk/go/doppler-native/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWorkplace(ctx *pulumi.Context, args *GetWorkplaceArgs, opts ...pulumi.InvokeOption) (*GetWorkplaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWorkplaceResult
	err := ctx.Invoke("doppler-native:workplace/v3:getWorkplace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWorkplaceArgs struct {
}

type GetWorkplaceResult struct {
	Items GetWorkplaceProperties `pulumi:"items"`
}

func GetWorkplaceOutput(ctx *pulumi.Context, args GetWorkplaceOutputArgs, opts ...pulumi.InvokeOption) GetWorkplaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWorkplaceResult, error) {
			args := v.(GetWorkplaceArgs)
			r, err := GetWorkplace(ctx, &args, opts...)
			var s GetWorkplaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWorkplaceResultOutput)
}

type GetWorkplaceOutputArgs struct {
}

func (GetWorkplaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkplaceArgs)(nil)).Elem()
}

type GetWorkplaceResultOutput struct{ *pulumi.OutputState }

func (GetWorkplaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkplaceResult)(nil)).Elem()
}

func (o GetWorkplaceResultOutput) ToGetWorkplaceResultOutput() GetWorkplaceResultOutput {
	return o
}

func (o GetWorkplaceResultOutput) ToGetWorkplaceResultOutputWithContext(ctx context.Context) GetWorkplaceResultOutput {
	return o
}

func (o GetWorkplaceResultOutput) Items() GetWorkplacePropertiesOutput {
	return o.ApplyT(func(v GetWorkplaceResult) GetWorkplaceProperties { return v.Items }).(GetWorkplacePropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWorkplaceResultOutput{})
}

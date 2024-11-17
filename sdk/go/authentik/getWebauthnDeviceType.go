// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWebauthnDeviceType(ctx *pulumi.Context, args *GetWebauthnDeviceTypeArgs, opts ...pulumi.InvokeOption) (*GetWebauthnDeviceTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWebauthnDeviceTypeResult
	err := ctx.Invoke("authentik:index/getWebauthnDeviceType:getWebauthnDeviceType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebauthnDeviceType.
type GetWebauthnDeviceTypeArgs struct {
	// Generated.
	Description *string `pulumi:"description"`
}

// A collection of values returned by getWebauthnDeviceType.
type GetWebauthnDeviceTypeResult struct {
	// Generated.
	Aaguid string `pulumi:"aaguid"`
	// Generated.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetWebauthnDeviceTypeOutput(ctx *pulumi.Context, args GetWebauthnDeviceTypeOutputArgs, opts ...pulumi.InvokeOption) GetWebauthnDeviceTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebauthnDeviceTypeResultOutput, error) {
			args := v.(GetWebauthnDeviceTypeArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetWebauthnDeviceTypeResult
			secret, err := ctx.InvokePackageRaw("authentik:index/getWebauthnDeviceType:getWebauthnDeviceType", args, &rv, "", opts...)
			if err != nil {
				return GetWebauthnDeviceTypeResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetWebauthnDeviceTypeResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetWebauthnDeviceTypeResultOutput), nil
			}
			return output, nil
		}).(GetWebauthnDeviceTypeResultOutput)
}

// A collection of arguments for invoking getWebauthnDeviceType.
type GetWebauthnDeviceTypeOutputArgs struct {
	// Generated.
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (GetWebauthnDeviceTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebauthnDeviceTypeArgs)(nil)).Elem()
}

// A collection of values returned by getWebauthnDeviceType.
type GetWebauthnDeviceTypeResultOutput struct{ *pulumi.OutputState }

func (GetWebauthnDeviceTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebauthnDeviceTypeResult)(nil)).Elem()
}

func (o GetWebauthnDeviceTypeResultOutput) ToGetWebauthnDeviceTypeResultOutput() GetWebauthnDeviceTypeResultOutput {
	return o
}

func (o GetWebauthnDeviceTypeResultOutput) ToGetWebauthnDeviceTypeResultOutputWithContext(ctx context.Context) GetWebauthnDeviceTypeResultOutput {
	return o
}

// Generated.
func (o GetWebauthnDeviceTypeResultOutput) Aaguid() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebauthnDeviceTypeResult) string { return v.Aaguid }).(pulumi.StringOutput)
}

// Generated.
func (o GetWebauthnDeviceTypeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebauthnDeviceTypeResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWebauthnDeviceTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebauthnDeviceTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebauthnDeviceTypeResultOutput{})
}

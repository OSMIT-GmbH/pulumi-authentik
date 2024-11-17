// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Source by name, slug or managed
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := authentik.GetSource(ctx, &authentik.GetSourceArgs{
//				Managed: pulumi.StringRef("goauthentik.io/sources/inbuilt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSource(ctx *pulumi.Context, args *GetSourceArgs, opts ...pulumi.InvokeOption) (*GetSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSourceResult
	err := ctx.Invoke("authentik:index/getSource:getSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSource.
type GetSourceArgs struct {
	// Generated.
	Managed *string `pulumi:"managed"`
	// Generated.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getSource.
type GetSourceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Generated.
	Managed string `pulumi:"managed"`
	// Generated.
	Name string `pulumi:"name"`
	// Generated.
	Slug string `pulumi:"slug"`
	// Generated.
	Uuid string `pulumi:"uuid"`
}

func GetSourceOutput(ctx *pulumi.Context, args GetSourceOutputArgs, opts ...pulumi.InvokeOption) GetSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSourceResultOutput, error) {
			args := v.(GetSourceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetSourceResult
			secret, err := ctx.InvokePackageRaw("authentik:index/getSource:getSource", args, &rv, "", opts...)
			if err != nil {
				return GetSourceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetSourceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetSourceResultOutput), nil
			}
			return output, nil
		}).(GetSourceResultOutput)
}

// A collection of arguments for invoking getSource.
type GetSourceOutputArgs struct {
	// Generated.
	Managed pulumi.StringPtrInput `pulumi:"managed"`
	// Generated.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (GetSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSourceArgs)(nil)).Elem()
}

// A collection of values returned by getSource.
type GetSourceResultOutput struct{ *pulumi.OutputState }

func (GetSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSourceResult)(nil)).Elem()
}

func (o GetSourceResultOutput) ToGetSourceResultOutput() GetSourceResultOutput {
	return o
}

func (o GetSourceResultOutput) ToGetSourceResultOutputWithContext(ctx context.Context) GetSourceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Generated.
func (o GetSourceResultOutput) Managed() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourceResult) string { return v.Managed }).(pulumi.StringOutput)
}

// Generated.
func (o GetSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Generated.
func (o GetSourceResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourceResult) string { return v.Slug }).(pulumi.StringOutput)
}

// Generated.
func (o GetSourceResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourceResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSourceResultOutput{})
}

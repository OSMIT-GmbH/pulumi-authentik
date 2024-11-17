// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get stages by name
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
//			// To get the ID of a stage by name
//			_, err := authentik.GetStage(ctx, &authentik.GetStageArgs{
//				Name: pulumi.StringRef("default-authentication-identification"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetStage(ctx *pulumi.Context, args *GetStageArgs, opts ...pulumi.InvokeOption) (*GetStageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStageResult
	err := ctx.Invoke("authentik:index/getStage:getStage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStage.
type GetStageArgs struct {
	// Generated.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getStage.
type GetStageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Generated.
	Name string `pulumi:"name"`
}

func GetStageOutput(ctx *pulumi.Context, args GetStageOutputArgs, opts ...pulumi.InvokeOption) GetStageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStageResultOutput, error) {
			args := v.(GetStageArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetStageResult
			secret, err := ctx.InvokePackageRaw("authentik:index/getStage:getStage", args, &rv, "", opts...)
			if err != nil {
				return GetStageResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetStageResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetStageResultOutput), nil
			}
			return output, nil
		}).(GetStageResultOutput)
}

// A collection of arguments for invoking getStage.
type GetStageOutputArgs struct {
	// Generated.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetStageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStageArgs)(nil)).Elem()
}

// A collection of values returned by getStage.
type GetStageResultOutput struct{ *pulumi.OutputState }

func (GetStageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStageResult)(nil)).Elem()
}

func (o GetStageResultOutput) ToGetStageResultOutput() GetStageResultOutput {
	return o
}

func (o GetStageResultOutput) ToGetStageResultOutputWithContext(ctx context.Context) GetStageResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetStageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Generated.
func (o GetStageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetStageResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStageResultOutput{})
}

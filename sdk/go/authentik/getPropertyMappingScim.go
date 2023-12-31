// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get SCIM Property mappings
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
//			_, err := authentik.LookupPropertyMappingScim(ctx, &authentik.LookupPropertyMappingScimArgs{
//				ManagedLists: []string{
//					"goauthentik.io/providers/scim/user",
//					"goauthentik.io/providers/scim/group",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPropertyMappingScim(ctx *pulumi.Context, args *LookupPropertyMappingScimArgs, opts ...pulumi.InvokeOption) (*LookupPropertyMappingScimResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPropertyMappingScimResult
	err := ctx.Invoke("authentik:index/getPropertyMappingScim:getPropertyMappingScim", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyMappingScim.
type LookupPropertyMappingScimArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     []string `pulumi:"ids"`
	Managed *string  `pulumi:"managed"`
	// Retrive multiple property mappings
	ManagedLists []string `pulumi:"managedLists"`
	Name         *string  `pulumi:"name"`
}

// A collection of values returned by getPropertyMappingScim.
type LookupPropertyMappingScimResult struct {
	// Generated.
	Expression string `pulumi:"expression"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of ids when `managedList` is set. Generated.
	Ids     []string `pulumi:"ids"`
	Managed *string  `pulumi:"managed"`
	// Retrive multiple property mappings
	ManagedLists []string `pulumi:"managedLists"`
	Name         *string  `pulumi:"name"`
}

func LookupPropertyMappingScimOutput(ctx *pulumi.Context, args LookupPropertyMappingScimOutputArgs, opts ...pulumi.InvokeOption) LookupPropertyMappingScimResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPropertyMappingScimResult, error) {
			args := v.(LookupPropertyMappingScimArgs)
			r, err := LookupPropertyMappingScim(ctx, &args, opts...)
			var s LookupPropertyMappingScimResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPropertyMappingScimResultOutput)
}

// A collection of arguments for invoking getPropertyMappingScim.
type LookupPropertyMappingScimOutputArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     pulumi.StringArrayInput `pulumi:"ids"`
	Managed pulumi.StringPtrInput   `pulumi:"managed"`
	// Retrive multiple property mappings
	ManagedLists pulumi.StringArrayInput `pulumi:"managedLists"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
}

func (LookupPropertyMappingScimOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingScimArgs)(nil)).Elem()
}

// A collection of values returned by getPropertyMappingScim.
type LookupPropertyMappingScimResultOutput struct{ *pulumi.OutputState }

func (LookupPropertyMappingScimResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingScimResult)(nil)).Elem()
}

func (o LookupPropertyMappingScimResultOutput) ToLookupPropertyMappingScimResultOutput() LookupPropertyMappingScimResultOutput {
	return o
}

func (o LookupPropertyMappingScimResultOutput) ToLookupPropertyMappingScimResultOutputWithContext(ctx context.Context) LookupPropertyMappingScimResultOutput {
	return o
}

func (o LookupPropertyMappingScimResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPropertyMappingScimResult] {
	return pulumix.Output[LookupPropertyMappingScimResult]{
		OutputState: o.OutputState,
	}
}

// Generated.
func (o LookupPropertyMappingScimResultOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingScimResult) string { return v.Expression }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPropertyMappingScimResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingScimResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of ids when `managedList` is set. Generated.
func (o LookupPropertyMappingScimResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingScimResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingScimResultOutput) Managed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingScimResult) *string { return v.Managed }).(pulumi.StringPtrOutput)
}

// Retrive multiple property mappings
func (o LookupPropertyMappingScimResultOutput) ManagedLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingScimResult) []string { return v.ManagedLists }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingScimResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingScimResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPropertyMappingScimResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get OAuth Provider Scope Property mappings
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := authentik.LookupPropertyMappingProviderScope(ctx, &authentik.LookupPropertyMappingProviderScopeArgs{
//				ManagedLists: []string{
//					"goauthentik.io/providers/oauth2/scope-email",
//					"goauthentik.io/providers/oauth2/scope-openid",
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
func LookupPropertyMappingProviderScope(ctx *pulumi.Context, args *LookupPropertyMappingProviderScopeArgs, opts ...pulumi.InvokeOption) (*LookupPropertyMappingProviderScopeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPropertyMappingProviderScopeResult
	err := ctx.Invoke("authentik:index/getPropertyMappingProviderScope:getPropertyMappingProviderScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyMappingProviderScope.
type LookupPropertyMappingProviderScopeArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     []string `pulumi:"ids"`
	Managed *string  `pulumi:"managed"`
	// Retrieve multiple property mappings
	ManagedLists []string `pulumi:"managedLists"`
	Name         *string  `pulumi:"name"`
	// Generated.
	ScopeName *string `pulumi:"scopeName"`
}

// A collection of values returned by getPropertyMappingProviderScope.
type LookupPropertyMappingProviderScopeResult struct {
	// Generated.
	Description string `pulumi:"description"`
	// Generated.
	Expression string `pulumi:"expression"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of ids when `managedList` is set. Generated.
	Ids     []string `pulumi:"ids"`
	Managed *string  `pulumi:"managed"`
	// Retrieve multiple property mappings
	ManagedLists []string `pulumi:"managedLists"`
	Name         *string  `pulumi:"name"`
	// Generated.
	ScopeName string `pulumi:"scopeName"`
}

func LookupPropertyMappingProviderScopeOutput(ctx *pulumi.Context, args LookupPropertyMappingProviderScopeOutputArgs, opts ...pulumi.InvokeOption) LookupPropertyMappingProviderScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPropertyMappingProviderScopeResultOutput, error) {
			args := v.(LookupPropertyMappingProviderScopeArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupPropertyMappingProviderScopeResult
			secret, err := ctx.InvokePackageRaw("authentik:index/getPropertyMappingProviderScope:getPropertyMappingProviderScope", args, &rv, "", opts...)
			if err != nil {
				return LookupPropertyMappingProviderScopeResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupPropertyMappingProviderScopeResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupPropertyMappingProviderScopeResultOutput), nil
			}
			return output, nil
		}).(LookupPropertyMappingProviderScopeResultOutput)
}

// A collection of arguments for invoking getPropertyMappingProviderScope.
type LookupPropertyMappingProviderScopeOutputArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     pulumi.StringArrayInput `pulumi:"ids"`
	Managed pulumi.StringPtrInput   `pulumi:"managed"`
	// Retrieve multiple property mappings
	ManagedLists pulumi.StringArrayInput `pulumi:"managedLists"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	// Generated.
	ScopeName pulumi.StringPtrInput `pulumi:"scopeName"`
}

func (LookupPropertyMappingProviderScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingProviderScopeArgs)(nil)).Elem()
}

// A collection of values returned by getPropertyMappingProviderScope.
type LookupPropertyMappingProviderScopeResultOutput struct{ *pulumi.OutputState }

func (LookupPropertyMappingProviderScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingProviderScopeResult)(nil)).Elem()
}

func (o LookupPropertyMappingProviderScopeResultOutput) ToLookupPropertyMappingProviderScopeResultOutput() LookupPropertyMappingProviderScopeResultOutput {
	return o
}

func (o LookupPropertyMappingProviderScopeResultOutput) ToLookupPropertyMappingProviderScopeResultOutputWithContext(ctx context.Context) LookupPropertyMappingProviderScopeResultOutput {
	return o
}

// Generated.
func (o LookupPropertyMappingProviderScopeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) string { return v.Description }).(pulumi.StringOutput)
}

// Generated.
func (o LookupPropertyMappingProviderScopeResultOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) string { return v.Expression }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPropertyMappingProviderScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of ids when `managedList` is set. Generated.
func (o LookupPropertyMappingProviderScopeResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingProviderScopeResultOutput) Managed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) *string { return v.Managed }).(pulumi.StringPtrOutput)
}

// Retrieve multiple property mappings
func (o LookupPropertyMappingProviderScopeResultOutput) ManagedLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) []string { return v.ManagedLists }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingProviderScopeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Generated.
func (o LookupPropertyMappingProviderScopeResultOutput) ScopeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingProviderScopeResult) string { return v.ScopeName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPropertyMappingProviderScopeResultOutput{})
}

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

// Get LDAP Property mappings
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
//			_, err := authentik.LookupPropertyMappingLdap(ctx, &authentik.LookupPropertyMappingLdapArgs{
//				ManagedLists: []string{
//					"goauthentik.io/sources/ldap/default-name",
//					"goauthentik.io/sources/ldap/default-mail",
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
func LookupPropertyMappingLdap(ctx *pulumi.Context, args *LookupPropertyMappingLdapArgs, opts ...pulumi.InvokeOption) (*LookupPropertyMappingLdapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPropertyMappingLdapResult
	err := ctx.Invoke("authentik:index/getPropertyMappingLdap:getPropertyMappingLdap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyMappingLdap.
type LookupPropertyMappingLdapArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     []string `pulumi:"ids"`
	Managed *string  `pulumi:"managed"`
	// Retrive multiple property mappings
	ManagedLists []string `pulumi:"managedLists"`
	Name         *string  `pulumi:"name"`
	// Generated.
	ObjectField *string `pulumi:"objectField"`
}

// A collection of values returned by getPropertyMappingLdap.
type LookupPropertyMappingLdapResult struct {
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
	// Generated.
	ObjectField string `pulumi:"objectField"`
}

func LookupPropertyMappingLdapOutput(ctx *pulumi.Context, args LookupPropertyMappingLdapOutputArgs, opts ...pulumi.InvokeOption) LookupPropertyMappingLdapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPropertyMappingLdapResult, error) {
			args := v.(LookupPropertyMappingLdapArgs)
			r, err := LookupPropertyMappingLdap(ctx, &args, opts...)
			var s LookupPropertyMappingLdapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPropertyMappingLdapResultOutput)
}

// A collection of arguments for invoking getPropertyMappingLdap.
type LookupPropertyMappingLdapOutputArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     pulumi.StringArrayInput `pulumi:"ids"`
	Managed pulumi.StringPtrInput   `pulumi:"managed"`
	// Retrive multiple property mappings
	ManagedLists pulumi.StringArrayInput `pulumi:"managedLists"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	// Generated.
	ObjectField pulumi.StringPtrInput `pulumi:"objectField"`
}

func (LookupPropertyMappingLdapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingLdapArgs)(nil)).Elem()
}

// A collection of values returned by getPropertyMappingLdap.
type LookupPropertyMappingLdapResultOutput struct{ *pulumi.OutputState }

func (LookupPropertyMappingLdapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingLdapResult)(nil)).Elem()
}

func (o LookupPropertyMappingLdapResultOutput) ToLookupPropertyMappingLdapResultOutput() LookupPropertyMappingLdapResultOutput {
	return o
}

func (o LookupPropertyMappingLdapResultOutput) ToLookupPropertyMappingLdapResultOutputWithContext(ctx context.Context) LookupPropertyMappingLdapResultOutput {
	return o
}

func (o LookupPropertyMappingLdapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPropertyMappingLdapResult] {
	return pulumix.Output[LookupPropertyMappingLdapResult]{
		OutputState: o.OutputState,
	}
}

// Generated.
func (o LookupPropertyMappingLdapResultOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) string { return v.Expression }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPropertyMappingLdapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of ids when `managedList` is set. Generated.
func (o LookupPropertyMappingLdapResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingLdapResultOutput) Managed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) *string { return v.Managed }).(pulumi.StringPtrOutput)
}

// Retrive multiple property mappings
func (o LookupPropertyMappingLdapResultOutput) ManagedLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) []string { return v.ManagedLists }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingLdapResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Generated.
func (o LookupPropertyMappingLdapResultOutput) ObjectField() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingLdapResult) string { return v.ObjectField }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPropertyMappingLdapResultOutput{})
}

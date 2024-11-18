// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get LDAP Source Property mappings
func LookupPropertyMappingSourceLdap(ctx *pulumi.Context, args *LookupPropertyMappingSourceLdapArgs, opts ...pulumi.InvokeOption) (*LookupPropertyMappingSourceLdapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPropertyMappingSourceLdapResult
	err := ctx.Invoke("authentik:index/getPropertyMappingSourceLdap:getPropertyMappingSourceLdap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyMappingSourceLdap.
type LookupPropertyMappingSourceLdapArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     []string `pulumi:"ids"`
	Managed *string  `pulumi:"managed"`
	// Retrieve multiple property mappings
	ManagedLists []string `pulumi:"managedLists"`
	Name         *string  `pulumi:"name"`
}

// A collection of values returned by getPropertyMappingSourceLdap.
type LookupPropertyMappingSourceLdapResult struct {
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
}

func LookupPropertyMappingSourceLdapOutput(ctx *pulumi.Context, args LookupPropertyMappingSourceLdapOutputArgs, opts ...pulumi.InvokeOption) LookupPropertyMappingSourceLdapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPropertyMappingSourceLdapResultOutput, error) {
			args := v.(LookupPropertyMappingSourceLdapArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupPropertyMappingSourceLdapResult
			secret, err := ctx.InvokePackageRaw("authentik:index/getPropertyMappingSourceLdap:getPropertyMappingSourceLdap", args, &rv, "", opts...)
			if err != nil {
				return LookupPropertyMappingSourceLdapResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupPropertyMappingSourceLdapResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupPropertyMappingSourceLdapResultOutput), nil
			}
			return output, nil
		}).(LookupPropertyMappingSourceLdapResultOutput)
}

// A collection of arguments for invoking getPropertyMappingSourceLdap.
type LookupPropertyMappingSourceLdapOutputArgs struct {
	// List of ids when `managedList` is set. Generated.
	Ids     pulumi.StringArrayInput `pulumi:"ids"`
	Managed pulumi.StringPtrInput   `pulumi:"managed"`
	// Retrieve multiple property mappings
	ManagedLists pulumi.StringArrayInput `pulumi:"managedLists"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
}

func (LookupPropertyMappingSourceLdapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingSourceLdapArgs)(nil)).Elem()
}

// A collection of values returned by getPropertyMappingSourceLdap.
type LookupPropertyMappingSourceLdapResultOutput struct{ *pulumi.OutputState }

func (LookupPropertyMappingSourceLdapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyMappingSourceLdapResult)(nil)).Elem()
}

func (o LookupPropertyMappingSourceLdapResultOutput) ToLookupPropertyMappingSourceLdapResultOutput() LookupPropertyMappingSourceLdapResultOutput {
	return o
}

func (o LookupPropertyMappingSourceLdapResultOutput) ToLookupPropertyMappingSourceLdapResultOutputWithContext(ctx context.Context) LookupPropertyMappingSourceLdapResultOutput {
	return o
}

// Generated.
func (o LookupPropertyMappingSourceLdapResultOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingSourceLdapResult) string { return v.Expression }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPropertyMappingSourceLdapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyMappingSourceLdapResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of ids when `managedList` is set. Generated.
func (o LookupPropertyMappingSourceLdapResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingSourceLdapResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingSourceLdapResultOutput) Managed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingSourceLdapResult) *string { return v.Managed }).(pulumi.StringPtrOutput)
}

// Retrieve multiple property mappings
func (o LookupPropertyMappingSourceLdapResultOutput) ManagedLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyMappingSourceLdapResult) []string { return v.ManagedLists }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyMappingSourceLdapResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPropertyMappingSourceLdapResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPropertyMappingSourceLdapResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			// Create an LDAP Provider
//			default_authentication_flow, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-authentication-flow"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			name, err := authentik.NewProviderLdap(ctx, "name", &authentik.ProviderLdapArgs{
//				Name:     pulumi.String("ldap-app"),
//				BaseDn:   pulumi.String("dc=ldap,dc=goauthentik,dc=io"),
//				BindFlow: pulumi.String(default_authentication_flow.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewApplication(ctx, "name", &authentik.ApplicationArgs{
//				Name:             pulumi.String("ldap-app"),
//				Slug:             pulumi.String("ldap-app"),
//				ProtocolProvider: name.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ProviderLdap struct {
	pulumi.CustomResourceState

	BaseDn   pulumi.StringOutput `pulumi:"baseDn"`
	BindFlow pulumi.StringOutput `pulumi:"bindFlow"`
	// Defaults to `direct`.
	BindMode    pulumi.StringPtrOutput `pulumi:"bindMode"`
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// Defaults to `4000`.
	GidStartNumber pulumi.IntPtrOutput `pulumi:"gidStartNumber"`
	// Defaults to `true`.
	MfaSupport pulumi.BoolPtrOutput `pulumi:"mfaSupport"`
	Name       pulumi.StringOutput  `pulumi:"name"`
	// Defaults to `direct`.
	SearchMode    pulumi.StringPtrOutput `pulumi:"searchMode"`
	TlsServerName pulumi.StringPtrOutput `pulumi:"tlsServerName"`
	// Defaults to `2000`.
	UidStartNumber pulumi.IntPtrOutput `pulumi:"uidStartNumber"`
	UnbindFlow     pulumi.StringOutput `pulumi:"unbindFlow"`
}

// NewProviderLdap registers a new resource with the given unique name, arguments, and options.
func NewProviderLdap(ctx *pulumi.Context,
	name string, args *ProviderLdapArgs, opts ...pulumi.ResourceOption) (*ProviderLdap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseDn == nil {
		return nil, errors.New("invalid value for required argument 'BaseDn'")
	}
	if args.BindFlow == nil {
		return nil, errors.New("invalid value for required argument 'BindFlow'")
	}
	if args.UnbindFlow == nil {
		return nil, errors.New("invalid value for required argument 'UnbindFlow'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderLdap
	err := ctx.RegisterResource("authentik:index/providerLdap:ProviderLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderLdap gets an existing ProviderLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderLdapState, opts ...pulumi.ResourceOption) (*ProviderLdap, error) {
	var resource ProviderLdap
	err := ctx.ReadResource("authentik:index/providerLdap:ProviderLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderLdap resources.
type providerLdapState struct {
	BaseDn   *string `pulumi:"baseDn"`
	BindFlow *string `pulumi:"bindFlow"`
	// Defaults to `direct`.
	BindMode    *string `pulumi:"bindMode"`
	Certificate *string `pulumi:"certificate"`
	// Defaults to `4000`.
	GidStartNumber *int `pulumi:"gidStartNumber"`
	// Defaults to `true`.
	MfaSupport *bool   `pulumi:"mfaSupport"`
	Name       *string `pulumi:"name"`
	// Defaults to `direct`.
	SearchMode    *string `pulumi:"searchMode"`
	TlsServerName *string `pulumi:"tlsServerName"`
	// Defaults to `2000`.
	UidStartNumber *int    `pulumi:"uidStartNumber"`
	UnbindFlow     *string `pulumi:"unbindFlow"`
}

type ProviderLdapState struct {
	BaseDn   pulumi.StringPtrInput
	BindFlow pulumi.StringPtrInput
	// Defaults to `direct`.
	BindMode    pulumi.StringPtrInput
	Certificate pulumi.StringPtrInput
	// Defaults to `4000`.
	GidStartNumber pulumi.IntPtrInput
	// Defaults to `true`.
	MfaSupport pulumi.BoolPtrInput
	Name       pulumi.StringPtrInput
	// Defaults to `direct`.
	SearchMode    pulumi.StringPtrInput
	TlsServerName pulumi.StringPtrInput
	// Defaults to `2000`.
	UidStartNumber pulumi.IntPtrInput
	UnbindFlow     pulumi.StringPtrInput
}

func (ProviderLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerLdapState)(nil)).Elem()
}

type providerLdapArgs struct {
	BaseDn   string `pulumi:"baseDn"`
	BindFlow string `pulumi:"bindFlow"`
	// Defaults to `direct`.
	BindMode    *string `pulumi:"bindMode"`
	Certificate *string `pulumi:"certificate"`
	// Defaults to `4000`.
	GidStartNumber *int `pulumi:"gidStartNumber"`
	// Defaults to `true`.
	MfaSupport *bool   `pulumi:"mfaSupport"`
	Name       *string `pulumi:"name"`
	// Defaults to `direct`.
	SearchMode    *string `pulumi:"searchMode"`
	TlsServerName *string `pulumi:"tlsServerName"`
	// Defaults to `2000`.
	UidStartNumber *int   `pulumi:"uidStartNumber"`
	UnbindFlow     string `pulumi:"unbindFlow"`
}

// The set of arguments for constructing a ProviderLdap resource.
type ProviderLdapArgs struct {
	BaseDn   pulumi.StringInput
	BindFlow pulumi.StringInput
	// Defaults to `direct`.
	BindMode    pulumi.StringPtrInput
	Certificate pulumi.StringPtrInput
	// Defaults to `4000`.
	GidStartNumber pulumi.IntPtrInput
	// Defaults to `true`.
	MfaSupport pulumi.BoolPtrInput
	Name       pulumi.StringPtrInput
	// Defaults to `direct`.
	SearchMode    pulumi.StringPtrInput
	TlsServerName pulumi.StringPtrInput
	// Defaults to `2000`.
	UidStartNumber pulumi.IntPtrInput
	UnbindFlow     pulumi.StringInput
}

func (ProviderLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerLdapArgs)(nil)).Elem()
}

type ProviderLdapInput interface {
	pulumi.Input

	ToProviderLdapOutput() ProviderLdapOutput
	ToProviderLdapOutputWithContext(ctx context.Context) ProviderLdapOutput
}

func (*ProviderLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderLdap)(nil)).Elem()
}

func (i *ProviderLdap) ToProviderLdapOutput() ProviderLdapOutput {
	return i.ToProviderLdapOutputWithContext(context.Background())
}

func (i *ProviderLdap) ToProviderLdapOutputWithContext(ctx context.Context) ProviderLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderLdapOutput)
}

// ProviderLdapArrayInput is an input type that accepts ProviderLdapArray and ProviderLdapArrayOutput values.
// You can construct a concrete instance of `ProviderLdapArrayInput` via:
//
//	ProviderLdapArray{ ProviderLdapArgs{...} }
type ProviderLdapArrayInput interface {
	pulumi.Input

	ToProviderLdapArrayOutput() ProviderLdapArrayOutput
	ToProviderLdapArrayOutputWithContext(context.Context) ProviderLdapArrayOutput
}

type ProviderLdapArray []ProviderLdapInput

func (ProviderLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderLdap)(nil)).Elem()
}

func (i ProviderLdapArray) ToProviderLdapArrayOutput() ProviderLdapArrayOutput {
	return i.ToProviderLdapArrayOutputWithContext(context.Background())
}

func (i ProviderLdapArray) ToProviderLdapArrayOutputWithContext(ctx context.Context) ProviderLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderLdapArrayOutput)
}

// ProviderLdapMapInput is an input type that accepts ProviderLdapMap and ProviderLdapMapOutput values.
// You can construct a concrete instance of `ProviderLdapMapInput` via:
//
//	ProviderLdapMap{ "key": ProviderLdapArgs{...} }
type ProviderLdapMapInput interface {
	pulumi.Input

	ToProviderLdapMapOutput() ProviderLdapMapOutput
	ToProviderLdapMapOutputWithContext(context.Context) ProviderLdapMapOutput
}

type ProviderLdapMap map[string]ProviderLdapInput

func (ProviderLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderLdap)(nil)).Elem()
}

func (i ProviderLdapMap) ToProviderLdapMapOutput() ProviderLdapMapOutput {
	return i.ToProviderLdapMapOutputWithContext(context.Background())
}

func (i ProviderLdapMap) ToProviderLdapMapOutputWithContext(ctx context.Context) ProviderLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderLdapMapOutput)
}

type ProviderLdapOutput struct{ *pulumi.OutputState }

func (ProviderLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderLdap)(nil)).Elem()
}

func (o ProviderLdapOutput) ToProviderLdapOutput() ProviderLdapOutput {
	return o
}

func (o ProviderLdapOutput) ToProviderLdapOutputWithContext(ctx context.Context) ProviderLdapOutput {
	return o
}

func (o ProviderLdapOutput) BaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringOutput { return v.BaseDn }).(pulumi.StringOutput)
}

func (o ProviderLdapOutput) BindFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringOutput { return v.BindFlow }).(pulumi.StringOutput)
}

// Defaults to `direct`.
func (o ProviderLdapOutput) BindMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringPtrOutput { return v.BindMode }).(pulumi.StringPtrOutput)
}

func (o ProviderLdapOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

// Defaults to `4000`.
func (o ProviderLdapOutput) GidStartNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.IntPtrOutput { return v.GidStartNumber }).(pulumi.IntPtrOutput)
}

// Defaults to `true`.
func (o ProviderLdapOutput) MfaSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.BoolPtrOutput { return v.MfaSupport }).(pulumi.BoolPtrOutput)
}

func (o ProviderLdapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defaults to `direct`.
func (o ProviderLdapOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringPtrOutput { return v.SearchMode }).(pulumi.StringPtrOutput)
}

func (o ProviderLdapOutput) TlsServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringPtrOutput { return v.TlsServerName }).(pulumi.StringPtrOutput)
}

// Defaults to `2000`.
func (o ProviderLdapOutput) UidStartNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.IntPtrOutput { return v.UidStartNumber }).(pulumi.IntPtrOutput)
}

func (o ProviderLdapOutput) UnbindFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderLdap) pulumi.StringOutput { return v.UnbindFlow }).(pulumi.StringOutput)
}

type ProviderLdapArrayOutput struct{ *pulumi.OutputState }

func (ProviderLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderLdap)(nil)).Elem()
}

func (o ProviderLdapArrayOutput) ToProviderLdapArrayOutput() ProviderLdapArrayOutput {
	return o
}

func (o ProviderLdapArrayOutput) ToProviderLdapArrayOutputWithContext(ctx context.Context) ProviderLdapArrayOutput {
	return o
}

func (o ProviderLdapArrayOutput) Index(i pulumi.IntInput) ProviderLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderLdap {
		return vs[0].([]*ProviderLdap)[vs[1].(int)]
	}).(ProviderLdapOutput)
}

type ProviderLdapMapOutput struct{ *pulumi.OutputState }

func (ProviderLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderLdap)(nil)).Elem()
}

func (o ProviderLdapMapOutput) ToProviderLdapMapOutput() ProviderLdapMapOutput {
	return o
}

func (o ProviderLdapMapOutput) ToProviderLdapMapOutputWithContext(ctx context.Context) ProviderLdapMapOutput {
	return o
}

func (o ProviderLdapMapOutput) MapIndex(k pulumi.StringInput) ProviderLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderLdap {
		return vs[0].(map[string]*ProviderLdap)[vs[1].(string)]
	}).(ProviderLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderLdapInput)(nil)).Elem(), &ProviderLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderLdapArrayInput)(nil)).Elem(), ProviderLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderLdapMapInput)(nil)).Elem(), ProviderLdapMap{})
	pulumi.RegisterOutputType(ProviderLdapOutput{})
	pulumi.RegisterOutputType(ProviderLdapArrayOutput{})
	pulumi.RegisterOutputType(ProviderLdapMapOutput{})
}

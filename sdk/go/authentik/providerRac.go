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

type ProviderRac struct {
	pulumi.CustomResourceState

	AuthenticationFlow pulumi.StringPtrOutput `pulumi:"authenticationFlow"`
	AuthorizationFlow  pulumi.StringOutput    `pulumi:"authorizationFlow"`
	// Defaults to `seconds=0`.
	ConnectionExpiry pulumi.StringPtrOutput   `pulumi:"connectionExpiry"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	PropertyMappings pulumi.StringArrayOutput `pulumi:"propertyMappings"`
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Settings pulumi.StringPtrOutput `pulumi:"settings"`
}

// NewProviderRac registers a new resource with the given unique name, arguments, and options.
func NewProviderRac(ctx *pulumi.Context,
	name string, args *ProviderRacArgs, opts ...pulumi.ResourceOption) (*ProviderRac, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationFlow == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationFlow'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderRac
	err := ctx.RegisterResource("authentik:index/providerRac:ProviderRac", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderRac gets an existing ProviderRac resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderRac(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderRacState, opts ...pulumi.ResourceOption) (*ProviderRac, error) {
	var resource ProviderRac
	err := ctx.ReadResource("authentik:index/providerRac:ProviderRac", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderRac resources.
type providerRacState struct {
	AuthenticationFlow *string `pulumi:"authenticationFlow"`
	AuthorizationFlow  *string `pulumi:"authorizationFlow"`
	// Defaults to `seconds=0`.
	ConnectionExpiry *string  `pulumi:"connectionExpiry"`
	Name             *string  `pulumi:"name"`
	PropertyMappings []string `pulumi:"propertyMappings"`
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Settings *string `pulumi:"settings"`
}

type ProviderRacState struct {
	AuthenticationFlow pulumi.StringPtrInput
	AuthorizationFlow  pulumi.StringPtrInput
	// Defaults to `seconds=0`.
	ConnectionExpiry pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	PropertyMappings pulumi.StringArrayInput
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Settings pulumi.StringPtrInput
}

func (ProviderRacState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRacState)(nil)).Elem()
}

type providerRacArgs struct {
	AuthenticationFlow *string `pulumi:"authenticationFlow"`
	AuthorizationFlow  string  `pulumi:"authorizationFlow"`
	// Defaults to `seconds=0`.
	ConnectionExpiry *string  `pulumi:"connectionExpiry"`
	Name             *string  `pulumi:"name"`
	PropertyMappings []string `pulumi:"propertyMappings"`
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Settings *string `pulumi:"settings"`
}

// The set of arguments for constructing a ProviderRac resource.
type ProviderRacArgs struct {
	AuthenticationFlow pulumi.StringPtrInput
	AuthorizationFlow  pulumi.StringInput
	// Defaults to `seconds=0`.
	ConnectionExpiry pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	PropertyMappings pulumi.StringArrayInput
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Settings pulumi.StringPtrInput
}

func (ProviderRacArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRacArgs)(nil)).Elem()
}

type ProviderRacInput interface {
	pulumi.Input

	ToProviderRacOutput() ProviderRacOutput
	ToProviderRacOutputWithContext(ctx context.Context) ProviderRacOutput
}

func (*ProviderRac) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderRac)(nil)).Elem()
}

func (i *ProviderRac) ToProviderRacOutput() ProviderRacOutput {
	return i.ToProviderRacOutputWithContext(context.Background())
}

func (i *ProviderRac) ToProviderRacOutputWithContext(ctx context.Context) ProviderRacOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRacOutput)
}

// ProviderRacArrayInput is an input type that accepts ProviderRacArray and ProviderRacArrayOutput values.
// You can construct a concrete instance of `ProviderRacArrayInput` via:
//
//	ProviderRacArray{ ProviderRacArgs{...} }
type ProviderRacArrayInput interface {
	pulumi.Input

	ToProviderRacArrayOutput() ProviderRacArrayOutput
	ToProviderRacArrayOutputWithContext(context.Context) ProviderRacArrayOutput
}

type ProviderRacArray []ProviderRacInput

func (ProviderRacArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderRac)(nil)).Elem()
}

func (i ProviderRacArray) ToProviderRacArrayOutput() ProviderRacArrayOutput {
	return i.ToProviderRacArrayOutputWithContext(context.Background())
}

func (i ProviderRacArray) ToProviderRacArrayOutputWithContext(ctx context.Context) ProviderRacArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRacArrayOutput)
}

// ProviderRacMapInput is an input type that accepts ProviderRacMap and ProviderRacMapOutput values.
// You can construct a concrete instance of `ProviderRacMapInput` via:
//
//	ProviderRacMap{ "key": ProviderRacArgs{...} }
type ProviderRacMapInput interface {
	pulumi.Input

	ToProviderRacMapOutput() ProviderRacMapOutput
	ToProviderRacMapOutputWithContext(context.Context) ProviderRacMapOutput
}

type ProviderRacMap map[string]ProviderRacInput

func (ProviderRacMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderRac)(nil)).Elem()
}

func (i ProviderRacMap) ToProviderRacMapOutput() ProviderRacMapOutput {
	return i.ToProviderRacMapOutputWithContext(context.Background())
}

func (i ProviderRacMap) ToProviderRacMapOutputWithContext(ctx context.Context) ProviderRacMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRacMapOutput)
}

type ProviderRacOutput struct{ *pulumi.OutputState }

func (ProviderRacOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderRac)(nil)).Elem()
}

func (o ProviderRacOutput) ToProviderRacOutput() ProviderRacOutput {
	return o
}

func (o ProviderRacOutput) ToProviderRacOutputWithContext(ctx context.Context) ProviderRacOutput {
	return o
}

func (o ProviderRacOutput) AuthenticationFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderRac) pulumi.StringPtrOutput { return v.AuthenticationFlow }).(pulumi.StringPtrOutput)
}

func (o ProviderRacOutput) AuthorizationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRac) pulumi.StringOutput { return v.AuthorizationFlow }).(pulumi.StringOutput)
}

// Defaults to `seconds=0`.
func (o ProviderRacOutput) ConnectionExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderRac) pulumi.StringPtrOutput { return v.ConnectionExpiry }).(pulumi.StringPtrOutput)
}

func (o ProviderRacOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRac) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderRacOutput) PropertyMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderRac) pulumi.StringArrayOutput { return v.PropertyMappings }).(pulumi.StringArrayOutput)
}

// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
func (o ProviderRacOutput) Settings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderRac) pulumi.StringPtrOutput { return v.Settings }).(pulumi.StringPtrOutput)
}

type ProviderRacArrayOutput struct{ *pulumi.OutputState }

func (ProviderRacArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderRac)(nil)).Elem()
}

func (o ProviderRacArrayOutput) ToProviderRacArrayOutput() ProviderRacArrayOutput {
	return o
}

func (o ProviderRacArrayOutput) ToProviderRacArrayOutputWithContext(ctx context.Context) ProviderRacArrayOutput {
	return o
}

func (o ProviderRacArrayOutput) Index(i pulumi.IntInput) ProviderRacOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderRac {
		return vs[0].([]*ProviderRac)[vs[1].(int)]
	}).(ProviderRacOutput)
}

type ProviderRacMapOutput struct{ *pulumi.OutputState }

func (ProviderRacMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderRac)(nil)).Elem()
}

func (o ProviderRacMapOutput) ToProviderRacMapOutput() ProviderRacMapOutput {
	return o
}

func (o ProviderRacMapOutput) ToProviderRacMapOutputWithContext(ctx context.Context) ProviderRacMapOutput {
	return o
}

func (o ProviderRacMapOutput) MapIndex(k pulumi.StringInput) ProviderRacOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderRac {
		return vs[0].(map[string]*ProviderRac)[vs[1].(string)]
	}).(ProviderRacOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRacInput)(nil)).Elem(), &ProviderRac{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRacArrayInput)(nil)).Elem(), ProviderRacArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRacMapInput)(nil)).Elem(), ProviderRacMap{})
	pulumi.RegisterOutputType(ProviderRacOutput{})
	pulumi.RegisterOutputType(ProviderRacArrayOutput{})
	pulumi.RegisterOutputType(ProviderRacMapOutput{})
}

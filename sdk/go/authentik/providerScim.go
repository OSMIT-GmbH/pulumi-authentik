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
type ProviderScim struct {
	pulumi.CustomResourceState

	ExcludeUsersServiceAccount pulumi.BoolPtrOutput     `pulumi:"excludeUsersServiceAccount"`
	FilterGroup                pulumi.StringPtrOutput   `pulumi:"filterGroup"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	PropertyMappings           pulumi.StringArrayOutput `pulumi:"propertyMappings"`
	PropertyMappingsGroups     pulumi.StringArrayOutput `pulumi:"propertyMappingsGroups"`
	Token                      pulumi.StringOutput      `pulumi:"token"`
	Url                        pulumi.StringOutput      `pulumi:"url"`
}

// NewProviderScim registers a new resource with the given unique name, arguments, and options.
func NewProviderScim(ctx *pulumi.Context,
	name string, args *ProviderScimArgs, opts ...pulumi.ResourceOption) (*ProviderScim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderScim
	err := ctx.RegisterResource("authentik:index/providerScim:ProviderScim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderScim gets an existing ProviderScim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderScim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderScimState, opts ...pulumi.ResourceOption) (*ProviderScim, error) {
	var resource ProviderScim
	err := ctx.ReadResource("authentik:index/providerScim:ProviderScim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderScim resources.
type providerScimState struct {
	ExcludeUsersServiceAccount *bool    `pulumi:"excludeUsersServiceAccount"`
	FilterGroup                *string  `pulumi:"filterGroup"`
	Name                       *string  `pulumi:"name"`
	PropertyMappings           []string `pulumi:"propertyMappings"`
	PropertyMappingsGroups     []string `pulumi:"propertyMappingsGroups"`
	Token                      *string  `pulumi:"token"`
	Url                        *string  `pulumi:"url"`
}

type ProviderScimState struct {
	ExcludeUsersServiceAccount pulumi.BoolPtrInput
	FilterGroup                pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	PropertyMappings           pulumi.StringArrayInput
	PropertyMappingsGroups     pulumi.StringArrayInput
	Token                      pulumi.StringPtrInput
	Url                        pulumi.StringPtrInput
}

func (ProviderScimState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerScimState)(nil)).Elem()
}

type providerScimArgs struct {
	ExcludeUsersServiceAccount *bool    `pulumi:"excludeUsersServiceAccount"`
	FilterGroup                *string  `pulumi:"filterGroup"`
	Name                       *string  `pulumi:"name"`
	PropertyMappings           []string `pulumi:"propertyMappings"`
	PropertyMappingsGroups     []string `pulumi:"propertyMappingsGroups"`
	Token                      string   `pulumi:"token"`
	Url                        string   `pulumi:"url"`
}

// The set of arguments for constructing a ProviderScim resource.
type ProviderScimArgs struct {
	ExcludeUsersServiceAccount pulumi.BoolPtrInput
	FilterGroup                pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	PropertyMappings           pulumi.StringArrayInput
	PropertyMappingsGroups     pulumi.StringArrayInput
	Token                      pulumi.StringInput
	Url                        pulumi.StringInput
}

func (ProviderScimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerScimArgs)(nil)).Elem()
}

type ProviderScimInput interface {
	pulumi.Input

	ToProviderScimOutput() ProviderScimOutput
	ToProviderScimOutputWithContext(ctx context.Context) ProviderScimOutput
}

func (*ProviderScim) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderScim)(nil)).Elem()
}

func (i *ProviderScim) ToProviderScimOutput() ProviderScimOutput {
	return i.ToProviderScimOutputWithContext(context.Background())
}

func (i *ProviderScim) ToProviderScimOutputWithContext(ctx context.Context) ProviderScimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderScimOutput)
}

// ProviderScimArrayInput is an input type that accepts ProviderScimArray and ProviderScimArrayOutput values.
// You can construct a concrete instance of `ProviderScimArrayInput` via:
//
//	ProviderScimArray{ ProviderScimArgs{...} }
type ProviderScimArrayInput interface {
	pulumi.Input

	ToProviderScimArrayOutput() ProviderScimArrayOutput
	ToProviderScimArrayOutputWithContext(context.Context) ProviderScimArrayOutput
}

type ProviderScimArray []ProviderScimInput

func (ProviderScimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderScim)(nil)).Elem()
}

func (i ProviderScimArray) ToProviderScimArrayOutput() ProviderScimArrayOutput {
	return i.ToProviderScimArrayOutputWithContext(context.Background())
}

func (i ProviderScimArray) ToProviderScimArrayOutputWithContext(ctx context.Context) ProviderScimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderScimArrayOutput)
}

// ProviderScimMapInput is an input type that accepts ProviderScimMap and ProviderScimMapOutput values.
// You can construct a concrete instance of `ProviderScimMapInput` via:
//
//	ProviderScimMap{ "key": ProviderScimArgs{...} }
type ProviderScimMapInput interface {
	pulumi.Input

	ToProviderScimMapOutput() ProviderScimMapOutput
	ToProviderScimMapOutputWithContext(context.Context) ProviderScimMapOutput
}

type ProviderScimMap map[string]ProviderScimInput

func (ProviderScimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderScim)(nil)).Elem()
}

func (i ProviderScimMap) ToProviderScimMapOutput() ProviderScimMapOutput {
	return i.ToProviderScimMapOutputWithContext(context.Background())
}

func (i ProviderScimMap) ToProviderScimMapOutputWithContext(ctx context.Context) ProviderScimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderScimMapOutput)
}

type ProviderScimOutput struct{ *pulumi.OutputState }

func (ProviderScimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderScim)(nil)).Elem()
}

func (o ProviderScimOutput) ToProviderScimOutput() ProviderScimOutput {
	return o
}

func (o ProviderScimOutput) ToProviderScimOutputWithContext(ctx context.Context) ProviderScimOutput {
	return o
}

func (o ProviderScimOutput) ExcludeUsersServiceAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.BoolPtrOutput { return v.ExcludeUsersServiceAccount }).(pulumi.BoolPtrOutput)
}

func (o ProviderScimOutput) FilterGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.StringPtrOutput { return v.FilterGroup }).(pulumi.StringPtrOutput)
}

func (o ProviderScimOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderScimOutput) PropertyMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.StringArrayOutput { return v.PropertyMappings }).(pulumi.StringArrayOutput)
}

func (o ProviderScimOutput) PropertyMappingsGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.StringArrayOutput { return v.PropertyMappingsGroups }).(pulumi.StringArrayOutput)
}

func (o ProviderScimOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ProviderScimOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderScim) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ProviderScimArrayOutput struct{ *pulumi.OutputState }

func (ProviderScimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderScim)(nil)).Elem()
}

func (o ProviderScimArrayOutput) ToProviderScimArrayOutput() ProviderScimArrayOutput {
	return o
}

func (o ProviderScimArrayOutput) ToProviderScimArrayOutputWithContext(ctx context.Context) ProviderScimArrayOutput {
	return o
}

func (o ProviderScimArrayOutput) Index(i pulumi.IntInput) ProviderScimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderScim {
		return vs[0].([]*ProviderScim)[vs[1].(int)]
	}).(ProviderScimOutput)
}

type ProviderScimMapOutput struct{ *pulumi.OutputState }

func (ProviderScimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderScim)(nil)).Elem()
}

func (o ProviderScimMapOutput) ToProviderScimMapOutput() ProviderScimMapOutput {
	return o
}

func (o ProviderScimMapOutput) ToProviderScimMapOutputWithContext(ctx context.Context) ProviderScimMapOutput {
	return o
}

func (o ProviderScimMapOutput) MapIndex(k pulumi.StringInput) ProviderScimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderScim {
		return vs[0].(map[string]*ProviderScim)[vs[1].(string)]
	}).(ProviderScimOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderScimInput)(nil)).Elem(), &ProviderScim{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderScimArrayInput)(nil)).Elem(), ProviderScimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderScimMapInput)(nil)).Elem(), ProviderScimMap{})
	pulumi.RegisterOutputType(ProviderScimOutput{})
	pulumi.RegisterOutputType(ProviderScimArrayOutput{})
	pulumi.RegisterOutputType(ProviderScimMapOutput{})
}

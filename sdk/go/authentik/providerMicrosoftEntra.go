// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderMicrosoftEntra struct {
	pulumi.CustomResourceState

	ClientId                   pulumi.StringOutput    `pulumi:"clientId"`
	ClientSecret               pulumi.StringOutput    `pulumi:"clientSecret"`
	ExcludeUsersServiceAccount pulumi.BoolPtrOutput   `pulumi:"excludeUsersServiceAccount"`
	FilterGroup                pulumi.StringPtrOutput `pulumi:"filterGroup"`
	// Allowed values: - `delete` - `doNothing`
	GroupDeleteAction      pulumi.StringPtrOutput   `pulumi:"groupDeleteAction"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	PropertyMappings       pulumi.StringArrayOutput `pulumi:"propertyMappings"`
	PropertyMappingsGroups pulumi.StringArrayOutput `pulumi:"propertyMappingsGroups"`
	TenantId               pulumi.StringOutput      `pulumi:"tenantId"`
	// Allowed values: - `delete` - `doNothing`
	UserDeleteAction pulumi.StringPtrOutput `pulumi:"userDeleteAction"`
}

// NewProviderMicrosoftEntra registers a new resource with the given unique name, arguments, and options.
func NewProviderMicrosoftEntra(ctx *pulumi.Context,
	name string, args *ProviderMicrosoftEntraArgs, opts ...pulumi.ResourceOption) (*ProviderMicrosoftEntra, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderMicrosoftEntra
	err := ctx.RegisterResource("authentik:index/providerMicrosoftEntra:ProviderMicrosoftEntra", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderMicrosoftEntra gets an existing ProviderMicrosoftEntra resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderMicrosoftEntra(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderMicrosoftEntraState, opts ...pulumi.ResourceOption) (*ProviderMicrosoftEntra, error) {
	var resource ProviderMicrosoftEntra
	err := ctx.ReadResource("authentik:index/providerMicrosoftEntra:ProviderMicrosoftEntra", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderMicrosoftEntra resources.
type providerMicrosoftEntraState struct {
	ClientId                   *string `pulumi:"clientId"`
	ClientSecret               *string `pulumi:"clientSecret"`
	ExcludeUsersServiceAccount *bool   `pulumi:"excludeUsersServiceAccount"`
	FilterGroup                *string `pulumi:"filterGroup"`
	// Allowed values: - `delete` - `doNothing`
	GroupDeleteAction      *string  `pulumi:"groupDeleteAction"`
	Name                   *string  `pulumi:"name"`
	PropertyMappings       []string `pulumi:"propertyMappings"`
	PropertyMappingsGroups []string `pulumi:"propertyMappingsGroups"`
	TenantId               *string  `pulumi:"tenantId"`
	// Allowed values: - `delete` - `doNothing`
	UserDeleteAction *string `pulumi:"userDeleteAction"`
}

type ProviderMicrosoftEntraState struct {
	ClientId                   pulumi.StringPtrInput
	ClientSecret               pulumi.StringPtrInput
	ExcludeUsersServiceAccount pulumi.BoolPtrInput
	FilterGroup                pulumi.StringPtrInput
	// Allowed values: - `delete` - `doNothing`
	GroupDeleteAction      pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	PropertyMappings       pulumi.StringArrayInput
	PropertyMappingsGroups pulumi.StringArrayInput
	TenantId               pulumi.StringPtrInput
	// Allowed values: - `delete` - `doNothing`
	UserDeleteAction pulumi.StringPtrInput
}

func (ProviderMicrosoftEntraState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerMicrosoftEntraState)(nil)).Elem()
}

type providerMicrosoftEntraArgs struct {
	ClientId                   string  `pulumi:"clientId"`
	ClientSecret               string  `pulumi:"clientSecret"`
	ExcludeUsersServiceAccount *bool   `pulumi:"excludeUsersServiceAccount"`
	FilterGroup                *string `pulumi:"filterGroup"`
	// Allowed values: - `delete` - `doNothing`
	GroupDeleteAction      *string  `pulumi:"groupDeleteAction"`
	Name                   *string  `pulumi:"name"`
	PropertyMappings       []string `pulumi:"propertyMappings"`
	PropertyMappingsGroups []string `pulumi:"propertyMappingsGroups"`
	TenantId               string   `pulumi:"tenantId"`
	// Allowed values: - `delete` - `doNothing`
	UserDeleteAction *string `pulumi:"userDeleteAction"`
}

// The set of arguments for constructing a ProviderMicrosoftEntra resource.
type ProviderMicrosoftEntraArgs struct {
	ClientId                   pulumi.StringInput
	ClientSecret               pulumi.StringInput
	ExcludeUsersServiceAccount pulumi.BoolPtrInput
	FilterGroup                pulumi.StringPtrInput
	// Allowed values: - `delete` - `doNothing`
	GroupDeleteAction      pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	PropertyMappings       pulumi.StringArrayInput
	PropertyMappingsGroups pulumi.StringArrayInput
	TenantId               pulumi.StringInput
	// Allowed values: - `delete` - `doNothing`
	UserDeleteAction pulumi.StringPtrInput
}

func (ProviderMicrosoftEntraArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerMicrosoftEntraArgs)(nil)).Elem()
}

type ProviderMicrosoftEntraInput interface {
	pulumi.Input

	ToProviderMicrosoftEntraOutput() ProviderMicrosoftEntraOutput
	ToProviderMicrosoftEntraOutputWithContext(ctx context.Context) ProviderMicrosoftEntraOutput
}

func (*ProviderMicrosoftEntra) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderMicrosoftEntra)(nil)).Elem()
}

func (i *ProviderMicrosoftEntra) ToProviderMicrosoftEntraOutput() ProviderMicrosoftEntraOutput {
	return i.ToProviderMicrosoftEntraOutputWithContext(context.Background())
}

func (i *ProviderMicrosoftEntra) ToProviderMicrosoftEntraOutputWithContext(ctx context.Context) ProviderMicrosoftEntraOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderMicrosoftEntraOutput)
}

// ProviderMicrosoftEntraArrayInput is an input type that accepts ProviderMicrosoftEntraArray and ProviderMicrosoftEntraArrayOutput values.
// You can construct a concrete instance of `ProviderMicrosoftEntraArrayInput` via:
//
//	ProviderMicrosoftEntraArray{ ProviderMicrosoftEntraArgs{...} }
type ProviderMicrosoftEntraArrayInput interface {
	pulumi.Input

	ToProviderMicrosoftEntraArrayOutput() ProviderMicrosoftEntraArrayOutput
	ToProviderMicrosoftEntraArrayOutputWithContext(context.Context) ProviderMicrosoftEntraArrayOutput
}

type ProviderMicrosoftEntraArray []ProviderMicrosoftEntraInput

func (ProviderMicrosoftEntraArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderMicrosoftEntra)(nil)).Elem()
}

func (i ProviderMicrosoftEntraArray) ToProviderMicrosoftEntraArrayOutput() ProviderMicrosoftEntraArrayOutput {
	return i.ToProviderMicrosoftEntraArrayOutputWithContext(context.Background())
}

func (i ProviderMicrosoftEntraArray) ToProviderMicrosoftEntraArrayOutputWithContext(ctx context.Context) ProviderMicrosoftEntraArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderMicrosoftEntraArrayOutput)
}

// ProviderMicrosoftEntraMapInput is an input type that accepts ProviderMicrosoftEntraMap and ProviderMicrosoftEntraMapOutput values.
// You can construct a concrete instance of `ProviderMicrosoftEntraMapInput` via:
//
//	ProviderMicrosoftEntraMap{ "key": ProviderMicrosoftEntraArgs{...} }
type ProviderMicrosoftEntraMapInput interface {
	pulumi.Input

	ToProviderMicrosoftEntraMapOutput() ProviderMicrosoftEntraMapOutput
	ToProviderMicrosoftEntraMapOutputWithContext(context.Context) ProviderMicrosoftEntraMapOutput
}

type ProviderMicrosoftEntraMap map[string]ProviderMicrosoftEntraInput

func (ProviderMicrosoftEntraMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderMicrosoftEntra)(nil)).Elem()
}

func (i ProviderMicrosoftEntraMap) ToProviderMicrosoftEntraMapOutput() ProviderMicrosoftEntraMapOutput {
	return i.ToProviderMicrosoftEntraMapOutputWithContext(context.Background())
}

func (i ProviderMicrosoftEntraMap) ToProviderMicrosoftEntraMapOutputWithContext(ctx context.Context) ProviderMicrosoftEntraMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderMicrosoftEntraMapOutput)
}

type ProviderMicrosoftEntraOutput struct{ *pulumi.OutputState }

func (ProviderMicrosoftEntraOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderMicrosoftEntra)(nil)).Elem()
}

func (o ProviderMicrosoftEntraOutput) ToProviderMicrosoftEntraOutput() ProviderMicrosoftEntraOutput {
	return o
}

func (o ProviderMicrosoftEntraOutput) ToProviderMicrosoftEntraOutputWithContext(ctx context.Context) ProviderMicrosoftEntraOutput {
	return o
}

func (o ProviderMicrosoftEntraOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o ProviderMicrosoftEntraOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o ProviderMicrosoftEntraOutput) ExcludeUsersServiceAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.BoolPtrOutput { return v.ExcludeUsersServiceAccount }).(pulumi.BoolPtrOutput)
}

func (o ProviderMicrosoftEntraOutput) FilterGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringPtrOutput { return v.FilterGroup }).(pulumi.StringPtrOutput)
}

// Allowed values: - `delete` - `doNothing`
func (o ProviderMicrosoftEntraOutput) GroupDeleteAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringPtrOutput { return v.GroupDeleteAction }).(pulumi.StringPtrOutput)
}

func (o ProviderMicrosoftEntraOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderMicrosoftEntraOutput) PropertyMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringArrayOutput { return v.PropertyMappings }).(pulumi.StringArrayOutput)
}

func (o ProviderMicrosoftEntraOutput) PropertyMappingsGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringArrayOutput { return v.PropertyMappingsGroups }).(pulumi.StringArrayOutput)
}

func (o ProviderMicrosoftEntraOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Allowed values: - `delete` - `doNothing`
func (o ProviderMicrosoftEntraOutput) UserDeleteAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderMicrosoftEntra) pulumi.StringPtrOutput { return v.UserDeleteAction }).(pulumi.StringPtrOutput)
}

type ProviderMicrosoftEntraArrayOutput struct{ *pulumi.OutputState }

func (ProviderMicrosoftEntraArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderMicrosoftEntra)(nil)).Elem()
}

func (o ProviderMicrosoftEntraArrayOutput) ToProviderMicrosoftEntraArrayOutput() ProviderMicrosoftEntraArrayOutput {
	return o
}

func (o ProviderMicrosoftEntraArrayOutput) ToProviderMicrosoftEntraArrayOutputWithContext(ctx context.Context) ProviderMicrosoftEntraArrayOutput {
	return o
}

func (o ProviderMicrosoftEntraArrayOutput) Index(i pulumi.IntInput) ProviderMicrosoftEntraOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderMicrosoftEntra {
		return vs[0].([]*ProviderMicrosoftEntra)[vs[1].(int)]
	}).(ProviderMicrosoftEntraOutput)
}

type ProviderMicrosoftEntraMapOutput struct{ *pulumi.OutputState }

func (ProviderMicrosoftEntraMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderMicrosoftEntra)(nil)).Elem()
}

func (o ProviderMicrosoftEntraMapOutput) ToProviderMicrosoftEntraMapOutput() ProviderMicrosoftEntraMapOutput {
	return o
}

func (o ProviderMicrosoftEntraMapOutput) ToProviderMicrosoftEntraMapOutputWithContext(ctx context.Context) ProviderMicrosoftEntraMapOutput {
	return o
}

func (o ProviderMicrosoftEntraMapOutput) MapIndex(k pulumi.StringInput) ProviderMicrosoftEntraOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderMicrosoftEntra {
		return vs[0].(map[string]*ProviderMicrosoftEntra)[vs[1].(string)]
	}).(ProviderMicrosoftEntraOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderMicrosoftEntraInput)(nil)).Elem(), &ProviderMicrosoftEntra{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderMicrosoftEntraArrayInput)(nil)).Elem(), ProviderMicrosoftEntraArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderMicrosoftEntraMapInput)(nil)).Elem(), ProviderMicrosoftEntraMap{})
	pulumi.RegisterOutputType(ProviderMicrosoftEntraOutput{})
	pulumi.RegisterOutputType(ProviderMicrosoftEntraArrayOutput{})
	pulumi.RegisterOutputType(ProviderMicrosoftEntraMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
type StageAuthenticatorWebauthn struct {
	pulumi.CustomResourceState

	// Allowed values: - `platform` - `cross-platform`
	AuthenticatorAttachment pulumi.StringPtrOutput   `pulumi:"authenticatorAttachment"`
	ConfigureFlow           pulumi.StringPtrOutput   `pulumi:"configureFlow"`
	DeviceTypeRestrictions  pulumi.StringArrayOutput `pulumi:"deviceTypeRestrictions"`
	FriendlyName            pulumi.StringPtrOutput   `pulumi:"friendlyName"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	// Allowed values: - `discouraged` - `preferred` - `required`
	ResidentKeyRequirement pulumi.StringPtrOutput `pulumi:"residentKeyRequirement"`
	// Allowed values: - `required` - `preferred` - `discouraged`
	UserVerification pulumi.StringPtrOutput `pulumi:"userVerification"`
}

// NewStageAuthenticatorWebauthn registers a new resource with the given unique name, arguments, and options.
func NewStageAuthenticatorWebauthn(ctx *pulumi.Context,
	name string, args *StageAuthenticatorWebauthnArgs, opts ...pulumi.ResourceOption) (*StageAuthenticatorWebauthn, error) {
	if args == nil {
		args = &StageAuthenticatorWebauthnArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageAuthenticatorWebauthn
	err := ctx.RegisterResource("authentik:index/stageAuthenticatorWebauthn:StageAuthenticatorWebauthn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageAuthenticatorWebauthn gets an existing StageAuthenticatorWebauthn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageAuthenticatorWebauthn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageAuthenticatorWebauthnState, opts ...pulumi.ResourceOption) (*StageAuthenticatorWebauthn, error) {
	var resource StageAuthenticatorWebauthn
	err := ctx.ReadResource("authentik:index/stageAuthenticatorWebauthn:StageAuthenticatorWebauthn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageAuthenticatorWebauthn resources.
type stageAuthenticatorWebauthnState struct {
	// Allowed values: - `platform` - `cross-platform`
	AuthenticatorAttachment *string  `pulumi:"authenticatorAttachment"`
	ConfigureFlow           *string  `pulumi:"configureFlow"`
	DeviceTypeRestrictions  []string `pulumi:"deviceTypeRestrictions"`
	FriendlyName            *string  `pulumi:"friendlyName"`
	Name                    *string  `pulumi:"name"`
	// Allowed values: - `discouraged` - `preferred` - `required`
	ResidentKeyRequirement *string `pulumi:"residentKeyRequirement"`
	// Allowed values: - `required` - `preferred` - `discouraged`
	UserVerification *string `pulumi:"userVerification"`
}

type StageAuthenticatorWebauthnState struct {
	// Allowed values: - `platform` - `cross-platform`
	AuthenticatorAttachment pulumi.StringPtrInput
	ConfigureFlow           pulumi.StringPtrInput
	DeviceTypeRestrictions  pulumi.StringArrayInput
	FriendlyName            pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	// Allowed values: - `discouraged` - `preferred` - `required`
	ResidentKeyRequirement pulumi.StringPtrInput
	// Allowed values: - `required` - `preferred` - `discouraged`
	UserVerification pulumi.StringPtrInput
}

func (StageAuthenticatorWebauthnState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageAuthenticatorWebauthnState)(nil)).Elem()
}

type stageAuthenticatorWebauthnArgs struct {
	// Allowed values: - `platform` - `cross-platform`
	AuthenticatorAttachment *string  `pulumi:"authenticatorAttachment"`
	ConfigureFlow           *string  `pulumi:"configureFlow"`
	DeviceTypeRestrictions  []string `pulumi:"deviceTypeRestrictions"`
	FriendlyName            *string  `pulumi:"friendlyName"`
	Name                    *string  `pulumi:"name"`
	// Allowed values: - `discouraged` - `preferred` - `required`
	ResidentKeyRequirement *string `pulumi:"residentKeyRequirement"`
	// Allowed values: - `required` - `preferred` - `discouraged`
	UserVerification *string `pulumi:"userVerification"`
}

// The set of arguments for constructing a StageAuthenticatorWebauthn resource.
type StageAuthenticatorWebauthnArgs struct {
	// Allowed values: - `platform` - `cross-platform`
	AuthenticatorAttachment pulumi.StringPtrInput
	ConfigureFlow           pulumi.StringPtrInput
	DeviceTypeRestrictions  pulumi.StringArrayInput
	FriendlyName            pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	// Allowed values: - `discouraged` - `preferred` - `required`
	ResidentKeyRequirement pulumi.StringPtrInput
	// Allowed values: - `required` - `preferred` - `discouraged`
	UserVerification pulumi.StringPtrInput
}

func (StageAuthenticatorWebauthnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageAuthenticatorWebauthnArgs)(nil)).Elem()
}

type StageAuthenticatorWebauthnInput interface {
	pulumi.Input

	ToStageAuthenticatorWebauthnOutput() StageAuthenticatorWebauthnOutput
	ToStageAuthenticatorWebauthnOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnOutput
}

func (*StageAuthenticatorWebauthn) ElementType() reflect.Type {
	return reflect.TypeOf((**StageAuthenticatorWebauthn)(nil)).Elem()
}

func (i *StageAuthenticatorWebauthn) ToStageAuthenticatorWebauthnOutput() StageAuthenticatorWebauthnOutput {
	return i.ToStageAuthenticatorWebauthnOutputWithContext(context.Background())
}

func (i *StageAuthenticatorWebauthn) ToStageAuthenticatorWebauthnOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorWebauthnOutput)
}

// StageAuthenticatorWebauthnArrayInput is an input type that accepts StageAuthenticatorWebauthnArray and StageAuthenticatorWebauthnArrayOutput values.
// You can construct a concrete instance of `StageAuthenticatorWebauthnArrayInput` via:
//
//	StageAuthenticatorWebauthnArray{ StageAuthenticatorWebauthnArgs{...} }
type StageAuthenticatorWebauthnArrayInput interface {
	pulumi.Input

	ToStageAuthenticatorWebauthnArrayOutput() StageAuthenticatorWebauthnArrayOutput
	ToStageAuthenticatorWebauthnArrayOutputWithContext(context.Context) StageAuthenticatorWebauthnArrayOutput
}

type StageAuthenticatorWebauthnArray []StageAuthenticatorWebauthnInput

func (StageAuthenticatorWebauthnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageAuthenticatorWebauthn)(nil)).Elem()
}

func (i StageAuthenticatorWebauthnArray) ToStageAuthenticatorWebauthnArrayOutput() StageAuthenticatorWebauthnArrayOutput {
	return i.ToStageAuthenticatorWebauthnArrayOutputWithContext(context.Background())
}

func (i StageAuthenticatorWebauthnArray) ToStageAuthenticatorWebauthnArrayOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorWebauthnArrayOutput)
}

// StageAuthenticatorWebauthnMapInput is an input type that accepts StageAuthenticatorWebauthnMap and StageAuthenticatorWebauthnMapOutput values.
// You can construct a concrete instance of `StageAuthenticatorWebauthnMapInput` via:
//
//	StageAuthenticatorWebauthnMap{ "key": StageAuthenticatorWebauthnArgs{...} }
type StageAuthenticatorWebauthnMapInput interface {
	pulumi.Input

	ToStageAuthenticatorWebauthnMapOutput() StageAuthenticatorWebauthnMapOutput
	ToStageAuthenticatorWebauthnMapOutputWithContext(context.Context) StageAuthenticatorWebauthnMapOutput
}

type StageAuthenticatorWebauthnMap map[string]StageAuthenticatorWebauthnInput

func (StageAuthenticatorWebauthnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageAuthenticatorWebauthn)(nil)).Elem()
}

func (i StageAuthenticatorWebauthnMap) ToStageAuthenticatorWebauthnMapOutput() StageAuthenticatorWebauthnMapOutput {
	return i.ToStageAuthenticatorWebauthnMapOutputWithContext(context.Background())
}

func (i StageAuthenticatorWebauthnMap) ToStageAuthenticatorWebauthnMapOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorWebauthnMapOutput)
}

type StageAuthenticatorWebauthnOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorWebauthnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageAuthenticatorWebauthn)(nil)).Elem()
}

func (o StageAuthenticatorWebauthnOutput) ToStageAuthenticatorWebauthnOutput() StageAuthenticatorWebauthnOutput {
	return o
}

func (o StageAuthenticatorWebauthnOutput) ToStageAuthenticatorWebauthnOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnOutput {
	return o
}

// Allowed values: - `platform` - `cross-platform`
func (o StageAuthenticatorWebauthnOutput) AuthenticatorAttachment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringPtrOutput { return v.AuthenticatorAttachment }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorWebauthnOutput) ConfigureFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringPtrOutput { return v.ConfigureFlow }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorWebauthnOutput) DeviceTypeRestrictions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringArrayOutput { return v.DeviceTypeRestrictions }).(pulumi.StringArrayOutput)
}

func (o StageAuthenticatorWebauthnOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorWebauthnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Allowed values: - `discouraged` - `preferred` - `required`
func (o StageAuthenticatorWebauthnOutput) ResidentKeyRequirement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringPtrOutput { return v.ResidentKeyRequirement }).(pulumi.StringPtrOutput)
}

// Allowed values: - `required` - `preferred` - `discouraged`
func (o StageAuthenticatorWebauthnOutput) UserVerification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorWebauthn) pulumi.StringPtrOutput { return v.UserVerification }).(pulumi.StringPtrOutput)
}

type StageAuthenticatorWebauthnArrayOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorWebauthnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageAuthenticatorWebauthn)(nil)).Elem()
}

func (o StageAuthenticatorWebauthnArrayOutput) ToStageAuthenticatorWebauthnArrayOutput() StageAuthenticatorWebauthnArrayOutput {
	return o
}

func (o StageAuthenticatorWebauthnArrayOutput) ToStageAuthenticatorWebauthnArrayOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnArrayOutput {
	return o
}

func (o StageAuthenticatorWebauthnArrayOutput) Index(i pulumi.IntInput) StageAuthenticatorWebauthnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageAuthenticatorWebauthn {
		return vs[0].([]*StageAuthenticatorWebauthn)[vs[1].(int)]
	}).(StageAuthenticatorWebauthnOutput)
}

type StageAuthenticatorWebauthnMapOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorWebauthnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageAuthenticatorWebauthn)(nil)).Elem()
}

func (o StageAuthenticatorWebauthnMapOutput) ToStageAuthenticatorWebauthnMapOutput() StageAuthenticatorWebauthnMapOutput {
	return o
}

func (o StageAuthenticatorWebauthnMapOutput) ToStageAuthenticatorWebauthnMapOutputWithContext(ctx context.Context) StageAuthenticatorWebauthnMapOutput {
	return o
}

func (o StageAuthenticatorWebauthnMapOutput) MapIndex(k pulumi.StringInput) StageAuthenticatorWebauthnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageAuthenticatorWebauthn {
		return vs[0].(map[string]*StageAuthenticatorWebauthn)[vs[1].(string)]
	}).(StageAuthenticatorWebauthnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorWebauthnInput)(nil)).Elem(), &StageAuthenticatorWebauthn{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorWebauthnArrayInput)(nil)).Elem(), StageAuthenticatorWebauthnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorWebauthnMapInput)(nil)).Elem(), StageAuthenticatorWebauthnMap{})
	pulumi.RegisterOutputType(StageAuthenticatorWebauthnOutput{})
	pulumi.RegisterOutputType(StageAuthenticatorWebauthnArrayOutput{})
	pulumi.RegisterOutputType(StageAuthenticatorWebauthnMapOutput{})
}

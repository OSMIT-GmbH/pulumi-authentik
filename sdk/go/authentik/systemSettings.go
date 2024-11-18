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
//			// Configure system settings
//			_, err := authentik.NewSystemSettings(ctx, "settings", &authentik.SystemSettingsArgs{
//				DefaultUserChangeUsername: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// {{codefile "shell" "examples/resources/authentik_system_settings/import.sh"}}
type SystemSettings struct {
	pulumi.CustomResourceState

	// Defaults to `gravatar,initials`.
	Avatars pulumi.StringPtrOutput `pulumi:"avatars"`
	// Defaults to `minutes=30`.
	DefaultTokenDuration pulumi.StringPtrOutput `pulumi:"defaultTokenDuration"`
	// Defaults to `60`.
	DefaultTokenLength pulumi.IntPtrOutput `pulumi:"defaultTokenLength"`
	// Defaults to `false`.
	DefaultUserChangeEmail pulumi.BoolPtrOutput `pulumi:"defaultUserChangeEmail"`
	// Defaults to `true`.
	DefaultUserChangeName pulumi.BoolPtrOutput `pulumi:"defaultUserChangeName"`
	// Defaults to `false`.
	DefaultUserChangeUsername pulumi.BoolPtrOutput `pulumi:"defaultUserChangeUsername"`
	// Defaults to `days=365`.
	EventRetention pulumi.StringPtrOutput      `pulumi:"eventRetention"`
	FooterLinks    pulumi.StringMapArrayOutput `pulumi:"footerLinks"`
	// Defaults to `true`.
	GdprCompliance pulumi.BoolPtrOutput `pulumi:"gdprCompliance"`
	// Defaults to `true`.
	Impersonation pulumi.BoolPtrOutput `pulumi:"impersonation"`
}

// NewSystemSettings registers a new resource with the given unique name, arguments, and options.
func NewSystemSettings(ctx *pulumi.Context,
	name string, args *SystemSettingsArgs, opts ...pulumi.ResourceOption) (*SystemSettings, error) {
	if args == nil {
		args = &SystemSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemSettings
	err := ctx.RegisterResource("authentik:index/systemSettings:SystemSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSettings gets an existing SystemSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSettingsState, opts ...pulumi.ResourceOption) (*SystemSettings, error) {
	var resource SystemSettings
	err := ctx.ReadResource("authentik:index/systemSettings:SystemSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSettings resources.
type systemSettingsState struct {
	// Defaults to `gravatar,initials`.
	Avatars *string `pulumi:"avatars"`
	// Defaults to `minutes=30`.
	DefaultTokenDuration *string `pulumi:"defaultTokenDuration"`
	// Defaults to `60`.
	DefaultTokenLength *int `pulumi:"defaultTokenLength"`
	// Defaults to `false`.
	DefaultUserChangeEmail *bool `pulumi:"defaultUserChangeEmail"`
	// Defaults to `true`.
	DefaultUserChangeName *bool `pulumi:"defaultUserChangeName"`
	// Defaults to `false`.
	DefaultUserChangeUsername *bool `pulumi:"defaultUserChangeUsername"`
	// Defaults to `days=365`.
	EventRetention *string             `pulumi:"eventRetention"`
	FooterLinks    []map[string]string `pulumi:"footerLinks"`
	// Defaults to `true`.
	GdprCompliance *bool `pulumi:"gdprCompliance"`
	// Defaults to `true`.
	Impersonation *bool `pulumi:"impersonation"`
}

type SystemSettingsState struct {
	// Defaults to `gravatar,initials`.
	Avatars pulumi.StringPtrInput
	// Defaults to `minutes=30`.
	DefaultTokenDuration pulumi.StringPtrInput
	// Defaults to `60`.
	DefaultTokenLength pulumi.IntPtrInput
	// Defaults to `false`.
	DefaultUserChangeEmail pulumi.BoolPtrInput
	// Defaults to `true`.
	DefaultUserChangeName pulumi.BoolPtrInput
	// Defaults to `false`.
	DefaultUserChangeUsername pulumi.BoolPtrInput
	// Defaults to `days=365`.
	EventRetention pulumi.StringPtrInput
	FooterLinks    pulumi.StringMapArrayInput
	// Defaults to `true`.
	GdprCompliance pulumi.BoolPtrInput
	// Defaults to `true`.
	Impersonation pulumi.BoolPtrInput
}

func (SystemSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSettingsState)(nil)).Elem()
}

type systemSettingsArgs struct {
	// Defaults to `gravatar,initials`.
	Avatars *string `pulumi:"avatars"`
	// Defaults to `minutes=30`.
	DefaultTokenDuration *string `pulumi:"defaultTokenDuration"`
	// Defaults to `60`.
	DefaultTokenLength *int `pulumi:"defaultTokenLength"`
	// Defaults to `false`.
	DefaultUserChangeEmail *bool `pulumi:"defaultUserChangeEmail"`
	// Defaults to `true`.
	DefaultUserChangeName *bool `pulumi:"defaultUserChangeName"`
	// Defaults to `false`.
	DefaultUserChangeUsername *bool `pulumi:"defaultUserChangeUsername"`
	// Defaults to `days=365`.
	EventRetention *string             `pulumi:"eventRetention"`
	FooterLinks    []map[string]string `pulumi:"footerLinks"`
	// Defaults to `true`.
	GdprCompliance *bool `pulumi:"gdprCompliance"`
	// Defaults to `true`.
	Impersonation *bool `pulumi:"impersonation"`
}

// The set of arguments for constructing a SystemSettings resource.
type SystemSettingsArgs struct {
	// Defaults to `gravatar,initials`.
	Avatars pulumi.StringPtrInput
	// Defaults to `minutes=30`.
	DefaultTokenDuration pulumi.StringPtrInput
	// Defaults to `60`.
	DefaultTokenLength pulumi.IntPtrInput
	// Defaults to `false`.
	DefaultUserChangeEmail pulumi.BoolPtrInput
	// Defaults to `true`.
	DefaultUserChangeName pulumi.BoolPtrInput
	// Defaults to `false`.
	DefaultUserChangeUsername pulumi.BoolPtrInput
	// Defaults to `days=365`.
	EventRetention pulumi.StringPtrInput
	FooterLinks    pulumi.StringMapArrayInput
	// Defaults to `true`.
	GdprCompliance pulumi.BoolPtrInput
	// Defaults to `true`.
	Impersonation pulumi.BoolPtrInput
}

func (SystemSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSettingsArgs)(nil)).Elem()
}

type SystemSettingsInput interface {
	pulumi.Input

	ToSystemSettingsOutput() SystemSettingsOutput
	ToSystemSettingsOutputWithContext(ctx context.Context) SystemSettingsOutput
}

func (*SystemSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSettings)(nil)).Elem()
}

func (i *SystemSettings) ToSystemSettingsOutput() SystemSettingsOutput {
	return i.ToSystemSettingsOutputWithContext(context.Background())
}

func (i *SystemSettings) ToSystemSettingsOutputWithContext(ctx context.Context) SystemSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingsOutput)
}

// SystemSettingsArrayInput is an input type that accepts SystemSettingsArray and SystemSettingsArrayOutput values.
// You can construct a concrete instance of `SystemSettingsArrayInput` via:
//
//	SystemSettingsArray{ SystemSettingsArgs{...} }
type SystemSettingsArrayInput interface {
	pulumi.Input

	ToSystemSettingsArrayOutput() SystemSettingsArrayOutput
	ToSystemSettingsArrayOutputWithContext(context.Context) SystemSettingsArrayOutput
}

type SystemSettingsArray []SystemSettingsInput

func (SystemSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSettings)(nil)).Elem()
}

func (i SystemSettingsArray) ToSystemSettingsArrayOutput() SystemSettingsArrayOutput {
	return i.ToSystemSettingsArrayOutputWithContext(context.Background())
}

func (i SystemSettingsArray) ToSystemSettingsArrayOutputWithContext(ctx context.Context) SystemSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingsArrayOutput)
}

// SystemSettingsMapInput is an input type that accepts SystemSettingsMap and SystemSettingsMapOutput values.
// You can construct a concrete instance of `SystemSettingsMapInput` via:
//
//	SystemSettingsMap{ "key": SystemSettingsArgs{...} }
type SystemSettingsMapInput interface {
	pulumi.Input

	ToSystemSettingsMapOutput() SystemSettingsMapOutput
	ToSystemSettingsMapOutputWithContext(context.Context) SystemSettingsMapOutput
}

type SystemSettingsMap map[string]SystemSettingsInput

func (SystemSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSettings)(nil)).Elem()
}

func (i SystemSettingsMap) ToSystemSettingsMapOutput() SystemSettingsMapOutput {
	return i.ToSystemSettingsMapOutputWithContext(context.Background())
}

func (i SystemSettingsMap) ToSystemSettingsMapOutputWithContext(ctx context.Context) SystemSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingsMapOutput)
}

type SystemSettingsOutput struct{ *pulumi.OutputState }

func (SystemSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSettings)(nil)).Elem()
}

func (o SystemSettingsOutput) ToSystemSettingsOutput() SystemSettingsOutput {
	return o
}

func (o SystemSettingsOutput) ToSystemSettingsOutputWithContext(ctx context.Context) SystemSettingsOutput {
	return o
}

// Defaults to `gravatar,initials`.
func (o SystemSettingsOutput) Avatars() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.StringPtrOutput { return v.Avatars }).(pulumi.StringPtrOutput)
}

// Defaults to `minutes=30`.
func (o SystemSettingsOutput) DefaultTokenDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.StringPtrOutput { return v.DefaultTokenDuration }).(pulumi.StringPtrOutput)
}

// Defaults to `60`.
func (o SystemSettingsOutput) DefaultTokenLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.IntPtrOutput { return v.DefaultTokenLength }).(pulumi.IntPtrOutput)
}

// Defaults to `false`.
func (o SystemSettingsOutput) DefaultUserChangeEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.BoolPtrOutput { return v.DefaultUserChangeEmail }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o SystemSettingsOutput) DefaultUserChangeName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.BoolPtrOutput { return v.DefaultUserChangeName }).(pulumi.BoolPtrOutput)
}

// Defaults to `false`.
func (o SystemSettingsOutput) DefaultUserChangeUsername() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.BoolPtrOutput { return v.DefaultUserChangeUsername }).(pulumi.BoolPtrOutput)
}

// Defaults to `days=365`.
func (o SystemSettingsOutput) EventRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.StringPtrOutput { return v.EventRetention }).(pulumi.StringPtrOutput)
}

func (o SystemSettingsOutput) FooterLinks() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.StringMapArrayOutput { return v.FooterLinks }).(pulumi.StringMapArrayOutput)
}

// Defaults to `true`.
func (o SystemSettingsOutput) GdprCompliance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.BoolPtrOutput { return v.GdprCompliance }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o SystemSettingsOutput) Impersonation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemSettings) pulumi.BoolPtrOutput { return v.Impersonation }).(pulumi.BoolPtrOutput)
}

type SystemSettingsArrayOutput struct{ *pulumi.OutputState }

func (SystemSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSettings)(nil)).Elem()
}

func (o SystemSettingsArrayOutput) ToSystemSettingsArrayOutput() SystemSettingsArrayOutput {
	return o
}

func (o SystemSettingsArrayOutput) ToSystemSettingsArrayOutputWithContext(ctx context.Context) SystemSettingsArrayOutput {
	return o
}

func (o SystemSettingsArrayOutput) Index(i pulumi.IntInput) SystemSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSettings {
		return vs[0].([]*SystemSettings)[vs[1].(int)]
	}).(SystemSettingsOutput)
}

type SystemSettingsMapOutput struct{ *pulumi.OutputState }

func (SystemSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSettings)(nil)).Elem()
}

func (o SystemSettingsMapOutput) ToSystemSettingsMapOutput() SystemSettingsMapOutput {
	return o
}

func (o SystemSettingsMapOutput) ToSystemSettingsMapOutputWithContext(ctx context.Context) SystemSettingsMapOutput {
	return o
}

func (o SystemSettingsMapOutput) MapIndex(k pulumi.StringInput) SystemSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSettings {
		return vs[0].(map[string]*SystemSettings)[vs[1].(string)]
	}).(SystemSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingsInput)(nil)).Elem(), &SystemSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingsArrayInput)(nil)).Elem(), SystemSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingsMapInput)(nil)).Elem(), SystemSettingsMap{})
	pulumi.RegisterOutputType(SystemSettingsOutput{})
	pulumi.RegisterOutputType(SystemSettingsArrayOutput{})
	pulumi.RegisterOutputType(SystemSettingsMapOutput{})
}

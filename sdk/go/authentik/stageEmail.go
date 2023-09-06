// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StageEmail struct {
	pulumi.CustomResourceState

	ActivateUserOnSuccess pulumi.BoolPtrOutput   `pulumi:"activateUserOnSuccess"`
	FromAddress           pulumi.StringPtrOutput `pulumi:"fromAddress"`
	Host                  pulumi.StringPtrOutput `pulumi:"host"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	Password              pulumi.StringPtrOutput `pulumi:"password"`
	Port                  pulumi.IntPtrOutput    `pulumi:"port"`
	Subject               pulumi.StringPtrOutput `pulumi:"subject"`
	Template              pulumi.StringPtrOutput `pulumi:"template"`
	Timeout               pulumi.IntPtrOutput    `pulumi:"timeout"`
	TokenExpiry           pulumi.IntPtrOutput    `pulumi:"tokenExpiry"`
	UseGlobalSettings     pulumi.BoolPtrOutput   `pulumi:"useGlobalSettings"`
	UseSsl                pulumi.BoolPtrOutput   `pulumi:"useSsl"`
	UseTls                pulumi.BoolPtrOutput   `pulumi:"useTls"`
	Username              pulumi.StringPtrOutput `pulumi:"username"`
}

// NewStageEmail registers a new resource with the given unique name, arguments, and options.
func NewStageEmail(ctx *pulumi.Context,
	name string, args *StageEmailArgs, opts ...pulumi.ResourceOption) (*StageEmail, error) {
	if args == nil {
		args = &StageEmailArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageEmail
	err := ctx.RegisterResource("authentik:index/stageEmail:StageEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageEmail gets an existing StageEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageEmailState, opts ...pulumi.ResourceOption) (*StageEmail, error) {
	var resource StageEmail
	err := ctx.ReadResource("authentik:index/stageEmail:StageEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageEmail resources.
type stageEmailState struct {
	ActivateUserOnSuccess *bool   `pulumi:"activateUserOnSuccess"`
	FromAddress           *string `pulumi:"fromAddress"`
	Host                  *string `pulumi:"host"`
	Name                  *string `pulumi:"name"`
	Password              *string `pulumi:"password"`
	Port                  *int    `pulumi:"port"`
	Subject               *string `pulumi:"subject"`
	Template              *string `pulumi:"template"`
	Timeout               *int    `pulumi:"timeout"`
	TokenExpiry           *int    `pulumi:"tokenExpiry"`
	UseGlobalSettings     *bool   `pulumi:"useGlobalSettings"`
	UseSsl                *bool   `pulumi:"useSsl"`
	UseTls                *bool   `pulumi:"useTls"`
	Username              *string `pulumi:"username"`
}

type StageEmailState struct {
	ActivateUserOnSuccess pulumi.BoolPtrInput
	FromAddress           pulumi.StringPtrInput
	Host                  pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	Subject               pulumi.StringPtrInput
	Template              pulumi.StringPtrInput
	Timeout               pulumi.IntPtrInput
	TokenExpiry           pulumi.IntPtrInput
	UseGlobalSettings     pulumi.BoolPtrInput
	UseSsl                pulumi.BoolPtrInput
	UseTls                pulumi.BoolPtrInput
	Username              pulumi.StringPtrInput
}

func (StageEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageEmailState)(nil)).Elem()
}

type stageEmailArgs struct {
	ActivateUserOnSuccess *bool   `pulumi:"activateUserOnSuccess"`
	FromAddress           *string `pulumi:"fromAddress"`
	Host                  *string `pulumi:"host"`
	Name                  *string `pulumi:"name"`
	Password              *string `pulumi:"password"`
	Port                  *int    `pulumi:"port"`
	Subject               *string `pulumi:"subject"`
	Template              *string `pulumi:"template"`
	Timeout               *int    `pulumi:"timeout"`
	TokenExpiry           *int    `pulumi:"tokenExpiry"`
	UseGlobalSettings     *bool   `pulumi:"useGlobalSettings"`
	UseSsl                *bool   `pulumi:"useSsl"`
	UseTls                *bool   `pulumi:"useTls"`
	Username              *string `pulumi:"username"`
}

// The set of arguments for constructing a StageEmail resource.
type StageEmailArgs struct {
	ActivateUserOnSuccess pulumi.BoolPtrInput
	FromAddress           pulumi.StringPtrInput
	Host                  pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	Subject               pulumi.StringPtrInput
	Template              pulumi.StringPtrInput
	Timeout               pulumi.IntPtrInput
	TokenExpiry           pulumi.IntPtrInput
	UseGlobalSettings     pulumi.BoolPtrInput
	UseSsl                pulumi.BoolPtrInput
	UseTls                pulumi.BoolPtrInput
	Username              pulumi.StringPtrInput
}

func (StageEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageEmailArgs)(nil)).Elem()
}

type StageEmailInput interface {
	pulumi.Input

	ToStageEmailOutput() StageEmailOutput
	ToStageEmailOutputWithContext(ctx context.Context) StageEmailOutput
}

func (*StageEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**StageEmail)(nil)).Elem()
}

func (i *StageEmail) ToStageEmailOutput() StageEmailOutput {
	return i.ToStageEmailOutputWithContext(context.Background())
}

func (i *StageEmail) ToStageEmailOutputWithContext(ctx context.Context) StageEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageEmailOutput)
}

// StageEmailArrayInput is an input type that accepts StageEmailArray and StageEmailArrayOutput values.
// You can construct a concrete instance of `StageEmailArrayInput` via:
//
//	StageEmailArray{ StageEmailArgs{...} }
type StageEmailArrayInput interface {
	pulumi.Input

	ToStageEmailArrayOutput() StageEmailArrayOutput
	ToStageEmailArrayOutputWithContext(context.Context) StageEmailArrayOutput
}

type StageEmailArray []StageEmailInput

func (StageEmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageEmail)(nil)).Elem()
}

func (i StageEmailArray) ToStageEmailArrayOutput() StageEmailArrayOutput {
	return i.ToStageEmailArrayOutputWithContext(context.Background())
}

func (i StageEmailArray) ToStageEmailArrayOutputWithContext(ctx context.Context) StageEmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageEmailArrayOutput)
}

// StageEmailMapInput is an input type that accepts StageEmailMap and StageEmailMapOutput values.
// You can construct a concrete instance of `StageEmailMapInput` via:
//
//	StageEmailMap{ "key": StageEmailArgs{...} }
type StageEmailMapInput interface {
	pulumi.Input

	ToStageEmailMapOutput() StageEmailMapOutput
	ToStageEmailMapOutputWithContext(context.Context) StageEmailMapOutput
}

type StageEmailMap map[string]StageEmailInput

func (StageEmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageEmail)(nil)).Elem()
}

func (i StageEmailMap) ToStageEmailMapOutput() StageEmailMapOutput {
	return i.ToStageEmailMapOutputWithContext(context.Background())
}

func (i StageEmailMap) ToStageEmailMapOutputWithContext(ctx context.Context) StageEmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageEmailMapOutput)
}

type StageEmailOutput struct{ *pulumi.OutputState }

func (StageEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageEmail)(nil)).Elem()
}

func (o StageEmailOutput) ToStageEmailOutput() StageEmailOutput {
	return o
}

func (o StageEmailOutput) ToStageEmailOutputWithContext(ctx context.Context) StageEmailOutput {
	return o
}

func (o StageEmailOutput) ActivateUserOnSuccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.BoolPtrOutput { return v.ActivateUserOnSuccess }).(pulumi.BoolPtrOutput)
}

func (o StageEmailOutput) FromAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringPtrOutput { return v.FromAddress }).(pulumi.StringPtrOutput)
}

func (o StageEmailOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

func (o StageEmailOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StageEmailOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o StageEmailOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o StageEmailOutput) Subject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringPtrOutput { return v.Subject }).(pulumi.StringPtrOutput)
}

func (o StageEmailOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

func (o StageEmailOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o StageEmailOutput) TokenExpiry() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.IntPtrOutput { return v.TokenExpiry }).(pulumi.IntPtrOutput)
}

func (o StageEmailOutput) UseGlobalSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.BoolPtrOutput { return v.UseGlobalSettings }).(pulumi.BoolPtrOutput)
}

func (o StageEmailOutput) UseSsl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.BoolPtrOutput { return v.UseSsl }).(pulumi.BoolPtrOutput)
}

func (o StageEmailOutput) UseTls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.BoolPtrOutput { return v.UseTls }).(pulumi.BoolPtrOutput)
}

func (o StageEmailOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageEmail) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type StageEmailArrayOutput struct{ *pulumi.OutputState }

func (StageEmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageEmail)(nil)).Elem()
}

func (o StageEmailArrayOutput) ToStageEmailArrayOutput() StageEmailArrayOutput {
	return o
}

func (o StageEmailArrayOutput) ToStageEmailArrayOutputWithContext(ctx context.Context) StageEmailArrayOutput {
	return o
}

func (o StageEmailArrayOutput) Index(i pulumi.IntInput) StageEmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageEmail {
		return vs[0].([]*StageEmail)[vs[1].(int)]
	}).(StageEmailOutput)
}

type StageEmailMapOutput struct{ *pulumi.OutputState }

func (StageEmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageEmail)(nil)).Elem()
}

func (o StageEmailMapOutput) ToStageEmailMapOutput() StageEmailMapOutput {
	return o
}

func (o StageEmailMapOutput) ToStageEmailMapOutputWithContext(ctx context.Context) StageEmailMapOutput {
	return o
}

func (o StageEmailMapOutput) MapIndex(k pulumi.StringInput) StageEmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageEmail {
		return vs[0].(map[string]*StageEmail)[vs[1].(string)]
	}).(StageEmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageEmailInput)(nil)).Elem(), &StageEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageEmailArrayInput)(nil)).Elem(), StageEmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageEmailMapInput)(nil)).Elem(), StageEmailMap{})
	pulumi.RegisterOutputType(StageEmailOutput{})
	pulumi.RegisterOutputType(StageEmailArrayOutput{})
	pulumi.RegisterOutputType(StageEmailMapOutput{})
}

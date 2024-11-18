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
//			// Create a TOTP Setup stage
//			_, err := authentik.NewStageAuthenticatorTotp(ctx, "name", &authentik.StageAuthenticatorTotpArgs{
//				Name: pulumi.String("totp-setup"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageAuthenticatorTotp struct {
	pulumi.CustomResourceState

	ConfigureFlow pulumi.StringPtrOutput `pulumi:"configureFlow"`
	// Allowed values: - `6` - `8`
	Digits       pulumi.StringPtrOutput `pulumi:"digits"`
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	Name         pulumi.StringOutput    `pulumi:"name"`
}

// NewStageAuthenticatorTotp registers a new resource with the given unique name, arguments, and options.
func NewStageAuthenticatorTotp(ctx *pulumi.Context,
	name string, args *StageAuthenticatorTotpArgs, opts ...pulumi.ResourceOption) (*StageAuthenticatorTotp, error) {
	if args == nil {
		args = &StageAuthenticatorTotpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageAuthenticatorTotp
	err := ctx.RegisterResource("authentik:index/stageAuthenticatorTotp:StageAuthenticatorTotp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageAuthenticatorTotp gets an existing StageAuthenticatorTotp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageAuthenticatorTotp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageAuthenticatorTotpState, opts ...pulumi.ResourceOption) (*StageAuthenticatorTotp, error) {
	var resource StageAuthenticatorTotp
	err := ctx.ReadResource("authentik:index/stageAuthenticatorTotp:StageAuthenticatorTotp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageAuthenticatorTotp resources.
type stageAuthenticatorTotpState struct {
	ConfigureFlow *string `pulumi:"configureFlow"`
	// Allowed values: - `6` - `8`
	Digits       *string `pulumi:"digits"`
	FriendlyName *string `pulumi:"friendlyName"`
	Name         *string `pulumi:"name"`
}

type StageAuthenticatorTotpState struct {
	ConfigureFlow pulumi.StringPtrInput
	// Allowed values: - `6` - `8`
	Digits       pulumi.StringPtrInput
	FriendlyName pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
}

func (StageAuthenticatorTotpState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageAuthenticatorTotpState)(nil)).Elem()
}

type stageAuthenticatorTotpArgs struct {
	ConfigureFlow *string `pulumi:"configureFlow"`
	// Allowed values: - `6` - `8`
	Digits       *string `pulumi:"digits"`
	FriendlyName *string `pulumi:"friendlyName"`
	Name         *string `pulumi:"name"`
}

// The set of arguments for constructing a StageAuthenticatorTotp resource.
type StageAuthenticatorTotpArgs struct {
	ConfigureFlow pulumi.StringPtrInput
	// Allowed values: - `6` - `8`
	Digits       pulumi.StringPtrInput
	FriendlyName pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
}

func (StageAuthenticatorTotpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageAuthenticatorTotpArgs)(nil)).Elem()
}

type StageAuthenticatorTotpInput interface {
	pulumi.Input

	ToStageAuthenticatorTotpOutput() StageAuthenticatorTotpOutput
	ToStageAuthenticatorTotpOutputWithContext(ctx context.Context) StageAuthenticatorTotpOutput
}

func (*StageAuthenticatorTotp) ElementType() reflect.Type {
	return reflect.TypeOf((**StageAuthenticatorTotp)(nil)).Elem()
}

func (i *StageAuthenticatorTotp) ToStageAuthenticatorTotpOutput() StageAuthenticatorTotpOutput {
	return i.ToStageAuthenticatorTotpOutputWithContext(context.Background())
}

func (i *StageAuthenticatorTotp) ToStageAuthenticatorTotpOutputWithContext(ctx context.Context) StageAuthenticatorTotpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorTotpOutput)
}

// StageAuthenticatorTotpArrayInput is an input type that accepts StageAuthenticatorTotpArray and StageAuthenticatorTotpArrayOutput values.
// You can construct a concrete instance of `StageAuthenticatorTotpArrayInput` via:
//
//	StageAuthenticatorTotpArray{ StageAuthenticatorTotpArgs{...} }
type StageAuthenticatorTotpArrayInput interface {
	pulumi.Input

	ToStageAuthenticatorTotpArrayOutput() StageAuthenticatorTotpArrayOutput
	ToStageAuthenticatorTotpArrayOutputWithContext(context.Context) StageAuthenticatorTotpArrayOutput
}

type StageAuthenticatorTotpArray []StageAuthenticatorTotpInput

func (StageAuthenticatorTotpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageAuthenticatorTotp)(nil)).Elem()
}

func (i StageAuthenticatorTotpArray) ToStageAuthenticatorTotpArrayOutput() StageAuthenticatorTotpArrayOutput {
	return i.ToStageAuthenticatorTotpArrayOutputWithContext(context.Background())
}

func (i StageAuthenticatorTotpArray) ToStageAuthenticatorTotpArrayOutputWithContext(ctx context.Context) StageAuthenticatorTotpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorTotpArrayOutput)
}

// StageAuthenticatorTotpMapInput is an input type that accepts StageAuthenticatorTotpMap and StageAuthenticatorTotpMapOutput values.
// You can construct a concrete instance of `StageAuthenticatorTotpMapInput` via:
//
//	StageAuthenticatorTotpMap{ "key": StageAuthenticatorTotpArgs{...} }
type StageAuthenticatorTotpMapInput interface {
	pulumi.Input

	ToStageAuthenticatorTotpMapOutput() StageAuthenticatorTotpMapOutput
	ToStageAuthenticatorTotpMapOutputWithContext(context.Context) StageAuthenticatorTotpMapOutput
}

type StageAuthenticatorTotpMap map[string]StageAuthenticatorTotpInput

func (StageAuthenticatorTotpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageAuthenticatorTotp)(nil)).Elem()
}

func (i StageAuthenticatorTotpMap) ToStageAuthenticatorTotpMapOutput() StageAuthenticatorTotpMapOutput {
	return i.ToStageAuthenticatorTotpMapOutputWithContext(context.Background())
}

func (i StageAuthenticatorTotpMap) ToStageAuthenticatorTotpMapOutputWithContext(ctx context.Context) StageAuthenticatorTotpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorTotpMapOutput)
}

type StageAuthenticatorTotpOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorTotpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageAuthenticatorTotp)(nil)).Elem()
}

func (o StageAuthenticatorTotpOutput) ToStageAuthenticatorTotpOutput() StageAuthenticatorTotpOutput {
	return o
}

func (o StageAuthenticatorTotpOutput) ToStageAuthenticatorTotpOutputWithContext(ctx context.Context) StageAuthenticatorTotpOutput {
	return o
}

func (o StageAuthenticatorTotpOutput) ConfigureFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorTotp) pulumi.StringPtrOutput { return v.ConfigureFlow }).(pulumi.StringPtrOutput)
}

// Allowed values: - `6` - `8`
func (o StageAuthenticatorTotpOutput) Digits() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorTotp) pulumi.StringPtrOutput { return v.Digits }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorTotpOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorTotp) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorTotpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageAuthenticatorTotp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type StageAuthenticatorTotpArrayOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorTotpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageAuthenticatorTotp)(nil)).Elem()
}

func (o StageAuthenticatorTotpArrayOutput) ToStageAuthenticatorTotpArrayOutput() StageAuthenticatorTotpArrayOutput {
	return o
}

func (o StageAuthenticatorTotpArrayOutput) ToStageAuthenticatorTotpArrayOutputWithContext(ctx context.Context) StageAuthenticatorTotpArrayOutput {
	return o
}

func (o StageAuthenticatorTotpArrayOutput) Index(i pulumi.IntInput) StageAuthenticatorTotpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageAuthenticatorTotp {
		return vs[0].([]*StageAuthenticatorTotp)[vs[1].(int)]
	}).(StageAuthenticatorTotpOutput)
}

type StageAuthenticatorTotpMapOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorTotpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageAuthenticatorTotp)(nil)).Elem()
}

func (o StageAuthenticatorTotpMapOutput) ToStageAuthenticatorTotpMapOutput() StageAuthenticatorTotpMapOutput {
	return o
}

func (o StageAuthenticatorTotpMapOutput) ToStageAuthenticatorTotpMapOutputWithContext(ctx context.Context) StageAuthenticatorTotpMapOutput {
	return o
}

func (o StageAuthenticatorTotpMapOutput) MapIndex(k pulumi.StringInput) StageAuthenticatorTotpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageAuthenticatorTotp {
		return vs[0].(map[string]*StageAuthenticatorTotp)[vs[1].(string)]
	}).(StageAuthenticatorTotpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorTotpInput)(nil)).Elem(), &StageAuthenticatorTotp{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorTotpArrayInput)(nil)).Elem(), StageAuthenticatorTotpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorTotpMapInput)(nil)).Elem(), StageAuthenticatorTotpMap{})
	pulumi.RegisterOutputType(StageAuthenticatorTotpOutput{})
	pulumi.RegisterOutputType(StageAuthenticatorTotpArrayOutput{})
	pulumi.RegisterOutputType(StageAuthenticatorTotpMapOutput{})
}

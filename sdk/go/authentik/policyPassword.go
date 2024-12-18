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
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create a password policy to require 8 chars
//			_, err := authentik.NewPolicyPassword(ctx, "name", &authentik.PolicyPasswordArgs{
//				Name:         pulumi.String("password"),
//				LengthMin:    pulumi.Int(8),
//				ErrorMessage: pulumi.String("foo"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PolicyPassword struct {
	pulumi.CustomResourceState

	AmountDigits    pulumi.IntPtrOutput `pulumi:"amountDigits"`
	AmountLowercase pulumi.IntPtrOutput `pulumi:"amountLowercase"`
	AmountSymbols   pulumi.IntPtrOutput `pulumi:"amountSymbols"`
	AmountUppercase pulumi.IntPtrOutput `pulumi:"amountUppercase"`
	// Defaults to `false`.
	CheckHaveIBeenPwned pulumi.BoolPtrOutput `pulumi:"checkHaveIBeenPwned"`
	// Defaults to `true`.
	CheckStaticRules pulumi.BoolPtrOutput `pulumi:"checkStaticRules"`
	// Defaults to `false`.
	CheckZxcvbn  pulumi.BoolPtrOutput `pulumi:"checkZxcvbn"`
	ErrorMessage pulumi.StringOutput  `pulumi:"errorMessage"`
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrOutput `pulumi:"executionLogging"`
	// Defaults to `1`.
	HibpAllowedCount pulumi.IntPtrOutput `pulumi:"hibpAllowedCount"`
	LengthMin        pulumi.IntPtrOutput `pulumi:"lengthMin"`
	Name             pulumi.StringOutput `pulumi:"name"`
	// Defaults to `password`.
	PasswordField pulumi.StringPtrOutput `pulumi:"passwordField"`
	// Defaults to `!\"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`.
	SymbolCharset pulumi.StringPtrOutput `pulumi:"symbolCharset"`
	// Defaults to `2`.
	ZxcvbnScoreThreshold pulumi.IntPtrOutput `pulumi:"zxcvbnScoreThreshold"`
}

// NewPolicyPassword registers a new resource with the given unique name, arguments, and options.
func NewPolicyPassword(ctx *pulumi.Context,
	name string, args *PolicyPasswordArgs, opts ...pulumi.ResourceOption) (*PolicyPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ErrorMessage == nil {
		return nil, errors.New("invalid value for required argument 'ErrorMessage'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyPassword
	err := ctx.RegisterResource("authentik:index/policyPassword:PolicyPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyPassword gets an existing PolicyPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyPasswordState, opts ...pulumi.ResourceOption) (*PolicyPassword, error) {
	var resource PolicyPassword
	err := ctx.ReadResource("authentik:index/policyPassword:PolicyPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyPassword resources.
type policyPasswordState struct {
	AmountDigits    *int `pulumi:"amountDigits"`
	AmountLowercase *int `pulumi:"amountLowercase"`
	AmountSymbols   *int `pulumi:"amountSymbols"`
	AmountUppercase *int `pulumi:"amountUppercase"`
	// Defaults to `false`.
	CheckHaveIBeenPwned *bool `pulumi:"checkHaveIBeenPwned"`
	// Defaults to `true`.
	CheckStaticRules *bool `pulumi:"checkStaticRules"`
	// Defaults to `false`.
	CheckZxcvbn  *bool   `pulumi:"checkZxcvbn"`
	ErrorMessage *string `pulumi:"errorMessage"`
	// Defaults to `false`.
	ExecutionLogging *bool `pulumi:"executionLogging"`
	// Defaults to `1`.
	HibpAllowedCount *int    `pulumi:"hibpAllowedCount"`
	LengthMin        *int    `pulumi:"lengthMin"`
	Name             *string `pulumi:"name"`
	// Defaults to `password`.
	PasswordField *string `pulumi:"passwordField"`
	// Defaults to `!\"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`.
	SymbolCharset *string `pulumi:"symbolCharset"`
	// Defaults to `2`.
	ZxcvbnScoreThreshold *int `pulumi:"zxcvbnScoreThreshold"`
}

type PolicyPasswordState struct {
	AmountDigits    pulumi.IntPtrInput
	AmountLowercase pulumi.IntPtrInput
	AmountSymbols   pulumi.IntPtrInput
	AmountUppercase pulumi.IntPtrInput
	// Defaults to `false`.
	CheckHaveIBeenPwned pulumi.BoolPtrInput
	// Defaults to `true`.
	CheckStaticRules pulumi.BoolPtrInput
	// Defaults to `false`.
	CheckZxcvbn  pulumi.BoolPtrInput
	ErrorMessage pulumi.StringPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	// Defaults to `1`.
	HibpAllowedCount pulumi.IntPtrInput
	LengthMin        pulumi.IntPtrInput
	Name             pulumi.StringPtrInput
	// Defaults to `password`.
	PasswordField pulumi.StringPtrInput
	// Defaults to `!\"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`.
	SymbolCharset pulumi.StringPtrInput
	// Defaults to `2`.
	ZxcvbnScoreThreshold pulumi.IntPtrInput
}

func (PolicyPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyPasswordState)(nil)).Elem()
}

type policyPasswordArgs struct {
	AmountDigits    *int `pulumi:"amountDigits"`
	AmountLowercase *int `pulumi:"amountLowercase"`
	AmountSymbols   *int `pulumi:"amountSymbols"`
	AmountUppercase *int `pulumi:"amountUppercase"`
	// Defaults to `false`.
	CheckHaveIBeenPwned *bool `pulumi:"checkHaveIBeenPwned"`
	// Defaults to `true`.
	CheckStaticRules *bool `pulumi:"checkStaticRules"`
	// Defaults to `false`.
	CheckZxcvbn  *bool  `pulumi:"checkZxcvbn"`
	ErrorMessage string `pulumi:"errorMessage"`
	// Defaults to `false`.
	ExecutionLogging *bool `pulumi:"executionLogging"`
	// Defaults to `1`.
	HibpAllowedCount *int    `pulumi:"hibpAllowedCount"`
	LengthMin        *int    `pulumi:"lengthMin"`
	Name             *string `pulumi:"name"`
	// Defaults to `password`.
	PasswordField *string `pulumi:"passwordField"`
	// Defaults to `!\"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`.
	SymbolCharset *string `pulumi:"symbolCharset"`
	// Defaults to `2`.
	ZxcvbnScoreThreshold *int `pulumi:"zxcvbnScoreThreshold"`
}

// The set of arguments for constructing a PolicyPassword resource.
type PolicyPasswordArgs struct {
	AmountDigits    pulumi.IntPtrInput
	AmountLowercase pulumi.IntPtrInput
	AmountSymbols   pulumi.IntPtrInput
	AmountUppercase pulumi.IntPtrInput
	// Defaults to `false`.
	CheckHaveIBeenPwned pulumi.BoolPtrInput
	// Defaults to `true`.
	CheckStaticRules pulumi.BoolPtrInput
	// Defaults to `false`.
	CheckZxcvbn  pulumi.BoolPtrInput
	ErrorMessage pulumi.StringInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	// Defaults to `1`.
	HibpAllowedCount pulumi.IntPtrInput
	LengthMin        pulumi.IntPtrInput
	Name             pulumi.StringPtrInput
	// Defaults to `password`.
	PasswordField pulumi.StringPtrInput
	// Defaults to `!\"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`.
	SymbolCharset pulumi.StringPtrInput
	// Defaults to `2`.
	ZxcvbnScoreThreshold pulumi.IntPtrInput
}

func (PolicyPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyPasswordArgs)(nil)).Elem()
}

type PolicyPasswordInput interface {
	pulumi.Input

	ToPolicyPasswordOutput() PolicyPasswordOutput
	ToPolicyPasswordOutputWithContext(ctx context.Context) PolicyPasswordOutput
}

func (*PolicyPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPassword)(nil)).Elem()
}

func (i *PolicyPassword) ToPolicyPasswordOutput() PolicyPasswordOutput {
	return i.ToPolicyPasswordOutputWithContext(context.Background())
}

func (i *PolicyPassword) ToPolicyPasswordOutputWithContext(ctx context.Context) PolicyPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPasswordOutput)
}

// PolicyPasswordArrayInput is an input type that accepts PolicyPasswordArray and PolicyPasswordArrayOutput values.
// You can construct a concrete instance of `PolicyPasswordArrayInput` via:
//
//	PolicyPasswordArray{ PolicyPasswordArgs{...} }
type PolicyPasswordArrayInput interface {
	pulumi.Input

	ToPolicyPasswordArrayOutput() PolicyPasswordArrayOutput
	ToPolicyPasswordArrayOutputWithContext(context.Context) PolicyPasswordArrayOutput
}

type PolicyPasswordArray []PolicyPasswordInput

func (PolicyPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyPassword)(nil)).Elem()
}

func (i PolicyPasswordArray) ToPolicyPasswordArrayOutput() PolicyPasswordArrayOutput {
	return i.ToPolicyPasswordArrayOutputWithContext(context.Background())
}

func (i PolicyPasswordArray) ToPolicyPasswordArrayOutputWithContext(ctx context.Context) PolicyPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPasswordArrayOutput)
}

// PolicyPasswordMapInput is an input type that accepts PolicyPasswordMap and PolicyPasswordMapOutput values.
// You can construct a concrete instance of `PolicyPasswordMapInput` via:
//
//	PolicyPasswordMap{ "key": PolicyPasswordArgs{...} }
type PolicyPasswordMapInput interface {
	pulumi.Input

	ToPolicyPasswordMapOutput() PolicyPasswordMapOutput
	ToPolicyPasswordMapOutputWithContext(context.Context) PolicyPasswordMapOutput
}

type PolicyPasswordMap map[string]PolicyPasswordInput

func (PolicyPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyPassword)(nil)).Elem()
}

func (i PolicyPasswordMap) ToPolicyPasswordMapOutput() PolicyPasswordMapOutput {
	return i.ToPolicyPasswordMapOutputWithContext(context.Background())
}

func (i PolicyPasswordMap) ToPolicyPasswordMapOutputWithContext(ctx context.Context) PolicyPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPasswordMapOutput)
}

type PolicyPasswordOutput struct{ *pulumi.OutputState }

func (PolicyPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPassword)(nil)).Elem()
}

func (o PolicyPasswordOutput) ToPolicyPasswordOutput() PolicyPasswordOutput {
	return o
}

func (o PolicyPasswordOutput) ToPolicyPasswordOutputWithContext(ctx context.Context) PolicyPasswordOutput {
	return o
}

func (o PolicyPasswordOutput) AmountDigits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.AmountDigits }).(pulumi.IntPtrOutput)
}

func (o PolicyPasswordOutput) AmountLowercase() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.AmountLowercase }).(pulumi.IntPtrOutput)
}

func (o PolicyPasswordOutput) AmountSymbols() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.AmountSymbols }).(pulumi.IntPtrOutput)
}

func (o PolicyPasswordOutput) AmountUppercase() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.AmountUppercase }).(pulumi.IntPtrOutput)
}

// Defaults to `false`.
func (o PolicyPasswordOutput) CheckHaveIBeenPwned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.BoolPtrOutput { return v.CheckHaveIBeenPwned }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o PolicyPasswordOutput) CheckStaticRules() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.BoolPtrOutput { return v.CheckStaticRules }).(pulumi.BoolPtrOutput)
}

// Defaults to `false`.
func (o PolicyPasswordOutput) CheckZxcvbn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.BoolPtrOutput { return v.CheckZxcvbn }).(pulumi.BoolPtrOutput)
}

func (o PolicyPasswordOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// Defaults to `false`.
func (o PolicyPasswordOutput) ExecutionLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.BoolPtrOutput { return v.ExecutionLogging }).(pulumi.BoolPtrOutput)
}

// Defaults to `1`.
func (o PolicyPasswordOutput) HibpAllowedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.HibpAllowedCount }).(pulumi.IntPtrOutput)
}

func (o PolicyPasswordOutput) LengthMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.LengthMin }).(pulumi.IntPtrOutput)
}

func (o PolicyPasswordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defaults to `password`.
func (o PolicyPasswordOutput) PasswordField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.StringPtrOutput { return v.PasswordField }).(pulumi.StringPtrOutput)
}

// Defaults to `!\"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`.
func (o PolicyPasswordOutput) SymbolCharset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.StringPtrOutput { return v.SymbolCharset }).(pulumi.StringPtrOutput)
}

// Defaults to `2`.
func (o PolicyPasswordOutput) ZxcvbnScoreThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyPassword) pulumi.IntPtrOutput { return v.ZxcvbnScoreThreshold }).(pulumi.IntPtrOutput)
}

type PolicyPasswordArrayOutput struct{ *pulumi.OutputState }

func (PolicyPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyPassword)(nil)).Elem()
}

func (o PolicyPasswordArrayOutput) ToPolicyPasswordArrayOutput() PolicyPasswordArrayOutput {
	return o
}

func (o PolicyPasswordArrayOutput) ToPolicyPasswordArrayOutputWithContext(ctx context.Context) PolicyPasswordArrayOutput {
	return o
}

func (o PolicyPasswordArrayOutput) Index(i pulumi.IntInput) PolicyPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyPassword {
		return vs[0].([]*PolicyPassword)[vs[1].(int)]
	}).(PolicyPasswordOutput)
}

type PolicyPasswordMapOutput struct{ *pulumi.OutputState }

func (PolicyPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyPassword)(nil)).Elem()
}

func (o PolicyPasswordMapOutput) ToPolicyPasswordMapOutput() PolicyPasswordMapOutput {
	return o
}

func (o PolicyPasswordMapOutput) ToPolicyPasswordMapOutputWithContext(ctx context.Context) PolicyPasswordMapOutput {
	return o
}

func (o PolicyPasswordMapOutput) MapIndex(k pulumi.StringInput) PolicyPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyPassword {
		return vs[0].(map[string]*PolicyPassword)[vs[1].(string)]
	}).(PolicyPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyPasswordInput)(nil)).Elem(), &PolicyPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyPasswordArrayInput)(nil)).Elem(), PolicyPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyPasswordMapInput)(nil)).Elem(), PolicyPasswordMap{})
	pulumi.RegisterOutputType(PolicyPasswordOutput{})
	pulumi.RegisterOutputType(PolicyPasswordArrayOutput{})
	pulumi.RegisterOutputType(PolicyPasswordMapOutput{})
}

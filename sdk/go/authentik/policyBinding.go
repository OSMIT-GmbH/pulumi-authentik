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
type PolicyBinding struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Defaults to `false`.
	FailureResult pulumi.BoolPtrOutput `pulumi:"failureResult"`
	// UUID of the group
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// Defaults to `false`.
	Negate pulumi.BoolPtrOutput `pulumi:"negate"`
	Order  pulumi.IntOutput     `pulumi:"order"`
	// UUID of the policy
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// ID of the object this binding should apply to
	Target pulumi.StringOutput `pulumi:"target"`
	// Defaults to `30`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// PK of the user
	User pulumi.IntPtrOutput `pulumi:"user"`
}

// NewPolicyBinding registers a new resource with the given unique name, arguments, and options.
func NewPolicyBinding(ctx *pulumi.Context,
	name string, args *PolicyBindingArgs, opts ...pulumi.ResourceOption) (*PolicyBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyBinding
	err := ctx.RegisterResource("authentik:index/policyBinding:PolicyBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyBinding gets an existing PolicyBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyBindingState, opts ...pulumi.ResourceOption) (*PolicyBinding, error) {
	var resource PolicyBinding
	err := ctx.ReadResource("authentik:index/policyBinding:PolicyBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyBinding resources.
type policyBindingState struct {
	// Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Defaults to `false`.
	FailureResult *bool `pulumi:"failureResult"`
	// UUID of the group
	Group *string `pulumi:"group"`
	// Defaults to `false`.
	Negate *bool `pulumi:"negate"`
	Order  *int  `pulumi:"order"`
	// UUID of the policy
	Policy *string `pulumi:"policy"`
	// ID of the object this binding should apply to
	Target *string `pulumi:"target"`
	// Defaults to `30`.
	Timeout *int `pulumi:"timeout"`
	// PK of the user
	User *int `pulumi:"user"`
}

type PolicyBindingState struct {
	// Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Defaults to `false`.
	FailureResult pulumi.BoolPtrInput
	// UUID of the group
	Group pulumi.StringPtrInput
	// Defaults to `false`.
	Negate pulumi.BoolPtrInput
	Order  pulumi.IntPtrInput
	// UUID of the policy
	Policy pulumi.StringPtrInput
	// ID of the object this binding should apply to
	Target pulumi.StringPtrInput
	// Defaults to `30`.
	Timeout pulumi.IntPtrInput
	// PK of the user
	User pulumi.IntPtrInput
}

func (PolicyBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyBindingState)(nil)).Elem()
}

type policyBindingArgs struct {
	// Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Defaults to `false`.
	FailureResult *bool `pulumi:"failureResult"`
	// UUID of the group
	Group *string `pulumi:"group"`
	// Defaults to `false`.
	Negate *bool `pulumi:"negate"`
	Order  int   `pulumi:"order"`
	// UUID of the policy
	Policy *string `pulumi:"policy"`
	// ID of the object this binding should apply to
	Target string `pulumi:"target"`
	// Defaults to `30`.
	Timeout *int `pulumi:"timeout"`
	// PK of the user
	User *int `pulumi:"user"`
}

// The set of arguments for constructing a PolicyBinding resource.
type PolicyBindingArgs struct {
	// Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Defaults to `false`.
	FailureResult pulumi.BoolPtrInput
	// UUID of the group
	Group pulumi.StringPtrInput
	// Defaults to `false`.
	Negate pulumi.BoolPtrInput
	Order  pulumi.IntInput
	// UUID of the policy
	Policy pulumi.StringPtrInput
	// ID of the object this binding should apply to
	Target pulumi.StringInput
	// Defaults to `30`.
	Timeout pulumi.IntPtrInput
	// PK of the user
	User pulumi.IntPtrInput
}

func (PolicyBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyBindingArgs)(nil)).Elem()
}

type PolicyBindingInput interface {
	pulumi.Input

	ToPolicyBindingOutput() PolicyBindingOutput
	ToPolicyBindingOutputWithContext(ctx context.Context) PolicyBindingOutput
}

func (*PolicyBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyBinding)(nil)).Elem()
}

func (i *PolicyBinding) ToPolicyBindingOutput() PolicyBindingOutput {
	return i.ToPolicyBindingOutputWithContext(context.Background())
}

func (i *PolicyBinding) ToPolicyBindingOutputWithContext(ctx context.Context) PolicyBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyBindingOutput)
}

// PolicyBindingArrayInput is an input type that accepts PolicyBindingArray and PolicyBindingArrayOutput values.
// You can construct a concrete instance of `PolicyBindingArrayInput` via:
//
//	PolicyBindingArray{ PolicyBindingArgs{...} }
type PolicyBindingArrayInput interface {
	pulumi.Input

	ToPolicyBindingArrayOutput() PolicyBindingArrayOutput
	ToPolicyBindingArrayOutputWithContext(context.Context) PolicyBindingArrayOutput
}

type PolicyBindingArray []PolicyBindingInput

func (PolicyBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyBinding)(nil)).Elem()
}

func (i PolicyBindingArray) ToPolicyBindingArrayOutput() PolicyBindingArrayOutput {
	return i.ToPolicyBindingArrayOutputWithContext(context.Background())
}

func (i PolicyBindingArray) ToPolicyBindingArrayOutputWithContext(ctx context.Context) PolicyBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyBindingArrayOutput)
}

// PolicyBindingMapInput is an input type that accepts PolicyBindingMap and PolicyBindingMapOutput values.
// You can construct a concrete instance of `PolicyBindingMapInput` via:
//
//	PolicyBindingMap{ "key": PolicyBindingArgs{...} }
type PolicyBindingMapInput interface {
	pulumi.Input

	ToPolicyBindingMapOutput() PolicyBindingMapOutput
	ToPolicyBindingMapOutputWithContext(context.Context) PolicyBindingMapOutput
}

type PolicyBindingMap map[string]PolicyBindingInput

func (PolicyBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyBinding)(nil)).Elem()
}

func (i PolicyBindingMap) ToPolicyBindingMapOutput() PolicyBindingMapOutput {
	return i.ToPolicyBindingMapOutputWithContext(context.Background())
}

func (i PolicyBindingMap) ToPolicyBindingMapOutputWithContext(ctx context.Context) PolicyBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyBindingMapOutput)
}

type PolicyBindingOutput struct{ *pulumi.OutputState }

func (PolicyBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyBinding)(nil)).Elem()
}

func (o PolicyBindingOutput) ToPolicyBindingOutput() PolicyBindingOutput {
	return o
}

func (o PolicyBindingOutput) ToPolicyBindingOutputWithContext(ctx context.Context) PolicyBindingOutput {
	return o
}

// Defaults to `true`.
func (o PolicyBindingOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Defaults to `false`.
func (o PolicyBindingOutput) FailureResult() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.BoolPtrOutput { return v.FailureResult }).(pulumi.BoolPtrOutput)
}

// UUID of the group
func (o PolicyBindingOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o PolicyBindingOutput) Negate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.BoolPtrOutput { return v.Negate }).(pulumi.BoolPtrOutput)
}

func (o PolicyBindingOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// UUID of the policy
func (o PolicyBindingOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// ID of the object this binding should apply to
func (o PolicyBindingOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

// Defaults to `30`.
func (o PolicyBindingOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// PK of the user
func (o PolicyBindingOutput) User() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyBinding) pulumi.IntPtrOutput { return v.User }).(pulumi.IntPtrOutput)
}

type PolicyBindingArrayOutput struct{ *pulumi.OutputState }

func (PolicyBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyBinding)(nil)).Elem()
}

func (o PolicyBindingArrayOutput) ToPolicyBindingArrayOutput() PolicyBindingArrayOutput {
	return o
}

func (o PolicyBindingArrayOutput) ToPolicyBindingArrayOutputWithContext(ctx context.Context) PolicyBindingArrayOutput {
	return o
}

func (o PolicyBindingArrayOutput) Index(i pulumi.IntInput) PolicyBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyBinding {
		return vs[0].([]*PolicyBinding)[vs[1].(int)]
	}).(PolicyBindingOutput)
}

type PolicyBindingMapOutput struct{ *pulumi.OutputState }

func (PolicyBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyBinding)(nil)).Elem()
}

func (o PolicyBindingMapOutput) ToPolicyBindingMapOutput() PolicyBindingMapOutput {
	return o
}

func (o PolicyBindingMapOutput) ToPolicyBindingMapOutputWithContext(ctx context.Context) PolicyBindingMapOutput {
	return o
}

func (o PolicyBindingMapOutput) MapIndex(k pulumi.StringInput) PolicyBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyBinding {
		return vs[0].(map[string]*PolicyBinding)[vs[1].(string)]
	}).(PolicyBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyBindingInput)(nil)).Elem(), &PolicyBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyBindingArrayInput)(nil)).Elem(), PolicyBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyBindingMapInput)(nil)).Elem(), PolicyBindingMap{})
	pulumi.RegisterOutputType(PolicyBindingOutput{})
	pulumi.RegisterOutputType(PolicyBindingArrayOutput{})
	pulumi.RegisterOutputType(PolicyBindingMapOutput{})
}

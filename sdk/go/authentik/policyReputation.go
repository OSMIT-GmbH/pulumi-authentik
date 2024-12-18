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
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create a reputation policy
//			_, err := authentik.NewPolicyReputation(ctx, "name", &authentik.PolicyReputationArgs{
//				Name: pulumi.String("reputation"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PolicyReputation struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	CheckIp pulumi.BoolPtrOutput `pulumi:"checkIp"`
	// Defaults to `true`.
	CheckUsername pulumi.BoolPtrOutput `pulumi:"checkUsername"`
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrOutput `pulumi:"executionLogging"`
	Name             pulumi.StringOutput  `pulumi:"name"`
	// Defaults to `10`.
	Threshold pulumi.IntPtrOutput `pulumi:"threshold"`
}

// NewPolicyReputation registers a new resource with the given unique name, arguments, and options.
func NewPolicyReputation(ctx *pulumi.Context,
	name string, args *PolicyReputationArgs, opts ...pulumi.ResourceOption) (*PolicyReputation, error) {
	if args == nil {
		args = &PolicyReputationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyReputation
	err := ctx.RegisterResource("authentik:index/policyReputation:PolicyReputation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyReputation gets an existing PolicyReputation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyReputation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyReputationState, opts ...pulumi.ResourceOption) (*PolicyReputation, error) {
	var resource PolicyReputation
	err := ctx.ReadResource("authentik:index/policyReputation:PolicyReputation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyReputation resources.
type policyReputationState struct {
	// Defaults to `true`.
	CheckIp *bool `pulumi:"checkIp"`
	// Defaults to `true`.
	CheckUsername *bool `pulumi:"checkUsername"`
	// Defaults to `false`.
	ExecutionLogging *bool   `pulumi:"executionLogging"`
	Name             *string `pulumi:"name"`
	// Defaults to `10`.
	Threshold *int `pulumi:"threshold"`
}

type PolicyReputationState struct {
	// Defaults to `true`.
	CheckIp pulumi.BoolPtrInput
	// Defaults to `true`.
	CheckUsername pulumi.BoolPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	Name             pulumi.StringPtrInput
	// Defaults to `10`.
	Threshold pulumi.IntPtrInput
}

func (PolicyReputationState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyReputationState)(nil)).Elem()
}

type policyReputationArgs struct {
	// Defaults to `true`.
	CheckIp *bool `pulumi:"checkIp"`
	// Defaults to `true`.
	CheckUsername *bool `pulumi:"checkUsername"`
	// Defaults to `false`.
	ExecutionLogging *bool   `pulumi:"executionLogging"`
	Name             *string `pulumi:"name"`
	// Defaults to `10`.
	Threshold *int `pulumi:"threshold"`
}

// The set of arguments for constructing a PolicyReputation resource.
type PolicyReputationArgs struct {
	// Defaults to `true`.
	CheckIp pulumi.BoolPtrInput
	// Defaults to `true`.
	CheckUsername pulumi.BoolPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	Name             pulumi.StringPtrInput
	// Defaults to `10`.
	Threshold pulumi.IntPtrInput
}

func (PolicyReputationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyReputationArgs)(nil)).Elem()
}

type PolicyReputationInput interface {
	pulumi.Input

	ToPolicyReputationOutput() PolicyReputationOutput
	ToPolicyReputationOutputWithContext(ctx context.Context) PolicyReputationOutput
}

func (*PolicyReputation) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyReputation)(nil)).Elem()
}

func (i *PolicyReputation) ToPolicyReputationOutput() PolicyReputationOutput {
	return i.ToPolicyReputationOutputWithContext(context.Background())
}

func (i *PolicyReputation) ToPolicyReputationOutputWithContext(ctx context.Context) PolicyReputationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyReputationOutput)
}

// PolicyReputationArrayInput is an input type that accepts PolicyReputationArray and PolicyReputationArrayOutput values.
// You can construct a concrete instance of `PolicyReputationArrayInput` via:
//
//	PolicyReputationArray{ PolicyReputationArgs{...} }
type PolicyReputationArrayInput interface {
	pulumi.Input

	ToPolicyReputationArrayOutput() PolicyReputationArrayOutput
	ToPolicyReputationArrayOutputWithContext(context.Context) PolicyReputationArrayOutput
}

type PolicyReputationArray []PolicyReputationInput

func (PolicyReputationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyReputation)(nil)).Elem()
}

func (i PolicyReputationArray) ToPolicyReputationArrayOutput() PolicyReputationArrayOutput {
	return i.ToPolicyReputationArrayOutputWithContext(context.Background())
}

func (i PolicyReputationArray) ToPolicyReputationArrayOutputWithContext(ctx context.Context) PolicyReputationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyReputationArrayOutput)
}

// PolicyReputationMapInput is an input type that accepts PolicyReputationMap and PolicyReputationMapOutput values.
// You can construct a concrete instance of `PolicyReputationMapInput` via:
//
//	PolicyReputationMap{ "key": PolicyReputationArgs{...} }
type PolicyReputationMapInput interface {
	pulumi.Input

	ToPolicyReputationMapOutput() PolicyReputationMapOutput
	ToPolicyReputationMapOutputWithContext(context.Context) PolicyReputationMapOutput
}

type PolicyReputationMap map[string]PolicyReputationInput

func (PolicyReputationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyReputation)(nil)).Elem()
}

func (i PolicyReputationMap) ToPolicyReputationMapOutput() PolicyReputationMapOutput {
	return i.ToPolicyReputationMapOutputWithContext(context.Background())
}

func (i PolicyReputationMap) ToPolicyReputationMapOutputWithContext(ctx context.Context) PolicyReputationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyReputationMapOutput)
}

type PolicyReputationOutput struct{ *pulumi.OutputState }

func (PolicyReputationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyReputation)(nil)).Elem()
}

func (o PolicyReputationOutput) ToPolicyReputationOutput() PolicyReputationOutput {
	return o
}

func (o PolicyReputationOutput) ToPolicyReputationOutputWithContext(ctx context.Context) PolicyReputationOutput {
	return o
}

// Defaults to `true`.
func (o PolicyReputationOutput) CheckIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyReputation) pulumi.BoolPtrOutput { return v.CheckIp }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o PolicyReputationOutput) CheckUsername() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyReputation) pulumi.BoolPtrOutput { return v.CheckUsername }).(pulumi.BoolPtrOutput)
}

// Defaults to `false`.
func (o PolicyReputationOutput) ExecutionLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyReputation) pulumi.BoolPtrOutput { return v.ExecutionLogging }).(pulumi.BoolPtrOutput)
}

func (o PolicyReputationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyReputation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defaults to `10`.
func (o PolicyReputationOutput) Threshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyReputation) pulumi.IntPtrOutput { return v.Threshold }).(pulumi.IntPtrOutput)
}

type PolicyReputationArrayOutput struct{ *pulumi.OutputState }

func (PolicyReputationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyReputation)(nil)).Elem()
}

func (o PolicyReputationArrayOutput) ToPolicyReputationArrayOutput() PolicyReputationArrayOutput {
	return o
}

func (o PolicyReputationArrayOutput) ToPolicyReputationArrayOutputWithContext(ctx context.Context) PolicyReputationArrayOutput {
	return o
}

func (o PolicyReputationArrayOutput) Index(i pulumi.IntInput) PolicyReputationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyReputation {
		return vs[0].([]*PolicyReputation)[vs[1].(int)]
	}).(PolicyReputationOutput)
}

type PolicyReputationMapOutput struct{ *pulumi.OutputState }

func (PolicyReputationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyReputation)(nil)).Elem()
}

func (o PolicyReputationMapOutput) ToPolicyReputationMapOutput() PolicyReputationMapOutput {
	return o
}

func (o PolicyReputationMapOutput) ToPolicyReputationMapOutputWithContext(ctx context.Context) PolicyReputationMapOutput {
	return o
}

func (o PolicyReputationMapOutput) MapIndex(k pulumi.StringInput) PolicyReputationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyReputation {
		return vs[0].(map[string]*PolicyReputation)[vs[1].(string)]
	}).(PolicyReputationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyReputationInput)(nil)).Elem(), &PolicyReputation{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyReputationArrayInput)(nil)).Elem(), PolicyReputationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyReputationMapInput)(nil)).Elem(), PolicyReputationMap{})
	pulumi.RegisterOutputType(PolicyReputationOutput{})
	pulumi.RegisterOutputType(PolicyReputationArrayOutput{})
	pulumi.RegisterOutputType(PolicyReputationMapOutput{})
}

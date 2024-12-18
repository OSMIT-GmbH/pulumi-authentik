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
//			// Create expiry policy
//			_, err := authentik.NewPolicyExpiry(ctx, "name", &authentik.PolicyExpiryArgs{
//				Name: pulumi.String("expiry"),
//				Days: pulumi.Int(3),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PolicyExpiry struct {
	pulumi.CustomResourceState

	Days pulumi.IntOutput `pulumi:"days"`
	// Defaults to `false`.
	DenyOnly pulumi.BoolPtrOutput `pulumi:"denyOnly"`
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrOutput `pulumi:"executionLogging"`
	Name             pulumi.StringOutput  `pulumi:"name"`
}

// NewPolicyExpiry registers a new resource with the given unique name, arguments, and options.
func NewPolicyExpiry(ctx *pulumi.Context,
	name string, args *PolicyExpiryArgs, opts ...pulumi.ResourceOption) (*PolicyExpiry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Days == nil {
		return nil, errors.New("invalid value for required argument 'Days'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyExpiry
	err := ctx.RegisterResource("authentik:index/policyExpiry:PolicyExpiry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyExpiry gets an existing PolicyExpiry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyExpiry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyExpiryState, opts ...pulumi.ResourceOption) (*PolicyExpiry, error) {
	var resource PolicyExpiry
	err := ctx.ReadResource("authentik:index/policyExpiry:PolicyExpiry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyExpiry resources.
type policyExpiryState struct {
	Days *int `pulumi:"days"`
	// Defaults to `false`.
	DenyOnly *bool `pulumi:"denyOnly"`
	// Defaults to `false`.
	ExecutionLogging *bool   `pulumi:"executionLogging"`
	Name             *string `pulumi:"name"`
}

type PolicyExpiryState struct {
	Days pulumi.IntPtrInput
	// Defaults to `false`.
	DenyOnly pulumi.BoolPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	Name             pulumi.StringPtrInput
}

func (PolicyExpiryState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyExpiryState)(nil)).Elem()
}

type policyExpiryArgs struct {
	Days int `pulumi:"days"`
	// Defaults to `false`.
	DenyOnly *bool `pulumi:"denyOnly"`
	// Defaults to `false`.
	ExecutionLogging *bool   `pulumi:"executionLogging"`
	Name             *string `pulumi:"name"`
}

// The set of arguments for constructing a PolicyExpiry resource.
type PolicyExpiryArgs struct {
	Days pulumi.IntInput
	// Defaults to `false`.
	DenyOnly pulumi.BoolPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	Name             pulumi.StringPtrInput
}

func (PolicyExpiryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyExpiryArgs)(nil)).Elem()
}

type PolicyExpiryInput interface {
	pulumi.Input

	ToPolicyExpiryOutput() PolicyExpiryOutput
	ToPolicyExpiryOutputWithContext(ctx context.Context) PolicyExpiryOutput
}

func (*PolicyExpiry) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyExpiry)(nil)).Elem()
}

func (i *PolicyExpiry) ToPolicyExpiryOutput() PolicyExpiryOutput {
	return i.ToPolicyExpiryOutputWithContext(context.Background())
}

func (i *PolicyExpiry) ToPolicyExpiryOutputWithContext(ctx context.Context) PolicyExpiryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyExpiryOutput)
}

// PolicyExpiryArrayInput is an input type that accepts PolicyExpiryArray and PolicyExpiryArrayOutput values.
// You can construct a concrete instance of `PolicyExpiryArrayInput` via:
//
//	PolicyExpiryArray{ PolicyExpiryArgs{...} }
type PolicyExpiryArrayInput interface {
	pulumi.Input

	ToPolicyExpiryArrayOutput() PolicyExpiryArrayOutput
	ToPolicyExpiryArrayOutputWithContext(context.Context) PolicyExpiryArrayOutput
}

type PolicyExpiryArray []PolicyExpiryInput

func (PolicyExpiryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyExpiry)(nil)).Elem()
}

func (i PolicyExpiryArray) ToPolicyExpiryArrayOutput() PolicyExpiryArrayOutput {
	return i.ToPolicyExpiryArrayOutputWithContext(context.Background())
}

func (i PolicyExpiryArray) ToPolicyExpiryArrayOutputWithContext(ctx context.Context) PolicyExpiryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyExpiryArrayOutput)
}

// PolicyExpiryMapInput is an input type that accepts PolicyExpiryMap and PolicyExpiryMapOutput values.
// You can construct a concrete instance of `PolicyExpiryMapInput` via:
//
//	PolicyExpiryMap{ "key": PolicyExpiryArgs{...} }
type PolicyExpiryMapInput interface {
	pulumi.Input

	ToPolicyExpiryMapOutput() PolicyExpiryMapOutput
	ToPolicyExpiryMapOutputWithContext(context.Context) PolicyExpiryMapOutput
}

type PolicyExpiryMap map[string]PolicyExpiryInput

func (PolicyExpiryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyExpiry)(nil)).Elem()
}

func (i PolicyExpiryMap) ToPolicyExpiryMapOutput() PolicyExpiryMapOutput {
	return i.ToPolicyExpiryMapOutputWithContext(context.Background())
}

func (i PolicyExpiryMap) ToPolicyExpiryMapOutputWithContext(ctx context.Context) PolicyExpiryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyExpiryMapOutput)
}

type PolicyExpiryOutput struct{ *pulumi.OutputState }

func (PolicyExpiryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyExpiry)(nil)).Elem()
}

func (o PolicyExpiryOutput) ToPolicyExpiryOutput() PolicyExpiryOutput {
	return o
}

func (o PolicyExpiryOutput) ToPolicyExpiryOutputWithContext(ctx context.Context) PolicyExpiryOutput {
	return o
}

func (o PolicyExpiryOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicyExpiry) pulumi.IntOutput { return v.Days }).(pulumi.IntOutput)
}

// Defaults to `false`.
func (o PolicyExpiryOutput) DenyOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyExpiry) pulumi.BoolPtrOutput { return v.DenyOnly }).(pulumi.BoolPtrOutput)
}

// Defaults to `false`.
func (o PolicyExpiryOutput) ExecutionLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyExpiry) pulumi.BoolPtrOutput { return v.ExecutionLogging }).(pulumi.BoolPtrOutput)
}

func (o PolicyExpiryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyExpiry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PolicyExpiryArrayOutput struct{ *pulumi.OutputState }

func (PolicyExpiryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyExpiry)(nil)).Elem()
}

func (o PolicyExpiryArrayOutput) ToPolicyExpiryArrayOutput() PolicyExpiryArrayOutput {
	return o
}

func (o PolicyExpiryArrayOutput) ToPolicyExpiryArrayOutputWithContext(ctx context.Context) PolicyExpiryArrayOutput {
	return o
}

func (o PolicyExpiryArrayOutput) Index(i pulumi.IntInput) PolicyExpiryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyExpiry {
		return vs[0].([]*PolicyExpiry)[vs[1].(int)]
	}).(PolicyExpiryOutput)
}

type PolicyExpiryMapOutput struct{ *pulumi.OutputState }

func (PolicyExpiryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyExpiry)(nil)).Elem()
}

func (o PolicyExpiryMapOutput) ToPolicyExpiryMapOutput() PolicyExpiryMapOutput {
	return o
}

func (o PolicyExpiryMapOutput) ToPolicyExpiryMapOutputWithContext(ctx context.Context) PolicyExpiryMapOutput {
	return o
}

func (o PolicyExpiryMapOutput) MapIndex(k pulumi.StringInput) PolicyExpiryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyExpiry {
		return vs[0].(map[string]*PolicyExpiry)[vs[1].(string)]
	}).(PolicyExpiryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyExpiryInput)(nil)).Elem(), &PolicyExpiry{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyExpiryArrayInput)(nil)).Elem(), PolicyExpiryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyExpiryMapInput)(nil)).Elem(), PolicyExpiryMap{})
	pulumi.RegisterOutputType(PolicyExpiryOutput{})
	pulumi.RegisterOutputType(PolicyExpiryArrayOutput{})
	pulumi.RegisterOutputType(PolicyExpiryMapOutput{})
}

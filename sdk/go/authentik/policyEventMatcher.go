// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//			_, err := authentik.NewPolicyEventMatcher(ctx, "name", &authentik.PolicyEventMatcherArgs{
//				Action:   pulumi.String("login"),
//				App:      pulumi.String("authentik.events"),
//				ClientIp: pulumi.String("1.2.3.4"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PolicyEventMatcher struct {
	pulumi.CustomResourceState

	Action   pulumi.StringPtrOutput `pulumi:"action"`
	App      pulumi.StringPtrOutput `pulumi:"app"`
	ClientIp pulumi.StringPtrOutput `pulumi:"clientIp"`
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrOutput   `pulumi:"executionLogging"`
	Model            pulumi.StringPtrOutput `pulumi:"model"`
	Name             pulumi.StringOutput    `pulumi:"name"`
}

// NewPolicyEventMatcher registers a new resource with the given unique name, arguments, and options.
func NewPolicyEventMatcher(ctx *pulumi.Context,
	name string, args *PolicyEventMatcherArgs, opts ...pulumi.ResourceOption) (*PolicyEventMatcher, error) {
	if args == nil {
		args = &PolicyEventMatcherArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyEventMatcher
	err := ctx.RegisterResource("authentik:index/policyEventMatcher:PolicyEventMatcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyEventMatcher gets an existing PolicyEventMatcher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyEventMatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyEventMatcherState, opts ...pulumi.ResourceOption) (*PolicyEventMatcher, error) {
	var resource PolicyEventMatcher
	err := ctx.ReadResource("authentik:index/policyEventMatcher:PolicyEventMatcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyEventMatcher resources.
type policyEventMatcherState struct {
	Action   *string `pulumi:"action"`
	App      *string `pulumi:"app"`
	ClientIp *string `pulumi:"clientIp"`
	// Defaults to `false`.
	ExecutionLogging *bool   `pulumi:"executionLogging"`
	Model            *string `pulumi:"model"`
	Name             *string `pulumi:"name"`
}

type PolicyEventMatcherState struct {
	Action   pulumi.StringPtrInput
	App      pulumi.StringPtrInput
	ClientIp pulumi.StringPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	Model            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
}

func (PolicyEventMatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyEventMatcherState)(nil)).Elem()
}

type policyEventMatcherArgs struct {
	Action   *string `pulumi:"action"`
	App      *string `pulumi:"app"`
	ClientIp *string `pulumi:"clientIp"`
	// Defaults to `false`.
	ExecutionLogging *bool   `pulumi:"executionLogging"`
	Model            *string `pulumi:"model"`
	Name             *string `pulumi:"name"`
}

// The set of arguments for constructing a PolicyEventMatcher resource.
type PolicyEventMatcherArgs struct {
	Action   pulumi.StringPtrInput
	App      pulumi.StringPtrInput
	ClientIp pulumi.StringPtrInput
	// Defaults to `false`.
	ExecutionLogging pulumi.BoolPtrInput
	Model            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
}

func (PolicyEventMatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyEventMatcherArgs)(nil)).Elem()
}

type PolicyEventMatcherInput interface {
	pulumi.Input

	ToPolicyEventMatcherOutput() PolicyEventMatcherOutput
	ToPolicyEventMatcherOutputWithContext(ctx context.Context) PolicyEventMatcherOutput
}

func (*PolicyEventMatcher) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyEventMatcher)(nil)).Elem()
}

func (i *PolicyEventMatcher) ToPolicyEventMatcherOutput() PolicyEventMatcherOutput {
	return i.ToPolicyEventMatcherOutputWithContext(context.Background())
}

func (i *PolicyEventMatcher) ToPolicyEventMatcherOutputWithContext(ctx context.Context) PolicyEventMatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyEventMatcherOutput)
}

func (i *PolicyEventMatcher) ToOutput(ctx context.Context) pulumix.Output[*PolicyEventMatcher] {
	return pulumix.Output[*PolicyEventMatcher]{
		OutputState: i.ToPolicyEventMatcherOutputWithContext(ctx).OutputState,
	}
}

// PolicyEventMatcherArrayInput is an input type that accepts PolicyEventMatcherArray and PolicyEventMatcherArrayOutput values.
// You can construct a concrete instance of `PolicyEventMatcherArrayInput` via:
//
//	PolicyEventMatcherArray{ PolicyEventMatcherArgs{...} }
type PolicyEventMatcherArrayInput interface {
	pulumi.Input

	ToPolicyEventMatcherArrayOutput() PolicyEventMatcherArrayOutput
	ToPolicyEventMatcherArrayOutputWithContext(context.Context) PolicyEventMatcherArrayOutput
}

type PolicyEventMatcherArray []PolicyEventMatcherInput

func (PolicyEventMatcherArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyEventMatcher)(nil)).Elem()
}

func (i PolicyEventMatcherArray) ToPolicyEventMatcherArrayOutput() PolicyEventMatcherArrayOutput {
	return i.ToPolicyEventMatcherArrayOutputWithContext(context.Background())
}

func (i PolicyEventMatcherArray) ToPolicyEventMatcherArrayOutputWithContext(ctx context.Context) PolicyEventMatcherArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyEventMatcherArrayOutput)
}

func (i PolicyEventMatcherArray) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyEventMatcher] {
	return pulumix.Output[[]*PolicyEventMatcher]{
		OutputState: i.ToPolicyEventMatcherArrayOutputWithContext(ctx).OutputState,
	}
}

// PolicyEventMatcherMapInput is an input type that accepts PolicyEventMatcherMap and PolicyEventMatcherMapOutput values.
// You can construct a concrete instance of `PolicyEventMatcherMapInput` via:
//
//	PolicyEventMatcherMap{ "key": PolicyEventMatcherArgs{...} }
type PolicyEventMatcherMapInput interface {
	pulumi.Input

	ToPolicyEventMatcherMapOutput() PolicyEventMatcherMapOutput
	ToPolicyEventMatcherMapOutputWithContext(context.Context) PolicyEventMatcherMapOutput
}

type PolicyEventMatcherMap map[string]PolicyEventMatcherInput

func (PolicyEventMatcherMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyEventMatcher)(nil)).Elem()
}

func (i PolicyEventMatcherMap) ToPolicyEventMatcherMapOutput() PolicyEventMatcherMapOutput {
	return i.ToPolicyEventMatcherMapOutputWithContext(context.Background())
}

func (i PolicyEventMatcherMap) ToPolicyEventMatcherMapOutputWithContext(ctx context.Context) PolicyEventMatcherMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyEventMatcherMapOutput)
}

func (i PolicyEventMatcherMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyEventMatcher] {
	return pulumix.Output[map[string]*PolicyEventMatcher]{
		OutputState: i.ToPolicyEventMatcherMapOutputWithContext(ctx).OutputState,
	}
}

type PolicyEventMatcherOutput struct{ *pulumi.OutputState }

func (PolicyEventMatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyEventMatcher)(nil)).Elem()
}

func (o PolicyEventMatcherOutput) ToPolicyEventMatcherOutput() PolicyEventMatcherOutput {
	return o
}

func (o PolicyEventMatcherOutput) ToPolicyEventMatcherOutputWithContext(ctx context.Context) PolicyEventMatcherOutput {
	return o
}

func (o PolicyEventMatcherOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyEventMatcher] {
	return pulumix.Output[*PolicyEventMatcher]{
		OutputState: o.OutputState,
	}
}

func (o PolicyEventMatcherOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyEventMatcher) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PolicyEventMatcherOutput) App() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyEventMatcher) pulumi.StringPtrOutput { return v.App }).(pulumi.StringPtrOutput)
}

func (o PolicyEventMatcherOutput) ClientIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyEventMatcher) pulumi.StringPtrOutput { return v.ClientIp }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o PolicyEventMatcherOutput) ExecutionLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyEventMatcher) pulumi.BoolPtrOutput { return v.ExecutionLogging }).(pulumi.BoolPtrOutput)
}

func (o PolicyEventMatcherOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyEventMatcher) pulumi.StringPtrOutput { return v.Model }).(pulumi.StringPtrOutput)
}

func (o PolicyEventMatcherOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyEventMatcher) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PolicyEventMatcherArrayOutput struct{ *pulumi.OutputState }

func (PolicyEventMatcherArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyEventMatcher)(nil)).Elem()
}

func (o PolicyEventMatcherArrayOutput) ToPolicyEventMatcherArrayOutput() PolicyEventMatcherArrayOutput {
	return o
}

func (o PolicyEventMatcherArrayOutput) ToPolicyEventMatcherArrayOutputWithContext(ctx context.Context) PolicyEventMatcherArrayOutput {
	return o
}

func (o PolicyEventMatcherArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyEventMatcher] {
	return pulumix.Output[[]*PolicyEventMatcher]{
		OutputState: o.OutputState,
	}
}

func (o PolicyEventMatcherArrayOutput) Index(i pulumi.IntInput) PolicyEventMatcherOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyEventMatcher {
		return vs[0].([]*PolicyEventMatcher)[vs[1].(int)]
	}).(PolicyEventMatcherOutput)
}

type PolicyEventMatcherMapOutput struct{ *pulumi.OutputState }

func (PolicyEventMatcherMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyEventMatcher)(nil)).Elem()
}

func (o PolicyEventMatcherMapOutput) ToPolicyEventMatcherMapOutput() PolicyEventMatcherMapOutput {
	return o
}

func (o PolicyEventMatcherMapOutput) ToPolicyEventMatcherMapOutputWithContext(ctx context.Context) PolicyEventMatcherMapOutput {
	return o
}

func (o PolicyEventMatcherMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyEventMatcher] {
	return pulumix.Output[map[string]*PolicyEventMatcher]{
		OutputState: o.OutputState,
	}
}

func (o PolicyEventMatcherMapOutput) MapIndex(k pulumi.StringInput) PolicyEventMatcherOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyEventMatcher {
		return vs[0].(map[string]*PolicyEventMatcher)[vs[1].(string)]
	}).(PolicyEventMatcherOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyEventMatcherInput)(nil)).Elem(), &PolicyEventMatcher{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyEventMatcherArrayInput)(nil)).Elem(), PolicyEventMatcherArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyEventMatcherMapInput)(nil)).Elem(), PolicyEventMatcherMap{})
	pulumi.RegisterOutputType(PolicyEventMatcherOutput{})
	pulumi.RegisterOutputType(PolicyEventMatcherArrayOutput{})
	pulumi.RegisterOutputType(PolicyEventMatcherMapOutput{})
}

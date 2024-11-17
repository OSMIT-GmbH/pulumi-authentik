// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
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
//			// Create a static TOTP Setup stage
//			_, err := authentik.NewStageAuthenticatorStatic(ctx, "name", &authentik.StageAuthenticatorStaticArgs{
//				Name: pulumi.String("static-totp-setup"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageAuthenticatorStatic struct {
	pulumi.CustomResourceState

	ConfigureFlow pulumi.StringPtrOutput `pulumi:"configureFlow"`
	FriendlyName  pulumi.StringPtrOutput `pulumi:"friendlyName"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	// Defaults to `6`.
	TokenCount pulumi.IntPtrOutput `pulumi:"tokenCount"`
	// Defaults to `12`.
	TokenLength pulumi.IntPtrOutput `pulumi:"tokenLength"`
}

// NewStageAuthenticatorStatic registers a new resource with the given unique name, arguments, and options.
func NewStageAuthenticatorStatic(ctx *pulumi.Context,
	name string, args *StageAuthenticatorStaticArgs, opts ...pulumi.ResourceOption) (*StageAuthenticatorStatic, error) {
	if args == nil {
		args = &StageAuthenticatorStaticArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageAuthenticatorStatic
	err := ctx.RegisterResource("authentik:index/stageAuthenticatorStatic:StageAuthenticatorStatic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageAuthenticatorStatic gets an existing StageAuthenticatorStatic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageAuthenticatorStatic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageAuthenticatorStaticState, opts ...pulumi.ResourceOption) (*StageAuthenticatorStatic, error) {
	var resource StageAuthenticatorStatic
	err := ctx.ReadResource("authentik:index/stageAuthenticatorStatic:StageAuthenticatorStatic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageAuthenticatorStatic resources.
type stageAuthenticatorStaticState struct {
	ConfigureFlow *string `pulumi:"configureFlow"`
	FriendlyName  *string `pulumi:"friendlyName"`
	Name          *string `pulumi:"name"`
	// Defaults to `6`.
	TokenCount *int `pulumi:"tokenCount"`
	// Defaults to `12`.
	TokenLength *int `pulumi:"tokenLength"`
}

type StageAuthenticatorStaticState struct {
	ConfigureFlow pulumi.StringPtrInput
	FriendlyName  pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	// Defaults to `6`.
	TokenCount pulumi.IntPtrInput
	// Defaults to `12`.
	TokenLength pulumi.IntPtrInput
}

func (StageAuthenticatorStaticState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageAuthenticatorStaticState)(nil)).Elem()
}

type stageAuthenticatorStaticArgs struct {
	ConfigureFlow *string `pulumi:"configureFlow"`
	FriendlyName  *string `pulumi:"friendlyName"`
	Name          *string `pulumi:"name"`
	// Defaults to `6`.
	TokenCount *int `pulumi:"tokenCount"`
	// Defaults to `12`.
	TokenLength *int `pulumi:"tokenLength"`
}

// The set of arguments for constructing a StageAuthenticatorStatic resource.
type StageAuthenticatorStaticArgs struct {
	ConfigureFlow pulumi.StringPtrInput
	FriendlyName  pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	// Defaults to `6`.
	TokenCount pulumi.IntPtrInput
	// Defaults to `12`.
	TokenLength pulumi.IntPtrInput
}

func (StageAuthenticatorStaticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageAuthenticatorStaticArgs)(nil)).Elem()
}

type StageAuthenticatorStaticInput interface {
	pulumi.Input

	ToStageAuthenticatorStaticOutput() StageAuthenticatorStaticOutput
	ToStageAuthenticatorStaticOutputWithContext(ctx context.Context) StageAuthenticatorStaticOutput
}

func (*StageAuthenticatorStatic) ElementType() reflect.Type {
	return reflect.TypeOf((**StageAuthenticatorStatic)(nil)).Elem()
}

func (i *StageAuthenticatorStatic) ToStageAuthenticatorStaticOutput() StageAuthenticatorStaticOutput {
	return i.ToStageAuthenticatorStaticOutputWithContext(context.Background())
}

func (i *StageAuthenticatorStatic) ToStageAuthenticatorStaticOutputWithContext(ctx context.Context) StageAuthenticatorStaticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorStaticOutput)
}

// StageAuthenticatorStaticArrayInput is an input type that accepts StageAuthenticatorStaticArray and StageAuthenticatorStaticArrayOutput values.
// You can construct a concrete instance of `StageAuthenticatorStaticArrayInput` via:
//
//	StageAuthenticatorStaticArray{ StageAuthenticatorStaticArgs{...} }
type StageAuthenticatorStaticArrayInput interface {
	pulumi.Input

	ToStageAuthenticatorStaticArrayOutput() StageAuthenticatorStaticArrayOutput
	ToStageAuthenticatorStaticArrayOutputWithContext(context.Context) StageAuthenticatorStaticArrayOutput
}

type StageAuthenticatorStaticArray []StageAuthenticatorStaticInput

func (StageAuthenticatorStaticArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageAuthenticatorStatic)(nil)).Elem()
}

func (i StageAuthenticatorStaticArray) ToStageAuthenticatorStaticArrayOutput() StageAuthenticatorStaticArrayOutput {
	return i.ToStageAuthenticatorStaticArrayOutputWithContext(context.Background())
}

func (i StageAuthenticatorStaticArray) ToStageAuthenticatorStaticArrayOutputWithContext(ctx context.Context) StageAuthenticatorStaticArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorStaticArrayOutput)
}

// StageAuthenticatorStaticMapInput is an input type that accepts StageAuthenticatorStaticMap and StageAuthenticatorStaticMapOutput values.
// You can construct a concrete instance of `StageAuthenticatorStaticMapInput` via:
//
//	StageAuthenticatorStaticMap{ "key": StageAuthenticatorStaticArgs{...} }
type StageAuthenticatorStaticMapInput interface {
	pulumi.Input

	ToStageAuthenticatorStaticMapOutput() StageAuthenticatorStaticMapOutput
	ToStageAuthenticatorStaticMapOutputWithContext(context.Context) StageAuthenticatorStaticMapOutput
}

type StageAuthenticatorStaticMap map[string]StageAuthenticatorStaticInput

func (StageAuthenticatorStaticMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageAuthenticatorStatic)(nil)).Elem()
}

func (i StageAuthenticatorStaticMap) ToStageAuthenticatorStaticMapOutput() StageAuthenticatorStaticMapOutput {
	return i.ToStageAuthenticatorStaticMapOutputWithContext(context.Background())
}

func (i StageAuthenticatorStaticMap) ToStageAuthenticatorStaticMapOutputWithContext(ctx context.Context) StageAuthenticatorStaticMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageAuthenticatorStaticMapOutput)
}

type StageAuthenticatorStaticOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorStaticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageAuthenticatorStatic)(nil)).Elem()
}

func (o StageAuthenticatorStaticOutput) ToStageAuthenticatorStaticOutput() StageAuthenticatorStaticOutput {
	return o
}

func (o StageAuthenticatorStaticOutput) ToStageAuthenticatorStaticOutputWithContext(ctx context.Context) StageAuthenticatorStaticOutput {
	return o
}

func (o StageAuthenticatorStaticOutput) ConfigureFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorStatic) pulumi.StringPtrOutput { return v.ConfigureFlow }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorStaticOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorStatic) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o StageAuthenticatorStaticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageAuthenticatorStatic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defaults to `6`.
func (o StageAuthenticatorStaticOutput) TokenCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorStatic) pulumi.IntPtrOutput { return v.TokenCount }).(pulumi.IntPtrOutput)
}

// Defaults to `12`.
func (o StageAuthenticatorStaticOutput) TokenLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StageAuthenticatorStatic) pulumi.IntPtrOutput { return v.TokenLength }).(pulumi.IntPtrOutput)
}

type StageAuthenticatorStaticArrayOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorStaticArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageAuthenticatorStatic)(nil)).Elem()
}

func (o StageAuthenticatorStaticArrayOutput) ToStageAuthenticatorStaticArrayOutput() StageAuthenticatorStaticArrayOutput {
	return o
}

func (o StageAuthenticatorStaticArrayOutput) ToStageAuthenticatorStaticArrayOutputWithContext(ctx context.Context) StageAuthenticatorStaticArrayOutput {
	return o
}

func (o StageAuthenticatorStaticArrayOutput) Index(i pulumi.IntInput) StageAuthenticatorStaticOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageAuthenticatorStatic {
		return vs[0].([]*StageAuthenticatorStatic)[vs[1].(int)]
	}).(StageAuthenticatorStaticOutput)
}

type StageAuthenticatorStaticMapOutput struct{ *pulumi.OutputState }

func (StageAuthenticatorStaticMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageAuthenticatorStatic)(nil)).Elem()
}

func (o StageAuthenticatorStaticMapOutput) ToStageAuthenticatorStaticMapOutput() StageAuthenticatorStaticMapOutput {
	return o
}

func (o StageAuthenticatorStaticMapOutput) ToStageAuthenticatorStaticMapOutputWithContext(ctx context.Context) StageAuthenticatorStaticMapOutput {
	return o
}

func (o StageAuthenticatorStaticMapOutput) MapIndex(k pulumi.StringInput) StageAuthenticatorStaticOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageAuthenticatorStatic {
		return vs[0].(map[string]*StageAuthenticatorStatic)[vs[1].(string)]
	}).(StageAuthenticatorStaticOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorStaticInput)(nil)).Elem(), &StageAuthenticatorStatic{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorStaticArrayInput)(nil)).Elem(), StageAuthenticatorStaticArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageAuthenticatorStaticMapInput)(nil)).Elem(), StageAuthenticatorStaticMap{})
	pulumi.RegisterOutputType(StageAuthenticatorStaticOutput{})
	pulumi.RegisterOutputType(StageAuthenticatorStaticArrayOutput{})
	pulumi.RegisterOutputType(StageAuthenticatorStaticMapOutput{})
}

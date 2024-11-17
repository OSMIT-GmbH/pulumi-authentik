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
//	"encoding/json"
//
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"foo": "bar",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = authentik.NewBlueprint(ctx, "instance", &authentik.BlueprintArgs{
//				Path:    pulumi.String("default/flow-default-authentication-flow.yaml"),
//				Context: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Blueprint struct {
	pulumi.CustomResourceState

	Content pulumi.StringPtrOutput `pulumi:"content"`
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Context pulumi.StringPtrOutput `pulumi:"context"`
	// Defaults to `true`.
	Enabled pulumi.BoolPtrOutput   `pulumi:"enabled"`
	Name    pulumi.StringOutput    `pulumi:"name"`
	Path    pulumi.StringPtrOutput `pulumi:"path"`
}

// NewBlueprint registers a new resource with the given unique name, arguments, and options.
func NewBlueprint(ctx *pulumi.Context,
	name string, args *BlueprintArgs, opts ...pulumi.ResourceOption) (*Blueprint, error) {
	if args == nil {
		args = &BlueprintArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Blueprint
	err := ctx.RegisterResource("authentik:index/blueprint:Blueprint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlueprint gets an existing Blueprint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlueprint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlueprintState, opts ...pulumi.ResourceOption) (*Blueprint, error) {
	var resource Blueprint
	err := ctx.ReadResource("authentik:index/blueprint:Blueprint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Blueprint resources.
type blueprintState struct {
	Content *string `pulumi:"content"`
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Context *string `pulumi:"context"`
	// Defaults to `true`.
	Enabled *bool   `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
	Path    *string `pulumi:"path"`
}

type BlueprintState struct {
	Content pulumi.StringPtrInput
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Context pulumi.StringPtrInput
	// Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	Name    pulumi.StringPtrInput
	Path    pulumi.StringPtrInput
}

func (BlueprintState) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintState)(nil)).Elem()
}

type blueprintArgs struct {
	Content *string `pulumi:"content"`
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Context *string `pulumi:"context"`
	// Defaults to `true`.
	Enabled *bool   `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
	Path    *string `pulumi:"path"`
}

// The set of arguments for constructing a Blueprint resource.
type BlueprintArgs struct {
	Content pulumi.StringPtrInput
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Context pulumi.StringPtrInput
	// Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	Name    pulumi.StringPtrInput
	Path    pulumi.StringPtrInput
}

func (BlueprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintArgs)(nil)).Elem()
}

type BlueprintInput interface {
	pulumi.Input

	ToBlueprintOutput() BlueprintOutput
	ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput
}

func (*Blueprint) ElementType() reflect.Type {
	return reflect.TypeOf((**Blueprint)(nil)).Elem()
}

func (i *Blueprint) ToBlueprintOutput() BlueprintOutput {
	return i.ToBlueprintOutputWithContext(context.Background())
}

func (i *Blueprint) ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintOutput)
}

// BlueprintArrayInput is an input type that accepts BlueprintArray and BlueprintArrayOutput values.
// You can construct a concrete instance of `BlueprintArrayInput` via:
//
//	BlueprintArray{ BlueprintArgs{...} }
type BlueprintArrayInput interface {
	pulumi.Input

	ToBlueprintArrayOutput() BlueprintArrayOutput
	ToBlueprintArrayOutputWithContext(context.Context) BlueprintArrayOutput
}

type BlueprintArray []BlueprintInput

func (BlueprintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Blueprint)(nil)).Elem()
}

func (i BlueprintArray) ToBlueprintArrayOutput() BlueprintArrayOutput {
	return i.ToBlueprintArrayOutputWithContext(context.Background())
}

func (i BlueprintArray) ToBlueprintArrayOutputWithContext(ctx context.Context) BlueprintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintArrayOutput)
}

// BlueprintMapInput is an input type that accepts BlueprintMap and BlueprintMapOutput values.
// You can construct a concrete instance of `BlueprintMapInput` via:
//
//	BlueprintMap{ "key": BlueprintArgs{...} }
type BlueprintMapInput interface {
	pulumi.Input

	ToBlueprintMapOutput() BlueprintMapOutput
	ToBlueprintMapOutputWithContext(context.Context) BlueprintMapOutput
}

type BlueprintMap map[string]BlueprintInput

func (BlueprintMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Blueprint)(nil)).Elem()
}

func (i BlueprintMap) ToBlueprintMapOutput() BlueprintMapOutput {
	return i.ToBlueprintMapOutputWithContext(context.Background())
}

func (i BlueprintMap) ToBlueprintMapOutputWithContext(ctx context.Context) BlueprintMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintMapOutput)
}

type BlueprintOutput struct{ *pulumi.OutputState }

func (BlueprintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Blueprint)(nil)).Elem()
}

func (o BlueprintOutput) ToBlueprintOutput() BlueprintOutput {
	return o
}

func (o BlueprintOutput) ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput {
	return o
}

func (o BlueprintOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
func (o BlueprintOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

// Defaults to `true`.
func (o BlueprintOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BlueprintOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BlueprintOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

type BlueprintArrayOutput struct{ *pulumi.OutputState }

func (BlueprintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Blueprint)(nil)).Elem()
}

func (o BlueprintArrayOutput) ToBlueprintArrayOutput() BlueprintArrayOutput {
	return o
}

func (o BlueprintArrayOutput) ToBlueprintArrayOutputWithContext(ctx context.Context) BlueprintArrayOutput {
	return o
}

func (o BlueprintArrayOutput) Index(i pulumi.IntInput) BlueprintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Blueprint {
		return vs[0].([]*Blueprint)[vs[1].(int)]
	}).(BlueprintOutput)
}

type BlueprintMapOutput struct{ *pulumi.OutputState }

func (BlueprintMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Blueprint)(nil)).Elem()
}

func (o BlueprintMapOutput) ToBlueprintMapOutput() BlueprintMapOutput {
	return o
}

func (o BlueprintMapOutput) ToBlueprintMapOutputWithContext(ctx context.Context) BlueprintMapOutput {
	return o
}

func (o BlueprintMapOutput) MapIndex(k pulumi.StringInput) BlueprintOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Blueprint {
		return vs[0].(map[string]*Blueprint)[vs[1].(string)]
	}).(BlueprintOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlueprintInput)(nil)).Elem(), &Blueprint{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlueprintArrayInput)(nil)).Elem(), BlueprintArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlueprintMapInput)(nil)).Elem(), BlueprintMap{})
	pulumi.RegisterOutputType(BlueprintOutput{})
	pulumi.RegisterOutputType(BlueprintArrayOutput{})
	pulumi.RegisterOutputType(BlueprintMapOutput{})
}

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
//			_, err := authentik.NewStageUserWrite(ctx, "name", &authentik.StageUserWriteArgs{
//				CreateUsersAsInactive: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageUserWrite struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	CreateUsersAsInactive pulumi.BoolPtrOutput   `pulumi:"createUsersAsInactive"`
	CreateUsersGroup      pulumi.StringPtrOutput `pulumi:"createUsersGroup"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	// Defaults to `createWhenRequired`.
	UserCreationMode pulumi.StringPtrOutput `pulumi:"userCreationMode"`
	// Defaults to ``.
	UserPathTemplate pulumi.StringPtrOutput `pulumi:"userPathTemplate"`
}

// NewStageUserWrite registers a new resource with the given unique name, arguments, and options.
func NewStageUserWrite(ctx *pulumi.Context,
	name string, args *StageUserWriteArgs, opts ...pulumi.ResourceOption) (*StageUserWrite, error) {
	if args == nil {
		args = &StageUserWriteArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageUserWrite
	err := ctx.RegisterResource("authentik:index/stageUserWrite:StageUserWrite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageUserWrite gets an existing StageUserWrite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageUserWrite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageUserWriteState, opts ...pulumi.ResourceOption) (*StageUserWrite, error) {
	var resource StageUserWrite
	err := ctx.ReadResource("authentik:index/stageUserWrite:StageUserWrite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageUserWrite resources.
type stageUserWriteState struct {
	// Defaults to `true`.
	CreateUsersAsInactive *bool   `pulumi:"createUsersAsInactive"`
	CreateUsersGroup      *string `pulumi:"createUsersGroup"`
	Name                  *string `pulumi:"name"`
	// Defaults to `createWhenRequired`.
	UserCreationMode *string `pulumi:"userCreationMode"`
	// Defaults to ``.
	UserPathTemplate *string `pulumi:"userPathTemplate"`
}

type StageUserWriteState struct {
	// Defaults to `true`.
	CreateUsersAsInactive pulumi.BoolPtrInput
	CreateUsersGroup      pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	// Defaults to `createWhenRequired`.
	UserCreationMode pulumi.StringPtrInput
	// Defaults to ``.
	UserPathTemplate pulumi.StringPtrInput
}

func (StageUserWriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageUserWriteState)(nil)).Elem()
}

type stageUserWriteArgs struct {
	// Defaults to `true`.
	CreateUsersAsInactive *bool   `pulumi:"createUsersAsInactive"`
	CreateUsersGroup      *string `pulumi:"createUsersGroup"`
	Name                  *string `pulumi:"name"`
	// Defaults to `createWhenRequired`.
	UserCreationMode *string `pulumi:"userCreationMode"`
	// Defaults to ``.
	UserPathTemplate *string `pulumi:"userPathTemplate"`
}

// The set of arguments for constructing a StageUserWrite resource.
type StageUserWriteArgs struct {
	// Defaults to `true`.
	CreateUsersAsInactive pulumi.BoolPtrInput
	CreateUsersGroup      pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	// Defaults to `createWhenRequired`.
	UserCreationMode pulumi.StringPtrInput
	// Defaults to ``.
	UserPathTemplate pulumi.StringPtrInput
}

func (StageUserWriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageUserWriteArgs)(nil)).Elem()
}

type StageUserWriteInput interface {
	pulumi.Input

	ToStageUserWriteOutput() StageUserWriteOutput
	ToStageUserWriteOutputWithContext(ctx context.Context) StageUserWriteOutput
}

func (*StageUserWrite) ElementType() reflect.Type {
	return reflect.TypeOf((**StageUserWrite)(nil)).Elem()
}

func (i *StageUserWrite) ToStageUserWriteOutput() StageUserWriteOutput {
	return i.ToStageUserWriteOutputWithContext(context.Background())
}

func (i *StageUserWrite) ToStageUserWriteOutputWithContext(ctx context.Context) StageUserWriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageUserWriteOutput)
}

func (i *StageUserWrite) ToOutput(ctx context.Context) pulumix.Output[*StageUserWrite] {
	return pulumix.Output[*StageUserWrite]{
		OutputState: i.ToStageUserWriteOutputWithContext(ctx).OutputState,
	}
}

// StageUserWriteArrayInput is an input type that accepts StageUserWriteArray and StageUserWriteArrayOutput values.
// You can construct a concrete instance of `StageUserWriteArrayInput` via:
//
//	StageUserWriteArray{ StageUserWriteArgs{...} }
type StageUserWriteArrayInput interface {
	pulumi.Input

	ToStageUserWriteArrayOutput() StageUserWriteArrayOutput
	ToStageUserWriteArrayOutputWithContext(context.Context) StageUserWriteArrayOutput
}

type StageUserWriteArray []StageUserWriteInput

func (StageUserWriteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageUserWrite)(nil)).Elem()
}

func (i StageUserWriteArray) ToStageUserWriteArrayOutput() StageUserWriteArrayOutput {
	return i.ToStageUserWriteArrayOutputWithContext(context.Background())
}

func (i StageUserWriteArray) ToStageUserWriteArrayOutputWithContext(ctx context.Context) StageUserWriteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageUserWriteArrayOutput)
}

func (i StageUserWriteArray) ToOutput(ctx context.Context) pulumix.Output[[]*StageUserWrite] {
	return pulumix.Output[[]*StageUserWrite]{
		OutputState: i.ToStageUserWriteArrayOutputWithContext(ctx).OutputState,
	}
}

// StageUserWriteMapInput is an input type that accepts StageUserWriteMap and StageUserWriteMapOutput values.
// You can construct a concrete instance of `StageUserWriteMapInput` via:
//
//	StageUserWriteMap{ "key": StageUserWriteArgs{...} }
type StageUserWriteMapInput interface {
	pulumi.Input

	ToStageUserWriteMapOutput() StageUserWriteMapOutput
	ToStageUserWriteMapOutputWithContext(context.Context) StageUserWriteMapOutput
}

type StageUserWriteMap map[string]StageUserWriteInput

func (StageUserWriteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageUserWrite)(nil)).Elem()
}

func (i StageUserWriteMap) ToStageUserWriteMapOutput() StageUserWriteMapOutput {
	return i.ToStageUserWriteMapOutputWithContext(context.Background())
}

func (i StageUserWriteMap) ToStageUserWriteMapOutputWithContext(ctx context.Context) StageUserWriteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageUserWriteMapOutput)
}

func (i StageUserWriteMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*StageUserWrite] {
	return pulumix.Output[map[string]*StageUserWrite]{
		OutputState: i.ToStageUserWriteMapOutputWithContext(ctx).OutputState,
	}
}

type StageUserWriteOutput struct{ *pulumi.OutputState }

func (StageUserWriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageUserWrite)(nil)).Elem()
}

func (o StageUserWriteOutput) ToStageUserWriteOutput() StageUserWriteOutput {
	return o
}

func (o StageUserWriteOutput) ToStageUserWriteOutputWithContext(ctx context.Context) StageUserWriteOutput {
	return o
}

func (o StageUserWriteOutput) ToOutput(ctx context.Context) pulumix.Output[*StageUserWrite] {
	return pulumix.Output[*StageUserWrite]{
		OutputState: o.OutputState,
	}
}

// Defaults to `true`.
func (o StageUserWriteOutput) CreateUsersAsInactive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageUserWrite) pulumi.BoolPtrOutput { return v.CreateUsersAsInactive }).(pulumi.BoolPtrOutput)
}

func (o StageUserWriteOutput) CreateUsersGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageUserWrite) pulumi.StringPtrOutput { return v.CreateUsersGroup }).(pulumi.StringPtrOutput)
}

func (o StageUserWriteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageUserWrite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defaults to `createWhenRequired`.
func (o StageUserWriteOutput) UserCreationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageUserWrite) pulumi.StringPtrOutput { return v.UserCreationMode }).(pulumi.StringPtrOutput)
}

// Defaults to “.
func (o StageUserWriteOutput) UserPathTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageUserWrite) pulumi.StringPtrOutput { return v.UserPathTemplate }).(pulumi.StringPtrOutput)
}

type StageUserWriteArrayOutput struct{ *pulumi.OutputState }

func (StageUserWriteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageUserWrite)(nil)).Elem()
}

func (o StageUserWriteArrayOutput) ToStageUserWriteArrayOutput() StageUserWriteArrayOutput {
	return o
}

func (o StageUserWriteArrayOutput) ToStageUserWriteArrayOutputWithContext(ctx context.Context) StageUserWriteArrayOutput {
	return o
}

func (o StageUserWriteArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*StageUserWrite] {
	return pulumix.Output[[]*StageUserWrite]{
		OutputState: o.OutputState,
	}
}

func (o StageUserWriteArrayOutput) Index(i pulumi.IntInput) StageUserWriteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageUserWrite {
		return vs[0].([]*StageUserWrite)[vs[1].(int)]
	}).(StageUserWriteOutput)
}

type StageUserWriteMapOutput struct{ *pulumi.OutputState }

func (StageUserWriteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageUserWrite)(nil)).Elem()
}

func (o StageUserWriteMapOutput) ToStageUserWriteMapOutput() StageUserWriteMapOutput {
	return o
}

func (o StageUserWriteMapOutput) ToStageUserWriteMapOutputWithContext(ctx context.Context) StageUserWriteMapOutput {
	return o
}

func (o StageUserWriteMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*StageUserWrite] {
	return pulumix.Output[map[string]*StageUserWrite]{
		OutputState: o.OutputState,
	}
}

func (o StageUserWriteMapOutput) MapIndex(k pulumi.StringInput) StageUserWriteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageUserWrite {
		return vs[0].(map[string]*StageUserWrite)[vs[1].(string)]
	}).(StageUserWriteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageUserWriteInput)(nil)).Elem(), &StageUserWrite{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageUserWriteArrayInput)(nil)).Elem(), StageUserWriteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageUserWriteMapInput)(nil)).Elem(), StageUserWriteMap{})
	pulumi.RegisterOutputType(StageUserWriteOutput{})
	pulumi.RegisterOutputType(StageUserWriteArrayOutput{})
	pulumi.RegisterOutputType(StageUserWriteMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StagePrompt struct {
	pulumi.CustomResourceState

	Fields             pulumi.StringArrayOutput `pulumi:"fields"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ValidationPolicies pulumi.StringArrayOutput `pulumi:"validationPolicies"`
}

// NewStagePrompt registers a new resource with the given unique name, arguments, and options.
func NewStagePrompt(ctx *pulumi.Context,
	name string, args *StagePromptArgs, opts ...pulumi.ResourceOption) (*StagePrompt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StagePrompt
	err := ctx.RegisterResource("authentik:index/stagePrompt:StagePrompt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStagePrompt gets an existing StagePrompt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStagePrompt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StagePromptState, opts ...pulumi.ResourceOption) (*StagePrompt, error) {
	var resource StagePrompt
	err := ctx.ReadResource("authentik:index/stagePrompt:StagePrompt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StagePrompt resources.
type stagePromptState struct {
	Fields             []string `pulumi:"fields"`
	Name               *string  `pulumi:"name"`
	ValidationPolicies []string `pulumi:"validationPolicies"`
}

type StagePromptState struct {
	Fields             pulumi.StringArrayInput
	Name               pulumi.StringPtrInput
	ValidationPolicies pulumi.StringArrayInput
}

func (StagePromptState) ElementType() reflect.Type {
	return reflect.TypeOf((*stagePromptState)(nil)).Elem()
}

type stagePromptArgs struct {
	Fields             []string `pulumi:"fields"`
	Name               *string  `pulumi:"name"`
	ValidationPolicies []string `pulumi:"validationPolicies"`
}

// The set of arguments for constructing a StagePrompt resource.
type StagePromptArgs struct {
	Fields             pulumi.StringArrayInput
	Name               pulumi.StringPtrInput
	ValidationPolicies pulumi.StringArrayInput
}

func (StagePromptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stagePromptArgs)(nil)).Elem()
}

type StagePromptInput interface {
	pulumi.Input

	ToStagePromptOutput() StagePromptOutput
	ToStagePromptOutputWithContext(ctx context.Context) StagePromptOutput
}

func (*StagePrompt) ElementType() reflect.Type {
	return reflect.TypeOf((**StagePrompt)(nil)).Elem()
}

func (i *StagePrompt) ToStagePromptOutput() StagePromptOutput {
	return i.ToStagePromptOutputWithContext(context.Background())
}

func (i *StagePrompt) ToStagePromptOutputWithContext(ctx context.Context) StagePromptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePromptOutput)
}

// StagePromptArrayInput is an input type that accepts StagePromptArray and StagePromptArrayOutput values.
// You can construct a concrete instance of `StagePromptArrayInput` via:
//
//	StagePromptArray{ StagePromptArgs{...} }
type StagePromptArrayInput interface {
	pulumi.Input

	ToStagePromptArrayOutput() StagePromptArrayOutput
	ToStagePromptArrayOutputWithContext(context.Context) StagePromptArrayOutput
}

type StagePromptArray []StagePromptInput

func (StagePromptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StagePrompt)(nil)).Elem()
}

func (i StagePromptArray) ToStagePromptArrayOutput() StagePromptArrayOutput {
	return i.ToStagePromptArrayOutputWithContext(context.Background())
}

func (i StagePromptArray) ToStagePromptArrayOutputWithContext(ctx context.Context) StagePromptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePromptArrayOutput)
}

// StagePromptMapInput is an input type that accepts StagePromptMap and StagePromptMapOutput values.
// You can construct a concrete instance of `StagePromptMapInput` via:
//
//	StagePromptMap{ "key": StagePromptArgs{...} }
type StagePromptMapInput interface {
	pulumi.Input

	ToStagePromptMapOutput() StagePromptMapOutput
	ToStagePromptMapOutputWithContext(context.Context) StagePromptMapOutput
}

type StagePromptMap map[string]StagePromptInput

func (StagePromptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StagePrompt)(nil)).Elem()
}

func (i StagePromptMap) ToStagePromptMapOutput() StagePromptMapOutput {
	return i.ToStagePromptMapOutputWithContext(context.Background())
}

func (i StagePromptMap) ToStagePromptMapOutputWithContext(ctx context.Context) StagePromptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePromptMapOutput)
}

type StagePromptOutput struct{ *pulumi.OutputState }

func (StagePromptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StagePrompt)(nil)).Elem()
}

func (o StagePromptOutput) ToStagePromptOutput() StagePromptOutput {
	return o
}

func (o StagePromptOutput) ToStagePromptOutputWithContext(ctx context.Context) StagePromptOutput {
	return o
}

func (o StagePromptOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StagePrompt) pulumi.StringArrayOutput { return v.Fields }).(pulumi.StringArrayOutput)
}

func (o StagePromptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StagePrompt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StagePromptOutput) ValidationPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StagePrompt) pulumi.StringArrayOutput { return v.ValidationPolicies }).(pulumi.StringArrayOutput)
}

type StagePromptArrayOutput struct{ *pulumi.OutputState }

func (StagePromptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StagePrompt)(nil)).Elem()
}

func (o StagePromptArrayOutput) ToStagePromptArrayOutput() StagePromptArrayOutput {
	return o
}

func (o StagePromptArrayOutput) ToStagePromptArrayOutputWithContext(ctx context.Context) StagePromptArrayOutput {
	return o
}

func (o StagePromptArrayOutput) Index(i pulumi.IntInput) StagePromptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StagePrompt {
		return vs[0].([]*StagePrompt)[vs[1].(int)]
	}).(StagePromptOutput)
}

type StagePromptMapOutput struct{ *pulumi.OutputState }

func (StagePromptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StagePrompt)(nil)).Elem()
}

func (o StagePromptMapOutput) ToStagePromptMapOutput() StagePromptMapOutput {
	return o
}

func (o StagePromptMapOutput) ToStagePromptMapOutputWithContext(ctx context.Context) StagePromptMapOutput {
	return o
}

func (o StagePromptMapOutput) MapIndex(k pulumi.StringInput) StagePromptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StagePrompt {
		return vs[0].(map[string]*StagePrompt)[vs[1].(string)]
	}).(StagePromptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StagePromptInput)(nil)).Elem(), &StagePrompt{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePromptArrayInput)(nil)).Elem(), StagePromptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePromptMapInput)(nil)).Elem(), StagePromptMap{})
	pulumi.RegisterOutputType(StagePromptOutput{})
	pulumi.RegisterOutputType(StagePromptArrayOutput{})
	pulumi.RegisterOutputType(StagePromptMapOutput{})
}

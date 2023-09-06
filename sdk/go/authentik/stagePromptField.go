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

type StagePromptField struct {
	pulumi.CustomResourceState

	FieldKey               pulumi.StringOutput    `pulumi:"fieldKey"`
	InitialValue           pulumi.StringPtrOutput `pulumi:"initialValue"`
	InitialValueExpression pulumi.BoolPtrOutput   `pulumi:"initialValueExpression"`
	Label                  pulumi.StringOutput    `pulumi:"label"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	Order                  pulumi.IntPtrOutput    `pulumi:"order"`
	Placeholder            pulumi.StringPtrOutput `pulumi:"placeholder"`
	PlaceholderExpression  pulumi.BoolPtrOutput   `pulumi:"placeholderExpression"`
	Required               pulumi.BoolPtrOutput   `pulumi:"required"`
	SubText                pulumi.StringPtrOutput `pulumi:"subText"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
}

// NewStagePromptField registers a new resource with the given unique name, arguments, and options.
func NewStagePromptField(ctx *pulumi.Context,
	name string, args *StagePromptFieldArgs, opts ...pulumi.ResourceOption) (*StagePromptField, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FieldKey == nil {
		return nil, errors.New("invalid value for required argument 'FieldKey'")
	}
	if args.Label == nil {
		return nil, errors.New("invalid value for required argument 'Label'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StagePromptField
	err := ctx.RegisterResource("authentik:index/stagePromptField:StagePromptField", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStagePromptField gets an existing StagePromptField resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStagePromptField(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StagePromptFieldState, opts ...pulumi.ResourceOption) (*StagePromptField, error) {
	var resource StagePromptField
	err := ctx.ReadResource("authentik:index/stagePromptField:StagePromptField", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StagePromptField resources.
type stagePromptFieldState struct {
	FieldKey               *string `pulumi:"fieldKey"`
	InitialValue           *string `pulumi:"initialValue"`
	InitialValueExpression *bool   `pulumi:"initialValueExpression"`
	Label                  *string `pulumi:"label"`
	Name                   *string `pulumi:"name"`
	Order                  *int    `pulumi:"order"`
	Placeholder            *string `pulumi:"placeholder"`
	PlaceholderExpression  *bool   `pulumi:"placeholderExpression"`
	Required               *bool   `pulumi:"required"`
	SubText                *string `pulumi:"subText"`
	Type                   *string `pulumi:"type"`
}

type StagePromptFieldState struct {
	FieldKey               pulumi.StringPtrInput
	InitialValue           pulumi.StringPtrInput
	InitialValueExpression pulumi.BoolPtrInput
	Label                  pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	Order                  pulumi.IntPtrInput
	Placeholder            pulumi.StringPtrInput
	PlaceholderExpression  pulumi.BoolPtrInput
	Required               pulumi.BoolPtrInput
	SubText                pulumi.StringPtrInput
	Type                   pulumi.StringPtrInput
}

func (StagePromptFieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*stagePromptFieldState)(nil)).Elem()
}

type stagePromptFieldArgs struct {
	FieldKey               string  `pulumi:"fieldKey"`
	InitialValue           *string `pulumi:"initialValue"`
	InitialValueExpression *bool   `pulumi:"initialValueExpression"`
	Label                  string  `pulumi:"label"`
	Name                   *string `pulumi:"name"`
	Order                  *int    `pulumi:"order"`
	Placeholder            *string `pulumi:"placeholder"`
	PlaceholderExpression  *bool   `pulumi:"placeholderExpression"`
	Required               *bool   `pulumi:"required"`
	SubText                *string `pulumi:"subText"`
	Type                   string  `pulumi:"type"`
}

// The set of arguments for constructing a StagePromptField resource.
type StagePromptFieldArgs struct {
	FieldKey               pulumi.StringInput
	InitialValue           pulumi.StringPtrInput
	InitialValueExpression pulumi.BoolPtrInput
	Label                  pulumi.StringInput
	Name                   pulumi.StringPtrInput
	Order                  pulumi.IntPtrInput
	Placeholder            pulumi.StringPtrInput
	PlaceholderExpression  pulumi.BoolPtrInput
	Required               pulumi.BoolPtrInput
	SubText                pulumi.StringPtrInput
	Type                   pulumi.StringInput
}

func (StagePromptFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stagePromptFieldArgs)(nil)).Elem()
}

type StagePromptFieldInput interface {
	pulumi.Input

	ToStagePromptFieldOutput() StagePromptFieldOutput
	ToStagePromptFieldOutputWithContext(ctx context.Context) StagePromptFieldOutput
}

func (*StagePromptField) ElementType() reflect.Type {
	return reflect.TypeOf((**StagePromptField)(nil)).Elem()
}

func (i *StagePromptField) ToStagePromptFieldOutput() StagePromptFieldOutput {
	return i.ToStagePromptFieldOutputWithContext(context.Background())
}

func (i *StagePromptField) ToStagePromptFieldOutputWithContext(ctx context.Context) StagePromptFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePromptFieldOutput)
}

// StagePromptFieldArrayInput is an input type that accepts StagePromptFieldArray and StagePromptFieldArrayOutput values.
// You can construct a concrete instance of `StagePromptFieldArrayInput` via:
//
//	StagePromptFieldArray{ StagePromptFieldArgs{...} }
type StagePromptFieldArrayInput interface {
	pulumi.Input

	ToStagePromptFieldArrayOutput() StagePromptFieldArrayOutput
	ToStagePromptFieldArrayOutputWithContext(context.Context) StagePromptFieldArrayOutput
}

type StagePromptFieldArray []StagePromptFieldInput

func (StagePromptFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StagePromptField)(nil)).Elem()
}

func (i StagePromptFieldArray) ToStagePromptFieldArrayOutput() StagePromptFieldArrayOutput {
	return i.ToStagePromptFieldArrayOutputWithContext(context.Background())
}

func (i StagePromptFieldArray) ToStagePromptFieldArrayOutputWithContext(ctx context.Context) StagePromptFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePromptFieldArrayOutput)
}

// StagePromptFieldMapInput is an input type that accepts StagePromptFieldMap and StagePromptFieldMapOutput values.
// You can construct a concrete instance of `StagePromptFieldMapInput` via:
//
//	StagePromptFieldMap{ "key": StagePromptFieldArgs{...} }
type StagePromptFieldMapInput interface {
	pulumi.Input

	ToStagePromptFieldMapOutput() StagePromptFieldMapOutput
	ToStagePromptFieldMapOutputWithContext(context.Context) StagePromptFieldMapOutput
}

type StagePromptFieldMap map[string]StagePromptFieldInput

func (StagePromptFieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StagePromptField)(nil)).Elem()
}

func (i StagePromptFieldMap) ToStagePromptFieldMapOutput() StagePromptFieldMapOutput {
	return i.ToStagePromptFieldMapOutputWithContext(context.Background())
}

func (i StagePromptFieldMap) ToStagePromptFieldMapOutputWithContext(ctx context.Context) StagePromptFieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePromptFieldMapOutput)
}

type StagePromptFieldOutput struct{ *pulumi.OutputState }

func (StagePromptFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StagePromptField)(nil)).Elem()
}

func (o StagePromptFieldOutput) ToStagePromptFieldOutput() StagePromptFieldOutput {
	return o
}

func (o StagePromptFieldOutput) ToStagePromptFieldOutputWithContext(ctx context.Context) StagePromptFieldOutput {
	return o
}

func (o StagePromptFieldOutput) FieldKey() pulumi.StringOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringOutput { return v.FieldKey }).(pulumi.StringOutput)
}

func (o StagePromptFieldOutput) InitialValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringPtrOutput { return v.InitialValue }).(pulumi.StringPtrOutput)
}

func (o StagePromptFieldOutput) InitialValueExpression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.BoolPtrOutput { return v.InitialValueExpression }).(pulumi.BoolPtrOutput)
}

func (o StagePromptFieldOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

func (o StagePromptFieldOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StagePromptFieldOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.IntPtrOutput { return v.Order }).(pulumi.IntPtrOutput)
}

func (o StagePromptFieldOutput) Placeholder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringPtrOutput { return v.Placeholder }).(pulumi.StringPtrOutput)
}

func (o StagePromptFieldOutput) PlaceholderExpression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.BoolPtrOutput { return v.PlaceholderExpression }).(pulumi.BoolPtrOutput)
}

func (o StagePromptFieldOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.BoolPtrOutput { return v.Required }).(pulumi.BoolPtrOutput)
}

func (o StagePromptFieldOutput) SubText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringPtrOutput { return v.SubText }).(pulumi.StringPtrOutput)
}

func (o StagePromptFieldOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StagePromptField) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type StagePromptFieldArrayOutput struct{ *pulumi.OutputState }

func (StagePromptFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StagePromptField)(nil)).Elem()
}

func (o StagePromptFieldArrayOutput) ToStagePromptFieldArrayOutput() StagePromptFieldArrayOutput {
	return o
}

func (o StagePromptFieldArrayOutput) ToStagePromptFieldArrayOutputWithContext(ctx context.Context) StagePromptFieldArrayOutput {
	return o
}

func (o StagePromptFieldArrayOutput) Index(i pulumi.IntInput) StagePromptFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StagePromptField {
		return vs[0].([]*StagePromptField)[vs[1].(int)]
	}).(StagePromptFieldOutput)
}

type StagePromptFieldMapOutput struct{ *pulumi.OutputState }

func (StagePromptFieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StagePromptField)(nil)).Elem()
}

func (o StagePromptFieldMapOutput) ToStagePromptFieldMapOutput() StagePromptFieldMapOutput {
	return o
}

func (o StagePromptFieldMapOutput) ToStagePromptFieldMapOutputWithContext(ctx context.Context) StagePromptFieldMapOutput {
	return o
}

func (o StagePromptFieldMapOutput) MapIndex(k pulumi.StringInput) StagePromptFieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StagePromptField {
		return vs[0].(map[string]*StagePromptField)[vs[1].(string)]
	}).(StagePromptFieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StagePromptFieldInput)(nil)).Elem(), &StagePromptField{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePromptFieldArrayInput)(nil)).Elem(), StagePromptFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePromptFieldMapInput)(nil)).Elem(), StagePromptFieldMap{})
	pulumi.RegisterOutputType(StagePromptFieldOutput{})
	pulumi.RegisterOutputType(StagePromptFieldArrayOutput{})
	pulumi.RegisterOutputType(StagePromptFieldMapOutput{})
}

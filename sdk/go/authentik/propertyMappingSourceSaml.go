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

// Manage SAML Source Property mappings
type PropertyMappingSourceSaml struct {
	pulumi.CustomResourceState

	Expression pulumi.StringOutput `pulumi:"expression"`
	Name       pulumi.StringOutput `pulumi:"name"`
}

// NewPropertyMappingSourceSaml registers a new resource with the given unique name, arguments, and options.
func NewPropertyMappingSourceSaml(ctx *pulumi.Context,
	name string, args *PropertyMappingSourceSamlArgs, opts ...pulumi.ResourceOption) (*PropertyMappingSourceSaml, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PropertyMappingSourceSaml
	err := ctx.RegisterResource("authentik:index/propertyMappingSourceSaml:PropertyMappingSourceSaml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyMappingSourceSaml gets an existing PropertyMappingSourceSaml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyMappingSourceSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyMappingSourceSamlState, opts ...pulumi.ResourceOption) (*PropertyMappingSourceSaml, error) {
	var resource PropertyMappingSourceSaml
	err := ctx.ReadResource("authentik:index/propertyMappingSourceSaml:PropertyMappingSourceSaml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyMappingSourceSaml resources.
type propertyMappingSourceSamlState struct {
	Expression *string `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

type PropertyMappingSourceSamlState struct {
	Expression pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingSourceSamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingSourceSamlState)(nil)).Elem()
}

type propertyMappingSourceSamlArgs struct {
	Expression string  `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

// The set of arguments for constructing a PropertyMappingSourceSaml resource.
type PropertyMappingSourceSamlArgs struct {
	Expression pulumi.StringInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingSourceSamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingSourceSamlArgs)(nil)).Elem()
}

type PropertyMappingSourceSamlInput interface {
	pulumi.Input

	ToPropertyMappingSourceSamlOutput() PropertyMappingSourceSamlOutput
	ToPropertyMappingSourceSamlOutputWithContext(ctx context.Context) PropertyMappingSourceSamlOutput
}

func (*PropertyMappingSourceSaml) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingSourceSaml)(nil)).Elem()
}

func (i *PropertyMappingSourceSaml) ToPropertyMappingSourceSamlOutput() PropertyMappingSourceSamlOutput {
	return i.ToPropertyMappingSourceSamlOutputWithContext(context.Background())
}

func (i *PropertyMappingSourceSaml) ToPropertyMappingSourceSamlOutputWithContext(ctx context.Context) PropertyMappingSourceSamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingSourceSamlOutput)
}

// PropertyMappingSourceSamlArrayInput is an input type that accepts PropertyMappingSourceSamlArray and PropertyMappingSourceSamlArrayOutput values.
// You can construct a concrete instance of `PropertyMappingSourceSamlArrayInput` via:
//
//	PropertyMappingSourceSamlArray{ PropertyMappingSourceSamlArgs{...} }
type PropertyMappingSourceSamlArrayInput interface {
	pulumi.Input

	ToPropertyMappingSourceSamlArrayOutput() PropertyMappingSourceSamlArrayOutput
	ToPropertyMappingSourceSamlArrayOutputWithContext(context.Context) PropertyMappingSourceSamlArrayOutput
}

type PropertyMappingSourceSamlArray []PropertyMappingSourceSamlInput

func (PropertyMappingSourceSamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingSourceSaml)(nil)).Elem()
}

func (i PropertyMappingSourceSamlArray) ToPropertyMappingSourceSamlArrayOutput() PropertyMappingSourceSamlArrayOutput {
	return i.ToPropertyMappingSourceSamlArrayOutputWithContext(context.Background())
}

func (i PropertyMappingSourceSamlArray) ToPropertyMappingSourceSamlArrayOutputWithContext(ctx context.Context) PropertyMappingSourceSamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingSourceSamlArrayOutput)
}

// PropertyMappingSourceSamlMapInput is an input type that accepts PropertyMappingSourceSamlMap and PropertyMappingSourceSamlMapOutput values.
// You can construct a concrete instance of `PropertyMappingSourceSamlMapInput` via:
//
//	PropertyMappingSourceSamlMap{ "key": PropertyMappingSourceSamlArgs{...} }
type PropertyMappingSourceSamlMapInput interface {
	pulumi.Input

	ToPropertyMappingSourceSamlMapOutput() PropertyMappingSourceSamlMapOutput
	ToPropertyMappingSourceSamlMapOutputWithContext(context.Context) PropertyMappingSourceSamlMapOutput
}

type PropertyMappingSourceSamlMap map[string]PropertyMappingSourceSamlInput

func (PropertyMappingSourceSamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingSourceSaml)(nil)).Elem()
}

func (i PropertyMappingSourceSamlMap) ToPropertyMappingSourceSamlMapOutput() PropertyMappingSourceSamlMapOutput {
	return i.ToPropertyMappingSourceSamlMapOutputWithContext(context.Background())
}

func (i PropertyMappingSourceSamlMap) ToPropertyMappingSourceSamlMapOutputWithContext(ctx context.Context) PropertyMappingSourceSamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingSourceSamlMapOutput)
}

type PropertyMappingSourceSamlOutput struct{ *pulumi.OutputState }

func (PropertyMappingSourceSamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingSourceSaml)(nil)).Elem()
}

func (o PropertyMappingSourceSamlOutput) ToPropertyMappingSourceSamlOutput() PropertyMappingSourceSamlOutput {
	return o
}

func (o PropertyMappingSourceSamlOutput) ToPropertyMappingSourceSamlOutputWithContext(ctx context.Context) PropertyMappingSourceSamlOutput {
	return o
}

func (o PropertyMappingSourceSamlOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingSourceSaml) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o PropertyMappingSourceSamlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingSourceSaml) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PropertyMappingSourceSamlArrayOutput struct{ *pulumi.OutputState }

func (PropertyMappingSourceSamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingSourceSaml)(nil)).Elem()
}

func (o PropertyMappingSourceSamlArrayOutput) ToPropertyMappingSourceSamlArrayOutput() PropertyMappingSourceSamlArrayOutput {
	return o
}

func (o PropertyMappingSourceSamlArrayOutput) ToPropertyMappingSourceSamlArrayOutputWithContext(ctx context.Context) PropertyMappingSourceSamlArrayOutput {
	return o
}

func (o PropertyMappingSourceSamlArrayOutput) Index(i pulumi.IntInput) PropertyMappingSourceSamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertyMappingSourceSaml {
		return vs[0].([]*PropertyMappingSourceSaml)[vs[1].(int)]
	}).(PropertyMappingSourceSamlOutput)
}

type PropertyMappingSourceSamlMapOutput struct{ *pulumi.OutputState }

func (PropertyMappingSourceSamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingSourceSaml)(nil)).Elem()
}

func (o PropertyMappingSourceSamlMapOutput) ToPropertyMappingSourceSamlMapOutput() PropertyMappingSourceSamlMapOutput {
	return o
}

func (o PropertyMappingSourceSamlMapOutput) ToPropertyMappingSourceSamlMapOutputWithContext(ctx context.Context) PropertyMappingSourceSamlMapOutput {
	return o
}

func (o PropertyMappingSourceSamlMapOutput) MapIndex(k pulumi.StringInput) PropertyMappingSourceSamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertyMappingSourceSaml {
		return vs[0].(map[string]*PropertyMappingSourceSaml)[vs[1].(string)]
	}).(PropertyMappingSourceSamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingSourceSamlInput)(nil)).Elem(), &PropertyMappingSourceSaml{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingSourceSamlArrayInput)(nil)).Elem(), PropertyMappingSourceSamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingSourceSamlMapInput)(nil)).Elem(), PropertyMappingSourceSamlMap{})
	pulumi.RegisterOutputType(PropertyMappingSourceSamlOutput{})
	pulumi.RegisterOutputType(PropertyMappingSourceSamlArrayOutput{})
	pulumi.RegisterOutputType(PropertyMappingSourceSamlMapOutput{})
}

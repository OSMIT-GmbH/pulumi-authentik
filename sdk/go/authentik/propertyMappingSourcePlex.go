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

// Manage Plex Source Property mappings
type PropertyMappingSourcePlex struct {
	pulumi.CustomResourceState

	Expression pulumi.StringOutput `pulumi:"expression"`
	Name       pulumi.StringOutput `pulumi:"name"`
}

// NewPropertyMappingSourcePlex registers a new resource with the given unique name, arguments, and options.
func NewPropertyMappingSourcePlex(ctx *pulumi.Context,
	name string, args *PropertyMappingSourcePlexArgs, opts ...pulumi.ResourceOption) (*PropertyMappingSourcePlex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PropertyMappingSourcePlex
	err := ctx.RegisterResource("authentik:index/propertyMappingSourcePlex:PropertyMappingSourcePlex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyMappingSourcePlex gets an existing PropertyMappingSourcePlex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyMappingSourcePlex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyMappingSourcePlexState, opts ...pulumi.ResourceOption) (*PropertyMappingSourcePlex, error) {
	var resource PropertyMappingSourcePlex
	err := ctx.ReadResource("authentik:index/propertyMappingSourcePlex:PropertyMappingSourcePlex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyMappingSourcePlex resources.
type propertyMappingSourcePlexState struct {
	Expression *string `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

type PropertyMappingSourcePlexState struct {
	Expression pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingSourcePlexState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingSourcePlexState)(nil)).Elem()
}

type propertyMappingSourcePlexArgs struct {
	Expression string  `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

// The set of arguments for constructing a PropertyMappingSourcePlex resource.
type PropertyMappingSourcePlexArgs struct {
	Expression pulumi.StringInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingSourcePlexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingSourcePlexArgs)(nil)).Elem()
}

type PropertyMappingSourcePlexInput interface {
	pulumi.Input

	ToPropertyMappingSourcePlexOutput() PropertyMappingSourcePlexOutput
	ToPropertyMappingSourcePlexOutputWithContext(ctx context.Context) PropertyMappingSourcePlexOutput
}

func (*PropertyMappingSourcePlex) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingSourcePlex)(nil)).Elem()
}

func (i *PropertyMappingSourcePlex) ToPropertyMappingSourcePlexOutput() PropertyMappingSourcePlexOutput {
	return i.ToPropertyMappingSourcePlexOutputWithContext(context.Background())
}

func (i *PropertyMappingSourcePlex) ToPropertyMappingSourcePlexOutputWithContext(ctx context.Context) PropertyMappingSourcePlexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingSourcePlexOutput)
}

// PropertyMappingSourcePlexArrayInput is an input type that accepts PropertyMappingSourcePlexArray and PropertyMappingSourcePlexArrayOutput values.
// You can construct a concrete instance of `PropertyMappingSourcePlexArrayInput` via:
//
//	PropertyMappingSourcePlexArray{ PropertyMappingSourcePlexArgs{...} }
type PropertyMappingSourcePlexArrayInput interface {
	pulumi.Input

	ToPropertyMappingSourcePlexArrayOutput() PropertyMappingSourcePlexArrayOutput
	ToPropertyMappingSourcePlexArrayOutputWithContext(context.Context) PropertyMappingSourcePlexArrayOutput
}

type PropertyMappingSourcePlexArray []PropertyMappingSourcePlexInput

func (PropertyMappingSourcePlexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingSourcePlex)(nil)).Elem()
}

func (i PropertyMappingSourcePlexArray) ToPropertyMappingSourcePlexArrayOutput() PropertyMappingSourcePlexArrayOutput {
	return i.ToPropertyMappingSourcePlexArrayOutputWithContext(context.Background())
}

func (i PropertyMappingSourcePlexArray) ToPropertyMappingSourcePlexArrayOutputWithContext(ctx context.Context) PropertyMappingSourcePlexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingSourcePlexArrayOutput)
}

// PropertyMappingSourcePlexMapInput is an input type that accepts PropertyMappingSourcePlexMap and PropertyMappingSourcePlexMapOutput values.
// You can construct a concrete instance of `PropertyMappingSourcePlexMapInput` via:
//
//	PropertyMappingSourcePlexMap{ "key": PropertyMappingSourcePlexArgs{...} }
type PropertyMappingSourcePlexMapInput interface {
	pulumi.Input

	ToPropertyMappingSourcePlexMapOutput() PropertyMappingSourcePlexMapOutput
	ToPropertyMappingSourcePlexMapOutputWithContext(context.Context) PropertyMappingSourcePlexMapOutput
}

type PropertyMappingSourcePlexMap map[string]PropertyMappingSourcePlexInput

func (PropertyMappingSourcePlexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingSourcePlex)(nil)).Elem()
}

func (i PropertyMappingSourcePlexMap) ToPropertyMappingSourcePlexMapOutput() PropertyMappingSourcePlexMapOutput {
	return i.ToPropertyMappingSourcePlexMapOutputWithContext(context.Background())
}

func (i PropertyMappingSourcePlexMap) ToPropertyMappingSourcePlexMapOutputWithContext(ctx context.Context) PropertyMappingSourcePlexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingSourcePlexMapOutput)
}

type PropertyMappingSourcePlexOutput struct{ *pulumi.OutputState }

func (PropertyMappingSourcePlexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingSourcePlex)(nil)).Elem()
}

func (o PropertyMappingSourcePlexOutput) ToPropertyMappingSourcePlexOutput() PropertyMappingSourcePlexOutput {
	return o
}

func (o PropertyMappingSourcePlexOutput) ToPropertyMappingSourcePlexOutputWithContext(ctx context.Context) PropertyMappingSourcePlexOutput {
	return o
}

func (o PropertyMappingSourcePlexOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingSourcePlex) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o PropertyMappingSourcePlexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingSourcePlex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PropertyMappingSourcePlexArrayOutput struct{ *pulumi.OutputState }

func (PropertyMappingSourcePlexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingSourcePlex)(nil)).Elem()
}

func (o PropertyMappingSourcePlexArrayOutput) ToPropertyMappingSourcePlexArrayOutput() PropertyMappingSourcePlexArrayOutput {
	return o
}

func (o PropertyMappingSourcePlexArrayOutput) ToPropertyMappingSourcePlexArrayOutputWithContext(ctx context.Context) PropertyMappingSourcePlexArrayOutput {
	return o
}

func (o PropertyMappingSourcePlexArrayOutput) Index(i pulumi.IntInput) PropertyMappingSourcePlexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertyMappingSourcePlex {
		return vs[0].([]*PropertyMappingSourcePlex)[vs[1].(int)]
	}).(PropertyMappingSourcePlexOutput)
}

type PropertyMappingSourcePlexMapOutput struct{ *pulumi.OutputState }

func (PropertyMappingSourcePlexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingSourcePlex)(nil)).Elem()
}

func (o PropertyMappingSourcePlexMapOutput) ToPropertyMappingSourcePlexMapOutput() PropertyMappingSourcePlexMapOutput {
	return o
}

func (o PropertyMappingSourcePlexMapOutput) ToPropertyMappingSourcePlexMapOutputWithContext(ctx context.Context) PropertyMappingSourcePlexMapOutput {
	return o
}

func (o PropertyMappingSourcePlexMapOutput) MapIndex(k pulumi.StringInput) PropertyMappingSourcePlexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertyMappingSourcePlex {
		return vs[0].(map[string]*PropertyMappingSourcePlex)[vs[1].(string)]
	}).(PropertyMappingSourcePlexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingSourcePlexInput)(nil)).Elem(), &PropertyMappingSourcePlex{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingSourcePlexArrayInput)(nil)).Elem(), PropertyMappingSourcePlexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingSourcePlexMapInput)(nil)).Elem(), PropertyMappingSourcePlexMap{})
	pulumi.RegisterOutputType(PropertyMappingSourcePlexOutput{})
	pulumi.RegisterOutputType(PropertyMappingSourcePlexArrayOutput{})
	pulumi.RegisterOutputType(PropertyMappingSourcePlexMapOutput{})
}

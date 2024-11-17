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

// Manage Radius Provider Property mappings
//
// > This resource is deprecated. Migrate to `PropertyMappingProviderRadius`.
type PropertyMappingRadius struct {
	pulumi.CustomResourceState

	Expression pulumi.StringOutput `pulumi:"expression"`
	Name       pulumi.StringOutput `pulumi:"name"`
}

// NewPropertyMappingRadius registers a new resource with the given unique name, arguments, and options.
func NewPropertyMappingRadius(ctx *pulumi.Context,
	name string, args *PropertyMappingRadiusArgs, opts ...pulumi.ResourceOption) (*PropertyMappingRadius, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PropertyMappingRadius
	err := ctx.RegisterResource("authentik:index/propertyMappingRadius:PropertyMappingRadius", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyMappingRadius gets an existing PropertyMappingRadius resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyMappingRadius(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyMappingRadiusState, opts ...pulumi.ResourceOption) (*PropertyMappingRadius, error) {
	var resource PropertyMappingRadius
	err := ctx.ReadResource("authentik:index/propertyMappingRadius:PropertyMappingRadius", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyMappingRadius resources.
type propertyMappingRadiusState struct {
	Expression *string `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

type PropertyMappingRadiusState struct {
	Expression pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingRadiusState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingRadiusState)(nil)).Elem()
}

type propertyMappingRadiusArgs struct {
	Expression string  `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

// The set of arguments for constructing a PropertyMappingRadius resource.
type PropertyMappingRadiusArgs struct {
	Expression pulumi.StringInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingRadiusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingRadiusArgs)(nil)).Elem()
}

type PropertyMappingRadiusInput interface {
	pulumi.Input

	ToPropertyMappingRadiusOutput() PropertyMappingRadiusOutput
	ToPropertyMappingRadiusOutputWithContext(ctx context.Context) PropertyMappingRadiusOutput
}

func (*PropertyMappingRadius) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingRadius)(nil)).Elem()
}

func (i *PropertyMappingRadius) ToPropertyMappingRadiusOutput() PropertyMappingRadiusOutput {
	return i.ToPropertyMappingRadiusOutputWithContext(context.Background())
}

func (i *PropertyMappingRadius) ToPropertyMappingRadiusOutputWithContext(ctx context.Context) PropertyMappingRadiusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingRadiusOutput)
}

// PropertyMappingRadiusArrayInput is an input type that accepts PropertyMappingRadiusArray and PropertyMappingRadiusArrayOutput values.
// You can construct a concrete instance of `PropertyMappingRadiusArrayInput` via:
//
//	PropertyMappingRadiusArray{ PropertyMappingRadiusArgs{...} }
type PropertyMappingRadiusArrayInput interface {
	pulumi.Input

	ToPropertyMappingRadiusArrayOutput() PropertyMappingRadiusArrayOutput
	ToPropertyMappingRadiusArrayOutputWithContext(context.Context) PropertyMappingRadiusArrayOutput
}

type PropertyMappingRadiusArray []PropertyMappingRadiusInput

func (PropertyMappingRadiusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingRadius)(nil)).Elem()
}

func (i PropertyMappingRadiusArray) ToPropertyMappingRadiusArrayOutput() PropertyMappingRadiusArrayOutput {
	return i.ToPropertyMappingRadiusArrayOutputWithContext(context.Background())
}

func (i PropertyMappingRadiusArray) ToPropertyMappingRadiusArrayOutputWithContext(ctx context.Context) PropertyMappingRadiusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingRadiusArrayOutput)
}

// PropertyMappingRadiusMapInput is an input type that accepts PropertyMappingRadiusMap and PropertyMappingRadiusMapOutput values.
// You can construct a concrete instance of `PropertyMappingRadiusMapInput` via:
//
//	PropertyMappingRadiusMap{ "key": PropertyMappingRadiusArgs{...} }
type PropertyMappingRadiusMapInput interface {
	pulumi.Input

	ToPropertyMappingRadiusMapOutput() PropertyMappingRadiusMapOutput
	ToPropertyMappingRadiusMapOutputWithContext(context.Context) PropertyMappingRadiusMapOutput
}

type PropertyMappingRadiusMap map[string]PropertyMappingRadiusInput

func (PropertyMappingRadiusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingRadius)(nil)).Elem()
}

func (i PropertyMappingRadiusMap) ToPropertyMappingRadiusMapOutput() PropertyMappingRadiusMapOutput {
	return i.ToPropertyMappingRadiusMapOutputWithContext(context.Background())
}

func (i PropertyMappingRadiusMap) ToPropertyMappingRadiusMapOutputWithContext(ctx context.Context) PropertyMappingRadiusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingRadiusMapOutput)
}

type PropertyMappingRadiusOutput struct{ *pulumi.OutputState }

func (PropertyMappingRadiusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingRadius)(nil)).Elem()
}

func (o PropertyMappingRadiusOutput) ToPropertyMappingRadiusOutput() PropertyMappingRadiusOutput {
	return o
}

func (o PropertyMappingRadiusOutput) ToPropertyMappingRadiusOutputWithContext(ctx context.Context) PropertyMappingRadiusOutput {
	return o
}

func (o PropertyMappingRadiusOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingRadius) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o PropertyMappingRadiusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingRadius) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PropertyMappingRadiusArrayOutput struct{ *pulumi.OutputState }

func (PropertyMappingRadiusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingRadius)(nil)).Elem()
}

func (o PropertyMappingRadiusArrayOutput) ToPropertyMappingRadiusArrayOutput() PropertyMappingRadiusArrayOutput {
	return o
}

func (o PropertyMappingRadiusArrayOutput) ToPropertyMappingRadiusArrayOutputWithContext(ctx context.Context) PropertyMappingRadiusArrayOutput {
	return o
}

func (o PropertyMappingRadiusArrayOutput) Index(i pulumi.IntInput) PropertyMappingRadiusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertyMappingRadius {
		return vs[0].([]*PropertyMappingRadius)[vs[1].(int)]
	}).(PropertyMappingRadiusOutput)
}

type PropertyMappingRadiusMapOutput struct{ *pulumi.OutputState }

func (PropertyMappingRadiusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingRadius)(nil)).Elem()
}

func (o PropertyMappingRadiusMapOutput) ToPropertyMappingRadiusMapOutput() PropertyMappingRadiusMapOutput {
	return o
}

func (o PropertyMappingRadiusMapOutput) ToPropertyMappingRadiusMapOutputWithContext(ctx context.Context) PropertyMappingRadiusMapOutput {
	return o
}

func (o PropertyMappingRadiusMapOutput) MapIndex(k pulumi.StringInput) PropertyMappingRadiusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertyMappingRadius {
		return vs[0].(map[string]*PropertyMappingRadius)[vs[1].(string)]
	}).(PropertyMappingRadiusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingRadiusInput)(nil)).Elem(), &PropertyMappingRadius{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingRadiusArrayInput)(nil)).Elem(), PropertyMappingRadiusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingRadiusMapInput)(nil)).Elem(), PropertyMappingRadiusMap{})
	pulumi.RegisterOutputType(PropertyMappingRadiusOutput{})
	pulumi.RegisterOutputType(PropertyMappingRadiusArrayOutput{})
	pulumi.RegisterOutputType(PropertyMappingRadiusMapOutput{})
}
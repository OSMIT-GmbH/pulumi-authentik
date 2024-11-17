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

// Manage SCIM Provider Property mappings
type PropertyMappingProviderScim struct {
	pulumi.CustomResourceState

	Expression pulumi.StringOutput `pulumi:"expression"`
	Name       pulumi.StringOutput `pulumi:"name"`
}

// NewPropertyMappingProviderScim registers a new resource with the given unique name, arguments, and options.
func NewPropertyMappingProviderScim(ctx *pulumi.Context,
	name string, args *PropertyMappingProviderScimArgs, opts ...pulumi.ResourceOption) (*PropertyMappingProviderScim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PropertyMappingProviderScim
	err := ctx.RegisterResource("authentik:index/propertyMappingProviderScim:PropertyMappingProviderScim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyMappingProviderScim gets an existing PropertyMappingProviderScim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyMappingProviderScim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyMappingProviderScimState, opts ...pulumi.ResourceOption) (*PropertyMappingProviderScim, error) {
	var resource PropertyMappingProviderScim
	err := ctx.ReadResource("authentik:index/propertyMappingProviderScim:PropertyMappingProviderScim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyMappingProviderScim resources.
type propertyMappingProviderScimState struct {
	Expression *string `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

type PropertyMappingProviderScimState struct {
	Expression pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingProviderScimState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingProviderScimState)(nil)).Elem()
}

type propertyMappingProviderScimArgs struct {
	Expression string  `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

// The set of arguments for constructing a PropertyMappingProviderScim resource.
type PropertyMappingProviderScimArgs struct {
	Expression pulumi.StringInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingProviderScimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingProviderScimArgs)(nil)).Elem()
}

type PropertyMappingProviderScimInput interface {
	pulumi.Input

	ToPropertyMappingProviderScimOutput() PropertyMappingProviderScimOutput
	ToPropertyMappingProviderScimOutputWithContext(ctx context.Context) PropertyMappingProviderScimOutput
}

func (*PropertyMappingProviderScim) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingProviderScim)(nil)).Elem()
}

func (i *PropertyMappingProviderScim) ToPropertyMappingProviderScimOutput() PropertyMappingProviderScimOutput {
	return i.ToPropertyMappingProviderScimOutputWithContext(context.Background())
}

func (i *PropertyMappingProviderScim) ToPropertyMappingProviderScimOutputWithContext(ctx context.Context) PropertyMappingProviderScimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingProviderScimOutput)
}

// PropertyMappingProviderScimArrayInput is an input type that accepts PropertyMappingProviderScimArray and PropertyMappingProviderScimArrayOutput values.
// You can construct a concrete instance of `PropertyMappingProviderScimArrayInput` via:
//
//	PropertyMappingProviderScimArray{ PropertyMappingProviderScimArgs{...} }
type PropertyMappingProviderScimArrayInput interface {
	pulumi.Input

	ToPropertyMappingProviderScimArrayOutput() PropertyMappingProviderScimArrayOutput
	ToPropertyMappingProviderScimArrayOutputWithContext(context.Context) PropertyMappingProviderScimArrayOutput
}

type PropertyMappingProviderScimArray []PropertyMappingProviderScimInput

func (PropertyMappingProviderScimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingProviderScim)(nil)).Elem()
}

func (i PropertyMappingProviderScimArray) ToPropertyMappingProviderScimArrayOutput() PropertyMappingProviderScimArrayOutput {
	return i.ToPropertyMappingProviderScimArrayOutputWithContext(context.Background())
}

func (i PropertyMappingProviderScimArray) ToPropertyMappingProviderScimArrayOutputWithContext(ctx context.Context) PropertyMappingProviderScimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingProviderScimArrayOutput)
}

// PropertyMappingProviderScimMapInput is an input type that accepts PropertyMappingProviderScimMap and PropertyMappingProviderScimMapOutput values.
// You can construct a concrete instance of `PropertyMappingProviderScimMapInput` via:
//
//	PropertyMappingProviderScimMap{ "key": PropertyMappingProviderScimArgs{...} }
type PropertyMappingProviderScimMapInput interface {
	pulumi.Input

	ToPropertyMappingProviderScimMapOutput() PropertyMappingProviderScimMapOutput
	ToPropertyMappingProviderScimMapOutputWithContext(context.Context) PropertyMappingProviderScimMapOutput
}

type PropertyMappingProviderScimMap map[string]PropertyMappingProviderScimInput

func (PropertyMappingProviderScimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingProviderScim)(nil)).Elem()
}

func (i PropertyMappingProviderScimMap) ToPropertyMappingProviderScimMapOutput() PropertyMappingProviderScimMapOutput {
	return i.ToPropertyMappingProviderScimMapOutputWithContext(context.Background())
}

func (i PropertyMappingProviderScimMap) ToPropertyMappingProviderScimMapOutputWithContext(ctx context.Context) PropertyMappingProviderScimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingProviderScimMapOutput)
}

type PropertyMappingProviderScimOutput struct{ *pulumi.OutputState }

func (PropertyMappingProviderScimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingProviderScim)(nil)).Elem()
}

func (o PropertyMappingProviderScimOutput) ToPropertyMappingProviderScimOutput() PropertyMappingProviderScimOutput {
	return o
}

func (o PropertyMappingProviderScimOutput) ToPropertyMappingProviderScimOutputWithContext(ctx context.Context) PropertyMappingProviderScimOutput {
	return o
}

func (o PropertyMappingProviderScimOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingProviderScim) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o PropertyMappingProviderScimOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingProviderScim) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PropertyMappingProviderScimArrayOutput struct{ *pulumi.OutputState }

func (PropertyMappingProviderScimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingProviderScim)(nil)).Elem()
}

func (o PropertyMappingProviderScimArrayOutput) ToPropertyMappingProviderScimArrayOutput() PropertyMappingProviderScimArrayOutput {
	return o
}

func (o PropertyMappingProviderScimArrayOutput) ToPropertyMappingProviderScimArrayOutputWithContext(ctx context.Context) PropertyMappingProviderScimArrayOutput {
	return o
}

func (o PropertyMappingProviderScimArrayOutput) Index(i pulumi.IntInput) PropertyMappingProviderScimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertyMappingProviderScim {
		return vs[0].([]*PropertyMappingProviderScim)[vs[1].(int)]
	}).(PropertyMappingProviderScimOutput)
}

type PropertyMappingProviderScimMapOutput struct{ *pulumi.OutputState }

func (PropertyMappingProviderScimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingProviderScim)(nil)).Elem()
}

func (o PropertyMappingProviderScimMapOutput) ToPropertyMappingProviderScimMapOutput() PropertyMappingProviderScimMapOutput {
	return o
}

func (o PropertyMappingProviderScimMapOutput) ToPropertyMappingProviderScimMapOutputWithContext(ctx context.Context) PropertyMappingProviderScimMapOutput {
	return o
}

func (o PropertyMappingProviderScimMapOutput) MapIndex(k pulumi.StringInput) PropertyMappingProviderScimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertyMappingProviderScim {
		return vs[0].(map[string]*PropertyMappingProviderScim)[vs[1].(string)]
	}).(PropertyMappingProviderScimOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingProviderScimInput)(nil)).Elem(), &PropertyMappingProviderScim{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingProviderScimArrayInput)(nil)).Elem(), PropertyMappingProviderScimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingProviderScimMapInput)(nil)).Elem(), PropertyMappingProviderScimMap{})
	pulumi.RegisterOutputType(PropertyMappingProviderScimOutput{})
	pulumi.RegisterOutputType(PropertyMappingProviderScimArrayOutput{})
	pulumi.RegisterOutputType(PropertyMappingProviderScimMapOutput{})
}

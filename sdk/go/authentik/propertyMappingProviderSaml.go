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

// Manage SAML Provider Property mappings
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := authentik.NewPropertyMappingProviderSaml(ctx, "saml-aws-rolessessionname", &authentik.PropertyMappingProviderSamlArgs{
//				Expression: pulumi.String("return user.email"),
//				SamlName:   pulumi.String("https://aws.amazon.com/SAML/Attributes/RoleSessionName"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PropertyMappingProviderSaml struct {
	pulumi.CustomResourceState

	Expression   pulumi.StringOutput    `pulumi:"expression"`
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	SamlName     pulumi.StringOutput    `pulumi:"samlName"`
}

// NewPropertyMappingProviderSaml registers a new resource with the given unique name, arguments, and options.
func NewPropertyMappingProviderSaml(ctx *pulumi.Context,
	name string, args *PropertyMappingProviderSamlArgs, opts ...pulumi.ResourceOption) (*PropertyMappingProviderSaml, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	if args.SamlName == nil {
		return nil, errors.New("invalid value for required argument 'SamlName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PropertyMappingProviderSaml
	err := ctx.RegisterResource("authentik:index/propertyMappingProviderSaml:PropertyMappingProviderSaml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyMappingProviderSaml gets an existing PropertyMappingProviderSaml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyMappingProviderSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyMappingProviderSamlState, opts ...pulumi.ResourceOption) (*PropertyMappingProviderSaml, error) {
	var resource PropertyMappingProviderSaml
	err := ctx.ReadResource("authentik:index/propertyMappingProviderSaml:PropertyMappingProviderSaml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyMappingProviderSaml resources.
type propertyMappingProviderSamlState struct {
	Expression   *string `pulumi:"expression"`
	FriendlyName *string `pulumi:"friendlyName"`
	Name         *string `pulumi:"name"`
	SamlName     *string `pulumi:"samlName"`
}

type PropertyMappingProviderSamlState struct {
	Expression   pulumi.StringPtrInput
	FriendlyName pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
	SamlName     pulumi.StringPtrInput
}

func (PropertyMappingProviderSamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingProviderSamlState)(nil)).Elem()
}

type propertyMappingProviderSamlArgs struct {
	Expression   string  `pulumi:"expression"`
	FriendlyName *string `pulumi:"friendlyName"`
	Name         *string `pulumi:"name"`
	SamlName     string  `pulumi:"samlName"`
}

// The set of arguments for constructing a PropertyMappingProviderSaml resource.
type PropertyMappingProviderSamlArgs struct {
	Expression   pulumi.StringInput
	FriendlyName pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
	SamlName     pulumi.StringInput
}

func (PropertyMappingProviderSamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingProviderSamlArgs)(nil)).Elem()
}

type PropertyMappingProviderSamlInput interface {
	pulumi.Input

	ToPropertyMappingProviderSamlOutput() PropertyMappingProviderSamlOutput
	ToPropertyMappingProviderSamlOutputWithContext(ctx context.Context) PropertyMappingProviderSamlOutput
}

func (*PropertyMappingProviderSaml) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingProviderSaml)(nil)).Elem()
}

func (i *PropertyMappingProviderSaml) ToPropertyMappingProviderSamlOutput() PropertyMappingProviderSamlOutput {
	return i.ToPropertyMappingProviderSamlOutputWithContext(context.Background())
}

func (i *PropertyMappingProviderSaml) ToPropertyMappingProviderSamlOutputWithContext(ctx context.Context) PropertyMappingProviderSamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingProviderSamlOutput)
}

// PropertyMappingProviderSamlArrayInput is an input type that accepts PropertyMappingProviderSamlArray and PropertyMappingProviderSamlArrayOutput values.
// You can construct a concrete instance of `PropertyMappingProviderSamlArrayInput` via:
//
//	PropertyMappingProviderSamlArray{ PropertyMappingProviderSamlArgs{...} }
type PropertyMappingProviderSamlArrayInput interface {
	pulumi.Input

	ToPropertyMappingProviderSamlArrayOutput() PropertyMappingProviderSamlArrayOutput
	ToPropertyMappingProviderSamlArrayOutputWithContext(context.Context) PropertyMappingProviderSamlArrayOutput
}

type PropertyMappingProviderSamlArray []PropertyMappingProviderSamlInput

func (PropertyMappingProviderSamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingProviderSaml)(nil)).Elem()
}

func (i PropertyMappingProviderSamlArray) ToPropertyMappingProviderSamlArrayOutput() PropertyMappingProviderSamlArrayOutput {
	return i.ToPropertyMappingProviderSamlArrayOutputWithContext(context.Background())
}

func (i PropertyMappingProviderSamlArray) ToPropertyMappingProviderSamlArrayOutputWithContext(ctx context.Context) PropertyMappingProviderSamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingProviderSamlArrayOutput)
}

// PropertyMappingProviderSamlMapInput is an input type that accepts PropertyMappingProviderSamlMap and PropertyMappingProviderSamlMapOutput values.
// You can construct a concrete instance of `PropertyMappingProviderSamlMapInput` via:
//
//	PropertyMappingProviderSamlMap{ "key": PropertyMappingProviderSamlArgs{...} }
type PropertyMappingProviderSamlMapInput interface {
	pulumi.Input

	ToPropertyMappingProviderSamlMapOutput() PropertyMappingProviderSamlMapOutput
	ToPropertyMappingProviderSamlMapOutputWithContext(context.Context) PropertyMappingProviderSamlMapOutput
}

type PropertyMappingProviderSamlMap map[string]PropertyMappingProviderSamlInput

func (PropertyMappingProviderSamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingProviderSaml)(nil)).Elem()
}

func (i PropertyMappingProviderSamlMap) ToPropertyMappingProviderSamlMapOutput() PropertyMappingProviderSamlMapOutput {
	return i.ToPropertyMappingProviderSamlMapOutputWithContext(context.Background())
}

func (i PropertyMappingProviderSamlMap) ToPropertyMappingProviderSamlMapOutputWithContext(ctx context.Context) PropertyMappingProviderSamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingProviderSamlMapOutput)
}

type PropertyMappingProviderSamlOutput struct{ *pulumi.OutputState }

func (PropertyMappingProviderSamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingProviderSaml)(nil)).Elem()
}

func (o PropertyMappingProviderSamlOutput) ToPropertyMappingProviderSamlOutput() PropertyMappingProviderSamlOutput {
	return o
}

func (o PropertyMappingProviderSamlOutput) ToPropertyMappingProviderSamlOutputWithContext(ctx context.Context) PropertyMappingProviderSamlOutput {
	return o
}

func (o PropertyMappingProviderSamlOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingProviderSaml) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o PropertyMappingProviderSamlOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PropertyMappingProviderSaml) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o PropertyMappingProviderSamlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingProviderSaml) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PropertyMappingProviderSamlOutput) SamlName() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingProviderSaml) pulumi.StringOutput { return v.SamlName }).(pulumi.StringOutput)
}

type PropertyMappingProviderSamlArrayOutput struct{ *pulumi.OutputState }

func (PropertyMappingProviderSamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingProviderSaml)(nil)).Elem()
}

func (o PropertyMappingProviderSamlArrayOutput) ToPropertyMappingProviderSamlArrayOutput() PropertyMappingProviderSamlArrayOutput {
	return o
}

func (o PropertyMappingProviderSamlArrayOutput) ToPropertyMappingProviderSamlArrayOutputWithContext(ctx context.Context) PropertyMappingProviderSamlArrayOutput {
	return o
}

func (o PropertyMappingProviderSamlArrayOutput) Index(i pulumi.IntInput) PropertyMappingProviderSamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertyMappingProviderSaml {
		return vs[0].([]*PropertyMappingProviderSaml)[vs[1].(int)]
	}).(PropertyMappingProviderSamlOutput)
}

type PropertyMappingProviderSamlMapOutput struct{ *pulumi.OutputState }

func (PropertyMappingProviderSamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingProviderSaml)(nil)).Elem()
}

func (o PropertyMappingProviderSamlMapOutput) ToPropertyMappingProviderSamlMapOutput() PropertyMappingProviderSamlMapOutput {
	return o
}

func (o PropertyMappingProviderSamlMapOutput) ToPropertyMappingProviderSamlMapOutputWithContext(ctx context.Context) PropertyMappingProviderSamlMapOutput {
	return o
}

func (o PropertyMappingProviderSamlMapOutput) MapIndex(k pulumi.StringInput) PropertyMappingProviderSamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertyMappingProviderSaml {
		return vs[0].(map[string]*PropertyMappingProviderSaml)[vs[1].(string)]
	}).(PropertyMappingProviderSamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingProviderSamlInput)(nil)).Elem(), &PropertyMappingProviderSaml{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingProviderSamlArrayInput)(nil)).Elem(), PropertyMappingProviderSamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingProviderSamlMapInput)(nil)).Elem(), PropertyMappingProviderSamlMap{})
	pulumi.RegisterOutputType(PropertyMappingProviderSamlOutput{})
	pulumi.RegisterOutputType(PropertyMappingProviderSamlArrayOutput{})
	pulumi.RegisterOutputType(PropertyMappingProviderSamlMapOutput{})
}

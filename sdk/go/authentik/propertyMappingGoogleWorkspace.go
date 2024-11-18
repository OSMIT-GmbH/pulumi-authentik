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

// Manage Google Workspace Provider Property mappings
//
// > This resource is deprecated. Migrate to `PropertyMappingProviderGoogleWorkspace`.
type PropertyMappingGoogleWorkspace struct {
	pulumi.CustomResourceState

	Expression pulumi.StringOutput `pulumi:"expression"`
	Name       pulumi.StringOutput `pulumi:"name"`
}

// NewPropertyMappingGoogleWorkspace registers a new resource with the given unique name, arguments, and options.
func NewPropertyMappingGoogleWorkspace(ctx *pulumi.Context,
	name string, args *PropertyMappingGoogleWorkspaceArgs, opts ...pulumi.ResourceOption) (*PropertyMappingGoogleWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PropertyMappingGoogleWorkspace
	err := ctx.RegisterResource("authentik:index/propertyMappingGoogleWorkspace:PropertyMappingGoogleWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyMappingGoogleWorkspace gets an existing PropertyMappingGoogleWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyMappingGoogleWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyMappingGoogleWorkspaceState, opts ...pulumi.ResourceOption) (*PropertyMappingGoogleWorkspace, error) {
	var resource PropertyMappingGoogleWorkspace
	err := ctx.ReadResource("authentik:index/propertyMappingGoogleWorkspace:PropertyMappingGoogleWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyMappingGoogleWorkspace resources.
type propertyMappingGoogleWorkspaceState struct {
	Expression *string `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

type PropertyMappingGoogleWorkspaceState struct {
	Expression pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingGoogleWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingGoogleWorkspaceState)(nil)).Elem()
}

type propertyMappingGoogleWorkspaceArgs struct {
	Expression string  `pulumi:"expression"`
	Name       *string `pulumi:"name"`
}

// The set of arguments for constructing a PropertyMappingGoogleWorkspace resource.
type PropertyMappingGoogleWorkspaceArgs struct {
	Expression pulumi.StringInput
	Name       pulumi.StringPtrInput
}

func (PropertyMappingGoogleWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyMappingGoogleWorkspaceArgs)(nil)).Elem()
}

type PropertyMappingGoogleWorkspaceInput interface {
	pulumi.Input

	ToPropertyMappingGoogleWorkspaceOutput() PropertyMappingGoogleWorkspaceOutput
	ToPropertyMappingGoogleWorkspaceOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceOutput
}

func (*PropertyMappingGoogleWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingGoogleWorkspace)(nil)).Elem()
}

func (i *PropertyMappingGoogleWorkspace) ToPropertyMappingGoogleWorkspaceOutput() PropertyMappingGoogleWorkspaceOutput {
	return i.ToPropertyMappingGoogleWorkspaceOutputWithContext(context.Background())
}

func (i *PropertyMappingGoogleWorkspace) ToPropertyMappingGoogleWorkspaceOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingGoogleWorkspaceOutput)
}

// PropertyMappingGoogleWorkspaceArrayInput is an input type that accepts PropertyMappingGoogleWorkspaceArray and PropertyMappingGoogleWorkspaceArrayOutput values.
// You can construct a concrete instance of `PropertyMappingGoogleWorkspaceArrayInput` via:
//
//	PropertyMappingGoogleWorkspaceArray{ PropertyMappingGoogleWorkspaceArgs{...} }
type PropertyMappingGoogleWorkspaceArrayInput interface {
	pulumi.Input

	ToPropertyMappingGoogleWorkspaceArrayOutput() PropertyMappingGoogleWorkspaceArrayOutput
	ToPropertyMappingGoogleWorkspaceArrayOutputWithContext(context.Context) PropertyMappingGoogleWorkspaceArrayOutput
}

type PropertyMappingGoogleWorkspaceArray []PropertyMappingGoogleWorkspaceInput

func (PropertyMappingGoogleWorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingGoogleWorkspace)(nil)).Elem()
}

func (i PropertyMappingGoogleWorkspaceArray) ToPropertyMappingGoogleWorkspaceArrayOutput() PropertyMappingGoogleWorkspaceArrayOutput {
	return i.ToPropertyMappingGoogleWorkspaceArrayOutputWithContext(context.Background())
}

func (i PropertyMappingGoogleWorkspaceArray) ToPropertyMappingGoogleWorkspaceArrayOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingGoogleWorkspaceArrayOutput)
}

// PropertyMappingGoogleWorkspaceMapInput is an input type that accepts PropertyMappingGoogleWorkspaceMap and PropertyMappingGoogleWorkspaceMapOutput values.
// You can construct a concrete instance of `PropertyMappingGoogleWorkspaceMapInput` via:
//
//	PropertyMappingGoogleWorkspaceMap{ "key": PropertyMappingGoogleWorkspaceArgs{...} }
type PropertyMappingGoogleWorkspaceMapInput interface {
	pulumi.Input

	ToPropertyMappingGoogleWorkspaceMapOutput() PropertyMappingGoogleWorkspaceMapOutput
	ToPropertyMappingGoogleWorkspaceMapOutputWithContext(context.Context) PropertyMappingGoogleWorkspaceMapOutput
}

type PropertyMappingGoogleWorkspaceMap map[string]PropertyMappingGoogleWorkspaceInput

func (PropertyMappingGoogleWorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingGoogleWorkspace)(nil)).Elem()
}

func (i PropertyMappingGoogleWorkspaceMap) ToPropertyMappingGoogleWorkspaceMapOutput() PropertyMappingGoogleWorkspaceMapOutput {
	return i.ToPropertyMappingGoogleWorkspaceMapOutputWithContext(context.Background())
}

func (i PropertyMappingGoogleWorkspaceMap) ToPropertyMappingGoogleWorkspaceMapOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyMappingGoogleWorkspaceMapOutput)
}

type PropertyMappingGoogleWorkspaceOutput struct{ *pulumi.OutputState }

func (PropertyMappingGoogleWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyMappingGoogleWorkspace)(nil)).Elem()
}

func (o PropertyMappingGoogleWorkspaceOutput) ToPropertyMappingGoogleWorkspaceOutput() PropertyMappingGoogleWorkspaceOutput {
	return o
}

func (o PropertyMappingGoogleWorkspaceOutput) ToPropertyMappingGoogleWorkspaceOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceOutput {
	return o
}

func (o PropertyMappingGoogleWorkspaceOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingGoogleWorkspace) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

func (o PropertyMappingGoogleWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertyMappingGoogleWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type PropertyMappingGoogleWorkspaceArrayOutput struct{ *pulumi.OutputState }

func (PropertyMappingGoogleWorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyMappingGoogleWorkspace)(nil)).Elem()
}

func (o PropertyMappingGoogleWorkspaceArrayOutput) ToPropertyMappingGoogleWorkspaceArrayOutput() PropertyMappingGoogleWorkspaceArrayOutput {
	return o
}

func (o PropertyMappingGoogleWorkspaceArrayOutput) ToPropertyMappingGoogleWorkspaceArrayOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceArrayOutput {
	return o
}

func (o PropertyMappingGoogleWorkspaceArrayOutput) Index(i pulumi.IntInput) PropertyMappingGoogleWorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertyMappingGoogleWorkspace {
		return vs[0].([]*PropertyMappingGoogleWorkspace)[vs[1].(int)]
	}).(PropertyMappingGoogleWorkspaceOutput)
}

type PropertyMappingGoogleWorkspaceMapOutput struct{ *pulumi.OutputState }

func (PropertyMappingGoogleWorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyMappingGoogleWorkspace)(nil)).Elem()
}

func (o PropertyMappingGoogleWorkspaceMapOutput) ToPropertyMappingGoogleWorkspaceMapOutput() PropertyMappingGoogleWorkspaceMapOutput {
	return o
}

func (o PropertyMappingGoogleWorkspaceMapOutput) ToPropertyMappingGoogleWorkspaceMapOutputWithContext(ctx context.Context) PropertyMappingGoogleWorkspaceMapOutput {
	return o
}

func (o PropertyMappingGoogleWorkspaceMapOutput) MapIndex(k pulumi.StringInput) PropertyMappingGoogleWorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertyMappingGoogleWorkspace {
		return vs[0].(map[string]*PropertyMappingGoogleWorkspace)[vs[1].(string)]
	}).(PropertyMappingGoogleWorkspaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingGoogleWorkspaceInput)(nil)).Elem(), &PropertyMappingGoogleWorkspace{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingGoogleWorkspaceArrayInput)(nil)).Elem(), PropertyMappingGoogleWorkspaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyMappingGoogleWorkspaceMapInput)(nil)).Elem(), PropertyMappingGoogleWorkspaceMap{})
	pulumi.RegisterOutputType(PropertyMappingGoogleWorkspaceOutput{})
	pulumi.RegisterOutputType(PropertyMappingGoogleWorkspaceArrayOutput{})
	pulumi.RegisterOutputType(PropertyMappingGoogleWorkspaceMapOutput{})
}

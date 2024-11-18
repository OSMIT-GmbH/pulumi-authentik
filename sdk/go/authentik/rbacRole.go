// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RbacRole struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
}

// NewRbacRole registers a new resource with the given unique name, arguments, and options.
func NewRbacRole(ctx *pulumi.Context,
	name string, args *RbacRoleArgs, opts ...pulumi.ResourceOption) (*RbacRole, error) {
	if args == nil {
		args = &RbacRoleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RbacRole
	err := ctx.RegisterResource("authentik:index/rbacRole:RbacRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRbacRole gets an existing RbacRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRbacRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RbacRoleState, opts ...pulumi.ResourceOption) (*RbacRole, error) {
	var resource RbacRole
	err := ctx.ReadResource("authentik:index/rbacRole:RbacRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RbacRole resources.
type rbacRoleState struct {
	Name *string `pulumi:"name"`
}

type RbacRoleState struct {
	Name pulumi.StringPtrInput
}

func (RbacRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacRoleState)(nil)).Elem()
}

type rbacRoleArgs struct {
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a RbacRole resource.
type RbacRoleArgs struct {
	Name pulumi.StringPtrInput
}

func (RbacRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacRoleArgs)(nil)).Elem()
}

type RbacRoleInput interface {
	pulumi.Input

	ToRbacRoleOutput() RbacRoleOutput
	ToRbacRoleOutputWithContext(ctx context.Context) RbacRoleOutput
}

func (*RbacRole) ElementType() reflect.Type {
	return reflect.TypeOf((**RbacRole)(nil)).Elem()
}

func (i *RbacRole) ToRbacRoleOutput() RbacRoleOutput {
	return i.ToRbacRoleOutputWithContext(context.Background())
}

func (i *RbacRole) ToRbacRoleOutputWithContext(ctx context.Context) RbacRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacRoleOutput)
}

// RbacRoleArrayInput is an input type that accepts RbacRoleArray and RbacRoleArrayOutput values.
// You can construct a concrete instance of `RbacRoleArrayInput` via:
//
//	RbacRoleArray{ RbacRoleArgs{...} }
type RbacRoleArrayInput interface {
	pulumi.Input

	ToRbacRoleArrayOutput() RbacRoleArrayOutput
	ToRbacRoleArrayOutputWithContext(context.Context) RbacRoleArrayOutput
}

type RbacRoleArray []RbacRoleInput

func (RbacRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RbacRole)(nil)).Elem()
}

func (i RbacRoleArray) ToRbacRoleArrayOutput() RbacRoleArrayOutput {
	return i.ToRbacRoleArrayOutputWithContext(context.Background())
}

func (i RbacRoleArray) ToRbacRoleArrayOutputWithContext(ctx context.Context) RbacRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacRoleArrayOutput)
}

// RbacRoleMapInput is an input type that accepts RbacRoleMap and RbacRoleMapOutput values.
// You can construct a concrete instance of `RbacRoleMapInput` via:
//
//	RbacRoleMap{ "key": RbacRoleArgs{...} }
type RbacRoleMapInput interface {
	pulumi.Input

	ToRbacRoleMapOutput() RbacRoleMapOutput
	ToRbacRoleMapOutputWithContext(context.Context) RbacRoleMapOutput
}

type RbacRoleMap map[string]RbacRoleInput

func (RbacRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RbacRole)(nil)).Elem()
}

func (i RbacRoleMap) ToRbacRoleMapOutput() RbacRoleMapOutput {
	return i.ToRbacRoleMapOutputWithContext(context.Background())
}

func (i RbacRoleMap) ToRbacRoleMapOutputWithContext(ctx context.Context) RbacRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacRoleMapOutput)
}

type RbacRoleOutput struct{ *pulumi.OutputState }

func (RbacRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RbacRole)(nil)).Elem()
}

func (o RbacRoleOutput) ToRbacRoleOutput() RbacRoleOutput {
	return o
}

func (o RbacRoleOutput) ToRbacRoleOutputWithContext(ctx context.Context) RbacRoleOutput {
	return o
}

func (o RbacRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type RbacRoleArrayOutput struct{ *pulumi.OutputState }

func (RbacRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RbacRole)(nil)).Elem()
}

func (o RbacRoleArrayOutput) ToRbacRoleArrayOutput() RbacRoleArrayOutput {
	return o
}

func (o RbacRoleArrayOutput) ToRbacRoleArrayOutputWithContext(ctx context.Context) RbacRoleArrayOutput {
	return o
}

func (o RbacRoleArrayOutput) Index(i pulumi.IntInput) RbacRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RbacRole {
		return vs[0].([]*RbacRole)[vs[1].(int)]
	}).(RbacRoleOutput)
}

type RbacRoleMapOutput struct{ *pulumi.OutputState }

func (RbacRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RbacRole)(nil)).Elem()
}

func (o RbacRoleMapOutput) ToRbacRoleMapOutput() RbacRoleMapOutput {
	return o
}

func (o RbacRoleMapOutput) ToRbacRoleMapOutputWithContext(ctx context.Context) RbacRoleMapOutput {
	return o
}

func (o RbacRoleMapOutput) MapIndex(k pulumi.StringInput) RbacRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RbacRole {
		return vs[0].(map[string]*RbacRole)[vs[1].(string)]
	}).(RbacRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RbacRoleInput)(nil)).Elem(), &RbacRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*RbacRoleArrayInput)(nil)).Elem(), RbacRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RbacRoleMapInput)(nil)).Elem(), RbacRoleMap{})
	pulumi.RegisterOutputType(RbacRoleOutput{})
	pulumi.RegisterOutputType(RbacRoleArrayOutput{})
	pulumi.RegisterOutputType(RbacRoleMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			// Create a super-user group with a user
//			name, err := authentik.NewUser(ctx, "name", &authentik.UserArgs{
//				Username: pulumi.String("user"),
//				Name:     pulumi.String("User"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewGroup(ctx, "group", &authentik.GroupArgs{
//				Name: pulumi.String("tf_admins"),
//				Users: pulumi.IntArray{
//					name.ID(),
//				},
//				IsSuperuser: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes pulumi.StringPtrOutput `pulumi:"attributes"`
	// Defaults to `false`.
	IsSuperuser pulumi.BoolPtrOutput     `pulumi:"isSuperuser"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	Parent      pulumi.StringPtrOutput   `pulumi:"parent"`
	Roles       pulumi.StringArrayOutput `pulumi:"roles"`
	// Generated.
	Users pulumi.IntArrayOutput `pulumi:"users"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("authentik:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("authentik:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes *string `pulumi:"attributes"`
	// Defaults to `false`.
	IsSuperuser *bool    `pulumi:"isSuperuser"`
	Name        *string  `pulumi:"name"`
	Parent      *string  `pulumi:"parent"`
	Roles       []string `pulumi:"roles"`
	// Generated.
	Users []int `pulumi:"users"`
}

type GroupState struct {
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes pulumi.StringPtrInput
	// Defaults to `false`.
	IsSuperuser pulumi.BoolPtrInput
	Name        pulumi.StringPtrInput
	Parent      pulumi.StringPtrInput
	Roles       pulumi.StringArrayInput
	// Generated.
	Users pulumi.IntArrayInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes *string `pulumi:"attributes"`
	// Defaults to `false`.
	IsSuperuser *bool    `pulumi:"isSuperuser"`
	Name        *string  `pulumi:"name"`
	Parent      *string  `pulumi:"parent"`
	Roles       []string `pulumi:"roles"`
	// Generated.
	Users []int `pulumi:"users"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes pulumi.StringPtrInput
	// Defaults to `false`.
	IsSuperuser pulumi.BoolPtrInput
	Name        pulumi.StringPtrInput
	Parent      pulumi.StringPtrInput
	Roles       pulumi.StringArrayInput
	// Generated.
	Users pulumi.IntArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
func (o GroupOutput) Attributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Attributes }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o GroupOutput) IsSuperuser() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.IsSuperuser }).(pulumi.BoolPtrOutput)
}

func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GroupOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

func (o GroupOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// Generated.
func (o GroupOutput) Users() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.IntArrayOutput { return v.Users }).(pulumi.IntArrayOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}

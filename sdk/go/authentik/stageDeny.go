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
//			_, err := authentik.NewStageDeny(ctx, "name", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageDeny struct {
	pulumi.CustomResourceState

	DenyMessage pulumi.StringPtrOutput `pulumi:"denyMessage"`
	Name        pulumi.StringOutput    `pulumi:"name"`
}

// NewStageDeny registers a new resource with the given unique name, arguments, and options.
func NewStageDeny(ctx *pulumi.Context,
	name string, args *StageDenyArgs, opts ...pulumi.ResourceOption) (*StageDeny, error) {
	if args == nil {
		args = &StageDenyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageDeny
	err := ctx.RegisterResource("authentik:index/stageDeny:StageDeny", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageDeny gets an existing StageDeny resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageDeny(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageDenyState, opts ...pulumi.ResourceOption) (*StageDeny, error) {
	var resource StageDeny
	err := ctx.ReadResource("authentik:index/stageDeny:StageDeny", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageDeny resources.
type stageDenyState struct {
	DenyMessage *string `pulumi:"denyMessage"`
	Name        *string `pulumi:"name"`
}

type StageDenyState struct {
	DenyMessage pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
}

func (StageDenyState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageDenyState)(nil)).Elem()
}

type stageDenyArgs struct {
	DenyMessage *string `pulumi:"denyMessage"`
	Name        *string `pulumi:"name"`
}

// The set of arguments for constructing a StageDeny resource.
type StageDenyArgs struct {
	DenyMessage pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
}

func (StageDenyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageDenyArgs)(nil)).Elem()
}

type StageDenyInput interface {
	pulumi.Input

	ToStageDenyOutput() StageDenyOutput
	ToStageDenyOutputWithContext(ctx context.Context) StageDenyOutput
}

func (*StageDeny) ElementType() reflect.Type {
	return reflect.TypeOf((**StageDeny)(nil)).Elem()
}

func (i *StageDeny) ToStageDenyOutput() StageDenyOutput {
	return i.ToStageDenyOutputWithContext(context.Background())
}

func (i *StageDeny) ToStageDenyOutputWithContext(ctx context.Context) StageDenyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDenyOutput)
}

// StageDenyArrayInput is an input type that accepts StageDenyArray and StageDenyArrayOutput values.
// You can construct a concrete instance of `StageDenyArrayInput` via:
//
//	StageDenyArray{ StageDenyArgs{...} }
type StageDenyArrayInput interface {
	pulumi.Input

	ToStageDenyArrayOutput() StageDenyArrayOutput
	ToStageDenyArrayOutputWithContext(context.Context) StageDenyArrayOutput
}

type StageDenyArray []StageDenyInput

func (StageDenyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageDeny)(nil)).Elem()
}

func (i StageDenyArray) ToStageDenyArrayOutput() StageDenyArrayOutput {
	return i.ToStageDenyArrayOutputWithContext(context.Background())
}

func (i StageDenyArray) ToStageDenyArrayOutputWithContext(ctx context.Context) StageDenyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDenyArrayOutput)
}

// StageDenyMapInput is an input type that accepts StageDenyMap and StageDenyMapOutput values.
// You can construct a concrete instance of `StageDenyMapInput` via:
//
//	StageDenyMap{ "key": StageDenyArgs{...} }
type StageDenyMapInput interface {
	pulumi.Input

	ToStageDenyMapOutput() StageDenyMapOutput
	ToStageDenyMapOutputWithContext(context.Context) StageDenyMapOutput
}

type StageDenyMap map[string]StageDenyInput

func (StageDenyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageDeny)(nil)).Elem()
}

func (i StageDenyMap) ToStageDenyMapOutput() StageDenyMapOutput {
	return i.ToStageDenyMapOutputWithContext(context.Background())
}

func (i StageDenyMap) ToStageDenyMapOutputWithContext(ctx context.Context) StageDenyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDenyMapOutput)
}

type StageDenyOutput struct{ *pulumi.OutputState }

func (StageDenyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageDeny)(nil)).Elem()
}

func (o StageDenyOutput) ToStageDenyOutput() StageDenyOutput {
	return o
}

func (o StageDenyOutput) ToStageDenyOutputWithContext(ctx context.Context) StageDenyOutput {
	return o
}

func (o StageDenyOutput) DenyMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageDeny) pulumi.StringPtrOutput { return v.DenyMessage }).(pulumi.StringPtrOutput)
}

func (o StageDenyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageDeny) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type StageDenyArrayOutput struct{ *pulumi.OutputState }

func (StageDenyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageDeny)(nil)).Elem()
}

func (o StageDenyArrayOutput) ToStageDenyArrayOutput() StageDenyArrayOutput {
	return o
}

func (o StageDenyArrayOutput) ToStageDenyArrayOutputWithContext(ctx context.Context) StageDenyArrayOutput {
	return o
}

func (o StageDenyArrayOutput) Index(i pulumi.IntInput) StageDenyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageDeny {
		return vs[0].([]*StageDeny)[vs[1].(int)]
	}).(StageDenyOutput)
}

type StageDenyMapOutput struct{ *pulumi.OutputState }

func (StageDenyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageDeny)(nil)).Elem()
}

func (o StageDenyMapOutput) ToStageDenyMapOutput() StageDenyMapOutput {
	return o
}

func (o StageDenyMapOutput) ToStageDenyMapOutputWithContext(ctx context.Context) StageDenyMapOutput {
	return o
}

func (o StageDenyMapOutput) MapIndex(k pulumi.StringInput) StageDenyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageDeny {
		return vs[0].(map[string]*StageDeny)[vs[1].(string)]
	}).(StageDenyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageDenyInput)(nil)).Elem(), &StageDeny{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDenyArrayInput)(nil)).Elem(), StageDenyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDenyMapInput)(nil)).Elem(), StageDenyMap{})
	pulumi.RegisterOutputType(StageDenyOutput{})
	pulumi.RegisterOutputType(StageDenyArrayOutput{})
	pulumi.RegisterOutputType(StageDenyMapOutput{})
}

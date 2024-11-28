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
//			// Create dummy stage
//			_, err := authentik.NewStageDummy(ctx, "name", &authentik.StageDummyArgs{
//				Name: pulumi.String("dummy"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageDummy struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
}

// NewStageDummy registers a new resource with the given unique name, arguments, and options.
func NewStageDummy(ctx *pulumi.Context,
	name string, args *StageDummyArgs, opts ...pulumi.ResourceOption) (*StageDummy, error) {
	if args == nil {
		args = &StageDummyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageDummy
	err := ctx.RegisterResource("authentik:index/stageDummy:StageDummy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageDummy gets an existing StageDummy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageDummy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageDummyState, opts ...pulumi.ResourceOption) (*StageDummy, error) {
	var resource StageDummy
	err := ctx.ReadResource("authentik:index/stageDummy:StageDummy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageDummy resources.
type stageDummyState struct {
	Name *string `pulumi:"name"`
}

type StageDummyState struct {
	Name pulumi.StringPtrInput
}

func (StageDummyState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageDummyState)(nil)).Elem()
}

type stageDummyArgs struct {
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a StageDummy resource.
type StageDummyArgs struct {
	Name pulumi.StringPtrInput
}

func (StageDummyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageDummyArgs)(nil)).Elem()
}

type StageDummyInput interface {
	pulumi.Input

	ToStageDummyOutput() StageDummyOutput
	ToStageDummyOutputWithContext(ctx context.Context) StageDummyOutput
}

func (*StageDummy) ElementType() reflect.Type {
	return reflect.TypeOf((**StageDummy)(nil)).Elem()
}

func (i *StageDummy) ToStageDummyOutput() StageDummyOutput {
	return i.ToStageDummyOutputWithContext(context.Background())
}

func (i *StageDummy) ToStageDummyOutputWithContext(ctx context.Context) StageDummyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDummyOutput)
}

// StageDummyArrayInput is an input type that accepts StageDummyArray and StageDummyArrayOutput values.
// You can construct a concrete instance of `StageDummyArrayInput` via:
//
//	StageDummyArray{ StageDummyArgs{...} }
type StageDummyArrayInput interface {
	pulumi.Input

	ToStageDummyArrayOutput() StageDummyArrayOutput
	ToStageDummyArrayOutputWithContext(context.Context) StageDummyArrayOutput
}

type StageDummyArray []StageDummyInput

func (StageDummyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageDummy)(nil)).Elem()
}

func (i StageDummyArray) ToStageDummyArrayOutput() StageDummyArrayOutput {
	return i.ToStageDummyArrayOutputWithContext(context.Background())
}

func (i StageDummyArray) ToStageDummyArrayOutputWithContext(ctx context.Context) StageDummyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDummyArrayOutput)
}

// StageDummyMapInput is an input type that accepts StageDummyMap and StageDummyMapOutput values.
// You can construct a concrete instance of `StageDummyMapInput` via:
//
//	StageDummyMap{ "key": StageDummyArgs{...} }
type StageDummyMapInput interface {
	pulumi.Input

	ToStageDummyMapOutput() StageDummyMapOutput
	ToStageDummyMapOutputWithContext(context.Context) StageDummyMapOutput
}

type StageDummyMap map[string]StageDummyInput

func (StageDummyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageDummy)(nil)).Elem()
}

func (i StageDummyMap) ToStageDummyMapOutput() StageDummyMapOutput {
	return i.ToStageDummyMapOutputWithContext(context.Background())
}

func (i StageDummyMap) ToStageDummyMapOutputWithContext(ctx context.Context) StageDummyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageDummyMapOutput)
}

type StageDummyOutput struct{ *pulumi.OutputState }

func (StageDummyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageDummy)(nil)).Elem()
}

func (o StageDummyOutput) ToStageDummyOutput() StageDummyOutput {
	return o
}

func (o StageDummyOutput) ToStageDummyOutputWithContext(ctx context.Context) StageDummyOutput {
	return o
}

func (o StageDummyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageDummy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type StageDummyArrayOutput struct{ *pulumi.OutputState }

func (StageDummyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageDummy)(nil)).Elem()
}

func (o StageDummyArrayOutput) ToStageDummyArrayOutput() StageDummyArrayOutput {
	return o
}

func (o StageDummyArrayOutput) ToStageDummyArrayOutputWithContext(ctx context.Context) StageDummyArrayOutput {
	return o
}

func (o StageDummyArrayOutput) Index(i pulumi.IntInput) StageDummyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageDummy {
		return vs[0].([]*StageDummy)[vs[1].(int)]
	}).(StageDummyOutput)
}

type StageDummyMapOutput struct{ *pulumi.OutputState }

func (StageDummyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageDummy)(nil)).Elem()
}

func (o StageDummyMapOutput) ToStageDummyMapOutput() StageDummyMapOutput {
	return o
}

func (o StageDummyMapOutput) ToStageDummyMapOutputWithContext(ctx context.Context) StageDummyMapOutput {
	return o
}

func (o StageDummyMapOutput) MapIndex(k pulumi.StringInput) StageDummyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageDummy {
		return vs[0].(map[string]*StageDummy)[vs[1].(string)]
	}).(StageDummyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageDummyInput)(nil)).Elem(), &StageDummy{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDummyArrayInput)(nil)).Elem(), StageDummyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageDummyMapInput)(nil)).Elem(), StageDummyMap{})
	pulumi.RegisterOutputType(StageDummyOutput{})
	pulumi.RegisterOutputType(StageDummyArrayOutput{})
	pulumi.RegisterOutputType(StageDummyMapOutput{})
}

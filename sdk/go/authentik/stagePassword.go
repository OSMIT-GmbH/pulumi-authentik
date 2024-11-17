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
//			_, err := authentik.NewStagePassword(ctx, "test", &authentik.StagePasswordArgs{
//				Backends: pulumi.StringArray{
//					pulumi.String("authentik.core.auth.InbuiltBackend"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StagePassword struct {
	pulumi.CustomResourceState

	// Defaults to `false`.
	AllowShowPassword pulumi.BoolPtrOutput     `pulumi:"allowShowPassword"`
	Backends          pulumi.StringArrayOutput `pulumi:"backends"`
	ConfigureFlow     pulumi.StringPtrOutput   `pulumi:"configureFlow"`
	// Defaults to `5`.
	FailedAttemptsBeforeCancel pulumi.IntPtrOutput `pulumi:"failedAttemptsBeforeCancel"`
	Name                       pulumi.StringOutput `pulumi:"name"`
}

// NewStagePassword registers a new resource with the given unique name, arguments, and options.
func NewStagePassword(ctx *pulumi.Context,
	name string, args *StagePasswordArgs, opts ...pulumi.ResourceOption) (*StagePassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backends == nil {
		return nil, errors.New("invalid value for required argument 'Backends'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StagePassword
	err := ctx.RegisterResource("authentik:index/stagePassword:StagePassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStagePassword gets an existing StagePassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStagePassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StagePasswordState, opts ...pulumi.ResourceOption) (*StagePassword, error) {
	var resource StagePassword
	err := ctx.ReadResource("authentik:index/stagePassword:StagePassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StagePassword resources.
type stagePasswordState struct {
	// Defaults to `false`.
	AllowShowPassword *bool    `pulumi:"allowShowPassword"`
	Backends          []string `pulumi:"backends"`
	ConfigureFlow     *string  `pulumi:"configureFlow"`
	// Defaults to `5`.
	FailedAttemptsBeforeCancel *int    `pulumi:"failedAttemptsBeforeCancel"`
	Name                       *string `pulumi:"name"`
}

type StagePasswordState struct {
	// Defaults to `false`.
	AllowShowPassword pulumi.BoolPtrInput
	Backends          pulumi.StringArrayInput
	ConfigureFlow     pulumi.StringPtrInput
	// Defaults to `5`.
	FailedAttemptsBeforeCancel pulumi.IntPtrInput
	Name                       pulumi.StringPtrInput
}

func (StagePasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*stagePasswordState)(nil)).Elem()
}

type stagePasswordArgs struct {
	// Defaults to `false`.
	AllowShowPassword *bool    `pulumi:"allowShowPassword"`
	Backends          []string `pulumi:"backends"`
	ConfigureFlow     *string  `pulumi:"configureFlow"`
	// Defaults to `5`.
	FailedAttemptsBeforeCancel *int    `pulumi:"failedAttemptsBeforeCancel"`
	Name                       *string `pulumi:"name"`
}

// The set of arguments for constructing a StagePassword resource.
type StagePasswordArgs struct {
	// Defaults to `false`.
	AllowShowPassword pulumi.BoolPtrInput
	Backends          pulumi.StringArrayInput
	ConfigureFlow     pulumi.StringPtrInput
	// Defaults to `5`.
	FailedAttemptsBeforeCancel pulumi.IntPtrInput
	Name                       pulumi.StringPtrInput
}

func (StagePasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stagePasswordArgs)(nil)).Elem()
}

type StagePasswordInput interface {
	pulumi.Input

	ToStagePasswordOutput() StagePasswordOutput
	ToStagePasswordOutputWithContext(ctx context.Context) StagePasswordOutput
}

func (*StagePassword) ElementType() reflect.Type {
	return reflect.TypeOf((**StagePassword)(nil)).Elem()
}

func (i *StagePassword) ToStagePasswordOutput() StagePasswordOutput {
	return i.ToStagePasswordOutputWithContext(context.Background())
}

func (i *StagePassword) ToStagePasswordOutputWithContext(ctx context.Context) StagePasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePasswordOutput)
}

// StagePasswordArrayInput is an input type that accepts StagePasswordArray and StagePasswordArrayOutput values.
// You can construct a concrete instance of `StagePasswordArrayInput` via:
//
//	StagePasswordArray{ StagePasswordArgs{...} }
type StagePasswordArrayInput interface {
	pulumi.Input

	ToStagePasswordArrayOutput() StagePasswordArrayOutput
	ToStagePasswordArrayOutputWithContext(context.Context) StagePasswordArrayOutput
}

type StagePasswordArray []StagePasswordInput

func (StagePasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StagePassword)(nil)).Elem()
}

func (i StagePasswordArray) ToStagePasswordArrayOutput() StagePasswordArrayOutput {
	return i.ToStagePasswordArrayOutputWithContext(context.Background())
}

func (i StagePasswordArray) ToStagePasswordArrayOutputWithContext(ctx context.Context) StagePasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePasswordArrayOutput)
}

// StagePasswordMapInput is an input type that accepts StagePasswordMap and StagePasswordMapOutput values.
// You can construct a concrete instance of `StagePasswordMapInput` via:
//
//	StagePasswordMap{ "key": StagePasswordArgs{...} }
type StagePasswordMapInput interface {
	pulumi.Input

	ToStagePasswordMapOutput() StagePasswordMapOutput
	ToStagePasswordMapOutputWithContext(context.Context) StagePasswordMapOutput
}

type StagePasswordMap map[string]StagePasswordInput

func (StagePasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StagePassword)(nil)).Elem()
}

func (i StagePasswordMap) ToStagePasswordMapOutput() StagePasswordMapOutput {
	return i.ToStagePasswordMapOutputWithContext(context.Background())
}

func (i StagePasswordMap) ToStagePasswordMapOutputWithContext(ctx context.Context) StagePasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePasswordMapOutput)
}

type StagePasswordOutput struct{ *pulumi.OutputState }

func (StagePasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StagePassword)(nil)).Elem()
}

func (o StagePasswordOutput) ToStagePasswordOutput() StagePasswordOutput {
	return o
}

func (o StagePasswordOutput) ToStagePasswordOutputWithContext(ctx context.Context) StagePasswordOutput {
	return o
}

// Defaults to `false`.
func (o StagePasswordOutput) AllowShowPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StagePassword) pulumi.BoolPtrOutput { return v.AllowShowPassword }).(pulumi.BoolPtrOutput)
}

func (o StagePasswordOutput) Backends() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StagePassword) pulumi.StringArrayOutput { return v.Backends }).(pulumi.StringArrayOutput)
}

func (o StagePasswordOutput) ConfigureFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StagePassword) pulumi.StringPtrOutput { return v.ConfigureFlow }).(pulumi.StringPtrOutput)
}

// Defaults to `5`.
func (o StagePasswordOutput) FailedAttemptsBeforeCancel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StagePassword) pulumi.IntPtrOutput { return v.FailedAttemptsBeforeCancel }).(pulumi.IntPtrOutput)
}

func (o StagePasswordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StagePassword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type StagePasswordArrayOutput struct{ *pulumi.OutputState }

func (StagePasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StagePassword)(nil)).Elem()
}

func (o StagePasswordArrayOutput) ToStagePasswordArrayOutput() StagePasswordArrayOutput {
	return o
}

func (o StagePasswordArrayOutput) ToStagePasswordArrayOutputWithContext(ctx context.Context) StagePasswordArrayOutput {
	return o
}

func (o StagePasswordArrayOutput) Index(i pulumi.IntInput) StagePasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StagePassword {
		return vs[0].([]*StagePassword)[vs[1].(int)]
	}).(StagePasswordOutput)
}

type StagePasswordMapOutput struct{ *pulumi.OutputState }

func (StagePasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StagePassword)(nil)).Elem()
}

func (o StagePasswordMapOutput) ToStagePasswordMapOutput() StagePasswordMapOutput {
	return o
}

func (o StagePasswordMapOutput) ToStagePasswordMapOutputWithContext(ctx context.Context) StagePasswordMapOutput {
	return o
}

func (o StagePasswordMapOutput) MapIndex(k pulumi.StringInput) StagePasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StagePassword {
		return vs[0].(map[string]*StagePassword)[vs[1].(string)]
	}).(StagePasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StagePasswordInput)(nil)).Elem(), &StagePassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePasswordArrayInput)(nil)).Elem(), StagePasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePasswordMapInput)(nil)).Elem(), StagePasswordMap{})
	pulumi.RegisterOutputType(StagePasswordOutput{})
	pulumi.RegisterOutputType(StagePasswordArrayOutput{})
	pulumi.RegisterOutputType(StagePasswordMapOutput{})
}

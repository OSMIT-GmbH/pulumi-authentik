// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create consent stage
//			_, err := authentik.NewStageConsent(ctx, "name", &authentik.StageConsentArgs{
//				Name: pulumi.String("consent"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageConsent struct {
	pulumi.CustomResourceState

	ConsentExpireIn pulumi.StringPtrOutput `pulumi:"consentExpireIn"`
	// Allowed values: - `alwaysRequire` - `permanent` - `expiring`
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	Name pulumi.StringOutput    `pulumi:"name"`
}

// NewStageConsent registers a new resource with the given unique name, arguments, and options.
func NewStageConsent(ctx *pulumi.Context,
	name string, args *StageConsentArgs, opts ...pulumi.ResourceOption) (*StageConsent, error) {
	if args == nil {
		args = &StageConsentArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageConsent
	err := ctx.RegisterResource("authentik:index/stageConsent:StageConsent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageConsent gets an existing StageConsent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageConsent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageConsentState, opts ...pulumi.ResourceOption) (*StageConsent, error) {
	var resource StageConsent
	err := ctx.ReadResource("authentik:index/stageConsent:StageConsent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageConsent resources.
type stageConsentState struct {
	ConsentExpireIn *string `pulumi:"consentExpireIn"`
	// Allowed values: - `alwaysRequire` - `permanent` - `expiring`
	Mode *string `pulumi:"mode"`
	Name *string `pulumi:"name"`
}

type StageConsentState struct {
	ConsentExpireIn pulumi.StringPtrInput
	// Allowed values: - `alwaysRequire` - `permanent` - `expiring`
	Mode pulumi.StringPtrInput
	Name pulumi.StringPtrInput
}

func (StageConsentState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageConsentState)(nil)).Elem()
}

type stageConsentArgs struct {
	ConsentExpireIn *string `pulumi:"consentExpireIn"`
	// Allowed values: - `alwaysRequire` - `permanent` - `expiring`
	Mode *string `pulumi:"mode"`
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a StageConsent resource.
type StageConsentArgs struct {
	ConsentExpireIn pulumi.StringPtrInput
	// Allowed values: - `alwaysRequire` - `permanent` - `expiring`
	Mode pulumi.StringPtrInput
	Name pulumi.StringPtrInput
}

func (StageConsentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageConsentArgs)(nil)).Elem()
}

type StageConsentInput interface {
	pulumi.Input

	ToStageConsentOutput() StageConsentOutput
	ToStageConsentOutputWithContext(ctx context.Context) StageConsentOutput
}

func (*StageConsent) ElementType() reflect.Type {
	return reflect.TypeOf((**StageConsent)(nil)).Elem()
}

func (i *StageConsent) ToStageConsentOutput() StageConsentOutput {
	return i.ToStageConsentOutputWithContext(context.Background())
}

func (i *StageConsent) ToStageConsentOutputWithContext(ctx context.Context) StageConsentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageConsentOutput)
}

// StageConsentArrayInput is an input type that accepts StageConsentArray and StageConsentArrayOutput values.
// You can construct a concrete instance of `StageConsentArrayInput` via:
//
//	StageConsentArray{ StageConsentArgs{...} }
type StageConsentArrayInput interface {
	pulumi.Input

	ToStageConsentArrayOutput() StageConsentArrayOutput
	ToStageConsentArrayOutputWithContext(context.Context) StageConsentArrayOutput
}

type StageConsentArray []StageConsentInput

func (StageConsentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageConsent)(nil)).Elem()
}

func (i StageConsentArray) ToStageConsentArrayOutput() StageConsentArrayOutput {
	return i.ToStageConsentArrayOutputWithContext(context.Background())
}

func (i StageConsentArray) ToStageConsentArrayOutputWithContext(ctx context.Context) StageConsentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageConsentArrayOutput)
}

// StageConsentMapInput is an input type that accepts StageConsentMap and StageConsentMapOutput values.
// You can construct a concrete instance of `StageConsentMapInput` via:
//
//	StageConsentMap{ "key": StageConsentArgs{...} }
type StageConsentMapInput interface {
	pulumi.Input

	ToStageConsentMapOutput() StageConsentMapOutput
	ToStageConsentMapOutputWithContext(context.Context) StageConsentMapOutput
}

type StageConsentMap map[string]StageConsentInput

func (StageConsentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageConsent)(nil)).Elem()
}

func (i StageConsentMap) ToStageConsentMapOutput() StageConsentMapOutput {
	return i.ToStageConsentMapOutputWithContext(context.Background())
}

func (i StageConsentMap) ToStageConsentMapOutputWithContext(ctx context.Context) StageConsentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageConsentMapOutput)
}

type StageConsentOutput struct{ *pulumi.OutputState }

func (StageConsentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageConsent)(nil)).Elem()
}

func (o StageConsentOutput) ToStageConsentOutput() StageConsentOutput {
	return o
}

func (o StageConsentOutput) ToStageConsentOutputWithContext(ctx context.Context) StageConsentOutput {
	return o
}

func (o StageConsentOutput) ConsentExpireIn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageConsent) pulumi.StringPtrOutput { return v.ConsentExpireIn }).(pulumi.StringPtrOutput)
}

// Allowed values: - `alwaysRequire` - `permanent` - `expiring`
func (o StageConsentOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageConsent) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o StageConsentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageConsent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type StageConsentArrayOutput struct{ *pulumi.OutputState }

func (StageConsentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageConsent)(nil)).Elem()
}

func (o StageConsentArrayOutput) ToStageConsentArrayOutput() StageConsentArrayOutput {
	return o
}

func (o StageConsentArrayOutput) ToStageConsentArrayOutputWithContext(ctx context.Context) StageConsentArrayOutput {
	return o
}

func (o StageConsentArrayOutput) Index(i pulumi.IntInput) StageConsentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageConsent {
		return vs[0].([]*StageConsent)[vs[1].(int)]
	}).(StageConsentOutput)
}

type StageConsentMapOutput struct{ *pulumi.OutputState }

func (StageConsentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageConsent)(nil)).Elem()
}

func (o StageConsentMapOutput) ToStageConsentMapOutput() StageConsentMapOutput {
	return o
}

func (o StageConsentMapOutput) ToStageConsentMapOutputWithContext(ctx context.Context) StageConsentMapOutput {
	return o
}

func (o StageConsentMapOutput) MapIndex(k pulumi.StringInput) StageConsentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageConsent {
		return vs[0].(map[string]*StageConsent)[vs[1].(string)]
	}).(StageConsentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageConsentInput)(nil)).Elem(), &StageConsent{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageConsentArrayInput)(nil)).Elem(), StageConsentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageConsentMapInput)(nil)).Elem(), StageConsentMap{})
	pulumi.RegisterOutputType(StageConsentOutput{})
	pulumi.RegisterOutputType(StageConsentArrayOutput{})
	pulumi.RegisterOutputType(StageConsentMapOutput{})
}

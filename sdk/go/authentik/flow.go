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
//			// Create a flow with a stage attached
//			name, err := authentik.NewStageDummy(ctx, "name", &authentik.StageDummyArgs{
//				Name: pulumi.String("test-stage"),
//			})
//			if err != nil {
//				return err
//			}
//			flow, err := authentik.NewFlow(ctx, "flow", &authentik.FlowArgs{
//				Name:        pulumi.String("test-flow"),
//				Title:       pulumi.String("Test flow"),
//				Slug:        pulumi.String("test-flow"),
//				Designation: pulumi.String("authorization"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewFlowStageBinding(ctx, "dummy-flow", &authentik.FlowStageBindingArgs{
//				Target: flow.Uuid,
//				Stage:  name.ID(),
//				Order:  pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Flow struct {
	pulumi.CustomResourceState

	// Allowed values: - `none` - `requireAuthenticated` - `requireUnauthenticated` - `requireSuperuser` - `requireOutpost`
	Authentication pulumi.StringPtrOutput `pulumi:"authentication"`
	// Optional URL to an image which will be used as the background during the flow.
	Background        pulumi.StringPtrOutput `pulumi:"background"`
	CompatibilityMode pulumi.BoolPtrOutput   `pulumi:"compatibilityMode"`
	DeniedAction      pulumi.StringPtrOutput `pulumi:"deniedAction"`
	// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
	// `stageConfiguration`
	Designation pulumi.StringOutput `pulumi:"designation"`
	// Allowed values: - `stacked` - `contentLeft` - `contentRight` - `sidebarLeft` - `sidebarRight`
	Layout pulumi.StringPtrOutput `pulumi:"layout"`
	Name   pulumi.StringOutput    `pulumi:"name"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrOutput `pulumi:"policyEngineMode"`
	Slug             pulumi.StringOutput    `pulumi:"slug"`
	Title            pulumi.StringOutput    `pulumi:"title"`
	Uuid             pulumi.StringOutput    `pulumi:"uuid"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Designation == nil {
		return nil, errors.New("invalid value for required argument 'Designation'")
	}
	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flow
	err := ctx.RegisterResource("authentik:index/flow:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("authentik:index/flow:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
	// Allowed values: - `none` - `requireAuthenticated` - `requireUnauthenticated` - `requireSuperuser` - `requireOutpost`
	Authentication *string `pulumi:"authentication"`
	// Optional URL to an image which will be used as the background during the flow.
	Background        *string `pulumi:"background"`
	CompatibilityMode *bool   `pulumi:"compatibilityMode"`
	DeniedAction      *string `pulumi:"deniedAction"`
	// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
	// `stageConfiguration`
	Designation *string `pulumi:"designation"`
	// Allowed values: - `stacked` - `contentLeft` - `contentRight` - `sidebarLeft` - `sidebarRight`
	Layout *string `pulumi:"layout"`
	Name   *string `pulumi:"name"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	Slug             *string `pulumi:"slug"`
	Title            *string `pulumi:"title"`
	Uuid             *string `pulumi:"uuid"`
}

type FlowState struct {
	// Allowed values: - `none` - `requireAuthenticated` - `requireUnauthenticated` - `requireSuperuser` - `requireOutpost`
	Authentication pulumi.StringPtrInput
	// Optional URL to an image which will be used as the background during the flow.
	Background        pulumi.StringPtrInput
	CompatibilityMode pulumi.BoolPtrInput
	DeniedAction      pulumi.StringPtrInput
	// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
	// `stageConfiguration`
	Designation pulumi.StringPtrInput
	// Allowed values: - `stacked` - `contentLeft` - `contentRight` - `sidebarLeft` - `sidebarRight`
	Layout pulumi.StringPtrInput
	Name   pulumi.StringPtrInput
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrInput
	Slug             pulumi.StringPtrInput
	Title            pulumi.StringPtrInput
	Uuid             pulumi.StringPtrInput
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	// Allowed values: - `none` - `requireAuthenticated` - `requireUnauthenticated` - `requireSuperuser` - `requireOutpost`
	Authentication *string `pulumi:"authentication"`
	// Optional URL to an image which will be used as the background during the flow.
	Background        *string `pulumi:"background"`
	CompatibilityMode *bool   `pulumi:"compatibilityMode"`
	DeniedAction      *string `pulumi:"deniedAction"`
	// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
	// `stageConfiguration`
	Designation string `pulumi:"designation"`
	// Allowed values: - `stacked` - `contentLeft` - `contentRight` - `sidebarLeft` - `sidebarRight`
	Layout *string `pulumi:"layout"`
	Name   *string `pulumi:"name"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	Slug             string  `pulumi:"slug"`
	Title            string  `pulumi:"title"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	// Allowed values: - `none` - `requireAuthenticated` - `requireUnauthenticated` - `requireSuperuser` - `requireOutpost`
	Authentication pulumi.StringPtrInput
	// Optional URL to an image which will be used as the background during the flow.
	Background        pulumi.StringPtrInput
	CompatibilityMode pulumi.BoolPtrInput
	DeniedAction      pulumi.StringPtrInput
	// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
	// `stageConfiguration`
	Designation pulumi.StringInput
	// Allowed values: - `stacked` - `contentLeft` - `contentRight` - `sidebarLeft` - `sidebarRight`
	Layout pulumi.StringPtrInput
	Name   pulumi.StringPtrInput
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrInput
	Slug             pulumi.StringInput
	Title            pulumi.StringInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}

type FlowInput interface {
	pulumi.Input

	ToFlowOutput() FlowOutput
	ToFlowOutputWithContext(ctx context.Context) FlowOutput
}

func (*Flow) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil)).Elem()
}

func (i *Flow) ToFlowOutput() FlowOutput {
	return i.ToFlowOutputWithContext(context.Background())
}

func (i *Flow) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowOutput)
}

// FlowArrayInput is an input type that accepts FlowArray and FlowArrayOutput values.
// You can construct a concrete instance of `FlowArrayInput` via:
//
//	FlowArray{ FlowArgs{...} }
type FlowArrayInput interface {
	pulumi.Input

	ToFlowArrayOutput() FlowArrayOutput
	ToFlowArrayOutputWithContext(context.Context) FlowArrayOutput
}

type FlowArray []FlowInput

func (FlowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flow)(nil)).Elem()
}

func (i FlowArray) ToFlowArrayOutput() FlowArrayOutput {
	return i.ToFlowArrayOutputWithContext(context.Background())
}

func (i FlowArray) ToFlowArrayOutputWithContext(ctx context.Context) FlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowArrayOutput)
}

// FlowMapInput is an input type that accepts FlowMap and FlowMapOutput values.
// You can construct a concrete instance of `FlowMapInput` via:
//
//	FlowMap{ "key": FlowArgs{...} }
type FlowMapInput interface {
	pulumi.Input

	ToFlowMapOutput() FlowMapOutput
	ToFlowMapOutputWithContext(context.Context) FlowMapOutput
}

type FlowMap map[string]FlowInput

func (FlowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flow)(nil)).Elem()
}

func (i FlowMap) ToFlowMapOutput() FlowMapOutput {
	return i.ToFlowMapOutputWithContext(context.Background())
}

func (i FlowMap) ToFlowMapOutputWithContext(ctx context.Context) FlowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowMapOutput)
}

type FlowOutput struct{ *pulumi.OutputState }

func (FlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil)).Elem()
}

func (o FlowOutput) ToFlowOutput() FlowOutput {
	return o
}

func (o FlowOutput) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return o
}

// Allowed values: - `none` - `requireAuthenticated` - `requireUnauthenticated` - `requireSuperuser` - `requireOutpost`
func (o FlowOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.Authentication }).(pulumi.StringPtrOutput)
}

// Optional URL to an image which will be used as the background during the flow.
func (o FlowOutput) Background() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.Background }).(pulumi.StringPtrOutput)
}

func (o FlowOutput) CompatibilityMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.BoolPtrOutput { return v.CompatibilityMode }).(pulumi.BoolPtrOutput)
}

func (o FlowOutput) DeniedAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.DeniedAction }).(pulumi.StringPtrOutput)
}

// Allowed values: - `authentication` - `authorization` - `invalidation` - `enrollment` - `unenrollment` - `recovery` -
// `stageConfiguration`
func (o FlowOutput) Designation() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Designation }).(pulumi.StringOutput)
}

// Allowed values: - `stacked` - `contentLeft` - `contentRight` - `sidebarLeft` - `sidebarRight`
func (o FlowOutput) Layout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.Layout }).(pulumi.StringPtrOutput)
}

func (o FlowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Allowed values: - `all` - `any`
func (o FlowOutput) PolicyEngineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.PolicyEngineMode }).(pulumi.StringPtrOutput)
}

func (o FlowOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o FlowOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

func (o FlowOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type FlowArrayOutput struct{ *pulumi.OutputState }

func (FlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flow)(nil)).Elem()
}

func (o FlowArrayOutput) ToFlowArrayOutput() FlowArrayOutput {
	return o
}

func (o FlowArrayOutput) ToFlowArrayOutputWithContext(ctx context.Context) FlowArrayOutput {
	return o
}

func (o FlowArrayOutput) Index(i pulumi.IntInput) FlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Flow {
		return vs[0].([]*Flow)[vs[1].(int)]
	}).(FlowOutput)
}

type FlowMapOutput struct{ *pulumi.OutputState }

func (FlowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flow)(nil)).Elem()
}

func (o FlowMapOutput) ToFlowMapOutput() FlowMapOutput {
	return o
}

func (o FlowMapOutput) ToFlowMapOutputWithContext(ctx context.Context) FlowMapOutput {
	return o
}

func (o FlowMapOutput) MapIndex(k pulumi.StringInput) FlowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Flow {
		return vs[0].(map[string]*Flow)[vs[1].(string)]
	}).(FlowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowInput)(nil)).Elem(), &Flow{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowArrayInput)(nil)).Elem(), FlowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowMapInput)(nil)).Elem(), FlowMap{})
	pulumi.RegisterOutputType(FlowOutput{})
	pulumi.RegisterOutputType(FlowArrayOutput{})
	pulumi.RegisterOutputType(FlowMapOutput{})
}

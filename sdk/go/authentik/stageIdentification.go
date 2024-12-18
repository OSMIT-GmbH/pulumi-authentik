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
//			// Create identification stage with sources and showing a password field
//			default_authorization_flow, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-provider-authorization-implicit-consent"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			name, err := authentik.NewSourceOauth(ctx, "name", &authentik.SourceOauthArgs{
//				Name:               pulumi.String("test"),
//				Slug:               pulumi.String("test"),
//				AuthenticationFlow: pulumi.String(default_authorization_flow.Id),
//				EnrollmentFlow:     pulumi.String(default_authorization_flow.Id),
//				ProviderType:       pulumi.String("discord"),
//				ConsumerKey:        pulumi.String("foo"),
//				ConsumerSecret:     pulumi.String("bar"),
//			})
//			if err != nil {
//				return err
//			}
//			nameStagePassword, err := authentik.NewStagePassword(ctx, "name", &authentik.StagePasswordArgs{
//				Name: pulumi.String("test-pass"),
//				Backends: pulumi.StringArray{
//					pulumi.String("authentik.core.auth.InbuiltBackend"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewStageIdentification(ctx, "name", &authentik.StageIdentificationArgs{
//				Name: pulumi.String("test-ident"),
//				UserFields: pulumi.StringArray{
//					pulumi.String("username"),
//				},
//				Sources: pulumi.StringArray{
//					name.Uuid,
//				},
//				PasswordStage: nameStagePassword.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StageIdentification struct {
	pulumi.CustomResourceState

	CaptchaStage            pulumi.StringPtrOutput `pulumi:"captchaStage"`
	CaseInsensitiveMatching pulumi.BoolPtrOutput   `pulumi:"caseInsensitiveMatching"`
	EnrollmentFlow          pulumi.StringPtrOutput `pulumi:"enrollmentFlow"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	PasswordStage           pulumi.StringPtrOutput `pulumi:"passwordStage"`
	PasswordlessFlow        pulumi.StringPtrOutput `pulumi:"passwordlessFlow"`
	// Defaults to `true`.
	PretendUserExists pulumi.BoolPtrOutput   `pulumi:"pretendUserExists"`
	RecoveryFlow      pulumi.StringPtrOutput `pulumi:"recoveryFlow"`
	// Defaults to `true`.
	ShowMatchedUser pulumi.BoolPtrOutput `pulumi:"showMatchedUser"`
	// Defaults to `false`.
	ShowSourceLabels pulumi.BoolPtrOutput     `pulumi:"showSourceLabels"`
	Sources          pulumi.StringArrayOutput `pulumi:"sources"`
	UserFields       pulumi.StringArrayOutput `pulumi:"userFields"`
}

// NewStageIdentification registers a new resource with the given unique name, arguments, and options.
func NewStageIdentification(ctx *pulumi.Context,
	name string, args *StageIdentificationArgs, opts ...pulumi.ResourceOption) (*StageIdentification, error) {
	if args == nil {
		args = &StageIdentificationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageIdentification
	err := ctx.RegisterResource("authentik:index/stageIdentification:StageIdentification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageIdentification gets an existing StageIdentification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageIdentification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageIdentificationState, opts ...pulumi.ResourceOption) (*StageIdentification, error) {
	var resource StageIdentification
	err := ctx.ReadResource("authentik:index/stageIdentification:StageIdentification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageIdentification resources.
type stageIdentificationState struct {
	CaptchaStage            *string `pulumi:"captchaStage"`
	CaseInsensitiveMatching *bool   `pulumi:"caseInsensitiveMatching"`
	EnrollmentFlow          *string `pulumi:"enrollmentFlow"`
	Name                    *string `pulumi:"name"`
	PasswordStage           *string `pulumi:"passwordStage"`
	PasswordlessFlow        *string `pulumi:"passwordlessFlow"`
	// Defaults to `true`.
	PretendUserExists *bool   `pulumi:"pretendUserExists"`
	RecoveryFlow      *string `pulumi:"recoveryFlow"`
	// Defaults to `true`.
	ShowMatchedUser *bool `pulumi:"showMatchedUser"`
	// Defaults to `false`.
	ShowSourceLabels *bool    `pulumi:"showSourceLabels"`
	Sources          []string `pulumi:"sources"`
	UserFields       []string `pulumi:"userFields"`
}

type StageIdentificationState struct {
	CaptchaStage            pulumi.StringPtrInput
	CaseInsensitiveMatching pulumi.BoolPtrInput
	EnrollmentFlow          pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	PasswordStage           pulumi.StringPtrInput
	PasswordlessFlow        pulumi.StringPtrInput
	// Defaults to `true`.
	PretendUserExists pulumi.BoolPtrInput
	RecoveryFlow      pulumi.StringPtrInput
	// Defaults to `true`.
	ShowMatchedUser pulumi.BoolPtrInput
	// Defaults to `false`.
	ShowSourceLabels pulumi.BoolPtrInput
	Sources          pulumi.StringArrayInput
	UserFields       pulumi.StringArrayInput
}

func (StageIdentificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageIdentificationState)(nil)).Elem()
}

type stageIdentificationArgs struct {
	CaptchaStage            *string `pulumi:"captchaStage"`
	CaseInsensitiveMatching *bool   `pulumi:"caseInsensitiveMatching"`
	EnrollmentFlow          *string `pulumi:"enrollmentFlow"`
	Name                    *string `pulumi:"name"`
	PasswordStage           *string `pulumi:"passwordStage"`
	PasswordlessFlow        *string `pulumi:"passwordlessFlow"`
	// Defaults to `true`.
	PretendUserExists *bool   `pulumi:"pretendUserExists"`
	RecoveryFlow      *string `pulumi:"recoveryFlow"`
	// Defaults to `true`.
	ShowMatchedUser *bool `pulumi:"showMatchedUser"`
	// Defaults to `false`.
	ShowSourceLabels *bool    `pulumi:"showSourceLabels"`
	Sources          []string `pulumi:"sources"`
	UserFields       []string `pulumi:"userFields"`
}

// The set of arguments for constructing a StageIdentification resource.
type StageIdentificationArgs struct {
	CaptchaStage            pulumi.StringPtrInput
	CaseInsensitiveMatching pulumi.BoolPtrInput
	EnrollmentFlow          pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	PasswordStage           pulumi.StringPtrInput
	PasswordlessFlow        pulumi.StringPtrInput
	// Defaults to `true`.
	PretendUserExists pulumi.BoolPtrInput
	RecoveryFlow      pulumi.StringPtrInput
	// Defaults to `true`.
	ShowMatchedUser pulumi.BoolPtrInput
	// Defaults to `false`.
	ShowSourceLabels pulumi.BoolPtrInput
	Sources          pulumi.StringArrayInput
	UserFields       pulumi.StringArrayInput
}

func (StageIdentificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageIdentificationArgs)(nil)).Elem()
}

type StageIdentificationInput interface {
	pulumi.Input

	ToStageIdentificationOutput() StageIdentificationOutput
	ToStageIdentificationOutputWithContext(ctx context.Context) StageIdentificationOutput
}

func (*StageIdentification) ElementType() reflect.Type {
	return reflect.TypeOf((**StageIdentification)(nil)).Elem()
}

func (i *StageIdentification) ToStageIdentificationOutput() StageIdentificationOutput {
	return i.ToStageIdentificationOutputWithContext(context.Background())
}

func (i *StageIdentification) ToStageIdentificationOutputWithContext(ctx context.Context) StageIdentificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageIdentificationOutput)
}

// StageIdentificationArrayInput is an input type that accepts StageIdentificationArray and StageIdentificationArrayOutput values.
// You can construct a concrete instance of `StageIdentificationArrayInput` via:
//
//	StageIdentificationArray{ StageIdentificationArgs{...} }
type StageIdentificationArrayInput interface {
	pulumi.Input

	ToStageIdentificationArrayOutput() StageIdentificationArrayOutput
	ToStageIdentificationArrayOutputWithContext(context.Context) StageIdentificationArrayOutput
}

type StageIdentificationArray []StageIdentificationInput

func (StageIdentificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageIdentification)(nil)).Elem()
}

func (i StageIdentificationArray) ToStageIdentificationArrayOutput() StageIdentificationArrayOutput {
	return i.ToStageIdentificationArrayOutputWithContext(context.Background())
}

func (i StageIdentificationArray) ToStageIdentificationArrayOutputWithContext(ctx context.Context) StageIdentificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageIdentificationArrayOutput)
}

// StageIdentificationMapInput is an input type that accepts StageIdentificationMap and StageIdentificationMapOutput values.
// You can construct a concrete instance of `StageIdentificationMapInput` via:
//
//	StageIdentificationMap{ "key": StageIdentificationArgs{...} }
type StageIdentificationMapInput interface {
	pulumi.Input

	ToStageIdentificationMapOutput() StageIdentificationMapOutput
	ToStageIdentificationMapOutputWithContext(context.Context) StageIdentificationMapOutput
}

type StageIdentificationMap map[string]StageIdentificationInput

func (StageIdentificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageIdentification)(nil)).Elem()
}

func (i StageIdentificationMap) ToStageIdentificationMapOutput() StageIdentificationMapOutput {
	return i.ToStageIdentificationMapOutputWithContext(context.Background())
}

func (i StageIdentificationMap) ToStageIdentificationMapOutputWithContext(ctx context.Context) StageIdentificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageIdentificationMapOutput)
}

type StageIdentificationOutput struct{ *pulumi.OutputState }

func (StageIdentificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageIdentification)(nil)).Elem()
}

func (o StageIdentificationOutput) ToStageIdentificationOutput() StageIdentificationOutput {
	return o
}

func (o StageIdentificationOutput) ToStageIdentificationOutputWithContext(ctx context.Context) StageIdentificationOutput {
	return o
}

func (o StageIdentificationOutput) CaptchaStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringPtrOutput { return v.CaptchaStage }).(pulumi.StringPtrOutput)
}

func (o StageIdentificationOutput) CaseInsensitiveMatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.BoolPtrOutput { return v.CaseInsensitiveMatching }).(pulumi.BoolPtrOutput)
}

func (o StageIdentificationOutput) EnrollmentFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringPtrOutput { return v.EnrollmentFlow }).(pulumi.StringPtrOutput)
}

func (o StageIdentificationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StageIdentificationOutput) PasswordStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringPtrOutput { return v.PasswordStage }).(pulumi.StringPtrOutput)
}

func (o StageIdentificationOutput) PasswordlessFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringPtrOutput { return v.PasswordlessFlow }).(pulumi.StringPtrOutput)
}

// Defaults to `true`.
func (o StageIdentificationOutput) PretendUserExists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.BoolPtrOutput { return v.PretendUserExists }).(pulumi.BoolPtrOutput)
}

func (o StageIdentificationOutput) RecoveryFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringPtrOutput { return v.RecoveryFlow }).(pulumi.StringPtrOutput)
}

// Defaults to `true`.
func (o StageIdentificationOutput) ShowMatchedUser() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.BoolPtrOutput { return v.ShowMatchedUser }).(pulumi.BoolPtrOutput)
}

// Defaults to `false`.
func (o StageIdentificationOutput) ShowSourceLabels() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.BoolPtrOutput { return v.ShowSourceLabels }).(pulumi.BoolPtrOutput)
}

func (o StageIdentificationOutput) Sources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringArrayOutput { return v.Sources }).(pulumi.StringArrayOutput)
}

func (o StageIdentificationOutput) UserFields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StageIdentification) pulumi.StringArrayOutput { return v.UserFields }).(pulumi.StringArrayOutput)
}

type StageIdentificationArrayOutput struct{ *pulumi.OutputState }

func (StageIdentificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageIdentification)(nil)).Elem()
}

func (o StageIdentificationArrayOutput) ToStageIdentificationArrayOutput() StageIdentificationArrayOutput {
	return o
}

func (o StageIdentificationArrayOutput) ToStageIdentificationArrayOutputWithContext(ctx context.Context) StageIdentificationArrayOutput {
	return o
}

func (o StageIdentificationArrayOutput) Index(i pulumi.IntInput) StageIdentificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageIdentification {
		return vs[0].([]*StageIdentification)[vs[1].(int)]
	}).(StageIdentificationOutput)
}

type StageIdentificationMapOutput struct{ *pulumi.OutputState }

func (StageIdentificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageIdentification)(nil)).Elem()
}

func (o StageIdentificationMapOutput) ToStageIdentificationMapOutput() StageIdentificationMapOutput {
	return o
}

func (o StageIdentificationMapOutput) ToStageIdentificationMapOutputWithContext(ctx context.Context) StageIdentificationMapOutput {
	return o
}

func (o StageIdentificationMapOutput) MapIndex(k pulumi.StringInput) StageIdentificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageIdentification {
		return vs[0].(map[string]*StageIdentification)[vs[1].(string)]
	}).(StageIdentificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageIdentificationInput)(nil)).Elem(), &StageIdentification{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageIdentificationArrayInput)(nil)).Elem(), StageIdentificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageIdentificationMapInput)(nil)).Elem(), StageIdentificationMap{})
	pulumi.RegisterOutputType(StageIdentificationOutput{})
	pulumi.RegisterOutputType(StageIdentificationArrayOutput{})
	pulumi.RegisterOutputType(StageIdentificationMapOutput{})
}

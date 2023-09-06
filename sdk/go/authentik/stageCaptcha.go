// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StageCaptcha struct {
	pulumi.CustomResourceState

	ApiUrl     pulumi.StringPtrOutput `pulumi:"apiUrl"`
	JsUrl      pulumi.StringPtrOutput `pulumi:"jsUrl"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	PrivateKey pulumi.StringOutput    `pulumi:"privateKey"`
	PublicKey  pulumi.StringOutput    `pulumi:"publicKey"`
}

// NewStageCaptcha registers a new resource with the given unique name, arguments, and options.
func NewStageCaptcha(ctx *pulumi.Context,
	name string, args *StageCaptchaArgs, opts ...pulumi.ResourceOption) (*StageCaptcha, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StageCaptcha
	err := ctx.RegisterResource("authentik:index/stageCaptcha:StageCaptcha", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStageCaptcha gets an existing StageCaptcha resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStageCaptcha(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageCaptchaState, opts ...pulumi.ResourceOption) (*StageCaptcha, error) {
	var resource StageCaptcha
	err := ctx.ReadResource("authentik:index/stageCaptcha:StageCaptcha", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StageCaptcha resources.
type stageCaptchaState struct {
	ApiUrl     *string `pulumi:"apiUrl"`
	JsUrl      *string `pulumi:"jsUrl"`
	Name       *string `pulumi:"name"`
	PrivateKey *string `pulumi:"privateKey"`
	PublicKey  *string `pulumi:"publicKey"`
}

type StageCaptchaState struct {
	ApiUrl     pulumi.StringPtrInput
	JsUrl      pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	PrivateKey pulumi.StringPtrInput
	PublicKey  pulumi.StringPtrInput
}

func (StageCaptchaState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageCaptchaState)(nil)).Elem()
}

type stageCaptchaArgs struct {
	ApiUrl     *string `pulumi:"apiUrl"`
	JsUrl      *string `pulumi:"jsUrl"`
	Name       *string `pulumi:"name"`
	PrivateKey string  `pulumi:"privateKey"`
	PublicKey  string  `pulumi:"publicKey"`
}

// The set of arguments for constructing a StageCaptcha resource.
type StageCaptchaArgs struct {
	ApiUrl     pulumi.StringPtrInput
	JsUrl      pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	PrivateKey pulumi.StringInput
	PublicKey  pulumi.StringInput
}

func (StageCaptchaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageCaptchaArgs)(nil)).Elem()
}

type StageCaptchaInput interface {
	pulumi.Input

	ToStageCaptchaOutput() StageCaptchaOutput
	ToStageCaptchaOutputWithContext(ctx context.Context) StageCaptchaOutput
}

func (*StageCaptcha) ElementType() reflect.Type {
	return reflect.TypeOf((**StageCaptcha)(nil)).Elem()
}

func (i *StageCaptcha) ToStageCaptchaOutput() StageCaptchaOutput {
	return i.ToStageCaptchaOutputWithContext(context.Background())
}

func (i *StageCaptcha) ToStageCaptchaOutputWithContext(ctx context.Context) StageCaptchaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageCaptchaOutput)
}

// StageCaptchaArrayInput is an input type that accepts StageCaptchaArray and StageCaptchaArrayOutput values.
// You can construct a concrete instance of `StageCaptchaArrayInput` via:
//
//	StageCaptchaArray{ StageCaptchaArgs{...} }
type StageCaptchaArrayInput interface {
	pulumi.Input

	ToStageCaptchaArrayOutput() StageCaptchaArrayOutput
	ToStageCaptchaArrayOutputWithContext(context.Context) StageCaptchaArrayOutput
}

type StageCaptchaArray []StageCaptchaInput

func (StageCaptchaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageCaptcha)(nil)).Elem()
}

func (i StageCaptchaArray) ToStageCaptchaArrayOutput() StageCaptchaArrayOutput {
	return i.ToStageCaptchaArrayOutputWithContext(context.Background())
}

func (i StageCaptchaArray) ToStageCaptchaArrayOutputWithContext(ctx context.Context) StageCaptchaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageCaptchaArrayOutput)
}

// StageCaptchaMapInput is an input type that accepts StageCaptchaMap and StageCaptchaMapOutput values.
// You can construct a concrete instance of `StageCaptchaMapInput` via:
//
//	StageCaptchaMap{ "key": StageCaptchaArgs{...} }
type StageCaptchaMapInput interface {
	pulumi.Input

	ToStageCaptchaMapOutput() StageCaptchaMapOutput
	ToStageCaptchaMapOutputWithContext(context.Context) StageCaptchaMapOutput
}

type StageCaptchaMap map[string]StageCaptchaInput

func (StageCaptchaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageCaptcha)(nil)).Elem()
}

func (i StageCaptchaMap) ToStageCaptchaMapOutput() StageCaptchaMapOutput {
	return i.ToStageCaptchaMapOutputWithContext(context.Background())
}

func (i StageCaptchaMap) ToStageCaptchaMapOutputWithContext(ctx context.Context) StageCaptchaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageCaptchaMapOutput)
}

type StageCaptchaOutput struct{ *pulumi.OutputState }

func (StageCaptchaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StageCaptcha)(nil)).Elem()
}

func (o StageCaptchaOutput) ToStageCaptchaOutput() StageCaptchaOutput {
	return o
}

func (o StageCaptchaOutput) ToStageCaptchaOutputWithContext(ctx context.Context) StageCaptchaOutput {
	return o
}

func (o StageCaptchaOutput) ApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageCaptcha) pulumi.StringPtrOutput { return v.ApiUrl }).(pulumi.StringPtrOutput)
}

func (o StageCaptchaOutput) JsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StageCaptcha) pulumi.StringPtrOutput { return v.JsUrl }).(pulumi.StringPtrOutput)
}

func (o StageCaptchaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StageCaptcha) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StageCaptchaOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *StageCaptcha) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o StageCaptchaOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *StageCaptcha) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

type StageCaptchaArrayOutput struct{ *pulumi.OutputState }

func (StageCaptchaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StageCaptcha)(nil)).Elem()
}

func (o StageCaptchaArrayOutput) ToStageCaptchaArrayOutput() StageCaptchaArrayOutput {
	return o
}

func (o StageCaptchaArrayOutput) ToStageCaptchaArrayOutputWithContext(ctx context.Context) StageCaptchaArrayOutput {
	return o
}

func (o StageCaptchaArrayOutput) Index(i pulumi.IntInput) StageCaptchaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StageCaptcha {
		return vs[0].([]*StageCaptcha)[vs[1].(int)]
	}).(StageCaptchaOutput)
}

type StageCaptchaMapOutput struct{ *pulumi.OutputState }

func (StageCaptchaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StageCaptcha)(nil)).Elem()
}

func (o StageCaptchaMapOutput) ToStageCaptchaMapOutput() StageCaptchaMapOutput {
	return o
}

func (o StageCaptchaMapOutput) ToStageCaptchaMapOutputWithContext(ctx context.Context) StageCaptchaMapOutput {
	return o
}

func (o StageCaptchaMapOutput) MapIndex(k pulumi.StringInput) StageCaptchaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StageCaptcha {
		return vs[0].(map[string]*StageCaptcha)[vs[1].(string)]
	}).(StageCaptchaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageCaptchaInput)(nil)).Elem(), &StageCaptcha{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageCaptchaArrayInput)(nil)).Elem(), StageCaptchaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageCaptchaMapInput)(nil)).Elem(), StageCaptchaMap{})
	pulumi.RegisterOutputType(StageCaptchaOutput{})
	pulumi.RegisterOutputType(StageCaptchaArrayOutput{})
	pulumi.RegisterOutputType(StageCaptchaMapOutput{})
}

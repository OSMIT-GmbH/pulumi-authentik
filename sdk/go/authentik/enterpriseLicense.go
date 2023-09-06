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

type EnterpriseLicense struct {
	pulumi.CustomResourceState

	Expiry        pulumi.StringOutput `pulumi:"expiry"`
	ExternalUsers pulumi.IntOutput    `pulumi:"externalUsers"`
	InternalUsers pulumi.IntOutput    `pulumi:"internalUsers"`
	Key           pulumi.StringOutput `pulumi:"key"`
	Name          pulumi.StringOutput `pulumi:"name"`
}

// NewEnterpriseLicense registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseLicense(ctx *pulumi.Context,
	name string, args *EnterpriseLicenseArgs, opts ...pulumi.ResourceOption) (*EnterpriseLicense, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Key != nil {
		args.Key = pulumi.ToSecret(args.Key).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseLicense
	err := ctx.RegisterResource("authentik:index/enterpriseLicense:EnterpriseLicense", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseLicense gets an existing EnterpriseLicense resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseLicenseState, opts ...pulumi.ResourceOption) (*EnterpriseLicense, error) {
	var resource EnterpriseLicense
	err := ctx.ReadResource("authentik:index/enterpriseLicense:EnterpriseLicense", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseLicense resources.
type enterpriseLicenseState struct {
	Expiry        *string `pulumi:"expiry"`
	ExternalUsers *int    `pulumi:"externalUsers"`
	InternalUsers *int    `pulumi:"internalUsers"`
	Key           *string `pulumi:"key"`
	Name          *string `pulumi:"name"`
}

type EnterpriseLicenseState struct {
	Expiry        pulumi.StringPtrInput
	ExternalUsers pulumi.IntPtrInput
	InternalUsers pulumi.IntPtrInput
	Key           pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
}

func (EnterpriseLicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseLicenseState)(nil)).Elem()
}

type enterpriseLicenseArgs struct {
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a EnterpriseLicense resource.
type EnterpriseLicenseArgs struct {
	Key pulumi.StringInput
}

func (EnterpriseLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseLicenseArgs)(nil)).Elem()
}

type EnterpriseLicenseInput interface {
	pulumi.Input

	ToEnterpriseLicenseOutput() EnterpriseLicenseOutput
	ToEnterpriseLicenseOutputWithContext(ctx context.Context) EnterpriseLicenseOutput
}

func (*EnterpriseLicense) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseLicense)(nil)).Elem()
}

func (i *EnterpriseLicense) ToEnterpriseLicenseOutput() EnterpriseLicenseOutput {
	return i.ToEnterpriseLicenseOutputWithContext(context.Background())
}

func (i *EnterpriseLicense) ToEnterpriseLicenseOutputWithContext(ctx context.Context) EnterpriseLicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseLicenseOutput)
}

// EnterpriseLicenseArrayInput is an input type that accepts EnterpriseLicenseArray and EnterpriseLicenseArrayOutput values.
// You can construct a concrete instance of `EnterpriseLicenseArrayInput` via:
//
//	EnterpriseLicenseArray{ EnterpriseLicenseArgs{...} }
type EnterpriseLicenseArrayInput interface {
	pulumi.Input

	ToEnterpriseLicenseArrayOutput() EnterpriseLicenseArrayOutput
	ToEnterpriseLicenseArrayOutputWithContext(context.Context) EnterpriseLicenseArrayOutput
}

type EnterpriseLicenseArray []EnterpriseLicenseInput

func (EnterpriseLicenseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseLicense)(nil)).Elem()
}

func (i EnterpriseLicenseArray) ToEnterpriseLicenseArrayOutput() EnterpriseLicenseArrayOutput {
	return i.ToEnterpriseLicenseArrayOutputWithContext(context.Background())
}

func (i EnterpriseLicenseArray) ToEnterpriseLicenseArrayOutputWithContext(ctx context.Context) EnterpriseLicenseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseLicenseArrayOutput)
}

// EnterpriseLicenseMapInput is an input type that accepts EnterpriseLicenseMap and EnterpriseLicenseMapOutput values.
// You can construct a concrete instance of `EnterpriseLicenseMapInput` via:
//
//	EnterpriseLicenseMap{ "key": EnterpriseLicenseArgs{...} }
type EnterpriseLicenseMapInput interface {
	pulumi.Input

	ToEnterpriseLicenseMapOutput() EnterpriseLicenseMapOutput
	ToEnterpriseLicenseMapOutputWithContext(context.Context) EnterpriseLicenseMapOutput
}

type EnterpriseLicenseMap map[string]EnterpriseLicenseInput

func (EnterpriseLicenseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseLicense)(nil)).Elem()
}

func (i EnterpriseLicenseMap) ToEnterpriseLicenseMapOutput() EnterpriseLicenseMapOutput {
	return i.ToEnterpriseLicenseMapOutputWithContext(context.Background())
}

func (i EnterpriseLicenseMap) ToEnterpriseLicenseMapOutputWithContext(ctx context.Context) EnterpriseLicenseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseLicenseMapOutput)
}

type EnterpriseLicenseOutput struct{ *pulumi.OutputState }

func (EnterpriseLicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseLicense)(nil)).Elem()
}

func (o EnterpriseLicenseOutput) ToEnterpriseLicenseOutput() EnterpriseLicenseOutput {
	return o
}

func (o EnterpriseLicenseOutput) ToEnterpriseLicenseOutputWithContext(ctx context.Context) EnterpriseLicenseOutput {
	return o
}

func (o EnterpriseLicenseOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLicense) pulumi.StringOutput { return v.Expiry }).(pulumi.StringOutput)
}

func (o EnterpriseLicenseOutput) ExternalUsers() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseLicense) pulumi.IntOutput { return v.ExternalUsers }).(pulumi.IntOutput)
}

func (o EnterpriseLicenseOutput) InternalUsers() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseLicense) pulumi.IntOutput { return v.InternalUsers }).(pulumi.IntOutput)
}

func (o EnterpriseLicenseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLicense) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o EnterpriseLicenseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLicense) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type EnterpriseLicenseArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseLicenseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseLicense)(nil)).Elem()
}

func (o EnterpriseLicenseArrayOutput) ToEnterpriseLicenseArrayOutput() EnterpriseLicenseArrayOutput {
	return o
}

func (o EnterpriseLicenseArrayOutput) ToEnterpriseLicenseArrayOutputWithContext(ctx context.Context) EnterpriseLicenseArrayOutput {
	return o
}

func (o EnterpriseLicenseArrayOutput) Index(i pulumi.IntInput) EnterpriseLicenseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseLicense {
		return vs[0].([]*EnterpriseLicense)[vs[1].(int)]
	}).(EnterpriseLicenseOutput)
}

type EnterpriseLicenseMapOutput struct{ *pulumi.OutputState }

func (EnterpriseLicenseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseLicense)(nil)).Elem()
}

func (o EnterpriseLicenseMapOutput) ToEnterpriseLicenseMapOutput() EnterpriseLicenseMapOutput {
	return o
}

func (o EnterpriseLicenseMapOutput) ToEnterpriseLicenseMapOutputWithContext(ctx context.Context) EnterpriseLicenseMapOutput {
	return o
}

func (o EnterpriseLicenseMapOutput) MapIndex(k pulumi.StringInput) EnterpriseLicenseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseLicense {
		return vs[0].(map[string]*EnterpriseLicense)[vs[1].(string)]
	}).(EnterpriseLicenseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseLicenseInput)(nil)).Elem(), &EnterpriseLicense{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseLicenseArrayInput)(nil)).Elem(), EnterpriseLicenseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseLicenseMapInput)(nil)).Elem(), EnterpriseLicenseMap{})
	pulumi.RegisterOutputType(EnterpriseLicenseOutput{})
	pulumi.RegisterOutputType(EnterpriseLicenseArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseLicenseMapOutput{})
}

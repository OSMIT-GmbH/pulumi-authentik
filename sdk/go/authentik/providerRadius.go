// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//			default_authentication_flow, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-authentication-flow"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			nameProviderRadius, err := authentik.NewProviderRadius(ctx, "nameProviderRadius", &authentik.ProviderRadiusArgs{
//				AuthorizationFlow: *pulumi.String(default_authentication_flow.Id),
//				ClientNetworks:    pulumi.String("10.10.0.0/24"),
//				SharedSecret:      pulumi.String("my-shared-secret"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewApplication(ctx, "nameApplication", &authentik.ApplicationArgs{
//				Slug:             pulumi.String("radius-app"),
//				ProtocolProvider: nameProviderRadius.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ProviderRadius struct {
	pulumi.CustomResourceState

	AuthorizationFlow pulumi.StringOutput `pulumi:"authorizationFlow"`
	// Defaults to `0.0.0.0/0, ::/0`.
	ClientNetworks pulumi.StringPtrOutput `pulumi:"clientNetworks"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	SharedSecret   pulumi.StringOutput    `pulumi:"sharedSecret"`
}

// NewProviderRadius registers a new resource with the given unique name, arguments, and options.
func NewProviderRadius(ctx *pulumi.Context,
	name string, args *ProviderRadiusArgs, opts ...pulumi.ResourceOption) (*ProviderRadius, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationFlow == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationFlow'")
	}
	if args.SharedSecret == nil {
		return nil, errors.New("invalid value for required argument 'SharedSecret'")
	}
	if args.SharedSecret != nil {
		args.SharedSecret = pulumi.ToSecret(args.SharedSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sharedSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderRadius
	err := ctx.RegisterResource("authentik:index/providerRadius:ProviderRadius", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderRadius gets an existing ProviderRadius resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderRadius(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderRadiusState, opts ...pulumi.ResourceOption) (*ProviderRadius, error) {
	var resource ProviderRadius
	err := ctx.ReadResource("authentik:index/providerRadius:ProviderRadius", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderRadius resources.
type providerRadiusState struct {
	AuthorizationFlow *string `pulumi:"authorizationFlow"`
	// Defaults to `0.0.0.0/0, ::/0`.
	ClientNetworks *string `pulumi:"clientNetworks"`
	Name           *string `pulumi:"name"`
	SharedSecret   *string `pulumi:"sharedSecret"`
}

type ProviderRadiusState struct {
	AuthorizationFlow pulumi.StringPtrInput
	// Defaults to `0.0.0.0/0, ::/0`.
	ClientNetworks pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	SharedSecret   pulumi.StringPtrInput
}

func (ProviderRadiusState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRadiusState)(nil)).Elem()
}

type providerRadiusArgs struct {
	AuthorizationFlow string `pulumi:"authorizationFlow"`
	// Defaults to `0.0.0.0/0, ::/0`.
	ClientNetworks *string `pulumi:"clientNetworks"`
	Name           *string `pulumi:"name"`
	SharedSecret   string  `pulumi:"sharedSecret"`
}

// The set of arguments for constructing a ProviderRadius resource.
type ProviderRadiusArgs struct {
	AuthorizationFlow pulumi.StringInput
	// Defaults to `0.0.0.0/0, ::/0`.
	ClientNetworks pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	SharedSecret   pulumi.StringInput
}

func (ProviderRadiusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRadiusArgs)(nil)).Elem()
}

type ProviderRadiusInput interface {
	pulumi.Input

	ToProviderRadiusOutput() ProviderRadiusOutput
	ToProviderRadiusOutputWithContext(ctx context.Context) ProviderRadiusOutput
}

func (*ProviderRadius) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderRadius)(nil)).Elem()
}

func (i *ProviderRadius) ToProviderRadiusOutput() ProviderRadiusOutput {
	return i.ToProviderRadiusOutputWithContext(context.Background())
}

func (i *ProviderRadius) ToProviderRadiusOutputWithContext(ctx context.Context) ProviderRadiusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRadiusOutput)
}

func (i *ProviderRadius) ToOutput(ctx context.Context) pulumix.Output[*ProviderRadius] {
	return pulumix.Output[*ProviderRadius]{
		OutputState: i.ToProviderRadiusOutputWithContext(ctx).OutputState,
	}
}

// ProviderRadiusArrayInput is an input type that accepts ProviderRadiusArray and ProviderRadiusArrayOutput values.
// You can construct a concrete instance of `ProviderRadiusArrayInput` via:
//
//	ProviderRadiusArray{ ProviderRadiusArgs{...} }
type ProviderRadiusArrayInput interface {
	pulumi.Input

	ToProviderRadiusArrayOutput() ProviderRadiusArrayOutput
	ToProviderRadiusArrayOutputWithContext(context.Context) ProviderRadiusArrayOutput
}

type ProviderRadiusArray []ProviderRadiusInput

func (ProviderRadiusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderRadius)(nil)).Elem()
}

func (i ProviderRadiusArray) ToProviderRadiusArrayOutput() ProviderRadiusArrayOutput {
	return i.ToProviderRadiusArrayOutputWithContext(context.Background())
}

func (i ProviderRadiusArray) ToProviderRadiusArrayOutputWithContext(ctx context.Context) ProviderRadiusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRadiusArrayOutput)
}

func (i ProviderRadiusArray) ToOutput(ctx context.Context) pulumix.Output[[]*ProviderRadius] {
	return pulumix.Output[[]*ProviderRadius]{
		OutputState: i.ToProviderRadiusArrayOutputWithContext(ctx).OutputState,
	}
}

// ProviderRadiusMapInput is an input type that accepts ProviderRadiusMap and ProviderRadiusMapOutput values.
// You can construct a concrete instance of `ProviderRadiusMapInput` via:
//
//	ProviderRadiusMap{ "key": ProviderRadiusArgs{...} }
type ProviderRadiusMapInput interface {
	pulumi.Input

	ToProviderRadiusMapOutput() ProviderRadiusMapOutput
	ToProviderRadiusMapOutputWithContext(context.Context) ProviderRadiusMapOutput
}

type ProviderRadiusMap map[string]ProviderRadiusInput

func (ProviderRadiusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderRadius)(nil)).Elem()
}

func (i ProviderRadiusMap) ToProviderRadiusMapOutput() ProviderRadiusMapOutput {
	return i.ToProviderRadiusMapOutputWithContext(context.Background())
}

func (i ProviderRadiusMap) ToProviderRadiusMapOutputWithContext(ctx context.Context) ProviderRadiusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRadiusMapOutput)
}

func (i ProviderRadiusMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProviderRadius] {
	return pulumix.Output[map[string]*ProviderRadius]{
		OutputState: i.ToProviderRadiusMapOutputWithContext(ctx).OutputState,
	}
}

type ProviderRadiusOutput struct{ *pulumi.OutputState }

func (ProviderRadiusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderRadius)(nil)).Elem()
}

func (o ProviderRadiusOutput) ToProviderRadiusOutput() ProviderRadiusOutput {
	return o
}

func (o ProviderRadiusOutput) ToProviderRadiusOutputWithContext(ctx context.Context) ProviderRadiusOutput {
	return o
}

func (o ProviderRadiusOutput) ToOutput(ctx context.Context) pulumix.Output[*ProviderRadius] {
	return pulumix.Output[*ProviderRadius]{
		OutputState: o.OutputState,
	}
}

func (o ProviderRadiusOutput) AuthorizationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRadius) pulumi.StringOutput { return v.AuthorizationFlow }).(pulumi.StringOutput)
}

// Defaults to `0.0.0.0/0, ::/0`.
func (o ProviderRadiusOutput) ClientNetworks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderRadius) pulumi.StringPtrOutput { return v.ClientNetworks }).(pulumi.StringPtrOutput)
}

func (o ProviderRadiusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRadius) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderRadiusOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRadius) pulumi.StringOutput { return v.SharedSecret }).(pulumi.StringOutput)
}

type ProviderRadiusArrayOutput struct{ *pulumi.OutputState }

func (ProviderRadiusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderRadius)(nil)).Elem()
}

func (o ProviderRadiusArrayOutput) ToProviderRadiusArrayOutput() ProviderRadiusArrayOutput {
	return o
}

func (o ProviderRadiusArrayOutput) ToProviderRadiusArrayOutputWithContext(ctx context.Context) ProviderRadiusArrayOutput {
	return o
}

func (o ProviderRadiusArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ProviderRadius] {
	return pulumix.Output[[]*ProviderRadius]{
		OutputState: o.OutputState,
	}
}

func (o ProviderRadiusArrayOutput) Index(i pulumi.IntInput) ProviderRadiusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderRadius {
		return vs[0].([]*ProviderRadius)[vs[1].(int)]
	}).(ProviderRadiusOutput)
}

type ProviderRadiusMapOutput struct{ *pulumi.OutputState }

func (ProviderRadiusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderRadius)(nil)).Elem()
}

func (o ProviderRadiusMapOutput) ToProviderRadiusMapOutput() ProviderRadiusMapOutput {
	return o
}

func (o ProviderRadiusMapOutput) ToProviderRadiusMapOutputWithContext(ctx context.Context) ProviderRadiusMapOutput {
	return o
}

func (o ProviderRadiusMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProviderRadius] {
	return pulumix.Output[map[string]*ProviderRadius]{
		OutputState: o.OutputState,
	}
}

func (o ProviderRadiusMapOutput) MapIndex(k pulumi.StringInput) ProviderRadiusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderRadius {
		return vs[0].(map[string]*ProviderRadius)[vs[1].(string)]
	}).(ProviderRadiusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRadiusInput)(nil)).Elem(), &ProviderRadius{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRadiusArrayInput)(nil)).Elem(), ProviderRadiusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRadiusMapInput)(nil)).Elem(), ProviderRadiusMap{})
	pulumi.RegisterOutputType(ProviderRadiusOutput{})
	pulumi.RegisterOutputType(ProviderRadiusArrayOutput{})
	pulumi.RegisterOutputType(ProviderRadiusMapOutput{})
}

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
//			default_authorization_flow, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-provider-authorization-implicit-consent"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			nameProviderOauth2, err := authentik.NewProviderOauth2(ctx, "nameProviderOauth2", &authentik.ProviderOauth2Args{
//				ClientId:          pulumi.String("example-app"),
//				ClientSecret:      pulumi.String("test"),
//				AuthorizationFlow: pulumi.String(default_authorization_flow.Id),
//			})
//			if err != nil {
//				return err
//			}
//			policy, err := authentik.NewPolicyExpression(ctx, "policy", &authentik.PolicyExpressionArgs{
//				Expression: pulumi.String("return True"),
//			})
//			if err != nil {
//				return err
//			}
//			nameApplication, err := authentik.NewApplication(ctx, "nameApplication", &authentik.ApplicationArgs{
//				Slug:             pulumi.String("example-app"),
//				ProtocolProvider: nameProviderOauth2.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewPolicyBinding(ctx, "app-access", &authentik.PolicyBindingArgs{
//				Target: nameApplication.Uuid,
//				Policy: policy.ID(),
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
type Application struct {
	pulumi.CustomResourceState

	BackchannelProviders pulumi.IntArrayOutput  `pulumi:"backchannelProviders"`
	Group                pulumi.StringPtrOutput `pulumi:"group"`
	MetaDescription      pulumi.StringPtrOutput `pulumi:"metaDescription"`
	MetaIcon             pulumi.StringPtrOutput `pulumi:"metaIcon"`
	MetaLaunchUrl        pulumi.StringPtrOutput `pulumi:"metaLaunchUrl"`
	MetaPublisher        pulumi.StringPtrOutput `pulumi:"metaPublisher"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	OpenInNewTab         pulumi.BoolPtrOutput   `pulumi:"openInNewTab"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrOutput `pulumi:"policyEngineMode"`
	ProtocolProvider pulumi.IntPtrOutput    `pulumi:"protocolProvider"`
	Slug             pulumi.StringOutput    `pulumi:"slug"`
	Uuid             pulumi.StringOutput    `pulumi:"uuid"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("authentik:index/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("authentik:index/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	BackchannelProviders []int   `pulumi:"backchannelProviders"`
	Group                *string `pulumi:"group"`
	MetaDescription      *string `pulumi:"metaDescription"`
	MetaIcon             *string `pulumi:"metaIcon"`
	MetaLaunchUrl        *string `pulumi:"metaLaunchUrl"`
	MetaPublisher        *string `pulumi:"metaPublisher"`
	Name                 *string `pulumi:"name"`
	OpenInNewTab         *bool   `pulumi:"openInNewTab"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	ProtocolProvider *int    `pulumi:"protocolProvider"`
	Slug             *string `pulumi:"slug"`
	Uuid             *string `pulumi:"uuid"`
}

type ApplicationState struct {
	BackchannelProviders pulumi.IntArrayInput
	Group                pulumi.StringPtrInput
	MetaDescription      pulumi.StringPtrInput
	MetaIcon             pulumi.StringPtrInput
	MetaLaunchUrl        pulumi.StringPtrInput
	MetaPublisher        pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	OpenInNewTab         pulumi.BoolPtrInput
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrInput
	ProtocolProvider pulumi.IntPtrInput
	Slug             pulumi.StringPtrInput
	Uuid             pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	BackchannelProviders []int   `pulumi:"backchannelProviders"`
	Group                *string `pulumi:"group"`
	MetaDescription      *string `pulumi:"metaDescription"`
	MetaIcon             *string `pulumi:"metaIcon"`
	MetaLaunchUrl        *string `pulumi:"metaLaunchUrl"`
	MetaPublisher        *string `pulumi:"metaPublisher"`
	Name                 *string `pulumi:"name"`
	OpenInNewTab         *bool   `pulumi:"openInNewTab"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	ProtocolProvider *int    `pulumi:"protocolProvider"`
	Slug             string  `pulumi:"slug"`
	Uuid             *string `pulumi:"uuid"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	BackchannelProviders pulumi.IntArrayInput
	Group                pulumi.StringPtrInput
	MetaDescription      pulumi.StringPtrInput
	MetaIcon             pulumi.StringPtrInput
	MetaLaunchUrl        pulumi.StringPtrInput
	MetaPublisher        pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	OpenInNewTab         pulumi.BoolPtrInput
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrInput
	ProtocolProvider pulumi.IntPtrInput
	Slug             pulumi.StringInput
	Uuid             pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) BackchannelProviders() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.IntArrayOutput { return v.BackchannelProviders }).(pulumi.IntArrayOutput)
}

func (o ApplicationOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) MetaDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.MetaDescription }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) MetaIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.MetaIcon }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) MetaLaunchUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.MetaLaunchUrl }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) MetaPublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.MetaPublisher }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) OpenInNewTab() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolPtrOutput { return v.OpenInNewTab }).(pulumi.BoolPtrOutput)
}

// Allowed values: - `all` - `any`
func (o ApplicationOutput) PolicyEngineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.PolicyEngineMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) ProtocolProvider() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.IntPtrOutput { return v.ProtocolProvider }).(pulumi.IntPtrOutput)
}

func (o ApplicationOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}

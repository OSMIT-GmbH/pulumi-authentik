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
//			// Create a proxy provider
//			default_authorization_flow, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-provider-authorization-implicit-consent"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			name, err := authentik.NewProviderProxy(ctx, "name", &authentik.ProviderProxyArgs{
//				Name:              pulumi.String("test-app"),
//				InternalHost:      pulumi.String("http://foo.bar.baz"),
//				ExternalHost:      pulumi.String("http://internal.service"),
//				AuthorizationFlow: pulumi.String(default_authorization_flow.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewApplication(ctx, "name", &authentik.ApplicationArgs{
//				Name:             pulumi.String("test-app"),
//				Slug:             pulumi.String("test-app"),
//				ProtocolProvider: name.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ProviderProxy struct {
	pulumi.CustomResourceState

	AccessTokenValidity        pulumi.StringPtrOutput `pulumi:"accessTokenValidity"`
	AuthenticationFlow         pulumi.StringPtrOutput `pulumi:"authenticationFlow"`
	AuthorizationFlow          pulumi.StringOutput    `pulumi:"authorizationFlow"`
	BasicAuthEnabled           pulumi.BoolPtrOutput   `pulumi:"basicAuthEnabled"`
	BasicAuthPasswordAttribute pulumi.StringPtrOutput `pulumi:"basicAuthPasswordAttribute"`
	BasicAuthUsernameAttribute pulumi.StringPtrOutput `pulumi:"basicAuthUsernameAttribute"`
	ClientId                   pulumi.StringOutput    `pulumi:"clientId"`
	CookieDomain               pulumi.StringPtrOutput `pulumi:"cookieDomain"`
	ExternalHost               pulumi.StringOutput    `pulumi:"externalHost"`
	InterceptHeaderAuth        pulumi.BoolPtrOutput   `pulumi:"interceptHeaderAuth"`
	InternalHost               pulumi.StringPtrOutput `pulumi:"internalHost"`
	InternalHostSslValidation  pulumi.BoolPtrOutput   `pulumi:"internalHostSslValidation"`
	InvalidationFlow           pulumi.StringOutput    `pulumi:"invalidationFlow"`
	// JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
	JwksSources pulumi.StringArrayOutput `pulumi:"jwksSources"`
	// Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
	Mode                 pulumi.StringPtrOutput   `pulumi:"mode"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	PropertyMappings     pulumi.StringArrayOutput `pulumi:"propertyMappings"`
	RefreshTokenValidity pulumi.StringPtrOutput   `pulumi:"refreshTokenValidity"`
	SkipPathRegex        pulumi.StringPtrOutput   `pulumi:"skipPathRegex"`
}

// NewProviderProxy registers a new resource with the given unique name, arguments, and options.
func NewProviderProxy(ctx *pulumi.Context,
	name string, args *ProviderProxyArgs, opts ...pulumi.ResourceOption) (*ProviderProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationFlow == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationFlow'")
	}
	if args.ExternalHost == nil {
		return nil, errors.New("invalid value for required argument 'ExternalHost'")
	}
	if args.InvalidationFlow == nil {
		return nil, errors.New("invalid value for required argument 'InvalidationFlow'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderProxy
	err := ctx.RegisterResource("authentik:index/providerProxy:ProviderProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderProxy gets an existing ProviderProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderProxyState, opts ...pulumi.ResourceOption) (*ProviderProxy, error) {
	var resource ProviderProxy
	err := ctx.ReadResource("authentik:index/providerProxy:ProviderProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderProxy resources.
type providerProxyState struct {
	AccessTokenValidity        *string `pulumi:"accessTokenValidity"`
	AuthenticationFlow         *string `pulumi:"authenticationFlow"`
	AuthorizationFlow          *string `pulumi:"authorizationFlow"`
	BasicAuthEnabled           *bool   `pulumi:"basicAuthEnabled"`
	BasicAuthPasswordAttribute *string `pulumi:"basicAuthPasswordAttribute"`
	BasicAuthUsernameAttribute *string `pulumi:"basicAuthUsernameAttribute"`
	ClientId                   *string `pulumi:"clientId"`
	CookieDomain               *string `pulumi:"cookieDomain"`
	ExternalHost               *string `pulumi:"externalHost"`
	InterceptHeaderAuth        *bool   `pulumi:"interceptHeaderAuth"`
	InternalHost               *string `pulumi:"internalHost"`
	InternalHostSslValidation  *bool   `pulumi:"internalHostSslValidation"`
	InvalidationFlow           *string `pulumi:"invalidationFlow"`
	// JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
	JwksSources []string `pulumi:"jwksSources"`
	// Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
	Mode                 *string  `pulumi:"mode"`
	Name                 *string  `pulumi:"name"`
	PropertyMappings     []string `pulumi:"propertyMappings"`
	RefreshTokenValidity *string  `pulumi:"refreshTokenValidity"`
	SkipPathRegex        *string  `pulumi:"skipPathRegex"`
}

type ProviderProxyState struct {
	AccessTokenValidity        pulumi.StringPtrInput
	AuthenticationFlow         pulumi.StringPtrInput
	AuthorizationFlow          pulumi.StringPtrInput
	BasicAuthEnabled           pulumi.BoolPtrInput
	BasicAuthPasswordAttribute pulumi.StringPtrInput
	BasicAuthUsernameAttribute pulumi.StringPtrInput
	ClientId                   pulumi.StringPtrInput
	CookieDomain               pulumi.StringPtrInput
	ExternalHost               pulumi.StringPtrInput
	InterceptHeaderAuth        pulumi.BoolPtrInput
	InternalHost               pulumi.StringPtrInput
	InternalHostSslValidation  pulumi.BoolPtrInput
	InvalidationFlow           pulumi.StringPtrInput
	// JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
	JwksSources pulumi.StringArrayInput
	// Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
	Mode                 pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	PropertyMappings     pulumi.StringArrayInput
	RefreshTokenValidity pulumi.StringPtrInput
	SkipPathRegex        pulumi.StringPtrInput
}

func (ProviderProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerProxyState)(nil)).Elem()
}

type providerProxyArgs struct {
	AccessTokenValidity        *string `pulumi:"accessTokenValidity"`
	AuthenticationFlow         *string `pulumi:"authenticationFlow"`
	AuthorizationFlow          string  `pulumi:"authorizationFlow"`
	BasicAuthEnabled           *bool   `pulumi:"basicAuthEnabled"`
	BasicAuthPasswordAttribute *string `pulumi:"basicAuthPasswordAttribute"`
	BasicAuthUsernameAttribute *string `pulumi:"basicAuthUsernameAttribute"`
	CookieDomain               *string `pulumi:"cookieDomain"`
	ExternalHost               string  `pulumi:"externalHost"`
	InterceptHeaderAuth        *bool   `pulumi:"interceptHeaderAuth"`
	InternalHost               *string `pulumi:"internalHost"`
	InternalHostSslValidation  *bool   `pulumi:"internalHostSslValidation"`
	InvalidationFlow           string  `pulumi:"invalidationFlow"`
	// JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
	JwksSources []string `pulumi:"jwksSources"`
	// Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
	Mode                 *string  `pulumi:"mode"`
	Name                 *string  `pulumi:"name"`
	PropertyMappings     []string `pulumi:"propertyMappings"`
	RefreshTokenValidity *string  `pulumi:"refreshTokenValidity"`
	SkipPathRegex        *string  `pulumi:"skipPathRegex"`
}

// The set of arguments for constructing a ProviderProxy resource.
type ProviderProxyArgs struct {
	AccessTokenValidity        pulumi.StringPtrInput
	AuthenticationFlow         pulumi.StringPtrInput
	AuthorizationFlow          pulumi.StringInput
	BasicAuthEnabled           pulumi.BoolPtrInput
	BasicAuthPasswordAttribute pulumi.StringPtrInput
	BasicAuthUsernameAttribute pulumi.StringPtrInput
	CookieDomain               pulumi.StringPtrInput
	ExternalHost               pulumi.StringInput
	InterceptHeaderAuth        pulumi.BoolPtrInput
	InternalHost               pulumi.StringPtrInput
	InternalHostSslValidation  pulumi.BoolPtrInput
	InvalidationFlow           pulumi.StringInput
	// JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
	JwksSources pulumi.StringArrayInput
	// Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
	Mode                 pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	PropertyMappings     pulumi.StringArrayInput
	RefreshTokenValidity pulumi.StringPtrInput
	SkipPathRegex        pulumi.StringPtrInput
}

func (ProviderProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerProxyArgs)(nil)).Elem()
}

type ProviderProxyInput interface {
	pulumi.Input

	ToProviderProxyOutput() ProviderProxyOutput
	ToProviderProxyOutputWithContext(ctx context.Context) ProviderProxyOutput
}

func (*ProviderProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderProxy)(nil)).Elem()
}

func (i *ProviderProxy) ToProviderProxyOutput() ProviderProxyOutput {
	return i.ToProviderProxyOutputWithContext(context.Background())
}

func (i *ProviderProxy) ToProviderProxyOutputWithContext(ctx context.Context) ProviderProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderProxyOutput)
}

// ProviderProxyArrayInput is an input type that accepts ProviderProxyArray and ProviderProxyArrayOutput values.
// You can construct a concrete instance of `ProviderProxyArrayInput` via:
//
//	ProviderProxyArray{ ProviderProxyArgs{...} }
type ProviderProxyArrayInput interface {
	pulumi.Input

	ToProviderProxyArrayOutput() ProviderProxyArrayOutput
	ToProviderProxyArrayOutputWithContext(context.Context) ProviderProxyArrayOutput
}

type ProviderProxyArray []ProviderProxyInput

func (ProviderProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderProxy)(nil)).Elem()
}

func (i ProviderProxyArray) ToProviderProxyArrayOutput() ProviderProxyArrayOutput {
	return i.ToProviderProxyArrayOutputWithContext(context.Background())
}

func (i ProviderProxyArray) ToProviderProxyArrayOutputWithContext(ctx context.Context) ProviderProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderProxyArrayOutput)
}

// ProviderProxyMapInput is an input type that accepts ProviderProxyMap and ProviderProxyMapOutput values.
// You can construct a concrete instance of `ProviderProxyMapInput` via:
//
//	ProviderProxyMap{ "key": ProviderProxyArgs{...} }
type ProviderProxyMapInput interface {
	pulumi.Input

	ToProviderProxyMapOutput() ProviderProxyMapOutput
	ToProviderProxyMapOutputWithContext(context.Context) ProviderProxyMapOutput
}

type ProviderProxyMap map[string]ProviderProxyInput

func (ProviderProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderProxy)(nil)).Elem()
}

func (i ProviderProxyMap) ToProviderProxyMapOutput() ProviderProxyMapOutput {
	return i.ToProviderProxyMapOutputWithContext(context.Background())
}

func (i ProviderProxyMap) ToProviderProxyMapOutputWithContext(ctx context.Context) ProviderProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderProxyMapOutput)
}

type ProviderProxyOutput struct{ *pulumi.OutputState }

func (ProviderProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderProxy)(nil)).Elem()
}

func (o ProviderProxyOutput) ToProviderProxyOutput() ProviderProxyOutput {
	return o
}

func (o ProviderProxyOutput) ToProviderProxyOutputWithContext(ctx context.Context) ProviderProxyOutput {
	return o
}

func (o ProviderProxyOutput) AccessTokenValidity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.AccessTokenValidity }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) AuthenticationFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.AuthenticationFlow }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) AuthorizationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringOutput { return v.AuthorizationFlow }).(pulumi.StringOutput)
}

func (o ProviderProxyOutput) BasicAuthEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.BoolPtrOutput { return v.BasicAuthEnabled }).(pulumi.BoolPtrOutput)
}

func (o ProviderProxyOutput) BasicAuthPasswordAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.BasicAuthPasswordAttribute }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) BasicAuthUsernameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.BasicAuthUsernameAttribute }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o ProviderProxyOutput) CookieDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.CookieDomain }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) ExternalHost() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringOutput { return v.ExternalHost }).(pulumi.StringOutput)
}

func (o ProviderProxyOutput) InterceptHeaderAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.BoolPtrOutput { return v.InterceptHeaderAuth }).(pulumi.BoolPtrOutput)
}

func (o ProviderProxyOutput) InternalHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.InternalHost }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) InternalHostSslValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.BoolPtrOutput { return v.InternalHostSslValidation }).(pulumi.BoolPtrOutput)
}

func (o ProviderProxyOutput) InvalidationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringOutput { return v.InvalidationFlow }).(pulumi.StringOutput)
}

// JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
func (o ProviderProxyOutput) JwksSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringArrayOutput { return v.JwksSources }).(pulumi.StringArrayOutput)
}

// Allowed values: - `proxy` - `forwardSingle` - `forwardDomain`
func (o ProviderProxyOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderProxyOutput) PropertyMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringArrayOutput { return v.PropertyMappings }).(pulumi.StringArrayOutput)
}

func (o ProviderProxyOutput) RefreshTokenValidity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.RefreshTokenValidity }).(pulumi.StringPtrOutput)
}

func (o ProviderProxyOutput) SkipPathRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) pulumi.StringPtrOutput { return v.SkipPathRegex }).(pulumi.StringPtrOutput)
}

type ProviderProxyArrayOutput struct{ *pulumi.OutputState }

func (ProviderProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderProxy)(nil)).Elem()
}

func (o ProviderProxyArrayOutput) ToProviderProxyArrayOutput() ProviderProxyArrayOutput {
	return o
}

func (o ProviderProxyArrayOutput) ToProviderProxyArrayOutputWithContext(ctx context.Context) ProviderProxyArrayOutput {
	return o
}

func (o ProviderProxyArrayOutput) Index(i pulumi.IntInput) ProviderProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderProxy {
		return vs[0].([]*ProviderProxy)[vs[1].(int)]
	}).(ProviderProxyOutput)
}

type ProviderProxyMapOutput struct{ *pulumi.OutputState }

func (ProviderProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderProxy)(nil)).Elem()
}

func (o ProviderProxyMapOutput) ToProviderProxyMapOutput() ProviderProxyMapOutput {
	return o
}

func (o ProviderProxyMapOutput) ToProviderProxyMapOutputWithContext(ctx context.Context) ProviderProxyMapOutput {
	return o
}

func (o ProviderProxyMapOutput) MapIndex(k pulumi.StringInput) ProviderProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderProxy {
		return vs[0].(map[string]*ProviderProxy)[vs[1].(string)]
	}).(ProviderProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderProxyInput)(nil)).Elem(), &ProviderProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderProxyArrayInput)(nil)).Elem(), ProviderProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderProxyMapInput)(nil)).Elem(), ProviderProxyMap{})
	pulumi.RegisterOutputType(ProviderProxyOutput{})
	pulumi.RegisterOutputType(ProviderProxyArrayOutput{})
	pulumi.RegisterOutputType(ProviderProxyMapOutput{})
}

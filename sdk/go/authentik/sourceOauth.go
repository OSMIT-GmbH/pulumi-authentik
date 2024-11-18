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
//	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create OAuth Source using an existing provider
//			default_source_authentication, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-source-authentication"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			default_source_enrollment, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-source-enrollment"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewSourceOauth(ctx, "name", &authentik.SourceOauthArgs{
//				Name:               pulumi.String("discord"),
//				Slug:               pulumi.String("discord"),
//				AuthenticationFlow: pulumi.String(default_source_authentication.Id),
//				EnrollmentFlow:     pulumi.String(default_source_enrollment.Id),
//				ProviderType:       pulumi.String("discord"),
//				ConsumerKey:        pulumi.String("foo"),
//				ConsumerSecret:     pulumi.String("bar"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SourceOauth struct {
	pulumi.CustomResourceState

	// Only required for OAuth1.
	AccessTokenUrl     pulumi.StringPtrOutput `pulumi:"accessTokenUrl"`
	AdditionalScopes   pulumi.StringPtrOutput `pulumi:"additionalScopes"`
	AuthenticationFlow pulumi.StringPtrOutput `pulumi:"authenticationFlow"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	AuthorizationUrl pulumi.StringPtrOutput `pulumi:"authorizationUrl"`
	CallbackUri      pulumi.StringOutput    `pulumi:"callbackUri"`
	ConsumerKey      pulumi.StringOutput    `pulumi:"consumerKey"`
	ConsumerSecret   pulumi.StringOutput    `pulumi:"consumerSecret"`
	Enabled          pulumi.BoolPtrOutput   `pulumi:"enabled"`
	EnrollmentFlow   pulumi.StringPtrOutput `pulumi:"enrollmentFlow"`
	// Allowed values: - `identifier` - `nameLink` - `nameDeny`
	GroupMatchingMode pulumi.StringPtrOutput `pulumi:"groupMatchingMode"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	// Manually configure JWKS keys for use with machine-to-machine authentication. JSON format expected. Use jsonencode() to
	// pass objects.
	OidcJwks pulumi.StringOutput `pulumi:"oidcJwks"`
	// Automatically configure JWKS if not specified by `oidcWellKnownUrl`.
	OidcJwksUrl pulumi.StringPtrOutput `pulumi:"oidcJwksUrl"`
	// Automatically configure source from OIDC well-known endpoint. URL is taken as is, and should end with
	// `.well-known/openid-configuration`.
	OidcWellKnownUrl pulumi.StringPtrOutput `pulumi:"oidcWellKnownUrl"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrOutput `pulumi:"policyEngineMode"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	ProfileUrl pulumi.StringPtrOutput `pulumi:"profileUrl"`
	// Allowed values: - `apple` - `openidconnect` - `azuread` - `discord` - `facebook` - `github` - `gitlab` - `google` -
	// `mailcow` - `okta` - `patreon` - `reddit` - `twitch` - `twitter`
	ProviderType pulumi.StringOutput `pulumi:"providerType"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	RequestTokenUrl pulumi.StringPtrOutput `pulumi:"requestTokenUrl"`
	Slug            pulumi.StringOutput    `pulumi:"slug"`
	// Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
	UserMatchingMode pulumi.StringPtrOutput `pulumi:"userMatchingMode"`
	UserPathTemplate pulumi.StringPtrOutput `pulumi:"userPathTemplate"`
	Uuid             pulumi.StringOutput    `pulumi:"uuid"`
}

// NewSourceOauth registers a new resource with the given unique name, arguments, and options.
func NewSourceOauth(ctx *pulumi.Context,
	name string, args *SourceOauthArgs, opts ...pulumi.ResourceOption) (*SourceOauth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerKey == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerKey'")
	}
	if args.ConsumerSecret == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerSecret'")
	}
	if args.ProviderType == nil {
		return nil, errors.New("invalid value for required argument 'ProviderType'")
	}
	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	if args.ConsumerSecret != nil {
		args.ConsumerSecret = pulumi.ToSecret(args.ConsumerSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"consumerSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SourceOauth
	err := ctx.RegisterResource("authentik:index/sourceOauth:SourceOauth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceOauth gets an existing SourceOauth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceOauth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceOauthState, opts ...pulumi.ResourceOption) (*SourceOauth, error) {
	var resource SourceOauth
	err := ctx.ReadResource("authentik:index/sourceOauth:SourceOauth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceOauth resources.
type sourceOauthState struct {
	// Only required for OAuth1.
	AccessTokenUrl     *string `pulumi:"accessTokenUrl"`
	AdditionalScopes   *string `pulumi:"additionalScopes"`
	AuthenticationFlow *string `pulumi:"authenticationFlow"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	AuthorizationUrl *string `pulumi:"authorizationUrl"`
	CallbackUri      *string `pulumi:"callbackUri"`
	ConsumerKey      *string `pulumi:"consumerKey"`
	ConsumerSecret   *string `pulumi:"consumerSecret"`
	Enabled          *bool   `pulumi:"enabled"`
	EnrollmentFlow   *string `pulumi:"enrollmentFlow"`
	// Allowed values: - `identifier` - `nameLink` - `nameDeny`
	GroupMatchingMode *string `pulumi:"groupMatchingMode"`
	Name              *string `pulumi:"name"`
	// Manually configure JWKS keys for use with machine-to-machine authentication. JSON format expected. Use jsonencode() to
	// pass objects.
	OidcJwks *string `pulumi:"oidcJwks"`
	// Automatically configure JWKS if not specified by `oidcWellKnownUrl`.
	OidcJwksUrl *string `pulumi:"oidcJwksUrl"`
	// Automatically configure source from OIDC well-known endpoint. URL is taken as is, and should end with
	// `.well-known/openid-configuration`.
	OidcWellKnownUrl *string `pulumi:"oidcWellKnownUrl"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	ProfileUrl *string `pulumi:"profileUrl"`
	// Allowed values: - `apple` - `openidconnect` - `azuread` - `discord` - `facebook` - `github` - `gitlab` - `google` -
	// `mailcow` - `okta` - `patreon` - `reddit` - `twitch` - `twitter`
	ProviderType *string `pulumi:"providerType"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	RequestTokenUrl *string `pulumi:"requestTokenUrl"`
	Slug            *string `pulumi:"slug"`
	// Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
	UserMatchingMode *string `pulumi:"userMatchingMode"`
	UserPathTemplate *string `pulumi:"userPathTemplate"`
	Uuid             *string `pulumi:"uuid"`
}

type SourceOauthState struct {
	// Only required for OAuth1.
	AccessTokenUrl     pulumi.StringPtrInput
	AdditionalScopes   pulumi.StringPtrInput
	AuthenticationFlow pulumi.StringPtrInput
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	AuthorizationUrl pulumi.StringPtrInput
	CallbackUri      pulumi.StringPtrInput
	ConsumerKey      pulumi.StringPtrInput
	ConsumerSecret   pulumi.StringPtrInput
	Enabled          pulumi.BoolPtrInput
	EnrollmentFlow   pulumi.StringPtrInput
	// Allowed values: - `identifier` - `nameLink` - `nameDeny`
	GroupMatchingMode pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	// Manually configure JWKS keys for use with machine-to-machine authentication. JSON format expected. Use jsonencode() to
	// pass objects.
	OidcJwks pulumi.StringPtrInput
	// Automatically configure JWKS if not specified by `oidcWellKnownUrl`.
	OidcJwksUrl pulumi.StringPtrInput
	// Automatically configure source from OIDC well-known endpoint. URL is taken as is, and should end with
	// `.well-known/openid-configuration`.
	OidcWellKnownUrl pulumi.StringPtrInput
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrInput
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	ProfileUrl pulumi.StringPtrInput
	// Allowed values: - `apple` - `openidconnect` - `azuread` - `discord` - `facebook` - `github` - `gitlab` - `google` -
	// `mailcow` - `okta` - `patreon` - `reddit` - `twitch` - `twitter`
	ProviderType pulumi.StringPtrInput
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	RequestTokenUrl pulumi.StringPtrInput
	Slug            pulumi.StringPtrInput
	// Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
	UserMatchingMode pulumi.StringPtrInput
	UserPathTemplate pulumi.StringPtrInput
	Uuid             pulumi.StringPtrInput
}

func (SourceOauthState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOauthState)(nil)).Elem()
}

type sourceOauthArgs struct {
	// Only required for OAuth1.
	AccessTokenUrl     *string `pulumi:"accessTokenUrl"`
	AdditionalScopes   *string `pulumi:"additionalScopes"`
	AuthenticationFlow *string `pulumi:"authenticationFlow"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	AuthorizationUrl *string `pulumi:"authorizationUrl"`
	ConsumerKey      string  `pulumi:"consumerKey"`
	ConsumerSecret   string  `pulumi:"consumerSecret"`
	Enabled          *bool   `pulumi:"enabled"`
	EnrollmentFlow   *string `pulumi:"enrollmentFlow"`
	// Allowed values: - `identifier` - `nameLink` - `nameDeny`
	GroupMatchingMode *string `pulumi:"groupMatchingMode"`
	Name              *string `pulumi:"name"`
	// Manually configure JWKS keys for use with machine-to-machine authentication. JSON format expected. Use jsonencode() to
	// pass objects.
	OidcJwks *string `pulumi:"oidcJwks"`
	// Automatically configure JWKS if not specified by `oidcWellKnownUrl`.
	OidcJwksUrl *string `pulumi:"oidcJwksUrl"`
	// Automatically configure source from OIDC well-known endpoint. URL is taken as is, and should end with
	// `.well-known/openid-configuration`.
	OidcWellKnownUrl *string `pulumi:"oidcWellKnownUrl"`
	// Allowed values: - `all` - `any`
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	ProfileUrl *string `pulumi:"profileUrl"`
	// Allowed values: - `apple` - `openidconnect` - `azuread` - `discord` - `facebook` - `github` - `gitlab` - `google` -
	// `mailcow` - `okta` - `patreon` - `reddit` - `twitch` - `twitter`
	ProviderType string `pulumi:"providerType"`
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	RequestTokenUrl *string `pulumi:"requestTokenUrl"`
	Slug            string  `pulumi:"slug"`
	// Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
	UserMatchingMode *string `pulumi:"userMatchingMode"`
	UserPathTemplate *string `pulumi:"userPathTemplate"`
	Uuid             *string `pulumi:"uuid"`
}

// The set of arguments for constructing a SourceOauth resource.
type SourceOauthArgs struct {
	// Only required for OAuth1.
	AccessTokenUrl     pulumi.StringPtrInput
	AdditionalScopes   pulumi.StringPtrInput
	AuthenticationFlow pulumi.StringPtrInput
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	AuthorizationUrl pulumi.StringPtrInput
	ConsumerKey      pulumi.StringInput
	ConsumerSecret   pulumi.StringInput
	Enabled          pulumi.BoolPtrInput
	EnrollmentFlow   pulumi.StringPtrInput
	// Allowed values: - `identifier` - `nameLink` - `nameDeny`
	GroupMatchingMode pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	// Manually configure JWKS keys for use with machine-to-machine authentication. JSON format expected. Use jsonencode() to
	// pass objects.
	OidcJwks pulumi.StringPtrInput
	// Automatically configure JWKS if not specified by `oidcWellKnownUrl`.
	OidcJwksUrl pulumi.StringPtrInput
	// Automatically configure source from OIDC well-known endpoint. URL is taken as is, and should end with
	// `.well-known/openid-configuration`.
	OidcWellKnownUrl pulumi.StringPtrInput
	// Allowed values: - `all` - `any`
	PolicyEngineMode pulumi.StringPtrInput
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	ProfileUrl pulumi.StringPtrInput
	// Allowed values: - `apple` - `openidconnect` - `azuread` - `discord` - `facebook` - `github` - `gitlab` - `google` -
	// `mailcow` - `okta` - `patreon` - `reddit` - `twitch` - `twitter`
	ProviderType pulumi.StringInput
	// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
	RequestTokenUrl pulumi.StringPtrInput
	Slug            pulumi.StringInput
	// Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
	UserMatchingMode pulumi.StringPtrInput
	UserPathTemplate pulumi.StringPtrInput
	Uuid             pulumi.StringPtrInput
}

func (SourceOauthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOauthArgs)(nil)).Elem()
}

type SourceOauthInput interface {
	pulumi.Input

	ToSourceOauthOutput() SourceOauthOutput
	ToSourceOauthOutputWithContext(ctx context.Context) SourceOauthOutput
}

func (*SourceOauth) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOauth)(nil)).Elem()
}

func (i *SourceOauth) ToSourceOauthOutput() SourceOauthOutput {
	return i.ToSourceOauthOutputWithContext(context.Background())
}

func (i *SourceOauth) ToSourceOauthOutputWithContext(ctx context.Context) SourceOauthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOauthOutput)
}

// SourceOauthArrayInput is an input type that accepts SourceOauthArray and SourceOauthArrayOutput values.
// You can construct a concrete instance of `SourceOauthArrayInput` via:
//
//	SourceOauthArray{ SourceOauthArgs{...} }
type SourceOauthArrayInput interface {
	pulumi.Input

	ToSourceOauthArrayOutput() SourceOauthArrayOutput
	ToSourceOauthArrayOutputWithContext(context.Context) SourceOauthArrayOutput
}

type SourceOauthArray []SourceOauthInput

func (SourceOauthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOauth)(nil)).Elem()
}

func (i SourceOauthArray) ToSourceOauthArrayOutput() SourceOauthArrayOutput {
	return i.ToSourceOauthArrayOutputWithContext(context.Background())
}

func (i SourceOauthArray) ToSourceOauthArrayOutputWithContext(ctx context.Context) SourceOauthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOauthArrayOutput)
}

// SourceOauthMapInput is an input type that accepts SourceOauthMap and SourceOauthMapOutput values.
// You can construct a concrete instance of `SourceOauthMapInput` via:
//
//	SourceOauthMap{ "key": SourceOauthArgs{...} }
type SourceOauthMapInput interface {
	pulumi.Input

	ToSourceOauthMapOutput() SourceOauthMapOutput
	ToSourceOauthMapOutputWithContext(context.Context) SourceOauthMapOutput
}

type SourceOauthMap map[string]SourceOauthInput

func (SourceOauthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOauth)(nil)).Elem()
}

func (i SourceOauthMap) ToSourceOauthMapOutput() SourceOauthMapOutput {
	return i.ToSourceOauthMapOutputWithContext(context.Background())
}

func (i SourceOauthMap) ToSourceOauthMapOutputWithContext(ctx context.Context) SourceOauthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOauthMapOutput)
}

type SourceOauthOutput struct{ *pulumi.OutputState }

func (SourceOauthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOauth)(nil)).Elem()
}

func (o SourceOauthOutput) ToSourceOauthOutput() SourceOauthOutput {
	return o
}

func (o SourceOauthOutput) ToSourceOauthOutputWithContext(ctx context.Context) SourceOauthOutput {
	return o
}

// Only required for OAuth1.
func (o SourceOauthOutput) AccessTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.AccessTokenUrl }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) AdditionalScopes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.AdditionalScopes }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) AuthenticationFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.AuthenticationFlow }).(pulumi.StringPtrOutput)
}

// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
func (o SourceOauthOutput) AuthorizationUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.AuthorizationUrl }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) CallbackUri() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.CallbackUri }).(pulumi.StringOutput)
}

func (o SourceOauthOutput) ConsumerKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.ConsumerKey }).(pulumi.StringOutput)
}

func (o SourceOauthOutput) ConsumerSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.ConsumerSecret }).(pulumi.StringOutput)
}

func (o SourceOauthOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SourceOauthOutput) EnrollmentFlow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.EnrollmentFlow }).(pulumi.StringPtrOutput)
}

// Allowed values: - `identifier` - `nameLink` - `nameDeny`
func (o SourceOauthOutput) GroupMatchingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.GroupMatchingMode }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Manually configure JWKS keys for use with machine-to-machine authentication. JSON format expected. Use jsonencode() to
// pass objects.
func (o SourceOauthOutput) OidcJwks() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.OidcJwks }).(pulumi.StringOutput)
}

// Automatically configure JWKS if not specified by `oidcWellKnownUrl`.
func (o SourceOauthOutput) OidcJwksUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.OidcJwksUrl }).(pulumi.StringPtrOutput)
}

// Automatically configure source from OIDC well-known endpoint. URL is taken as is, and should end with
// `.well-known/openid-configuration`.
func (o SourceOauthOutput) OidcWellKnownUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.OidcWellKnownUrl }).(pulumi.StringPtrOutput)
}

// Allowed values: - `all` - `any`
func (o SourceOauthOutput) PolicyEngineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.PolicyEngineMode }).(pulumi.StringPtrOutput)
}

// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
func (o SourceOauthOutput) ProfileUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.ProfileUrl }).(pulumi.StringPtrOutput)
}

// Allowed values: - `apple` - `openidconnect` - `azuread` - `discord` - `facebook` - `github` - `gitlab` - `google` -
// `mailcow` - `okta` - `patreon` - `reddit` - `twitch` - `twitter`
func (o SourceOauthOutput) ProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.ProviderType }).(pulumi.StringOutput)
}

// Manually configure OAuth2 URLs when `oidcWellKnownUrl` is not set.
func (o SourceOauthOutput) RequestTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.RequestTokenUrl }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Allowed values: - `identifier` - `emailLink` - `emailDeny` - `usernameLink` - `usernameDeny`
func (o SourceOauthOutput) UserMatchingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.UserMatchingMode }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) UserPathTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringPtrOutput { return v.UserPathTemplate }).(pulumi.StringPtrOutput)
}

func (o SourceOauthOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOauth) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type SourceOauthArrayOutput struct{ *pulumi.OutputState }

func (SourceOauthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOauth)(nil)).Elem()
}

func (o SourceOauthArrayOutput) ToSourceOauthArrayOutput() SourceOauthArrayOutput {
	return o
}

func (o SourceOauthArrayOutput) ToSourceOauthArrayOutputWithContext(ctx context.Context) SourceOauthArrayOutput {
	return o
}

func (o SourceOauthArrayOutput) Index(i pulumi.IntInput) SourceOauthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceOauth {
		return vs[0].([]*SourceOauth)[vs[1].(int)]
	}).(SourceOauthOutput)
}

type SourceOauthMapOutput struct{ *pulumi.OutputState }

func (SourceOauthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOauth)(nil)).Elem()
}

func (o SourceOauthMapOutput) ToSourceOauthMapOutput() SourceOauthMapOutput {
	return o
}

func (o SourceOauthMapOutput) ToSourceOauthMapOutputWithContext(ctx context.Context) SourceOauthMapOutput {
	return o
}

func (o SourceOauthMapOutput) MapIndex(k pulumi.StringInput) SourceOauthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceOauth {
		return vs[0].(map[string]*SourceOauth)[vs[1].(string)]
	}).(SourceOauthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOauthInput)(nil)).Elem(), &SourceOauth{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOauthArrayInput)(nil)).Elem(), SourceOauthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOauthMapInput)(nil)).Elem(), SourceOauthMap{})
	pulumi.RegisterOutputType(SourceOauthOutput{})
	pulumi.RegisterOutputType(SourceOauthArrayOutput{})
	pulumi.RegisterOutputType(SourceOauthMapOutput{})
}

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
//			default_authorization_flow, err := authentik.LookupFlow(ctx, &authentik.LookupFlowArgs{
//				Slug: pulumi.StringRef("default-provider-authorization-implicit-consent"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewSourcePlex(ctx, "name", &authentik.SourcePlexArgs{
//				Slug:               pulumi.String("plex"),
//				AuthenticationFlow: *pulumi.String(default_authorization_flow.Id),
//				EnrollmentFlow:     *pulumi.String(default_authorization_flow.Id),
//				ClientId:           pulumi.String("foo-bar-baz"),
//				PlexToken:          pulumi.String("foo"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SourcePlex struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	AllowFriends       pulumi.BoolPtrOutput     `pulumi:"allowFriends"`
	AllowedServers     pulumi.StringArrayOutput `pulumi:"allowedServers"`
	AuthenticationFlow pulumi.StringOutput      `pulumi:"authenticationFlow"`
	ClientId           pulumi.StringOutput      `pulumi:"clientId"`
	// Defaults to `true`.
	Enabled        pulumi.BoolPtrOutput `pulumi:"enabled"`
	EnrollmentFlow pulumi.StringOutput  `pulumi:"enrollmentFlow"`
	Name           pulumi.StringOutput  `pulumi:"name"`
	PlexToken      pulumi.StringOutput  `pulumi:"plexToken"`
	// Defaults to `any`.
	PolicyEngineMode pulumi.StringPtrOutput `pulumi:"policyEngineMode"`
	Slug             pulumi.StringOutput    `pulumi:"slug"`
	// Defaults to `identifier`.
	UserMatchingMode pulumi.StringPtrOutput `pulumi:"userMatchingMode"`
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate pulumi.StringPtrOutput `pulumi:"userPathTemplate"`
	// Generated.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewSourcePlex registers a new resource with the given unique name, arguments, and options.
func NewSourcePlex(ctx *pulumi.Context,
	name string, args *SourcePlexArgs, opts ...pulumi.ResourceOption) (*SourcePlex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationFlow == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationFlow'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.EnrollmentFlow == nil {
		return nil, errors.New("invalid value for required argument 'EnrollmentFlow'")
	}
	if args.PlexToken == nil {
		return nil, errors.New("invalid value for required argument 'PlexToken'")
	}
	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	if args.PlexToken != nil {
		args.PlexToken = pulumi.ToSecret(args.PlexToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"plexToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SourcePlex
	err := ctx.RegisterResource("authentik:index/sourcePlex:SourcePlex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePlex gets an existing SourcePlex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePlex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePlexState, opts ...pulumi.ResourceOption) (*SourcePlex, error) {
	var resource SourcePlex
	err := ctx.ReadResource("authentik:index/sourcePlex:SourcePlex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePlex resources.
type sourcePlexState struct {
	// Defaults to `true`.
	AllowFriends       *bool    `pulumi:"allowFriends"`
	AllowedServers     []string `pulumi:"allowedServers"`
	AuthenticationFlow *string  `pulumi:"authenticationFlow"`
	ClientId           *string  `pulumi:"clientId"`
	// Defaults to `true`.
	Enabled        *bool   `pulumi:"enabled"`
	EnrollmentFlow *string `pulumi:"enrollmentFlow"`
	Name           *string `pulumi:"name"`
	PlexToken      *string `pulumi:"plexToken"`
	// Defaults to `any`.
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	Slug             *string `pulumi:"slug"`
	// Defaults to `identifier`.
	UserMatchingMode *string `pulumi:"userMatchingMode"`
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `pulumi:"userPathTemplate"`
	// Generated.
	Uuid *string `pulumi:"uuid"`
}

type SourcePlexState struct {
	// Defaults to `true`.
	AllowFriends       pulumi.BoolPtrInput
	AllowedServers     pulumi.StringArrayInput
	AuthenticationFlow pulumi.StringPtrInput
	ClientId           pulumi.StringPtrInput
	// Defaults to `true`.
	Enabled        pulumi.BoolPtrInput
	EnrollmentFlow pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	PlexToken      pulumi.StringPtrInput
	// Defaults to `any`.
	PolicyEngineMode pulumi.StringPtrInput
	Slug             pulumi.StringPtrInput
	// Defaults to `identifier`.
	UserMatchingMode pulumi.StringPtrInput
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate pulumi.StringPtrInput
	// Generated.
	Uuid pulumi.StringPtrInput
}

func (SourcePlexState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePlexState)(nil)).Elem()
}

type sourcePlexArgs struct {
	// Defaults to `true`.
	AllowFriends       *bool    `pulumi:"allowFriends"`
	AllowedServers     []string `pulumi:"allowedServers"`
	AuthenticationFlow string   `pulumi:"authenticationFlow"`
	ClientId           string   `pulumi:"clientId"`
	// Defaults to `true`.
	Enabled        *bool   `pulumi:"enabled"`
	EnrollmentFlow string  `pulumi:"enrollmentFlow"`
	Name           *string `pulumi:"name"`
	PlexToken      string  `pulumi:"plexToken"`
	// Defaults to `any`.
	PolicyEngineMode *string `pulumi:"policyEngineMode"`
	Slug             string  `pulumi:"slug"`
	// Defaults to `identifier`.
	UserMatchingMode *string `pulumi:"userMatchingMode"`
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `pulumi:"userPathTemplate"`
	// Generated.
	Uuid *string `pulumi:"uuid"`
}

// The set of arguments for constructing a SourcePlex resource.
type SourcePlexArgs struct {
	// Defaults to `true`.
	AllowFriends       pulumi.BoolPtrInput
	AllowedServers     pulumi.StringArrayInput
	AuthenticationFlow pulumi.StringInput
	ClientId           pulumi.StringInput
	// Defaults to `true`.
	Enabled        pulumi.BoolPtrInput
	EnrollmentFlow pulumi.StringInput
	Name           pulumi.StringPtrInput
	PlexToken      pulumi.StringInput
	// Defaults to `any`.
	PolicyEngineMode pulumi.StringPtrInput
	Slug             pulumi.StringInput
	// Defaults to `identifier`.
	UserMatchingMode pulumi.StringPtrInput
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate pulumi.StringPtrInput
	// Generated.
	Uuid pulumi.StringPtrInput
}

func (SourcePlexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePlexArgs)(nil)).Elem()
}

type SourcePlexInput interface {
	pulumi.Input

	ToSourcePlexOutput() SourcePlexOutput
	ToSourcePlexOutputWithContext(ctx context.Context) SourcePlexOutput
}

func (*SourcePlex) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePlex)(nil)).Elem()
}

func (i *SourcePlex) ToSourcePlexOutput() SourcePlexOutput {
	return i.ToSourcePlexOutputWithContext(context.Background())
}

func (i *SourcePlex) ToSourcePlexOutputWithContext(ctx context.Context) SourcePlexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePlexOutput)
}

func (i *SourcePlex) ToOutput(ctx context.Context) pulumix.Output[*SourcePlex] {
	return pulumix.Output[*SourcePlex]{
		OutputState: i.ToSourcePlexOutputWithContext(ctx).OutputState,
	}
}

// SourcePlexArrayInput is an input type that accepts SourcePlexArray and SourcePlexArrayOutput values.
// You can construct a concrete instance of `SourcePlexArrayInput` via:
//
//	SourcePlexArray{ SourcePlexArgs{...} }
type SourcePlexArrayInput interface {
	pulumi.Input

	ToSourcePlexArrayOutput() SourcePlexArrayOutput
	ToSourcePlexArrayOutputWithContext(context.Context) SourcePlexArrayOutput
}

type SourcePlexArray []SourcePlexInput

func (SourcePlexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePlex)(nil)).Elem()
}

func (i SourcePlexArray) ToSourcePlexArrayOutput() SourcePlexArrayOutput {
	return i.ToSourcePlexArrayOutputWithContext(context.Background())
}

func (i SourcePlexArray) ToSourcePlexArrayOutputWithContext(ctx context.Context) SourcePlexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePlexArrayOutput)
}

func (i SourcePlexArray) ToOutput(ctx context.Context) pulumix.Output[[]*SourcePlex] {
	return pulumix.Output[[]*SourcePlex]{
		OutputState: i.ToSourcePlexArrayOutputWithContext(ctx).OutputState,
	}
}

// SourcePlexMapInput is an input type that accepts SourcePlexMap and SourcePlexMapOutput values.
// You can construct a concrete instance of `SourcePlexMapInput` via:
//
//	SourcePlexMap{ "key": SourcePlexArgs{...} }
type SourcePlexMapInput interface {
	pulumi.Input

	ToSourcePlexMapOutput() SourcePlexMapOutput
	ToSourcePlexMapOutputWithContext(context.Context) SourcePlexMapOutput
}

type SourcePlexMap map[string]SourcePlexInput

func (SourcePlexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePlex)(nil)).Elem()
}

func (i SourcePlexMap) ToSourcePlexMapOutput() SourcePlexMapOutput {
	return i.ToSourcePlexMapOutputWithContext(context.Background())
}

func (i SourcePlexMap) ToSourcePlexMapOutputWithContext(ctx context.Context) SourcePlexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePlexMapOutput)
}

func (i SourcePlexMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SourcePlex] {
	return pulumix.Output[map[string]*SourcePlex]{
		OutputState: i.ToSourcePlexMapOutputWithContext(ctx).OutputState,
	}
}

type SourcePlexOutput struct{ *pulumi.OutputState }

func (SourcePlexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePlex)(nil)).Elem()
}

func (o SourcePlexOutput) ToSourcePlexOutput() SourcePlexOutput {
	return o
}

func (o SourcePlexOutput) ToSourcePlexOutputWithContext(ctx context.Context) SourcePlexOutput {
	return o
}

func (o SourcePlexOutput) ToOutput(ctx context.Context) pulumix.Output[*SourcePlex] {
	return pulumix.Output[*SourcePlex]{
		OutputState: o.OutputState,
	}
}

// Defaults to `true`.
func (o SourcePlexOutput) AllowFriends() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.BoolPtrOutput { return v.AllowFriends }).(pulumi.BoolPtrOutput)
}

func (o SourcePlexOutput) AllowedServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringArrayOutput { return v.AllowedServers }).(pulumi.StringArrayOutput)
}

func (o SourcePlexOutput) AuthenticationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.AuthenticationFlow }).(pulumi.StringOutput)
}

func (o SourcePlexOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// Defaults to `true`.
func (o SourcePlexOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SourcePlexOutput) EnrollmentFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.EnrollmentFlow }).(pulumi.StringOutput)
}

func (o SourcePlexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SourcePlexOutput) PlexToken() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.PlexToken }).(pulumi.StringOutput)
}

// Defaults to `any`.
func (o SourcePlexOutput) PolicyEngineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringPtrOutput { return v.PolicyEngineMode }).(pulumi.StringPtrOutput)
}

func (o SourcePlexOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Defaults to `identifier`.
func (o SourcePlexOutput) UserMatchingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringPtrOutput { return v.UserMatchingMode }).(pulumi.StringPtrOutput)
}

// Defaults to `goauthentik.io/sources/%(slug)s`.
func (o SourcePlexOutput) UserPathTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringPtrOutput { return v.UserPathTemplate }).(pulumi.StringPtrOutput)
}

// Generated.
func (o SourcePlexOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePlex) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type SourcePlexArrayOutput struct{ *pulumi.OutputState }

func (SourcePlexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePlex)(nil)).Elem()
}

func (o SourcePlexArrayOutput) ToSourcePlexArrayOutput() SourcePlexArrayOutput {
	return o
}

func (o SourcePlexArrayOutput) ToSourcePlexArrayOutputWithContext(ctx context.Context) SourcePlexArrayOutput {
	return o
}

func (o SourcePlexArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SourcePlex] {
	return pulumix.Output[[]*SourcePlex]{
		OutputState: o.OutputState,
	}
}

func (o SourcePlexArrayOutput) Index(i pulumi.IntInput) SourcePlexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourcePlex {
		return vs[0].([]*SourcePlex)[vs[1].(int)]
	}).(SourcePlexOutput)
}

type SourcePlexMapOutput struct{ *pulumi.OutputState }

func (SourcePlexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePlex)(nil)).Elem()
}

func (o SourcePlexMapOutput) ToSourcePlexMapOutput() SourcePlexMapOutput {
	return o
}

func (o SourcePlexMapOutput) ToSourcePlexMapOutputWithContext(ctx context.Context) SourcePlexMapOutput {
	return o
}

func (o SourcePlexMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SourcePlex] {
	return pulumix.Output[map[string]*SourcePlex]{
		OutputState: o.OutputState,
	}
}

func (o SourcePlexMapOutput) MapIndex(k pulumi.StringInput) SourcePlexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourcePlex {
		return vs[0].(map[string]*SourcePlex)[vs[1].(string)]
	}).(SourcePlexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePlexInput)(nil)).Elem(), &SourcePlex{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePlexArrayInput)(nil)).Elem(), SourcePlexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePlexMapInput)(nil)).Elem(), SourcePlexMap{})
	pulumi.RegisterOutputType(SourcePlexOutput{})
	pulumi.RegisterOutputType(SourcePlexArrayOutput{})
	pulumi.RegisterOutputType(SourcePlexMapOutput{})
}

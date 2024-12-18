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
//			// Create LDAP Source
//			_, err := authentik.NewSourceLdap(ctx, "name", &authentik.SourceLdapArgs{
//				Name:         pulumi.String("ldap-test"),
//				Slug:         pulumi.String("ldap-test"),
//				ServerUri:    pulumi.String("ldaps://1.2.3.4"),
//				BindCn:       pulumi.String("foo"),
//				BindPassword: pulumi.String("bar"),
//				BaseDn:       pulumi.String("dn=foo"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SourceLdap struct {
	pulumi.CustomResourceState

	// Defaults to ``.
	AdditionalGroupDn pulumi.StringPtrOutput `pulumi:"additionalGroupDn"`
	// Defaults to ``.
	AdditionalUserDn pulumi.StringPtrOutput `pulumi:"additionalUserDn"`
	BaseDn           pulumi.StringOutput    `pulumi:"baseDn"`
	BindCn           pulumi.StringOutput    `pulumi:"bindCn"`
	BindPassword     pulumi.StringOutput    `pulumi:"bindPassword"`
	// Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Defaults to `member`.
	GroupMembershipField pulumi.StringPtrOutput `pulumi:"groupMembershipField"`
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter pulumi.StringPtrOutput `pulumi:"groupObjectFilter"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	// Defaults to `objectSid`.
	ObjectUniquenessField pulumi.StringPtrOutput `pulumi:"objectUniquenessField"`
	// Defaults to `false`.
	PasswordLoginUpdateInternalPassword pulumi.BoolPtrOutput     `pulumi:"passwordLoginUpdateInternalPassword"`
	PropertyMappings                    pulumi.StringArrayOutput `pulumi:"propertyMappings"`
	PropertyMappingsGroups              pulumi.StringArrayOutput `pulumi:"propertyMappingsGroups"`
	ServerUri                           pulumi.StringOutput      `pulumi:"serverUri"`
	Slug                                pulumi.StringOutput      `pulumi:"slug"`
	// Defaults to `false`.
	Sni pulumi.BoolPtrOutput `pulumi:"sni"`
	// Defaults to `true`.
	StartTls pulumi.BoolPtrOutput `pulumi:"startTls"`
	// Defaults to `true`.
	SyncGroups      pulumi.BoolPtrOutput   `pulumi:"syncGroups"`
	SyncParentGroup pulumi.StringPtrOutput `pulumi:"syncParentGroup"`
	// Defaults to `true`.
	SyncUsers pulumi.BoolPtrOutput `pulumi:"syncUsers"`
	// Defaults to `true`.
	SyncUsersPassword pulumi.BoolPtrOutput `pulumi:"syncUsersPassword"`
	// Defaults to `(objectClass=person)`.
	UserObjectFilter pulumi.StringPtrOutput `pulumi:"userObjectFilter"`
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate pulumi.StringPtrOutput `pulumi:"userPathTemplate"`
	// Generated.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewSourceLdap registers a new resource with the given unique name, arguments, and options.
func NewSourceLdap(ctx *pulumi.Context,
	name string, args *SourceLdapArgs, opts ...pulumi.ResourceOption) (*SourceLdap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseDn == nil {
		return nil, errors.New("invalid value for required argument 'BaseDn'")
	}
	if args.BindCn == nil {
		return nil, errors.New("invalid value for required argument 'BindCn'")
	}
	if args.BindPassword == nil {
		return nil, errors.New("invalid value for required argument 'BindPassword'")
	}
	if args.ServerUri == nil {
		return nil, errors.New("invalid value for required argument 'ServerUri'")
	}
	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	if args.BindPassword != nil {
		args.BindPassword = pulumi.ToSecret(args.BindPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bindPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SourceLdap
	err := ctx.RegisterResource("authentik:index/sourceLdap:SourceLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLdap gets an existing SourceLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLdapState, opts ...pulumi.ResourceOption) (*SourceLdap, error) {
	var resource SourceLdap
	err := ctx.ReadResource("authentik:index/sourceLdap:SourceLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLdap resources.
type sourceLdapState struct {
	// Defaults to ``.
	AdditionalGroupDn *string `pulumi:"additionalGroupDn"`
	// Defaults to ``.
	AdditionalUserDn *string `pulumi:"additionalUserDn"`
	BaseDn           *string `pulumi:"baseDn"`
	BindCn           *string `pulumi:"bindCn"`
	BindPassword     *string `pulumi:"bindPassword"`
	// Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Defaults to `member`.
	GroupMembershipField *string `pulumi:"groupMembershipField"`
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter *string `pulumi:"groupObjectFilter"`
	Name              *string `pulumi:"name"`
	// Defaults to `objectSid`.
	ObjectUniquenessField *string `pulumi:"objectUniquenessField"`
	// Defaults to `false`.
	PasswordLoginUpdateInternalPassword *bool    `pulumi:"passwordLoginUpdateInternalPassword"`
	PropertyMappings                    []string `pulumi:"propertyMappings"`
	PropertyMappingsGroups              []string `pulumi:"propertyMappingsGroups"`
	ServerUri                           *string  `pulumi:"serverUri"`
	Slug                                *string  `pulumi:"slug"`
	// Defaults to `false`.
	Sni *bool `pulumi:"sni"`
	// Defaults to `true`.
	StartTls *bool `pulumi:"startTls"`
	// Defaults to `true`.
	SyncGroups      *bool   `pulumi:"syncGroups"`
	SyncParentGroup *string `pulumi:"syncParentGroup"`
	// Defaults to `true`.
	SyncUsers *bool `pulumi:"syncUsers"`
	// Defaults to `true`.
	SyncUsersPassword *bool `pulumi:"syncUsersPassword"`
	// Defaults to `(objectClass=person)`.
	UserObjectFilter *string `pulumi:"userObjectFilter"`
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `pulumi:"userPathTemplate"`
	// Generated.
	Uuid *string `pulumi:"uuid"`
}

type SourceLdapState struct {
	// Defaults to ``.
	AdditionalGroupDn pulumi.StringPtrInput
	// Defaults to ``.
	AdditionalUserDn pulumi.StringPtrInput
	BaseDn           pulumi.StringPtrInput
	BindCn           pulumi.StringPtrInput
	BindPassword     pulumi.StringPtrInput
	// Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Defaults to `member`.
	GroupMembershipField pulumi.StringPtrInput
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	// Defaults to `objectSid`.
	ObjectUniquenessField pulumi.StringPtrInput
	// Defaults to `false`.
	PasswordLoginUpdateInternalPassword pulumi.BoolPtrInput
	PropertyMappings                    pulumi.StringArrayInput
	PropertyMappingsGroups              pulumi.StringArrayInput
	ServerUri                           pulumi.StringPtrInput
	Slug                                pulumi.StringPtrInput
	// Defaults to `false`.
	Sni pulumi.BoolPtrInput
	// Defaults to `true`.
	StartTls pulumi.BoolPtrInput
	// Defaults to `true`.
	SyncGroups      pulumi.BoolPtrInput
	SyncParentGroup pulumi.StringPtrInput
	// Defaults to `true`.
	SyncUsers pulumi.BoolPtrInput
	// Defaults to `true`.
	SyncUsersPassword pulumi.BoolPtrInput
	// Defaults to `(objectClass=person)`.
	UserObjectFilter pulumi.StringPtrInput
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate pulumi.StringPtrInput
	// Generated.
	Uuid pulumi.StringPtrInput
}

func (SourceLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLdapState)(nil)).Elem()
}

type sourceLdapArgs struct {
	// Defaults to ``.
	AdditionalGroupDn *string `pulumi:"additionalGroupDn"`
	// Defaults to ``.
	AdditionalUserDn *string `pulumi:"additionalUserDn"`
	BaseDn           string  `pulumi:"baseDn"`
	BindCn           string  `pulumi:"bindCn"`
	BindPassword     string  `pulumi:"bindPassword"`
	// Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Defaults to `member`.
	GroupMembershipField *string `pulumi:"groupMembershipField"`
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter *string `pulumi:"groupObjectFilter"`
	Name              *string `pulumi:"name"`
	// Defaults to `objectSid`.
	ObjectUniquenessField *string `pulumi:"objectUniquenessField"`
	// Defaults to `false`.
	PasswordLoginUpdateInternalPassword *bool    `pulumi:"passwordLoginUpdateInternalPassword"`
	PropertyMappings                    []string `pulumi:"propertyMappings"`
	PropertyMappingsGroups              []string `pulumi:"propertyMappingsGroups"`
	ServerUri                           string   `pulumi:"serverUri"`
	Slug                                string   `pulumi:"slug"`
	// Defaults to `false`.
	Sni *bool `pulumi:"sni"`
	// Defaults to `true`.
	StartTls *bool `pulumi:"startTls"`
	// Defaults to `true`.
	SyncGroups      *bool   `pulumi:"syncGroups"`
	SyncParentGroup *string `pulumi:"syncParentGroup"`
	// Defaults to `true`.
	SyncUsers *bool `pulumi:"syncUsers"`
	// Defaults to `true`.
	SyncUsersPassword *bool `pulumi:"syncUsersPassword"`
	// Defaults to `(objectClass=person)`.
	UserObjectFilter *string `pulumi:"userObjectFilter"`
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `pulumi:"userPathTemplate"`
	// Generated.
	Uuid *string `pulumi:"uuid"`
}

// The set of arguments for constructing a SourceLdap resource.
type SourceLdapArgs struct {
	// Defaults to ``.
	AdditionalGroupDn pulumi.StringPtrInput
	// Defaults to ``.
	AdditionalUserDn pulumi.StringPtrInput
	BaseDn           pulumi.StringInput
	BindCn           pulumi.StringInput
	BindPassword     pulumi.StringInput
	// Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Defaults to `member`.
	GroupMembershipField pulumi.StringPtrInput
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	// Defaults to `objectSid`.
	ObjectUniquenessField pulumi.StringPtrInput
	// Defaults to `false`.
	PasswordLoginUpdateInternalPassword pulumi.BoolPtrInput
	PropertyMappings                    pulumi.StringArrayInput
	PropertyMappingsGroups              pulumi.StringArrayInput
	ServerUri                           pulumi.StringInput
	Slug                                pulumi.StringInput
	// Defaults to `false`.
	Sni pulumi.BoolPtrInput
	// Defaults to `true`.
	StartTls pulumi.BoolPtrInput
	// Defaults to `true`.
	SyncGroups      pulumi.BoolPtrInput
	SyncParentGroup pulumi.StringPtrInput
	// Defaults to `true`.
	SyncUsers pulumi.BoolPtrInput
	// Defaults to `true`.
	SyncUsersPassword pulumi.BoolPtrInput
	// Defaults to `(objectClass=person)`.
	UserObjectFilter pulumi.StringPtrInput
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate pulumi.StringPtrInput
	// Generated.
	Uuid pulumi.StringPtrInput
}

func (SourceLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLdapArgs)(nil)).Elem()
}

type SourceLdapInput interface {
	pulumi.Input

	ToSourceLdapOutput() SourceLdapOutput
	ToSourceLdapOutputWithContext(ctx context.Context) SourceLdapOutput
}

func (*SourceLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLdap)(nil)).Elem()
}

func (i *SourceLdap) ToSourceLdapOutput() SourceLdapOutput {
	return i.ToSourceLdapOutputWithContext(context.Background())
}

func (i *SourceLdap) ToSourceLdapOutputWithContext(ctx context.Context) SourceLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLdapOutput)
}

// SourceLdapArrayInput is an input type that accepts SourceLdapArray and SourceLdapArrayOutput values.
// You can construct a concrete instance of `SourceLdapArrayInput` via:
//
//	SourceLdapArray{ SourceLdapArgs{...} }
type SourceLdapArrayInput interface {
	pulumi.Input

	ToSourceLdapArrayOutput() SourceLdapArrayOutput
	ToSourceLdapArrayOutputWithContext(context.Context) SourceLdapArrayOutput
}

type SourceLdapArray []SourceLdapInput

func (SourceLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLdap)(nil)).Elem()
}

func (i SourceLdapArray) ToSourceLdapArrayOutput() SourceLdapArrayOutput {
	return i.ToSourceLdapArrayOutputWithContext(context.Background())
}

func (i SourceLdapArray) ToSourceLdapArrayOutputWithContext(ctx context.Context) SourceLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLdapArrayOutput)
}

// SourceLdapMapInput is an input type that accepts SourceLdapMap and SourceLdapMapOutput values.
// You can construct a concrete instance of `SourceLdapMapInput` via:
//
//	SourceLdapMap{ "key": SourceLdapArgs{...} }
type SourceLdapMapInput interface {
	pulumi.Input

	ToSourceLdapMapOutput() SourceLdapMapOutput
	ToSourceLdapMapOutputWithContext(context.Context) SourceLdapMapOutput
}

type SourceLdapMap map[string]SourceLdapInput

func (SourceLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLdap)(nil)).Elem()
}

func (i SourceLdapMap) ToSourceLdapMapOutput() SourceLdapMapOutput {
	return i.ToSourceLdapMapOutputWithContext(context.Background())
}

func (i SourceLdapMap) ToSourceLdapMapOutputWithContext(ctx context.Context) SourceLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLdapMapOutput)
}

type SourceLdapOutput struct{ *pulumi.OutputState }

func (SourceLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLdap)(nil)).Elem()
}

func (o SourceLdapOutput) ToSourceLdapOutput() SourceLdapOutput {
	return o
}

func (o SourceLdapOutput) ToSourceLdapOutputWithContext(ctx context.Context) SourceLdapOutput {
	return o
}

// Defaults to “.
func (o SourceLdapOutput) AdditionalGroupDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.AdditionalGroupDn }).(pulumi.StringPtrOutput)
}

// Defaults to “.
func (o SourceLdapOutput) AdditionalUserDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.AdditionalUserDn }).(pulumi.StringPtrOutput)
}

func (o SourceLdapOutput) BaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.BaseDn }).(pulumi.StringOutput)
}

func (o SourceLdapOutput) BindCn() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.BindCn }).(pulumi.StringOutput)
}

func (o SourceLdapOutput) BindPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.BindPassword }).(pulumi.StringOutput)
}

// Defaults to `true`.
func (o SourceLdapOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Defaults to `member`.
func (o SourceLdapOutput) GroupMembershipField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.GroupMembershipField }).(pulumi.StringPtrOutput)
}

// Defaults to `(objectClass=group)`.
func (o SourceLdapOutput) GroupObjectFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.GroupObjectFilter }).(pulumi.StringPtrOutput)
}

func (o SourceLdapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defaults to `objectSid`.
func (o SourceLdapOutput) ObjectUniquenessField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.ObjectUniquenessField }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o SourceLdapOutput) PasswordLoginUpdateInternalPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.PasswordLoginUpdateInternalPassword }).(pulumi.BoolPtrOutput)
}

func (o SourceLdapOutput) PropertyMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringArrayOutput { return v.PropertyMappings }).(pulumi.StringArrayOutput)
}

func (o SourceLdapOutput) PropertyMappingsGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringArrayOutput { return v.PropertyMappingsGroups }).(pulumi.StringArrayOutput)
}

func (o SourceLdapOutput) ServerUri() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.ServerUri }).(pulumi.StringOutput)
}

func (o SourceLdapOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Defaults to `false`.
func (o SourceLdapOutput) Sni() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.Sni }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o SourceLdapOutput) StartTls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.StartTls }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o SourceLdapOutput) SyncGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.SyncGroups }).(pulumi.BoolPtrOutput)
}

func (o SourceLdapOutput) SyncParentGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.SyncParentGroup }).(pulumi.StringPtrOutput)
}

// Defaults to `true`.
func (o SourceLdapOutput) SyncUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.SyncUsers }).(pulumi.BoolPtrOutput)
}

// Defaults to `true`.
func (o SourceLdapOutput) SyncUsersPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.BoolPtrOutput { return v.SyncUsersPassword }).(pulumi.BoolPtrOutput)
}

// Defaults to `(objectClass=person)`.
func (o SourceLdapOutput) UserObjectFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.UserObjectFilter }).(pulumi.StringPtrOutput)
}

// Defaults to `goauthentik.io/sources/%(slug)s`.
func (o SourceLdapOutput) UserPathTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringPtrOutput { return v.UserPathTemplate }).(pulumi.StringPtrOutput)
}

// Generated.
func (o SourceLdapOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLdap) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type SourceLdapArrayOutput struct{ *pulumi.OutputState }

func (SourceLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLdap)(nil)).Elem()
}

func (o SourceLdapArrayOutput) ToSourceLdapArrayOutput() SourceLdapArrayOutput {
	return o
}

func (o SourceLdapArrayOutput) ToSourceLdapArrayOutputWithContext(ctx context.Context) SourceLdapArrayOutput {
	return o
}

func (o SourceLdapArrayOutput) Index(i pulumi.IntInput) SourceLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceLdap {
		return vs[0].([]*SourceLdap)[vs[1].(int)]
	}).(SourceLdapOutput)
}

type SourceLdapMapOutput struct{ *pulumi.OutputState }

func (SourceLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLdap)(nil)).Elem()
}

func (o SourceLdapMapOutput) ToSourceLdapMapOutput() SourceLdapMapOutput {
	return o
}

func (o SourceLdapMapOutput) ToSourceLdapMapOutputWithContext(ctx context.Context) SourceLdapMapOutput {
	return o
}

func (o SourceLdapMapOutput) MapIndex(k pulumi.StringInput) SourceLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceLdap {
		return vs[0].(map[string]*SourceLdap)[vs[1].(string)]
	}).(SourceLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLdapInput)(nil)).Elem(), &SourceLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLdapArrayInput)(nil)).Elem(), SourceLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLdapMapInput)(nil)).Elem(), SourceLdapMap{})
	pulumi.RegisterOutputType(SourceLdapOutput{})
	pulumi.RegisterOutputType(SourceLdapArrayOutput{})
	pulumi.RegisterOutputType(SourceLdapMapOutput{})
}

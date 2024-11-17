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
//			// Create a user
//			_, err := authentik.NewUser(ctx, "nameUser", &authentik.UserArgs{
//				Username: pulumi.String("user"),
//			})
//			if err != nil {
//				return err
//			}
//			group, err := authentik.NewGroup(ctx, "group", nil)
//			if err != nil {
//				return err
//			}
//			_, err = authentik.NewUser(ctx, "nameIndex/userUser", &authentik.UserArgs{
//				Username: pulumi.String("user"),
//				Groups: pulumi.StringArray{
//					group.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type User struct {
	pulumi.CustomResourceState

	// JSON format expected. Use jsonencode() to pass objects.
	Attributes pulumi.StringPtrOutput   `pulumi:"attributes"`
	Email      pulumi.StringPtrOutput   `pulumi:"email"`
	Groups     pulumi.StringArrayOutput `pulumi:"groups"`
	IsActive   pulumi.BoolPtrOutput     `pulumi:"isActive"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	// Optionally set the user's password. Changing the password in authentik will not trigger an update here.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	Path     pulumi.StringPtrOutput `pulumi:"path"`
	// Allowed values: - `internal` - `external` - `serviceAccount` - `internalServiceAccount`
	Type     pulumi.StringPtrOutput `pulumi:"type"`
	Username pulumi.StringOutput    `pulumi:"username"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("authentik:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("authentik:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// JSON format expected. Use jsonencode() to pass objects.
	Attributes *string  `pulumi:"attributes"`
	Email      *string  `pulumi:"email"`
	Groups     []string `pulumi:"groups"`
	IsActive   *bool    `pulumi:"isActive"`
	Name       *string  `pulumi:"name"`
	// Optionally set the user's password. Changing the password in authentik will not trigger an update here.
	Password *string `pulumi:"password"`
	Path     *string `pulumi:"path"`
	// Allowed values: - `internal` - `external` - `serviceAccount` - `internalServiceAccount`
	Type     *string `pulumi:"type"`
	Username *string `pulumi:"username"`
}

type UserState struct {
	// JSON format expected. Use jsonencode() to pass objects.
	Attributes pulumi.StringPtrInput
	Email      pulumi.StringPtrInput
	Groups     pulumi.StringArrayInput
	IsActive   pulumi.BoolPtrInput
	Name       pulumi.StringPtrInput
	// Optionally set the user's password. Changing the password in authentik will not trigger an update here.
	Password pulumi.StringPtrInput
	Path     pulumi.StringPtrInput
	// Allowed values: - `internal` - `external` - `serviceAccount` - `internalServiceAccount`
	Type     pulumi.StringPtrInput
	Username pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// JSON format expected. Use jsonencode() to pass objects.
	Attributes *string  `pulumi:"attributes"`
	Email      *string  `pulumi:"email"`
	Groups     []string `pulumi:"groups"`
	IsActive   *bool    `pulumi:"isActive"`
	Name       *string  `pulumi:"name"`
	// Optionally set the user's password. Changing the password in authentik will not trigger an update here.
	Password *string `pulumi:"password"`
	Path     *string `pulumi:"path"`
	// Allowed values: - `internal` - `external` - `serviceAccount` - `internalServiceAccount`
	Type     *string `pulumi:"type"`
	Username string  `pulumi:"username"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// JSON format expected. Use jsonencode() to pass objects.
	Attributes pulumi.StringPtrInput
	Email      pulumi.StringPtrInput
	Groups     pulumi.StringArrayInput
	IsActive   pulumi.BoolPtrInput
	Name       pulumi.StringPtrInput
	// Optionally set the user's password. Changing the password in authentik will not trigger an update here.
	Password pulumi.StringPtrInput
	Path     pulumi.StringPtrInput
	// Allowed values: - `internal` - `external` - `serviceAccount` - `internalServiceAccount`
	Type     pulumi.StringPtrInput
	Username pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// JSON format expected. Use jsonencode() to pass objects.
func (o UserOutput) Attributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Attributes }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o UserOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optionally set the user's password. Changing the password in authentik will not trigger an update here.
func (o UserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Allowed values: - `internal` - `external` - `serviceAccount` - `internalServiceAccount`
func (o UserOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}

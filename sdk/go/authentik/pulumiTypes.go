// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetGroupUsersObj struct {
	Attributes string `pulumi:"attributes"`
	Email      string `pulumi:"email"`
	IsActive   bool   `pulumi:"isActive"`
	LastLogin  string `pulumi:"lastLogin"`
	Name       string `pulumi:"name"`
	Pk         int    `pulumi:"pk"`
	Uid        string `pulumi:"uid"`
	Username   string `pulumi:"username"`
}

// GetGroupUsersObjInput is an input type that accepts GetGroupUsersObjArgs and GetGroupUsersObjOutput values.
// You can construct a concrete instance of `GetGroupUsersObjInput` via:
//
//	GetGroupUsersObjArgs{...}
type GetGroupUsersObjInput interface {
	pulumi.Input

	ToGetGroupUsersObjOutput() GetGroupUsersObjOutput
	ToGetGroupUsersObjOutputWithContext(context.Context) GetGroupUsersObjOutput
}

type GetGroupUsersObjArgs struct {
	Attributes pulumi.StringInput `pulumi:"attributes"`
	Email      pulumi.StringInput `pulumi:"email"`
	IsActive   pulumi.BoolInput   `pulumi:"isActive"`
	LastLogin  pulumi.StringInput `pulumi:"lastLogin"`
	Name       pulumi.StringInput `pulumi:"name"`
	Pk         pulumi.IntInput    `pulumi:"pk"`
	Uid        pulumi.StringInput `pulumi:"uid"`
	Username   pulumi.StringInput `pulumi:"username"`
}

func (GetGroupUsersObjArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUsersObj)(nil)).Elem()
}

func (i GetGroupUsersObjArgs) ToGetGroupUsersObjOutput() GetGroupUsersObjOutput {
	return i.ToGetGroupUsersObjOutputWithContext(context.Background())
}

func (i GetGroupUsersObjArgs) ToGetGroupUsersObjOutputWithContext(ctx context.Context) GetGroupUsersObjOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupUsersObjOutput)
}

// GetGroupUsersObjArrayInput is an input type that accepts GetGroupUsersObjArray and GetGroupUsersObjArrayOutput values.
// You can construct a concrete instance of `GetGroupUsersObjArrayInput` via:
//
//	GetGroupUsersObjArray{ GetGroupUsersObjArgs{...} }
type GetGroupUsersObjArrayInput interface {
	pulumi.Input

	ToGetGroupUsersObjArrayOutput() GetGroupUsersObjArrayOutput
	ToGetGroupUsersObjArrayOutputWithContext(context.Context) GetGroupUsersObjArrayOutput
}

type GetGroupUsersObjArray []GetGroupUsersObjInput

func (GetGroupUsersObjArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupUsersObj)(nil)).Elem()
}

func (i GetGroupUsersObjArray) ToGetGroupUsersObjArrayOutput() GetGroupUsersObjArrayOutput {
	return i.ToGetGroupUsersObjArrayOutputWithContext(context.Background())
}

func (i GetGroupUsersObjArray) ToGetGroupUsersObjArrayOutputWithContext(ctx context.Context) GetGroupUsersObjArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupUsersObjArrayOutput)
}

type GetGroupUsersObjOutput struct{ *pulumi.OutputState }

func (GetGroupUsersObjOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUsersObj)(nil)).Elem()
}

func (o GetGroupUsersObjOutput) ToGetGroupUsersObjOutput() GetGroupUsersObjOutput {
	return o
}

func (o GetGroupUsersObjOutput) ToGetGroupUsersObjOutputWithContext(ctx context.Context) GetGroupUsersObjOutput {
	return o
}

func (o GetGroupUsersObjOutput) Attributes() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersObj) string { return v.Attributes }).(pulumi.StringOutput)
}

func (o GetGroupUsersObjOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersObj) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetGroupUsersObjOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupUsersObj) bool { return v.IsActive }).(pulumi.BoolOutput)
}

func (o GetGroupUsersObjOutput) LastLogin() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersObj) string { return v.LastLogin }).(pulumi.StringOutput)
}

func (o GetGroupUsersObjOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersObj) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetGroupUsersObjOutput) Pk() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupUsersObj) int { return v.Pk }).(pulumi.IntOutput)
}

func (o GetGroupUsersObjOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersObj) string { return v.Uid }).(pulumi.StringOutput)
}

func (o GetGroupUsersObjOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersObj) string { return v.Username }).(pulumi.StringOutput)
}

type GetGroupUsersObjArrayOutput struct{ *pulumi.OutputState }

func (GetGroupUsersObjArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupUsersObj)(nil)).Elem()
}

func (o GetGroupUsersObjArrayOutput) ToGetGroupUsersObjArrayOutput() GetGroupUsersObjArrayOutput {
	return o
}

func (o GetGroupUsersObjArrayOutput) ToGetGroupUsersObjArrayOutputWithContext(ctx context.Context) GetGroupUsersObjArrayOutput {
	return o
}

func (o GetGroupUsersObjArrayOutput) Index(i pulumi.IntInput) GetGroupUsersObjOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGroupUsersObj {
		return vs[0].([]GetGroupUsersObj)[vs[1].(int)]
	}).(GetGroupUsersObjOutput)
}

type GetGroupsGroup struct {
	Attributes  string                   `pulumi:"attributes"`
	IsSuperuser bool                     `pulumi:"isSuperuser"`
	Name        string                   `pulumi:"name"`
	NumPk       int                      `pulumi:"numPk"`
	Parent      string                   `pulumi:"parent"`
	ParentName  string                   `pulumi:"parentName"`
	Pk          string                   `pulumi:"pk"`
	Users       []int                    `pulumi:"users"`
	UsersObjs   []GetGroupsGroupUsersObj `pulumi:"usersObjs"`
}

// GetGroupsGroupInput is an input type that accepts GetGroupsGroupArgs and GetGroupsGroupOutput values.
// You can construct a concrete instance of `GetGroupsGroupInput` via:
//
//	GetGroupsGroupArgs{...}
type GetGroupsGroupInput interface {
	pulumi.Input

	ToGetGroupsGroupOutput() GetGroupsGroupOutput
	ToGetGroupsGroupOutputWithContext(context.Context) GetGroupsGroupOutput
}

type GetGroupsGroupArgs struct {
	Attributes  pulumi.StringInput               `pulumi:"attributes"`
	IsSuperuser pulumi.BoolInput                 `pulumi:"isSuperuser"`
	Name        pulumi.StringInput               `pulumi:"name"`
	NumPk       pulumi.IntInput                  `pulumi:"numPk"`
	Parent      pulumi.StringInput               `pulumi:"parent"`
	ParentName  pulumi.StringInput               `pulumi:"parentName"`
	Pk          pulumi.StringInput               `pulumi:"pk"`
	Users       pulumi.IntArrayInput             `pulumi:"users"`
	UsersObjs   GetGroupsGroupUsersObjArrayInput `pulumi:"usersObjs"`
}

func (GetGroupsGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsGroup)(nil)).Elem()
}

func (i GetGroupsGroupArgs) ToGetGroupsGroupOutput() GetGroupsGroupOutput {
	return i.ToGetGroupsGroupOutputWithContext(context.Background())
}

func (i GetGroupsGroupArgs) ToGetGroupsGroupOutputWithContext(ctx context.Context) GetGroupsGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupsGroupOutput)
}

// GetGroupsGroupArrayInput is an input type that accepts GetGroupsGroupArray and GetGroupsGroupArrayOutput values.
// You can construct a concrete instance of `GetGroupsGroupArrayInput` via:
//
//	GetGroupsGroupArray{ GetGroupsGroupArgs{...} }
type GetGroupsGroupArrayInput interface {
	pulumi.Input

	ToGetGroupsGroupArrayOutput() GetGroupsGroupArrayOutput
	ToGetGroupsGroupArrayOutputWithContext(context.Context) GetGroupsGroupArrayOutput
}

type GetGroupsGroupArray []GetGroupsGroupInput

func (GetGroupsGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupsGroup)(nil)).Elem()
}

func (i GetGroupsGroupArray) ToGetGroupsGroupArrayOutput() GetGroupsGroupArrayOutput {
	return i.ToGetGroupsGroupArrayOutputWithContext(context.Background())
}

func (i GetGroupsGroupArray) ToGetGroupsGroupArrayOutputWithContext(ctx context.Context) GetGroupsGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupsGroupArrayOutput)
}

type GetGroupsGroupOutput struct{ *pulumi.OutputState }

func (GetGroupsGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsGroup)(nil)).Elem()
}

func (o GetGroupsGroupOutput) ToGetGroupsGroupOutput() GetGroupsGroupOutput {
	return o
}

func (o GetGroupsGroupOutput) ToGetGroupsGroupOutputWithContext(ctx context.Context) GetGroupsGroupOutput {
	return o
}

func (o GetGroupsGroupOutput) Attributes() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroup) string { return v.Attributes }).(pulumi.StringOutput)
}

func (o GetGroupsGroupOutput) IsSuperuser() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupsGroup) bool { return v.IsSuperuser }).(pulumi.BoolOutput)
}

func (o GetGroupsGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroup) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetGroupsGroupOutput) NumPk() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupsGroup) int { return v.NumPk }).(pulumi.IntOutput)
}

func (o GetGroupsGroupOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroup) string { return v.Parent }).(pulumi.StringOutput)
}

func (o GetGroupsGroupOutput) ParentName() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroup) string { return v.ParentName }).(pulumi.StringOutput)
}

func (o GetGroupsGroupOutput) Pk() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroup) string { return v.Pk }).(pulumi.StringOutput)
}

func (o GetGroupsGroupOutput) Users() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetGroupsGroup) []int { return v.Users }).(pulumi.IntArrayOutput)
}

func (o GetGroupsGroupOutput) UsersObjs() GetGroupsGroupUsersObjArrayOutput {
	return o.ApplyT(func(v GetGroupsGroup) []GetGroupsGroupUsersObj { return v.UsersObjs }).(GetGroupsGroupUsersObjArrayOutput)
}

type GetGroupsGroupArrayOutput struct{ *pulumi.OutputState }

func (GetGroupsGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupsGroup)(nil)).Elem()
}

func (o GetGroupsGroupArrayOutput) ToGetGroupsGroupArrayOutput() GetGroupsGroupArrayOutput {
	return o
}

func (o GetGroupsGroupArrayOutput) ToGetGroupsGroupArrayOutputWithContext(ctx context.Context) GetGroupsGroupArrayOutput {
	return o
}

func (o GetGroupsGroupArrayOutput) Index(i pulumi.IntInput) GetGroupsGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGroupsGroup {
		return vs[0].([]GetGroupsGroup)[vs[1].(int)]
	}).(GetGroupsGroupOutput)
}

type GetGroupsGroupUsersObj struct {
	Attributes string `pulumi:"attributes"`
	Email      string `pulumi:"email"`
	IsActive   bool   `pulumi:"isActive"`
	LastLogin  string `pulumi:"lastLogin"`
	Name       string `pulumi:"name"`
	Pk         int    `pulumi:"pk"`
	Uid        string `pulumi:"uid"`
	Username   string `pulumi:"username"`
}

// GetGroupsGroupUsersObjInput is an input type that accepts GetGroupsGroupUsersObjArgs and GetGroupsGroupUsersObjOutput values.
// You can construct a concrete instance of `GetGroupsGroupUsersObjInput` via:
//
//	GetGroupsGroupUsersObjArgs{...}
type GetGroupsGroupUsersObjInput interface {
	pulumi.Input

	ToGetGroupsGroupUsersObjOutput() GetGroupsGroupUsersObjOutput
	ToGetGroupsGroupUsersObjOutputWithContext(context.Context) GetGroupsGroupUsersObjOutput
}

type GetGroupsGroupUsersObjArgs struct {
	Attributes pulumi.StringInput `pulumi:"attributes"`
	Email      pulumi.StringInput `pulumi:"email"`
	IsActive   pulumi.BoolInput   `pulumi:"isActive"`
	LastLogin  pulumi.StringInput `pulumi:"lastLogin"`
	Name       pulumi.StringInput `pulumi:"name"`
	Pk         pulumi.IntInput    `pulumi:"pk"`
	Uid        pulumi.StringInput `pulumi:"uid"`
	Username   pulumi.StringInput `pulumi:"username"`
}

func (GetGroupsGroupUsersObjArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsGroupUsersObj)(nil)).Elem()
}

func (i GetGroupsGroupUsersObjArgs) ToGetGroupsGroupUsersObjOutput() GetGroupsGroupUsersObjOutput {
	return i.ToGetGroupsGroupUsersObjOutputWithContext(context.Background())
}

func (i GetGroupsGroupUsersObjArgs) ToGetGroupsGroupUsersObjOutputWithContext(ctx context.Context) GetGroupsGroupUsersObjOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupsGroupUsersObjOutput)
}

// GetGroupsGroupUsersObjArrayInput is an input type that accepts GetGroupsGroupUsersObjArray and GetGroupsGroupUsersObjArrayOutput values.
// You can construct a concrete instance of `GetGroupsGroupUsersObjArrayInput` via:
//
//	GetGroupsGroupUsersObjArray{ GetGroupsGroupUsersObjArgs{...} }
type GetGroupsGroupUsersObjArrayInput interface {
	pulumi.Input

	ToGetGroupsGroupUsersObjArrayOutput() GetGroupsGroupUsersObjArrayOutput
	ToGetGroupsGroupUsersObjArrayOutputWithContext(context.Context) GetGroupsGroupUsersObjArrayOutput
}

type GetGroupsGroupUsersObjArray []GetGroupsGroupUsersObjInput

func (GetGroupsGroupUsersObjArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupsGroupUsersObj)(nil)).Elem()
}

func (i GetGroupsGroupUsersObjArray) ToGetGroupsGroupUsersObjArrayOutput() GetGroupsGroupUsersObjArrayOutput {
	return i.ToGetGroupsGroupUsersObjArrayOutputWithContext(context.Background())
}

func (i GetGroupsGroupUsersObjArray) ToGetGroupsGroupUsersObjArrayOutputWithContext(ctx context.Context) GetGroupsGroupUsersObjArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupsGroupUsersObjArrayOutput)
}

type GetGroupsGroupUsersObjOutput struct{ *pulumi.OutputState }

func (GetGroupsGroupUsersObjOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsGroupUsersObj)(nil)).Elem()
}

func (o GetGroupsGroupUsersObjOutput) ToGetGroupsGroupUsersObjOutput() GetGroupsGroupUsersObjOutput {
	return o
}

func (o GetGroupsGroupUsersObjOutput) ToGetGroupsGroupUsersObjOutputWithContext(ctx context.Context) GetGroupsGroupUsersObjOutput {
	return o
}

func (o GetGroupsGroupUsersObjOutput) Attributes() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) string { return v.Attributes }).(pulumi.StringOutput)
}

func (o GetGroupsGroupUsersObjOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetGroupsGroupUsersObjOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) bool { return v.IsActive }).(pulumi.BoolOutput)
}

func (o GetGroupsGroupUsersObjOutput) LastLogin() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) string { return v.LastLogin }).(pulumi.StringOutput)
}

func (o GetGroupsGroupUsersObjOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetGroupsGroupUsersObjOutput) Pk() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) int { return v.Pk }).(pulumi.IntOutput)
}

func (o GetGroupsGroupUsersObjOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) string { return v.Uid }).(pulumi.StringOutput)
}

func (o GetGroupsGroupUsersObjOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsGroupUsersObj) string { return v.Username }).(pulumi.StringOutput)
}

type GetGroupsGroupUsersObjArrayOutput struct{ *pulumi.OutputState }

func (GetGroupsGroupUsersObjArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupsGroupUsersObj)(nil)).Elem()
}

func (o GetGroupsGroupUsersObjArrayOutput) ToGetGroupsGroupUsersObjArrayOutput() GetGroupsGroupUsersObjArrayOutput {
	return o
}

func (o GetGroupsGroupUsersObjArrayOutput) ToGetGroupsGroupUsersObjArrayOutputWithContext(ctx context.Context) GetGroupsGroupUsersObjArrayOutput {
	return o
}

func (o GetGroupsGroupUsersObjArrayOutput) Index(i pulumi.IntInput) GetGroupsGroupUsersObjOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGroupsGroupUsersObj {
		return vs[0].([]GetGroupsGroupUsersObj)[vs[1].(int)]
	}).(GetGroupsGroupUsersObjOutput)
}

type GetUsersUser struct {
	Attributes  string   `pulumi:"attributes"`
	Avatar      string   `pulumi:"avatar"`
	Email       string   `pulumi:"email"`
	Groups      []string `pulumi:"groups"`
	IsActive    bool     `pulumi:"isActive"`
	IsSuperuser bool     `pulumi:"isSuperuser"`
	LastLogin   string   `pulumi:"lastLogin"`
	Name        string   `pulumi:"name"`
	Path        string   `pulumi:"path"`
	Pk          int      `pulumi:"pk"`
	Type        string   `pulumi:"type"`
	Uid         string   `pulumi:"uid"`
	Username    string   `pulumi:"username"`
	Uuid        string   `pulumi:"uuid"`
}

// GetUsersUserInput is an input type that accepts GetUsersUserArgs and GetUsersUserOutput values.
// You can construct a concrete instance of `GetUsersUserInput` via:
//
//	GetUsersUserArgs{...}
type GetUsersUserInput interface {
	pulumi.Input

	ToGetUsersUserOutput() GetUsersUserOutput
	ToGetUsersUserOutputWithContext(context.Context) GetUsersUserOutput
}

type GetUsersUserArgs struct {
	Attributes  pulumi.StringInput      `pulumi:"attributes"`
	Avatar      pulumi.StringInput      `pulumi:"avatar"`
	Email       pulumi.StringInput      `pulumi:"email"`
	Groups      pulumi.StringArrayInput `pulumi:"groups"`
	IsActive    pulumi.BoolInput        `pulumi:"isActive"`
	IsSuperuser pulumi.BoolInput        `pulumi:"isSuperuser"`
	LastLogin   pulumi.StringInput      `pulumi:"lastLogin"`
	Name        pulumi.StringInput      `pulumi:"name"`
	Path        pulumi.StringInput      `pulumi:"path"`
	Pk          pulumi.IntInput         `pulumi:"pk"`
	Type        pulumi.StringInput      `pulumi:"type"`
	Uid         pulumi.StringInput      `pulumi:"uid"`
	Username    pulumi.StringInput      `pulumi:"username"`
	Uuid        pulumi.StringInput      `pulumi:"uuid"`
}

func (GetUsersUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersUser)(nil)).Elem()
}

func (i GetUsersUserArgs) ToGetUsersUserOutput() GetUsersUserOutput {
	return i.ToGetUsersUserOutputWithContext(context.Background())
}

func (i GetUsersUserArgs) ToGetUsersUserOutputWithContext(ctx context.Context) GetUsersUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUsersUserOutput)
}

// GetUsersUserArrayInput is an input type that accepts GetUsersUserArray and GetUsersUserArrayOutput values.
// You can construct a concrete instance of `GetUsersUserArrayInput` via:
//
//	GetUsersUserArray{ GetUsersUserArgs{...} }
type GetUsersUserArrayInput interface {
	pulumi.Input

	ToGetUsersUserArrayOutput() GetUsersUserArrayOutput
	ToGetUsersUserArrayOutputWithContext(context.Context) GetUsersUserArrayOutput
}

type GetUsersUserArray []GetUsersUserInput

func (GetUsersUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUsersUser)(nil)).Elem()
}

func (i GetUsersUserArray) ToGetUsersUserArrayOutput() GetUsersUserArrayOutput {
	return i.ToGetUsersUserArrayOutputWithContext(context.Background())
}

func (i GetUsersUserArray) ToGetUsersUserArrayOutputWithContext(ctx context.Context) GetUsersUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUsersUserArrayOutput)
}

type GetUsersUserOutput struct{ *pulumi.OutputState }

func (GetUsersUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersUser)(nil)).Elem()
}

func (o GetUsersUserOutput) ToGetUsersUserOutput() GetUsersUserOutput {
	return o
}

func (o GetUsersUserOutput) ToGetUsersUserOutputWithContext(ctx context.Context) GetUsersUserOutput {
	return o
}

func (o GetUsersUserOutput) Attributes() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Attributes }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Avatar() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Avatar }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersUser) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o GetUsersUserOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v GetUsersUser) bool { return v.IsActive }).(pulumi.BoolOutput)
}

func (o GetUsersUserOutput) IsSuperuser() pulumi.BoolOutput {
	return o.ApplyT(func(v GetUsersUser) bool { return v.IsSuperuser }).(pulumi.BoolOutput)
}

func (o GetUsersUserOutput) LastLogin() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.LastLogin }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Pk() pulumi.IntOutput {
	return o.ApplyT(func(v GetUsersUser) int { return v.Pk }).(pulumi.IntOutput)
}

func (o GetUsersUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Uid }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Username }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Uuid }).(pulumi.StringOutput)
}

type GetUsersUserArrayOutput struct{ *pulumi.OutputState }

func (GetUsersUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUsersUser)(nil)).Elem()
}

func (o GetUsersUserArrayOutput) ToGetUsersUserArrayOutput() GetUsersUserArrayOutput {
	return o
}

func (o GetUsersUserArrayOutput) ToGetUsersUserArrayOutputWithContext(ctx context.Context) GetUsersUserArrayOutput {
	return o
}

func (o GetUsersUserArrayOutput) Index(i pulumi.IntInput) GetUsersUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetUsersUser {
		return vs[0].([]GetUsersUser)[vs[1].(int)]
	}).(GetUsersUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupUsersObjInput)(nil)).Elem(), GetGroupUsersObjArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupUsersObjArrayInput)(nil)).Elem(), GetGroupUsersObjArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupsGroupInput)(nil)).Elem(), GetGroupsGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupsGroupArrayInput)(nil)).Elem(), GetGroupsGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupsGroupUsersObjInput)(nil)).Elem(), GetGroupsGroupUsersObjArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupsGroupUsersObjArrayInput)(nil)).Elem(), GetGroupsGroupUsersObjArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUsersUserInput)(nil)).Elem(), GetUsersUserArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUsersUserArrayInput)(nil)).Elem(), GetUsersUserArray{})
	pulumi.RegisterOutputType(GetGroupUsersObjOutput{})
	pulumi.RegisterOutputType(GetGroupUsersObjArrayOutput{})
	pulumi.RegisterOutputType(GetGroupsGroupOutput{})
	pulumi.RegisterOutputType(GetGroupsGroupArrayOutput{})
	pulumi.RegisterOutputType(GetGroupsGroupUsersObjOutput{})
	pulumi.RegisterOutputType(GetGroupsGroupUsersObjArrayOutput{})
	pulumi.RegisterOutputType(GetUsersUserOutput{})
	pulumi.RegisterOutputType(GetUsersUserArrayOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get users list
//
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
//			_, err := authentik.GetUsers(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = authentik.GetUsers(ctx, &authentik.GetUsersArgs{
//				IsSuperuser: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("authentik:index/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	Attributes     *string  `pulumi:"attributes"`
	Email          *string  `pulumi:"email"`
	GroupsByNames  []string `pulumi:"groupsByNames"`
	GroupsByPks    []string `pulumi:"groupsByPks"`
	IsActive       *bool    `pulumi:"isActive"`
	IsSuperuser    *bool    `pulumi:"isSuperuser"`
	Name           *string  `pulumi:"name"`
	Ordering       *string  `pulumi:"ordering"`
	Path           *string  `pulumi:"path"`
	PathStartswith *string  `pulumi:"pathStartswith"`
	Search         *string  `pulumi:"search"`
	Username       *string  `pulumi:"username"`
	Uuid           *string  `pulumi:"uuid"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	Attributes    *string  `pulumi:"attributes"`
	Email         *string  `pulumi:"email"`
	GroupsByNames []string `pulumi:"groupsByNames"`
	GroupsByPks   []string `pulumi:"groupsByPks"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	IsActive       *bool   `pulumi:"isActive"`
	IsSuperuser    *bool   `pulumi:"isSuperuser"`
	Name           *string `pulumi:"name"`
	Ordering       *string `pulumi:"ordering"`
	Path           *string `pulumi:"path"`
	PathStartswith *string `pulumi:"pathStartswith"`
	Search         *string `pulumi:"search"`
	Username       *string `pulumi:"username"`
	// Generated.
	Users []GetUsersUser `pulumi:"users"`
	Uuid  *string        `pulumi:"uuid"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			var s GetUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	Attributes     pulumi.StringPtrInput   `pulumi:"attributes"`
	Email          pulumi.StringPtrInput   `pulumi:"email"`
	GroupsByNames  pulumi.StringArrayInput `pulumi:"groupsByNames"`
	GroupsByPks    pulumi.StringArrayInput `pulumi:"groupsByPks"`
	IsActive       pulumi.BoolPtrInput     `pulumi:"isActive"`
	IsSuperuser    pulumi.BoolPtrInput     `pulumi:"isSuperuser"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	Ordering       pulumi.StringPtrInput   `pulumi:"ordering"`
	Path           pulumi.StringPtrInput   `pulumi:"path"`
	PathStartswith pulumi.StringPtrInput   `pulumi:"pathStartswith"`
	Search         pulumi.StringPtrInput   `pulumi:"search"`
	Username       pulumi.StringPtrInput   `pulumi:"username"`
	Uuid           pulumi.StringPtrInput   `pulumi:"uuid"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetUsersResult] {
	return pulumix.Output[GetUsersResult]{
		OutputState: o.OutputState,
	}
}

func (o GetUsersResultOutput) Attributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Attributes }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) GroupsByNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.GroupsByNames }).(pulumi.StringArrayOutput)
}

func (o GetUsersResultOutput) GroupsByPks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.GroupsByPks }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o GetUsersResultOutput) IsSuperuser() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.IsSuperuser }).(pulumi.BoolPtrOutput)
}

func (o GetUsersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Ordering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Ordering }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) PathStartswith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.PathStartswith }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// Generated.
func (o GetUsersResultOutput) Users() GetUsersUserArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUser { return v.Users }).(GetUsersUserArrayOutput)
}

func (o GetUsersResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}

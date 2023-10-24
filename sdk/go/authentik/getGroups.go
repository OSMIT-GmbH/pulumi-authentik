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

// Get groups list
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
//			_, err := authentik.GetGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = authentik.GetGroups(ctx, &authentik.GetGroupsArgs{
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
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupsResult
	err := ctx.Invoke("authentik:index/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	Attributes         *string  `pulumi:"attributes"`
	IsSuperuser        *bool    `pulumi:"isSuperuser"`
	MembersByPks       []int    `pulumi:"membersByPks"`
	MembersByUsernames []string `pulumi:"membersByUsernames"`
	Name               *string  `pulumi:"name"`
	Ordering           *string  `pulumi:"ordering"`
	Search             *string  `pulumi:"search"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	Attributes *string `pulumi:"attributes"`
	// Generated.
	Groups []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	IsSuperuser        *bool    `pulumi:"isSuperuser"`
	MembersByPks       []int    `pulumi:"membersByPks"`
	MembersByUsernames []string `pulumi:"membersByUsernames"`
	Name               *string  `pulumi:"name"`
	Ordering           *string  `pulumi:"ordering"`
	Search             *string  `pulumi:"search"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			var s GetGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	Attributes         pulumi.StringPtrInput   `pulumi:"attributes"`
	IsSuperuser        pulumi.BoolPtrInput     `pulumi:"isSuperuser"`
	MembersByPks       pulumi.IntArrayInput    `pulumi:"membersByPks"`
	MembersByUsernames pulumi.StringArrayInput `pulumi:"membersByUsernames"`
	Name               pulumi.StringPtrInput   `pulumi:"name"`
	Ordering           pulumi.StringPtrInput   `pulumi:"ordering"`
	Search             pulumi.StringPtrInput   `pulumi:"search"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetGroupsResult] {
	return pulumix.Output[GetGroupsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetGroupsResultOutput) Attributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Attributes }).(pulumi.StringPtrOutput)
}

// Generated.
func (o GetGroupsResultOutput) Groups() GetGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroup { return v.Groups }).(GetGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) IsSuperuser() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *bool { return v.IsSuperuser }).(pulumi.BoolPtrOutput)
}

func (o GetGroupsResultOutput) MembersByPks() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []int { return v.MembersByPks }).(pulumi.IntArrayOutput)
}

func (o GetGroupsResultOutput) MembersByUsernames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.MembersByUsernames }).(pulumi.StringArrayOutput)
}

func (o GetGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) Ordering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Ordering }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}

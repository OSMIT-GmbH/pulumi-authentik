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

// Get users by pk or username
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
//			_, err := authentik.LookupUser(ctx, &authentik.LookupUserArgs{
//				Username: pulumi.StringRef("akadmin"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("authentik:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// Generated.
	Pk *int `pulumi:"pk"`
	// Generated.
	Username *string `pulumi:"username"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// Generated.
	Attributes string `pulumi:"attributes"`
	// Generated.
	Avatar string `pulumi:"avatar"`
	// Generated.
	Email string `pulumi:"email"`
	// Generated.
	Groups []string `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Generated.
	IsActive bool `pulumi:"isActive"`
	// Generated.
	IsSuperuser bool `pulumi:"isSuperuser"`
	// Generated.
	LastLogin string `pulumi:"lastLogin"`
	// Generated.
	Name string `pulumi:"name"`
	// Generated.
	Path string `pulumi:"path"`
	// Generated.
	Pk int `pulumi:"pk"`
	// Generated.
	Type string `pulumi:"type"`
	// Generated.
	Uid string `pulumi:"uid"`
	// Generated.
	Username string `pulumi:"username"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// Generated.
	Pk pulumi.IntPtrInput `pulumi:"pk"`
	// Generated.
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserResult] {
	return pulumix.Output[LookupUserResult]{
		OutputState: o.OutputState,
	}
}

// Generated.
func (o LookupUserResultOutput) Attributes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Attributes }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Avatar() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Avatar }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.IsActive }).(pulumi.BoolOutput)
}

// Generated.
func (o LookupUserResultOutput) IsSuperuser() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.IsSuperuser }).(pulumi.BoolOutput)
}

// Generated.
func (o LookupUserResultOutput) LastLogin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.LastLogin }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Path }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Pk() pulumi.IntOutput {
	return o.ApplyT(func(v LookupUserResult) int { return v.Pk }).(pulumi.IntOutput)
}

// Generated.
func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Generated.
func (o LookupUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}

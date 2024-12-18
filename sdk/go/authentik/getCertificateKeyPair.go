// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get certificate-key pairs by name
//
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
//			// To get the the ID and other info about a certificate
//			_, err := authentik.LookupCertificateKeyPair(ctx, &authentik.LookupCertificateKeyPairArgs{
//				Name: "authentik Self-signed Certificate",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCertificateKeyPair(ctx *pulumi.Context, args *LookupCertificateKeyPairArgs, opts ...pulumi.InvokeOption) (*LookupCertificateKeyPairResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateKeyPairResult
	err := ctx.Invoke("authentik:index/getCertificateKeyPair:getCertificateKeyPair", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificateKeyPair.
type LookupCertificateKeyPairArgs struct {
	// If set to true, certificate data will be fetched. Defaults to `true`.
	FetchCertificate *bool `pulumi:"fetchCertificate"`
	// If set to true, private key data will be fetched. Defaults to `true`.
	FetchKey *bool `pulumi:"fetchKey"`
	// Generated.
	KeyData *string `pulumi:"keyData"`
	Name    string  `pulumi:"name"`
}

// A collection of values returned by getCertificateKeyPair.
type LookupCertificateKeyPairResult struct {
	// Generated.
	CertificateData string `pulumi:"certificateData"`
	// Generated.
	Expiry string `pulumi:"expiry"`
	// If set to true, certificate data will be fetched. Defaults to `true`.
	FetchCertificate *bool `pulumi:"fetchCertificate"`
	// If set to true, private key data will be fetched. Defaults to `true`.
	FetchKey *bool `pulumi:"fetchKey"`
	// SHA1-hashed certificate fingerprint Generated.
	Fingerprint1 string `pulumi:"fingerprint1"`
	// SHA256-hashed certificate fingerprint Generated.
	Fingerprint256 string `pulumi:"fingerprint256"`
	// Generated.
	Id string `pulumi:"id"`
	// Generated.
	KeyData string `pulumi:"keyData"`
	Name    string `pulumi:"name"`
	// Generated.
	Subject string `pulumi:"subject"`
}

func LookupCertificateKeyPairOutput(ctx *pulumi.Context, args LookupCertificateKeyPairOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateKeyPairResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateKeyPairResultOutput, error) {
			args := v.(LookupCertificateKeyPairArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupCertificateKeyPairResult
			secret, err := ctx.InvokePackageRaw("authentik:index/getCertificateKeyPair:getCertificateKeyPair", args, &rv, "", opts...)
			if err != nil {
				return LookupCertificateKeyPairResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupCertificateKeyPairResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupCertificateKeyPairResultOutput), nil
			}
			return output, nil
		}).(LookupCertificateKeyPairResultOutput)
}

// A collection of arguments for invoking getCertificateKeyPair.
type LookupCertificateKeyPairOutputArgs struct {
	// If set to true, certificate data will be fetched. Defaults to `true`.
	FetchCertificate pulumi.BoolPtrInput `pulumi:"fetchCertificate"`
	// If set to true, private key data will be fetched. Defaults to `true`.
	FetchKey pulumi.BoolPtrInput `pulumi:"fetchKey"`
	// Generated.
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Name    pulumi.StringInput    `pulumi:"name"`
}

func (LookupCertificateKeyPairOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateKeyPairArgs)(nil)).Elem()
}

// A collection of values returned by getCertificateKeyPair.
type LookupCertificateKeyPairResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateKeyPairResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateKeyPairResult)(nil)).Elem()
}

func (o LookupCertificateKeyPairResultOutput) ToLookupCertificateKeyPairResultOutput() LookupCertificateKeyPairResultOutput {
	return o
}

func (o LookupCertificateKeyPairResultOutput) ToLookupCertificateKeyPairResultOutputWithContext(ctx context.Context) LookupCertificateKeyPairResultOutput {
	return o
}

// Generated.
func (o LookupCertificateKeyPairResultOutput) CertificateData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.CertificateData }).(pulumi.StringOutput)
}

// Generated.
func (o LookupCertificateKeyPairResultOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.Expiry }).(pulumi.StringOutput)
}

// If set to true, certificate data will be fetched. Defaults to `true`.
func (o LookupCertificateKeyPairResultOutput) FetchCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) *bool { return v.FetchCertificate }).(pulumi.BoolPtrOutput)
}

// If set to true, private key data will be fetched. Defaults to `true`.
func (o LookupCertificateKeyPairResultOutput) FetchKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) *bool { return v.FetchKey }).(pulumi.BoolPtrOutput)
}

// SHA1-hashed certificate fingerprint Generated.
func (o LookupCertificateKeyPairResultOutput) Fingerprint1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.Fingerprint1 }).(pulumi.StringOutput)
}

// SHA256-hashed certificate fingerprint Generated.
func (o LookupCertificateKeyPairResultOutput) Fingerprint256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.Fingerprint256 }).(pulumi.StringOutput)
}

// Generated.
func (o LookupCertificateKeyPairResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.Id }).(pulumi.StringOutput)
}

// Generated.
func (o LookupCertificateKeyPairResultOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.KeyData }).(pulumi.StringOutput)
}

func (o LookupCertificateKeyPairResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.Name }).(pulumi.StringOutput)
}

// Generated.
func (o LookupCertificateKeyPairResultOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateKeyPairResult) string { return v.Subject }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateKeyPairResultOutput{})
}

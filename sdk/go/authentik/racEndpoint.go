// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RacEndpoint struct {
	pulumi.CustomResourceState

	Host               pulumi.StringOutput      `pulumi:"host"`
	MaximumConnections pulumi.IntPtrOutput      `pulumi:"maximumConnections"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	PropertyMappings   pulumi.StringArrayOutput `pulumi:"propertyMappings"`
	// Allowed values: - `rdp` - `vnc` - `ssh`
	Protocol         pulumi.StringOutput `pulumi:"protocol"`
	ProtocolProvider pulumi.IntOutput    `pulumi:"protocolProvider"`
	// JSON format expected. Use jsonencode() to pass objects.
	Settings pulumi.StringPtrOutput `pulumi:"settings"`
}

// NewRacEndpoint registers a new resource with the given unique name, arguments, and options.
func NewRacEndpoint(ctx *pulumi.Context,
	name string, args *RacEndpointArgs, opts ...pulumi.ResourceOption) (*RacEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ProtocolProvider == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolProvider'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RacEndpoint
	err := ctx.RegisterResource("authentik:index/racEndpoint:RacEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRacEndpoint gets an existing RacEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRacEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RacEndpointState, opts ...pulumi.ResourceOption) (*RacEndpoint, error) {
	var resource RacEndpoint
	err := ctx.ReadResource("authentik:index/racEndpoint:RacEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RacEndpoint resources.
type racEndpointState struct {
	Host               *string  `pulumi:"host"`
	MaximumConnections *int     `pulumi:"maximumConnections"`
	Name               *string  `pulumi:"name"`
	PropertyMappings   []string `pulumi:"propertyMappings"`
	// Allowed values: - `rdp` - `vnc` - `ssh`
	Protocol         *string `pulumi:"protocol"`
	ProtocolProvider *int    `pulumi:"protocolProvider"`
	// JSON format expected. Use jsonencode() to pass objects.
	Settings *string `pulumi:"settings"`
}

type RacEndpointState struct {
	Host               pulumi.StringPtrInput
	MaximumConnections pulumi.IntPtrInput
	Name               pulumi.StringPtrInput
	PropertyMappings   pulumi.StringArrayInput
	// Allowed values: - `rdp` - `vnc` - `ssh`
	Protocol         pulumi.StringPtrInput
	ProtocolProvider pulumi.IntPtrInput
	// JSON format expected. Use jsonencode() to pass objects.
	Settings pulumi.StringPtrInput
}

func (RacEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*racEndpointState)(nil)).Elem()
}

type racEndpointArgs struct {
	Host               string   `pulumi:"host"`
	MaximumConnections *int     `pulumi:"maximumConnections"`
	Name               *string  `pulumi:"name"`
	PropertyMappings   []string `pulumi:"propertyMappings"`
	// Allowed values: - `rdp` - `vnc` - `ssh`
	Protocol         string `pulumi:"protocol"`
	ProtocolProvider int    `pulumi:"protocolProvider"`
	// JSON format expected. Use jsonencode() to pass objects.
	Settings *string `pulumi:"settings"`
}

// The set of arguments for constructing a RacEndpoint resource.
type RacEndpointArgs struct {
	Host               pulumi.StringInput
	MaximumConnections pulumi.IntPtrInput
	Name               pulumi.StringPtrInput
	PropertyMappings   pulumi.StringArrayInput
	// Allowed values: - `rdp` - `vnc` - `ssh`
	Protocol         pulumi.StringInput
	ProtocolProvider pulumi.IntInput
	// JSON format expected. Use jsonencode() to pass objects.
	Settings pulumi.StringPtrInput
}

func (RacEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*racEndpointArgs)(nil)).Elem()
}

type RacEndpointInput interface {
	pulumi.Input

	ToRacEndpointOutput() RacEndpointOutput
	ToRacEndpointOutputWithContext(ctx context.Context) RacEndpointOutput
}

func (*RacEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**RacEndpoint)(nil)).Elem()
}

func (i *RacEndpoint) ToRacEndpointOutput() RacEndpointOutput {
	return i.ToRacEndpointOutputWithContext(context.Background())
}

func (i *RacEndpoint) ToRacEndpointOutputWithContext(ctx context.Context) RacEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RacEndpointOutput)
}

// RacEndpointArrayInput is an input type that accepts RacEndpointArray and RacEndpointArrayOutput values.
// You can construct a concrete instance of `RacEndpointArrayInput` via:
//
//	RacEndpointArray{ RacEndpointArgs{...} }
type RacEndpointArrayInput interface {
	pulumi.Input

	ToRacEndpointArrayOutput() RacEndpointArrayOutput
	ToRacEndpointArrayOutputWithContext(context.Context) RacEndpointArrayOutput
}

type RacEndpointArray []RacEndpointInput

func (RacEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RacEndpoint)(nil)).Elem()
}

func (i RacEndpointArray) ToRacEndpointArrayOutput() RacEndpointArrayOutput {
	return i.ToRacEndpointArrayOutputWithContext(context.Background())
}

func (i RacEndpointArray) ToRacEndpointArrayOutputWithContext(ctx context.Context) RacEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RacEndpointArrayOutput)
}

// RacEndpointMapInput is an input type that accepts RacEndpointMap and RacEndpointMapOutput values.
// You can construct a concrete instance of `RacEndpointMapInput` via:
//
//	RacEndpointMap{ "key": RacEndpointArgs{...} }
type RacEndpointMapInput interface {
	pulumi.Input

	ToRacEndpointMapOutput() RacEndpointMapOutput
	ToRacEndpointMapOutputWithContext(context.Context) RacEndpointMapOutput
}

type RacEndpointMap map[string]RacEndpointInput

func (RacEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RacEndpoint)(nil)).Elem()
}

func (i RacEndpointMap) ToRacEndpointMapOutput() RacEndpointMapOutput {
	return i.ToRacEndpointMapOutputWithContext(context.Background())
}

func (i RacEndpointMap) ToRacEndpointMapOutputWithContext(ctx context.Context) RacEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RacEndpointMapOutput)
}

type RacEndpointOutput struct{ *pulumi.OutputState }

func (RacEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RacEndpoint)(nil)).Elem()
}

func (o RacEndpointOutput) ToRacEndpointOutput() RacEndpointOutput {
	return o
}

func (o RacEndpointOutput) ToRacEndpointOutputWithContext(ctx context.Context) RacEndpointOutput {
	return o
}

func (o RacEndpointOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

func (o RacEndpointOutput) MaximumConnections() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.IntPtrOutput { return v.MaximumConnections }).(pulumi.IntPtrOutput)
}

func (o RacEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RacEndpointOutput) PropertyMappings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.StringArrayOutput { return v.PropertyMappings }).(pulumi.StringArrayOutput)
}

// Allowed values: - `rdp` - `vnc` - `ssh`
func (o RacEndpointOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o RacEndpointOutput) ProtocolProvider() pulumi.IntOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.IntOutput { return v.ProtocolProvider }).(pulumi.IntOutput)
}

// JSON format expected. Use jsonencode() to pass objects.
func (o RacEndpointOutput) Settings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RacEndpoint) pulumi.StringPtrOutput { return v.Settings }).(pulumi.StringPtrOutput)
}

type RacEndpointArrayOutput struct{ *pulumi.OutputState }

func (RacEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RacEndpoint)(nil)).Elem()
}

func (o RacEndpointArrayOutput) ToRacEndpointArrayOutput() RacEndpointArrayOutput {
	return o
}

func (o RacEndpointArrayOutput) ToRacEndpointArrayOutputWithContext(ctx context.Context) RacEndpointArrayOutput {
	return o
}

func (o RacEndpointArrayOutput) Index(i pulumi.IntInput) RacEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RacEndpoint {
		return vs[0].([]*RacEndpoint)[vs[1].(int)]
	}).(RacEndpointOutput)
}

type RacEndpointMapOutput struct{ *pulumi.OutputState }

func (RacEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RacEndpoint)(nil)).Elem()
}

func (o RacEndpointMapOutput) ToRacEndpointMapOutput() RacEndpointMapOutput {
	return o
}

func (o RacEndpointMapOutput) ToRacEndpointMapOutputWithContext(ctx context.Context) RacEndpointMapOutput {
	return o
}

func (o RacEndpointMapOutput) MapIndex(k pulumi.StringInput) RacEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RacEndpoint {
		return vs[0].(map[string]*RacEndpoint)[vs[1].(string)]
	}).(RacEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RacEndpointInput)(nil)).Elem(), &RacEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*RacEndpointArrayInput)(nil)).Elem(), RacEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RacEndpointMapInput)(nil)).Elem(), RacEndpointMap{})
	pulumi.RegisterOutputType(RacEndpointOutput{})
	pulumi.RegisterOutputType(RacEndpointArrayOutput{})
	pulumi.RegisterOutputType(RacEndpointMapOutput{})
}

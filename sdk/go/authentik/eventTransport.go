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

type EventTransport struct {
	pulumi.CustomResourceState

	Mode           pulumi.StringOutput    `pulumi:"mode"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	SendOnce       pulumi.BoolPtrOutput   `pulumi:"sendOnce"`
	WebhookMapping pulumi.StringPtrOutput `pulumi:"webhookMapping"`
	WebhookUrl     pulumi.StringPtrOutput `pulumi:"webhookUrl"`
}

// NewEventTransport registers a new resource with the given unique name, arguments, and options.
func NewEventTransport(ctx *pulumi.Context,
	name string, args *EventTransportArgs, opts ...pulumi.ResourceOption) (*EventTransport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventTransport
	err := ctx.RegisterResource("authentik:index/eventTransport:EventTransport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventTransport gets an existing EventTransport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventTransport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventTransportState, opts ...pulumi.ResourceOption) (*EventTransport, error) {
	var resource EventTransport
	err := ctx.ReadResource("authentik:index/eventTransport:EventTransport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventTransport resources.
type eventTransportState struct {
	Mode           *string `pulumi:"mode"`
	Name           *string `pulumi:"name"`
	SendOnce       *bool   `pulumi:"sendOnce"`
	WebhookMapping *string `pulumi:"webhookMapping"`
	WebhookUrl     *string `pulumi:"webhookUrl"`
}

type EventTransportState struct {
	Mode           pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	SendOnce       pulumi.BoolPtrInput
	WebhookMapping pulumi.StringPtrInput
	WebhookUrl     pulumi.StringPtrInput
}

func (EventTransportState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTransportState)(nil)).Elem()
}

type eventTransportArgs struct {
	Mode           string  `pulumi:"mode"`
	Name           *string `pulumi:"name"`
	SendOnce       *bool   `pulumi:"sendOnce"`
	WebhookMapping *string `pulumi:"webhookMapping"`
	WebhookUrl     *string `pulumi:"webhookUrl"`
}

// The set of arguments for constructing a EventTransport resource.
type EventTransportArgs struct {
	Mode           pulumi.StringInput
	Name           pulumi.StringPtrInput
	SendOnce       pulumi.BoolPtrInput
	WebhookMapping pulumi.StringPtrInput
	WebhookUrl     pulumi.StringPtrInput
}

func (EventTransportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTransportArgs)(nil)).Elem()
}

type EventTransportInput interface {
	pulumi.Input

	ToEventTransportOutput() EventTransportOutput
	ToEventTransportOutputWithContext(ctx context.Context) EventTransportOutput
}

func (*EventTransport) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTransport)(nil)).Elem()
}

func (i *EventTransport) ToEventTransportOutput() EventTransportOutput {
	return i.ToEventTransportOutputWithContext(context.Background())
}

func (i *EventTransport) ToEventTransportOutputWithContext(ctx context.Context) EventTransportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTransportOutput)
}

// EventTransportArrayInput is an input type that accepts EventTransportArray and EventTransportArrayOutput values.
// You can construct a concrete instance of `EventTransportArrayInput` via:
//
//	EventTransportArray{ EventTransportArgs{...} }
type EventTransportArrayInput interface {
	pulumi.Input

	ToEventTransportArrayOutput() EventTransportArrayOutput
	ToEventTransportArrayOutputWithContext(context.Context) EventTransportArrayOutput
}

type EventTransportArray []EventTransportInput

func (EventTransportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventTransport)(nil)).Elem()
}

func (i EventTransportArray) ToEventTransportArrayOutput() EventTransportArrayOutput {
	return i.ToEventTransportArrayOutputWithContext(context.Background())
}

func (i EventTransportArray) ToEventTransportArrayOutputWithContext(ctx context.Context) EventTransportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTransportArrayOutput)
}

// EventTransportMapInput is an input type that accepts EventTransportMap and EventTransportMapOutput values.
// You can construct a concrete instance of `EventTransportMapInput` via:
//
//	EventTransportMap{ "key": EventTransportArgs{...} }
type EventTransportMapInput interface {
	pulumi.Input

	ToEventTransportMapOutput() EventTransportMapOutput
	ToEventTransportMapOutputWithContext(context.Context) EventTransportMapOutput
}

type EventTransportMap map[string]EventTransportInput

func (EventTransportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventTransport)(nil)).Elem()
}

func (i EventTransportMap) ToEventTransportMapOutput() EventTransportMapOutput {
	return i.ToEventTransportMapOutputWithContext(context.Background())
}

func (i EventTransportMap) ToEventTransportMapOutputWithContext(ctx context.Context) EventTransportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTransportMapOutput)
}

type EventTransportOutput struct{ *pulumi.OutputState }

func (EventTransportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTransport)(nil)).Elem()
}

func (o EventTransportOutput) ToEventTransportOutput() EventTransportOutput {
	return o
}

func (o EventTransportOutput) ToEventTransportOutputWithContext(ctx context.Context) EventTransportOutput {
	return o
}

func (o EventTransportOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTransport) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o EventTransportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTransport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventTransportOutput) SendOnce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventTransport) pulumi.BoolPtrOutput { return v.SendOnce }).(pulumi.BoolPtrOutput)
}

func (o EventTransportOutput) WebhookMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventTransport) pulumi.StringPtrOutput { return v.WebhookMapping }).(pulumi.StringPtrOutput)
}

func (o EventTransportOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventTransport) pulumi.StringPtrOutput { return v.WebhookUrl }).(pulumi.StringPtrOutput)
}

type EventTransportArrayOutput struct{ *pulumi.OutputState }

func (EventTransportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventTransport)(nil)).Elem()
}

func (o EventTransportArrayOutput) ToEventTransportArrayOutput() EventTransportArrayOutput {
	return o
}

func (o EventTransportArrayOutput) ToEventTransportArrayOutputWithContext(ctx context.Context) EventTransportArrayOutput {
	return o
}

func (o EventTransportArrayOutput) Index(i pulumi.IntInput) EventTransportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventTransport {
		return vs[0].([]*EventTransport)[vs[1].(int)]
	}).(EventTransportOutput)
}

type EventTransportMapOutput struct{ *pulumi.OutputState }

func (EventTransportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventTransport)(nil)).Elem()
}

func (o EventTransportMapOutput) ToEventTransportMapOutput() EventTransportMapOutput {
	return o
}

func (o EventTransportMapOutput) ToEventTransportMapOutputWithContext(ctx context.Context) EventTransportMapOutput {
	return o
}

func (o EventTransportMapOutput) MapIndex(k pulumi.StringInput) EventTransportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventTransport {
		return vs[0].(map[string]*EventTransport)[vs[1].(string)]
	}).(EventTransportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventTransportInput)(nil)).Elem(), &EventTransport{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventTransportArrayInput)(nil)).Elem(), EventTransportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventTransportMapInput)(nil)).Elem(), EventTransportMap{})
	pulumi.RegisterOutputType(EventTransportOutput{})
	pulumi.RegisterOutputType(EventTransportArrayOutput{})
	pulumi.RegisterOutputType(EventTransportMapOutput{})
}

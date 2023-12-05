// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package runpod

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type NetworkStorage struct {
	pulumi.CustomResourceState

	DataCenterId   pulumi.StringOutput      `pulumi:"dataCenterId"`
	Name           pulumi.StringOutput      `pulumi:"name"`
	NetworkStorage NetworkStorageTypeOutput `pulumi:"networkStorage"`
	Size           pulumi.IntOutput         `pulumi:"size"`
}

// NewNetworkStorage registers a new resource with the given unique name, arguments, and options.
func NewNetworkStorage(ctx *pulumi.Context,
	name string, args *NetworkStorageArgs, opts ...pulumi.ResourceOption) (*NetworkStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataCenterId == nil {
		return nil, errors.New("invalid value for required argument 'DataCenterId'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkStorage
	err := ctx.RegisterResource("runpod:index:NetworkStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkStorage gets an existing NetworkStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkStorageState, opts ...pulumi.ResourceOption) (*NetworkStorage, error) {
	var resource NetworkStorage
	err := ctx.ReadResource("runpod:index:NetworkStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkStorage resources.
type networkStorageState struct {
}

type NetworkStorageState struct {
}

func (NetworkStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkStorageState)(nil)).Elem()
}

type networkStorageArgs struct {
	DataCenterId string `pulumi:"dataCenterId"`
	Name         string `pulumi:"name"`
	Size         int    `pulumi:"size"`
}

// The set of arguments for constructing a NetworkStorage resource.
type NetworkStorageArgs struct {
	DataCenterId pulumi.StringInput
	Name         pulumi.StringInput
	Size         pulumi.IntInput
}

func (NetworkStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkStorageArgs)(nil)).Elem()
}

type NetworkStorageInput interface {
	pulumi.Input

	ToNetworkStorageOutput() NetworkStorageOutput
	ToNetworkStorageOutputWithContext(ctx context.Context) NetworkStorageOutput
}

func (*NetworkStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkStorage)(nil)).Elem()
}

func (i *NetworkStorage) ToNetworkStorageOutput() NetworkStorageOutput {
	return i.ToNetworkStorageOutputWithContext(context.Background())
}

func (i *NetworkStorage) ToNetworkStorageOutputWithContext(ctx context.Context) NetworkStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkStorageOutput)
}

type NetworkStorageOutput struct{ *pulumi.OutputState }

func (NetworkStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkStorage)(nil)).Elem()
}

func (o NetworkStorageOutput) ToNetworkStorageOutput() NetworkStorageOutput {
	return o
}

func (o NetworkStorageOutput) ToNetworkStorageOutputWithContext(ctx context.Context) NetworkStorageOutput {
	return o
}

func (o NetworkStorageOutput) DataCenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkStorage) pulumi.StringOutput { return v.DataCenterId }).(pulumi.StringOutput)
}

func (o NetworkStorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkStorage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkStorageOutput) NetworkStorage() NetworkStorageTypeOutput {
	return o.ApplyT(func(v *NetworkStorage) NetworkStorageTypeOutput { return v.NetworkStorage }).(NetworkStorageTypeOutput)
}

func (o NetworkStorageOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkStorage) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkStorageInput)(nil)).Elem(), &NetworkStorage{})
	pulumi.RegisterOutputType(NetworkStorageOutput{})
}

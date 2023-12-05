// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package runpod

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

var _ = internal.GetEnvOrDefault

type DataCenter struct {
	Id             string `pulumi:"id"`
	Location       string `pulumi:"location"`
	Name           string `pulumi:"name"`
	StorageSupport bool   `pulumi:"storageSupport"`
}

type DataCenterOutput struct{ *pulumi.OutputState }

func (DataCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCenter)(nil)).Elem()
}

func (o DataCenterOutput) ToDataCenterOutput() DataCenterOutput {
	return o
}

func (o DataCenterOutput) ToDataCenterOutputWithContext(ctx context.Context) DataCenterOutput {
	return o
}

func (o DataCenterOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataCenter) string { return v.Id }).(pulumi.StringOutput)
}

func (o DataCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v DataCenter) string { return v.Location }).(pulumi.StringOutput)
}

func (o DataCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataCenter) string { return v.Name }).(pulumi.StringOutput)
}

func (o DataCenterOutput) StorageSupport() pulumi.BoolOutput {
	return o.ApplyT(func(v DataCenter) bool { return v.StorageSupport }).(pulumi.BoolOutput)
}

type Gpu struct {
	Id    string `pulumi:"id"`
	PodId string `pulumi:"podId"`
}

type GpuOutput struct{ *pulumi.OutputState }

func (GpuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gpu)(nil)).Elem()
}

func (o GpuOutput) ToGpuOutput() GpuOutput {
	return o
}

func (o GpuOutput) ToGpuOutputWithContext(ctx context.Context) GpuOutput {
	return o
}

func (o GpuOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Gpu) string { return v.Id }).(pulumi.StringOutput)
}

func (o GpuOutput) PodId() pulumi.StringOutput {
	return o.ApplyT(func(v Gpu) string { return v.PodId }).(pulumi.StringOutput)
}

type GpuArrayOutput struct{ *pulumi.OutputState }

func (GpuArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Gpu)(nil)).Elem()
}

func (o GpuArrayOutput) ToGpuArrayOutput() GpuArrayOutput {
	return o
}

func (o GpuArrayOutput) ToGpuArrayOutputWithContext(ctx context.Context) GpuArrayOutput {
	return o
}

func (o GpuArrayOutput) Index(i pulumi.IntInput) GpuOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Gpu {
		return vs[0].([]Gpu)[vs[1].(int)]
	}).(GpuOutput)
}

type NetworkStorageType struct {
	DataCenter   DataCenter `pulumi:"dataCenter"`
	DataCenterId string     `pulumi:"dataCenterId"`
	Id           string     `pulumi:"id"`
	Name         string     `pulumi:"name"`
	Size         int        `pulumi:"size"`
}

type NetworkStorageTypeOutput struct{ *pulumi.OutputState }

func (NetworkStorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkStorageType)(nil)).Elem()
}

func (o NetworkStorageTypeOutput) ToNetworkStorageTypeOutput() NetworkStorageTypeOutput {
	return o
}

func (o NetworkStorageTypeOutput) ToNetworkStorageTypeOutputWithContext(ctx context.Context) NetworkStorageTypeOutput {
	return o
}

func (o NetworkStorageTypeOutput) DataCenter() DataCenterOutput {
	return o.ApplyT(func(v NetworkStorageType) DataCenter { return v.DataCenter }).(DataCenterOutput)
}

func (o NetworkStorageTypeOutput) DataCenterId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkStorageType) string { return v.DataCenterId }).(pulumi.StringOutput)
}

func (o NetworkStorageTypeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkStorageType) string { return v.Id }).(pulumi.StringOutput)
}

func (o NetworkStorageTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkStorageType) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkStorageTypeOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkStorageType) int { return v.Size }).(pulumi.IntOutput)
}

type PodType struct {
	AdjustedCostPerHr       float64     `pulumi:"adjustedCostPerHr"`
	AiApiId                 string      `pulumi:"aiApiId"`
	ApiKey                  string      `pulumi:"apiKey"`
	ConsumerUserId          string      `pulumi:"consumerUserId"`
	ContainerDiskInGb       int         `pulumi:"containerDiskInGb"`
	ContainerRegistryAuthId string      `pulumi:"containerRegistryAuthId"`
	CostMultiplier          float64     `pulumi:"costMultiplier"`
	CostPerHr               float64     `pulumi:"costPerHr"`
	CreatedAt               string      `pulumi:"createdAt"`
	DesiredStatus           string      `pulumi:"desiredStatus"`
	DockerArgs              string      `pulumi:"dockerArgs"`
	DockerId                string      `pulumi:"dockerId"`
	Env                     []string    `pulumi:"env"`
	GpuCount                int         `pulumi:"gpuCount"`
	GpuPowerLimitPercent    int         `pulumi:"gpuPowerLimitPercent"`
	Gpus                    []Gpu       `pulumi:"gpus"`
	Id                      string      `pulumi:"id"`
	ImageName               string      `pulumi:"imageName"`
	LastStartedAt           string      `pulumi:"lastStartedAt"`
	LastStatusChange        string      `pulumi:"lastStatusChange"`
	Locked                  bool        `pulumi:"locked"`
	MachineId               string      `pulumi:"machineId"`
	MemoryInGb              float64     `pulumi:"memoryInGb"`
	Name                    string      `pulumi:"name"`
	PodType                 string      `pulumi:"podType"`
	Port                    int         `pulumi:"port"`
	Ports                   string      `pulumi:"ports"`
	Registry                PodRegistry `pulumi:"registry"`
	TemplateId              string      `pulumi:"templateId"`
	UptimeSeconds           int         `pulumi:"uptimeSeconds"`
	VcpuCount               float64     `pulumi:"vcpuCount"`
	Version                 int         `pulumi:"version"`
	VolumeEncrypted         bool        `pulumi:"volumeEncrypted"`
	VolumeInGb              float64     `pulumi:"volumeInGb"`
	VolumeKey               string      `pulumi:"volumeKey"`
	VolumeMountPath         string      `pulumi:"volumeMountPath"`
}

type PodTypeOutput struct{ *pulumi.OutputState }

func (PodTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodType)(nil)).Elem()
}

func (o PodTypeOutput) ToPodTypeOutput() PodTypeOutput {
	return o
}

func (o PodTypeOutput) ToPodTypeOutputWithContext(ctx context.Context) PodTypeOutput {
	return o
}

func (o PodTypeOutput) AdjustedCostPerHr() pulumi.Float64Output {
	return o.ApplyT(func(v PodType) float64 { return v.AdjustedCostPerHr }).(pulumi.Float64Output)
}

func (o PodTypeOutput) AiApiId() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.AiApiId }).(pulumi.StringOutput)
}

func (o PodTypeOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.ApiKey }).(pulumi.StringOutput)
}

func (o PodTypeOutput) ConsumerUserId() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.ConsumerUserId }).(pulumi.StringOutput)
}

func (o PodTypeOutput) ContainerDiskInGb() pulumi.IntOutput {
	return o.ApplyT(func(v PodType) int { return v.ContainerDiskInGb }).(pulumi.IntOutput)
}

func (o PodTypeOutput) ContainerRegistryAuthId() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.ContainerRegistryAuthId }).(pulumi.StringOutput)
}

func (o PodTypeOutput) CostMultiplier() pulumi.Float64Output {
	return o.ApplyT(func(v PodType) float64 { return v.CostMultiplier }).(pulumi.Float64Output)
}

func (o PodTypeOutput) CostPerHr() pulumi.Float64Output {
	return o.ApplyT(func(v PodType) float64 { return v.CostPerHr }).(pulumi.Float64Output)
}

func (o PodTypeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o PodTypeOutput) DesiredStatus() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.DesiredStatus }).(pulumi.StringOutput)
}

func (o PodTypeOutput) DockerArgs() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.DockerArgs }).(pulumi.StringOutput)
}

func (o PodTypeOutput) DockerId() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.DockerId }).(pulumi.StringOutput)
}

func (o PodTypeOutput) Env() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PodType) []string { return v.Env }).(pulumi.StringArrayOutput)
}

func (o PodTypeOutput) GpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v PodType) int { return v.GpuCount }).(pulumi.IntOutput)
}

func (o PodTypeOutput) GpuPowerLimitPercent() pulumi.IntOutput {
	return o.ApplyT(func(v PodType) int { return v.GpuPowerLimitPercent }).(pulumi.IntOutput)
}

func (o PodTypeOutput) Gpus() GpuArrayOutput {
	return o.ApplyT(func(v PodType) []Gpu { return v.Gpus }).(GpuArrayOutput)
}

func (o PodTypeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.Id }).(pulumi.StringOutput)
}

func (o PodTypeOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o PodTypeOutput) LastStartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.LastStartedAt }).(pulumi.StringOutput)
}

func (o PodTypeOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o PodTypeOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v PodType) bool { return v.Locked }).(pulumi.BoolOutput)
}

func (o PodTypeOutput) MachineId() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.MachineId }).(pulumi.StringOutput)
}

func (o PodTypeOutput) MemoryInGb() pulumi.Float64Output {
	return o.ApplyT(func(v PodType) float64 { return v.MemoryInGb }).(pulumi.Float64Output)
}

func (o PodTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.Name }).(pulumi.StringOutput)
}

func (o PodTypeOutput) PodType() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.PodType }).(pulumi.StringOutput)
}

func (o PodTypeOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v PodType) int { return v.Port }).(pulumi.IntOutput)
}

func (o PodTypeOutput) Ports() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.Ports }).(pulumi.StringOutput)
}

func (o PodTypeOutput) Registry() PodRegistryOutput {
	return o.ApplyT(func(v PodType) PodRegistry { return v.Registry }).(PodRegistryOutput)
}

func (o PodTypeOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.TemplateId }).(pulumi.StringOutput)
}

func (o PodTypeOutput) UptimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v PodType) int { return v.UptimeSeconds }).(pulumi.IntOutput)
}

func (o PodTypeOutput) VcpuCount() pulumi.Float64Output {
	return o.ApplyT(func(v PodType) float64 { return v.VcpuCount }).(pulumi.Float64Output)
}

func (o PodTypeOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v PodType) int { return v.Version }).(pulumi.IntOutput)
}

func (o PodTypeOutput) VolumeEncrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v PodType) bool { return v.VolumeEncrypted }).(pulumi.BoolOutput)
}

func (o PodTypeOutput) VolumeInGb() pulumi.Float64Output {
	return o.ApplyT(func(v PodType) float64 { return v.VolumeInGb }).(pulumi.Float64Output)
}

func (o PodTypeOutput) VolumeKey() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.VolumeKey }).(pulumi.StringOutput)
}

func (o PodTypeOutput) VolumeMountPath() pulumi.StringOutput {
	return o.ApplyT(func(v PodType) string { return v.VolumeMountPath }).(pulumi.StringOutput)
}

type PodEnv struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// PodEnvInput is an input type that accepts PodEnvArgs and PodEnvOutput values.
// You can construct a concrete instance of `PodEnvInput` via:
//
//	PodEnvArgs{...}
type PodEnvInput interface {
	pulumi.Input

	ToPodEnvOutput() PodEnvOutput
	ToPodEnvOutputWithContext(context.Context) PodEnvOutput
}

type PodEnvArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (PodEnvArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PodEnv)(nil)).Elem()
}

func (i PodEnvArgs) ToPodEnvOutput() PodEnvOutput {
	return i.ToPodEnvOutputWithContext(context.Background())
}

func (i PodEnvArgs) ToPodEnvOutputWithContext(ctx context.Context) PodEnvOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodEnvOutput)
}

// PodEnvArrayInput is an input type that accepts PodEnvArray and PodEnvArrayOutput values.
// You can construct a concrete instance of `PodEnvArrayInput` via:
//
//	PodEnvArray{ PodEnvArgs{...} }
type PodEnvArrayInput interface {
	pulumi.Input

	ToPodEnvArrayOutput() PodEnvArrayOutput
	ToPodEnvArrayOutputWithContext(context.Context) PodEnvArrayOutput
}

type PodEnvArray []PodEnvInput

func (PodEnvArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PodEnv)(nil)).Elem()
}

func (i PodEnvArray) ToPodEnvArrayOutput() PodEnvArrayOutput {
	return i.ToPodEnvArrayOutputWithContext(context.Background())
}

func (i PodEnvArray) ToPodEnvArrayOutputWithContext(ctx context.Context) PodEnvArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodEnvArrayOutput)
}

type PodEnvOutput struct{ *pulumi.OutputState }

func (PodEnvOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodEnv)(nil)).Elem()
}

func (o PodEnvOutput) ToPodEnvOutput() PodEnvOutput {
	return o
}

func (o PodEnvOutput) ToPodEnvOutputWithContext(ctx context.Context) PodEnvOutput {
	return o
}

func (o PodEnvOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v PodEnv) string { return v.Key }).(pulumi.StringOutput)
}

func (o PodEnvOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v PodEnv) string { return v.Value }).(pulumi.StringOutput)
}

type PodEnvArrayOutput struct{ *pulumi.OutputState }

func (PodEnvArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PodEnv)(nil)).Elem()
}

func (o PodEnvArrayOutput) ToPodEnvArrayOutput() PodEnvArrayOutput {
	return o
}

func (o PodEnvArrayOutput) ToPodEnvArrayOutputWithContext(ctx context.Context) PodEnvArrayOutput {
	return o
}

func (o PodEnvArrayOutput) Index(i pulumi.IntInput) PodEnvOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PodEnv {
		return vs[0].([]PodEnv)[vs[1].(int)]
	}).(PodEnvOutput)
}

type PodRegistry struct {
	Auth     string `pulumi:"auth"`
	Pass     string `pulumi:"pass"`
	Url      string `pulumi:"url"`
	User     string `pulumi:"user"`
	Username string `pulumi:"username"`
}

type PodRegistryOutput struct{ *pulumi.OutputState }

func (PodRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodRegistry)(nil)).Elem()
}

func (o PodRegistryOutput) ToPodRegistryOutput() PodRegistryOutput {
	return o
}

func (o PodRegistryOutput) ToPodRegistryOutputWithContext(ctx context.Context) PodRegistryOutput {
	return o
}

func (o PodRegistryOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v PodRegistry) string { return v.Auth }).(pulumi.StringOutput)
}

func (o PodRegistryOutput) Pass() pulumi.StringOutput {
	return o.ApplyT(func(v PodRegistry) string { return v.Pass }).(pulumi.StringOutput)
}

func (o PodRegistryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v PodRegistry) string { return v.Url }).(pulumi.StringOutput)
}

func (o PodRegistryOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v PodRegistry) string { return v.User }).(pulumi.StringOutput)
}

func (o PodRegistryOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v PodRegistry) string { return v.Username }).(pulumi.StringOutput)
}

type SavingsPlanInput struct {
	PlanLength  string  `pulumi:"planLength"`
	UpfrontCost float64 `pulumi:"upfrontCost"`
}

// SavingsPlanInputInput is an input type that accepts SavingsPlanInputArgs and SavingsPlanInputOutput values.
// You can construct a concrete instance of `SavingsPlanInputInput` via:
//
//	SavingsPlanInputArgs{...}
type SavingsPlanInputInput interface {
	pulumi.Input

	ToSavingsPlanInputOutput() SavingsPlanInputOutput
	ToSavingsPlanInputOutputWithContext(context.Context) SavingsPlanInputOutput
}

type SavingsPlanInputArgs struct {
	PlanLength  pulumi.StringInput  `pulumi:"planLength"`
	UpfrontCost pulumi.Float64Input `pulumi:"upfrontCost"`
}

func (SavingsPlanInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SavingsPlanInput)(nil)).Elem()
}

func (i SavingsPlanInputArgs) ToSavingsPlanInputOutput() SavingsPlanInputOutput {
	return i.ToSavingsPlanInputOutputWithContext(context.Background())
}

func (i SavingsPlanInputArgs) ToSavingsPlanInputOutputWithContext(ctx context.Context) SavingsPlanInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavingsPlanInputOutput)
}

func (i SavingsPlanInputArgs) ToSavingsPlanInputPtrOutput() SavingsPlanInputPtrOutput {
	return i.ToSavingsPlanInputPtrOutputWithContext(context.Background())
}

func (i SavingsPlanInputArgs) ToSavingsPlanInputPtrOutputWithContext(ctx context.Context) SavingsPlanInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavingsPlanInputOutput).ToSavingsPlanInputPtrOutputWithContext(ctx)
}

// SavingsPlanInputPtrInput is an input type that accepts SavingsPlanInputArgs, SavingsPlanInputPtr and SavingsPlanInputPtrOutput values.
// You can construct a concrete instance of `SavingsPlanInputPtrInput` via:
//
//	        SavingsPlanInputArgs{...}
//
//	or:
//
//	        nil
type SavingsPlanInputPtrInput interface {
	pulumi.Input

	ToSavingsPlanInputPtrOutput() SavingsPlanInputPtrOutput
	ToSavingsPlanInputPtrOutputWithContext(context.Context) SavingsPlanInputPtrOutput
}

type savingsPlanInputPtrType SavingsPlanInputArgs

func SavingsPlanInputPtr(v *SavingsPlanInputArgs) SavingsPlanInputPtrInput {
	return (*savingsPlanInputPtrType)(v)
}

func (*savingsPlanInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SavingsPlanInput)(nil)).Elem()
}

func (i *savingsPlanInputPtrType) ToSavingsPlanInputPtrOutput() SavingsPlanInputPtrOutput {
	return i.ToSavingsPlanInputPtrOutputWithContext(context.Background())
}

func (i *savingsPlanInputPtrType) ToSavingsPlanInputPtrOutputWithContext(ctx context.Context) SavingsPlanInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavingsPlanInputPtrOutput)
}

type SavingsPlanInputOutput struct{ *pulumi.OutputState }

func (SavingsPlanInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SavingsPlanInput)(nil)).Elem()
}

func (o SavingsPlanInputOutput) ToSavingsPlanInputOutput() SavingsPlanInputOutput {
	return o
}

func (o SavingsPlanInputOutput) ToSavingsPlanInputOutputWithContext(ctx context.Context) SavingsPlanInputOutput {
	return o
}

func (o SavingsPlanInputOutput) ToSavingsPlanInputPtrOutput() SavingsPlanInputPtrOutput {
	return o.ToSavingsPlanInputPtrOutputWithContext(context.Background())
}

func (o SavingsPlanInputOutput) ToSavingsPlanInputPtrOutputWithContext(ctx context.Context) SavingsPlanInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SavingsPlanInput) *SavingsPlanInput {
		return &v
	}).(SavingsPlanInputPtrOutput)
}

func (o SavingsPlanInputOutput) PlanLength() pulumi.StringOutput {
	return o.ApplyT(func(v SavingsPlanInput) string { return v.PlanLength }).(pulumi.StringOutput)
}

func (o SavingsPlanInputOutput) UpfrontCost() pulumi.Float64Output {
	return o.ApplyT(func(v SavingsPlanInput) float64 { return v.UpfrontCost }).(pulumi.Float64Output)
}

type SavingsPlanInputPtrOutput struct{ *pulumi.OutputState }

func (SavingsPlanInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SavingsPlanInput)(nil)).Elem()
}

func (o SavingsPlanInputPtrOutput) ToSavingsPlanInputPtrOutput() SavingsPlanInputPtrOutput {
	return o
}

func (o SavingsPlanInputPtrOutput) ToSavingsPlanInputPtrOutputWithContext(ctx context.Context) SavingsPlanInputPtrOutput {
	return o
}

func (o SavingsPlanInputPtrOutput) Elem() SavingsPlanInputOutput {
	return o.ApplyT(func(v *SavingsPlanInput) SavingsPlanInput {
		if v != nil {
			return *v
		}
		var ret SavingsPlanInput
		return ret
	}).(SavingsPlanInputOutput)
}

func (o SavingsPlanInputPtrOutput) PlanLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SavingsPlanInput) *string {
		if v == nil {
			return nil
		}
		return &v.PlanLength
	}).(pulumi.StringPtrOutput)
}

func (o SavingsPlanInputPtrOutput) UpfrontCost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SavingsPlanInput) *float64 {
		if v == nil {
			return nil
		}
		return &v.UpfrontCost
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodEnvInput)(nil)).Elem(), PodEnvArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodEnvArrayInput)(nil)).Elem(), PodEnvArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SavingsPlanInputInput)(nil)).Elem(), SavingsPlanInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SavingsPlanInputPtrInput)(nil)).Elem(), SavingsPlanInputArgs{})
	pulumi.RegisterOutputType(DataCenterOutput{})
	pulumi.RegisterOutputType(GpuOutput{})
	pulumi.RegisterOutputType(GpuArrayOutput{})
	pulumi.RegisterOutputType(NetworkStorageTypeOutput{})
	pulumi.RegisterOutputType(PodTypeOutput{})
	pulumi.RegisterOutputType(PodEnvOutput{})
	pulumi.RegisterOutputType(PodEnvArrayOutput{})
	pulumi.RegisterOutputType(PodRegistryOutput{})
	pulumi.RegisterOutputType(SavingsPlanInputOutput{})
	pulumi.RegisterOutputType(SavingsPlanInputPtrOutput{})
}

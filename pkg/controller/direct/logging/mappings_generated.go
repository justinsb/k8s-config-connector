package logging

// import (
// 	pb "cloud.google.com/go/logging/apiv2/loggingpb"

// 	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resources/logging/v1beta1"
// )

// func LoggingLogMetricSpec_FromProto(ctx *MapContext, in *pb.LogMetric) *krm.LoggingLogMetricSpec {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.LoggingLogMetricSpec{}
// 	// MISSING: Name
// 	out.Description = LazyPtr(in.GetDescription())
// 	out.Filter = LazyPtr(in.GetFilter())
// 	// MISSING: BucketName
// 	out.Disabled = LazyPtr(in.GetDisabled())
// 	out.MetricDescriptor = LogmetricMetricDescriptor_FromProto(ctx, in.GetMetricDescriptor())
// 	out.ValueExtractor = LazyPtr(in.GetValueExtractor())
// 	out.LabelExtractors = Slice_FromProto(ctx, in.LabelExtractors, map[string]string_FromProto)
// 	out.BucketOptions = LogmetricBucketOptions_FromProto(ctx, in.GetBucketOptions())
// 	// MISSING: CreateTime
// 	// MISSING: UpdateTime
// 	// MISSING: Version
// 	return out
// }
// func LoggingLogMetricSpec_ToProto(ctx *MapContext, in *krm.LoggingLogMetricSpec) *pb.LogMetric {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.LogMetric{}
// 	// MISSING: Name
// 	out.Description = ValueOf(in.Description)
// 	out.Filter = LoggingLogMetricSpec_Filter_ToProto(ctx, in.Filter)
// 	// MISSING: BucketName
// 	out.Disabled = ValueOf(in.Disabled)
// 	out.MetricDescriptor = LogmetricMetricDescriptor_ToProto(ctx, in.MetricDescriptor)
// 	out.ValueExtractor = ValueOf(in.ValueExtractor)
// 	out.LabelExtractors = Slice_ToProto(ctx, in.LabelExtractors, map[string]string_ToProto)
// 	out.BucketOptions = LogmetricBucketOptions_ToProto(ctx, in.BucketOptions)
// 	// MISSING: CreateTime
// 	// MISSING: UpdateTime
// 	// MISSING: Version
// 	return out
// }
// func LoggingLogMetricStatus_FromProto(ctx *MapContext, in *pb.LogMetric) *krm.LoggingLogMetricStatus {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.LoggingLogMetricStatus{}
// 	// MISSING: Name
// 	// MISSING: Description
// 	// MISSING: Filter
// 	// MISSING: BucketName
// 	// MISSING: Disabled
// 	out.MetricDescriptor = LogmetricMetricDescriptorStatus_FromProto(ctx, in.GetMetricDescriptor())
// 	// MISSING: ValueExtractor
// 	// MISSING: LabelExtractors
// 	// MISSING: BucketOptions
// 	out.CreateTime = LogMetric_CreateTime_FromProto(ctx, in.GetCreateTime())
// 	out.UpdateTime = LogMetric_UpdateTime_FromProto(ctx, in.GetUpdateTime())
// 	// MISSING: Version
// 	return out
// }
// func LoggingLogMetricStatus_ToProto(ctx *MapContext, in *krm.LoggingLogMetricStatus) *pb.LogMetric {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.LogMetric{}
// 	// MISSING: Name
// 	// MISSING: Description
// 	// MISSING: Filter
// 	// MISSING: BucketName
// 	// MISSING: Disabled
// 	out.MetricDescriptor = LogmetricMetricDescriptorStatus_ToProto(ctx, in.MetricDescriptor)
// 	// MISSING: ValueExtractor
// 	// MISSING: LabelExtractors
// 	// MISSING: BucketOptions
// 	out.CreateTime = LogMetric_CreateTime_ToProto(ctx, in.CreateTime)
// 	out.UpdateTime = LogMetric_UpdateTime_ToProto(ctx, in.UpdateTime)
// 	// MISSING: Version
// 	return out
// }
// func LogmetricMetricDescriptor_FromProto(ctx *MapContext, in *pb.MetricDescriptor) *krm.LogmetricMetricDescriptor {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.LogmetricMetricDescriptor{}
// 	// MISSING: Name
// 	// MISSING: Type
// 	out.Labels = Slice_FromProto(ctx, in.Labels, LogmetricLabels_FromProto)
// 	out.MetricKind = Enum_FromProto(ctx, in.MetricKind)
// 	out.ValueType = Enum_FromProto(ctx, in.ValueType)
// 	out.Unit = LazyPtr(in.GetUnit())
// 	// MISSING: Description
// 	out.DisplayName = LazyPtr(in.GetDisplayName())
// 	out.Metadata = LogmetricMetadata_FromProto(ctx, in.GetMetadata())
// 	out.LaunchStage = Enum_FromProto(ctx, in.LaunchStage)
// 	// MISSING: MonitoredResourceTypes
// 	return out
// }
// func LogmetricMetricDescriptor_ToProto(ctx *MapContext, in *krm.LogmetricMetricDescriptor) *pb.MetricDescriptor {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.MetricDescriptor{}
// 	// MISSING: Name
// 	// MISSING: Type
// 	out.Labels = Slice_ToProto(ctx, in.Labels, LogmetricLabels_ToProto)
// 	out.MetricKind = Enum_ToProto[pb.MetricDescriptor_MetricKind](ctx, in.MetricKind)
// 	out.ValueType = Enum_ToProto[pb.MetricDescriptor_ValueType](ctx, in.ValueType)
// 	out.Unit = ValueOf(in.Unit)
// 	// MISSING: Description
// 	out.DisplayName = ValueOf(in.DisplayName)
// 	out.Metadata = LogmetricMetadata_ToProto(ctx, in.Metadata)
// 	out.LaunchStage = Enum_ToProto[pb.LaunchStage](ctx, in.LaunchStage)
// 	// MISSING: MonitoredResourceTypes
// 	return out
// }

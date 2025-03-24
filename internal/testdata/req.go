package testdata

import (
	collecctorpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logpb "go.opentelemetry.io/proto/otlp/logs/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
)

var TestReq = &collecctorpb.ExportLogsServiceRequest{
	ResourceLogs: []*logpb.ResourceLogs{
		{
			Resource: &resourcepb.Resource{
				Attributes: []*commonpb.KeyValue{
					{
						Key: "foo",
						Value: &commonpb.AnyValue{
							Value: &commonpb.AnyValue_BoolValue{BoolValue: true},
						},
					},
				},
			},
			ScopeLogs: []*logpb.ScopeLogs{
				{
					Scope: &commonpb.InstrumentationScope{
						Attributes: []*commonpb.KeyValue{
							{
								Key: "foo",
								Value: &commonpb.AnyValue{
									Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 1.23},
								},
							},
						},
					},
					LogRecords: []*logpb.LogRecord{
						{
							Attributes: []*commonpb.KeyValue{
								{
									Key: "foo",
									Value: &commonpb.AnyValue{
										Value: &commonpb.AnyValue_StringValue{StringValue: "bar"},
									},
								},
							},
						},
						{
							Attributes: []*commonpb.KeyValue{
								{
									Key: "foo",
									Value: &commonpb.AnyValue{
										Value: &commonpb.AnyValue_StringValue{StringValue: "baz"},
									},
								},
							},
						},
						{
							Attributes: []*commonpb.KeyValue{
								{
									Key: "foo",
									Value: &commonpb.AnyValue{
										Value: &commonpb.AnyValue_StringValue{StringValue: "bar"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Resource: &resourcepb.Resource{
				Attributes: []*commonpb.KeyValue{
					{
						Key: "foo",
						Value: &commonpb.AnyValue{
							Value: &commonpb.AnyValue_BoolValue{BoolValue: true},
						},
					},
				},
			},
			ScopeLogs: []*logpb.ScopeLogs{
				{
					Scope: &commonpb.InstrumentationScope{
						Attributes: []*commonpb.KeyValue{
							{
								Key: "foo",
								Value: &commonpb.AnyValue{
									Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 1.23},
								},
							},
						},
					},
					LogRecords: []*logpb.LogRecord{
						{
							Attributes: []*commonpb.KeyValue{
								{
									Key: "foo",
									Value: &commonpb.AnyValue{
										Value: &commonpb.AnyValue_StringValue{StringValue: "bar"},
									},
								},
							},
						},
						{
							Attributes: []*commonpb.KeyValue{
								{
									Key: "foo",
									Value: &commonpb.AnyValue{
										Value: &commonpb.AnyValue_StringValue{StringValue: "baz"},
									},
								},
							},
						},
						{
							Attributes: []*commonpb.KeyValue{
								{
									Key: "foo",
									Value: &commonpb.AnyValue{
										Value: &commonpb.AnyValue_StringValue{StringValue: "bar"},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

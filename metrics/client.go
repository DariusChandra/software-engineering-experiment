package metrics

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func StartTracer(opts ...tracer.StartOption) {
	tracer.Start(opts...)
}

func StopTracer() {
	tracer.Stop()
}

func TagsToOpts(tags map[string]interface{}) []tracer.StartOption {
	opts := make([]tracer.StartOption, 0, len(tags))
	for k, v := range tags {
		opts = append(opts, tracer.WithGlobalTag(k, v))
	}
	return opts
}

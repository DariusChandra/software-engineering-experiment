package main

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"reflect"
	"testing"
)

func Test_getTracerOpts(t *testing.T) {
	tests := []struct {
		name string
		want []tracer.StartOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTracerOpts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTracerOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

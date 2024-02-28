package main

import (
	"fmt"
	"github.com/DariusChandra/software-engineering-experiment/math"
	"github.com/DariusChandra/software-engineering-experiment/metrics"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"time"
)

var opts struct {
	MongoDB struct {
		URI            string        `long:"uri" env:"MONGODB_URI" description:"MongoDB URI" default:"mongodb://localhost:27018" required:"true"`
		DBName         string        `long:"db-name" env:"DATABASE_NAME" description:"MongoDB database name" default:"adtech-adcredits" required:"true"`
		ConnectTimeout time.Duration `long:"connect-timeout" description:"MongoDB connection timeout" default:"10s"`
		MaxPoolSize    uint64        `long:"max-pool-size" description:"MongoDB max pool size" default:"10"`
	} `group:"MongoDB" namespace:"mongodb"`
	GCP struct {
		ProjectID   string `long:"project-id" env:"GCP_PROJECT_ID" description:"GCP project ID" default:"dh-adtech" required:"true"`
		PubsubTopic string `long:"pubsub-topic" env:"PUBSUB_TOPIC" description:"GCP pubsub topic" default:"adtech-adcredits-credit-events-stg-euw3-v1" required:"true"`
	} `group:"GCP" namespace:"gcp"`
}

func main() {
	// start datadog tracing
	tracerOpts := getTracerOpts()
	metrics.StartTracer(tracerOpts...)
	defer metrics.StopTracer()

	fmt.Println("hello world")
	a := math.Sum(1, 2)
	fmt.Println(a)
}

func getTracerOpts() []tracer.StartOption {
	tags := map[string]interface{}{
		"dh_cc_id":         "100160131",
		"dh_squad":         "adtech-advertiser",
		"dh_platform":      "vendor-tech",
		"dh_tribe":         "adtech",
		"dh_slack_channel": "adtech-tech",
	}

	tracerOpts := metrics.TagsToOpts(tags)
	tracerOpts = append(tracerOpts,
		tracer.WithRuntimeMetrics(),
		tracer.WithDogstatsdAddress("unix:///var/run/datadog/dsd.socket"),
		tracer.WithLogStartup(false),
	)

	return tracerOpts
}

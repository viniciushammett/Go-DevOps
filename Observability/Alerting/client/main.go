package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

func main() {
	shutdown := initTraceAndMetricsProvider()
	defer shutdown()

	continuouslySendRequests()
}

func initTraceAndMetricsProvider() func() {
	ctx := context.Background()

	otelAgentAddr, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if !ok {
		otelAgentAddr = "0.0.0.0:4317"
	}

	closeMetrics := initMetrics(ctx, otelAgentAddr)
	closeTraces := initTracer(ctx, otelAgentAddr)

	return func() {
		doneCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		closeTraces(doneCtx)
		closeMetrics(doneCtx)
	}
}

func initTracer(ctx context.Context, otelAgentAddr string) func(context.Context) {
	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otelAgentAddr),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))
	traceExp, err := otlptrace.New(ctx, traceClient)
	handleErr(err, "Failed to create the collector trace exporter")

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("demo-client"),
		),
	)
	handleErr(err, "failed to create resource")

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	return func(doneCtx context.Context) {
		if err := traceExp.Shutdown(doneCtx); err != nil {
			otel.Handle(err)
		}
	}
}

func initMetrics(ctx context.Context, otelAgentAddr string) func(context.Context) {
	metricClient := otlpmetricgrpc.NewClient(
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(otelAgentAddr))
	metricExp, err := otlpmetric.New(ctx, metricClient)
	handleErr(err, "Failed to create the collector metric exporter")

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(attribute.String("service", "demo-client")),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("demo-client"),
		),
	)
	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithHistogramDistribution(),
			metricExp,
		),
		controller.WithExporter(metricExp),
		controller.WithCollectPeriod(2*time.Second),
		controller.WithResource(res),
	)
	global.SetMeterProvider(pusher)

	err = pusher.Start(ctx)
	handleErr(err, "Failed to start metric pusher")

	return func(doneCtx context.Context) {
		if err := pusher.Stop(doneCtx); err != nil {
			otel.Handle(err)
		}
	}
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func continuouslySendRequests() {
	var (
		tracer       = otel.Tracer("demo-client-tracer")
		meter        = global.Meter("demo-client-meter")
		instruments  = NewClientInstruments(meter)
		commonLabels = []attribute.KeyValue{
			attribute.String("method", "repl"),
			attribute.String("client", "cli"),
		}
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	for {
		startTime := time.Now()
		ctx, span := tracer.Start(context.Background(), "ExecuteRequest")
		makeRequest(ctx)
		span.End()
		latencyMs := float64(time.Since(startTime)) / 1e6
		nr := int(rng.Int31n(7))
		for i := 0; i < nr; i++ {
			randLineLength := rng.Int63n(999)
			meter.RecordBatch(
				ctx,
				commonLabels,
				instruments.LineCounts.Measurement(1),
				instruments.LineLengths.Measurement(randLineLength),
			)
			fmt.Printf("#%d: LineLength: %dBy\n", i, randLineLength)
		}

		meter.RecordBatch(
			ctx,
			commonLabels,
			instruments.RequestLatency.Measurement(latencyMs),
			instruments.RequestCount.Measurement(1),
		)

		fmt.Printf("Latency: %.3fms\n", latencyMs)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func makeRequest(ctx context.Context) {

	demoServerAddr, ok := os.LookupEnv("DEMO_SERVER_ENDPOINT")
	if !ok {
		demoServerAddr = "http://0.0.0.0:7080/hello"
	}

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", demoServerAddr, nil)
	if err != nil {
		handleErr(err, "failed to http request")
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
}

type ClientInstruments struct {
	RequestLatency metric.Float64Histogram
	RequestCount   metric.Int64Counter
	LineLengths    metric.Int64Histogram
	LineCounts     metric.Int64Counter
}

func NewClientInstruments(meter metric.Meter) ClientInstruments {
	return ClientInstruments{
		RequestLatency: metric.Must(meter).
			NewFloat64Histogram(
				"demo_client/request_latency",
				metric.WithDescription("The latency of requests processed"),
			),
		RequestCount: metric.Must(meter).
			NewInt64Counter(
				"demo_client/request_counts",
				metric.WithDescription("The number of requests processed"),
			),
		LineLengths: metric.Must(meter).
			NewInt64Histogram(
				"demo_client/line_lengths",
				metric.WithDescription("The lengths of the various lines in"),
			),
		LineCounts: metric.Must(meter).
			NewInt64Counter(
				"demo_client/line_counts",
				metric.WithDescription("The counts of the lines in"),
			),
	}
}

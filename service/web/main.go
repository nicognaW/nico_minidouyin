// Code generated by hertz generator.

package main

import (
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"

	"nico_minidouyin/config"
	"nico_minidouyin/mw"
)

func DevEnv() bool {
	return os.Getenv("ENV") == "dev"
}

func hlogInit() {
	logger := hertzlogrus.NewLogger()
	// hlog init
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelDebug)
}

func otelInit() {
	if !DevEnv() {
		p := provider.NewOpenTelemetryProvider(
			provider.WithServiceName(config.FeedServiceName),
			// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
			provider.WithInsecure(),
		)
		defer func() { _ = p.Shutdown(context.Background()) }()
	}
}

func main() {
	hlogInit()
	otelInit()
	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		tracer,
		server.WithHostPorts(config.ServiceAddress),
		*WithConsul(),
	)
	h.Use(mw.ProtoJsonMiddleware())

	// use pprof mw
	pprof.Register(h)
	// use otel mw
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}

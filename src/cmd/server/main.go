package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/caiowWillian/first-crud-golang/src/cmd/server/route"
	"github.com/caiowWillian/first-crud-golang/src/pkg/configuration"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println(basepath)
	configApp := configuration.Settings{
		ConsulAddress: "http://localhost:8500",
		ConsulKey:     "teste",
		LocalPath:     basepath,
	}

	appsettings, configErr := configuration.NewConfigService(configApp)

	if configErr != nil {
		panic("config not found")
	}

	jaegerCfg := config.Configuration{
		ServiceName: "first-crud-golang",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := jaegerCfg.NewTracer()

	if err != nil {
		os.Exit(0)
	}

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	ctx := context.Background()

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		address := fmt.Sprintf(":%d", appsettings.GetInt("port"))
		r := mux.NewRouter()
		r.Handle("/metrics", promhttp.Handler())
		route.MakeRoutes(ctx, r)

		fmt.Println("listening on port", address)
		errs <- http.ListenAndServe(address, r)
	}()

	level.Error(logger).Log("exit", <-errs)
}

package bootstrap

import (
	"context"
	"errors"
	"github.com/samber/do/v2"
	"go.uber.org/zap"
	"image-optimization-api/internal/app"
	"image-optimization-api/internal/app/provider"
	"image-optimization-api/internal/service/compression"
	"image-optimization-api/pkg/server"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

func New() *Bootstrap {
	inj := do.New()
	do.Provide(inj, provider.ProvideConfig)

	mustInitLogger()

	return &Bootstrap{
		inj: inj,
	}
}

type Bootstrap struct {
	inj *do.RootScope
}

func (b *Bootstrap) Website() {
	do.ProvideValue(b.inj, do.MustInvoke[*app.Config](b.inj).Env)

	do.Provide(b.inj, provider.ProvideProtoValidator)
	do.Provide(b.inj, provider.ProvideRabbitMQConnection)
	do.Provide(b.inj, provider.ProvideFilestorage)

	do.Provide(b.inj, provider.ProvideImageRepository)

	do.Provide(b.inj, provider.ProvideImageService)
	do.Provide(b.inj, provider.ProvideCompressionService)

	do.Provide(b.inj, provider.ProvideQueueListener)
	do.Provide(b.inj, provider.ProvideWebsiteServer)

	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer cancelFunc()

	zap.L().Info("Starting application")

	waitForTheEnd := &sync.WaitGroup{}

	go func() {
		waitForTheEnd.Add(1)
		defer waitForTheEnd.Done()

		compressionService := do.MustInvoke[*compression.QueueListener](b.inj)
		zap.L().Info("Starting listening RabbitMQ")

		if err := compressionService.ListenUpdates(ctx); err != nil {
			zap.L().Error("RabbitMQ listening has been stopped, ", zap.Error(err))
		}
	}()

	go func() {
		waitForTheEnd.Add(1)
		defer waitForTheEnd.Done()

		srv := do.MustInvoke[*server.Server](b.inj)
		zap.L().Info("Starting server", zap.String("port", srv.GetPort()))
		go func() {
			<-ctx.Done()
			if err := srv.Shutdown(ctx); err != nil {
				zap.L().Error(err.Error())
			}
		}()
		if err := srv.Start(srv.GetPort()); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				zap.L().Error(err.Error())
			}
			zap.L().Info("Server has been stopped")
		}
	}()

	<-ctx.Done()
	waitForTheEnd.Wait()

	zap.L().Info("Closing application")

	if err := b.inj.Shutdown(); err != nil {
		zap.L().Error(err.Error())
	}
}

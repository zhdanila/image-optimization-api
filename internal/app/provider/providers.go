package provider

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/samber/do/v2"
	"github.com/streadway/amqp"
	"image-optimization-api/internal/app"
	"image-optimization-api/internal/app/interface/http/website"
	"image-optimization-api/pkg/server"
)

func ProvideConfig(_ do.Injector) (*app.Config, error) {
	return app.NewConfig()
}

func ProvideProtoValidator(_ do.Injector) (*protovalidate.Validator, error) {
	return protovalidate.New()
}

func ProvideRabbitMQConnection(inj do.Injector) (*amqp.Connection, error) {
	cnf := do.MustInvoke[*app.Config](inj)

	connection, err := amqp.Dial(cnf.RabbitMQHost)
	if err != nil {
		panic(err)
	}

	return connection, nil
}

func ProvideWebsiteServer(inj do.Injector) (*server.Server, error) {
	cnf := do.MustInvoke[*app.Config](inj)

	return website.NewServer(cnf, inj), nil
}

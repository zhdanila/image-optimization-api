package provider

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bufbuild/protovalidate-go"
	"github.com/samber/do/v2"
	"github.com/streadway/amqp"
	"image-optimization-api/internal/app"
	"image-optimization-api/internal/app/interface/http/website"
	"image-optimization-api/internal/service/compression"
	"image-optimization-api/pkg/db"
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

func ProvideFilestorage(inj do.Injector) (*s3.S3, error) {
	cnf := do.MustInvoke[*app.Config](inj)

	amazonConfig := db.AmazonConfig{
		Region:    cnf.S3Region,
		AccessKey: cnf.S3AccessKey,
		SecretKey: cnf.S3SecretKey,
	}

	s3db, err := db.NewS3DB(amazonConfig)
	if err != nil {
		panic(err)
	}

	return s3db, nil
}

func ProvideQueueListener(inj do.Injector) (*compression.QueueListener, error) {
	return compression.NewListener(
		do.MustInvoke[*amqp.Connection](inj),
		do.MustInvoke[*compression.Service](inj),
	), nil
}

func ProvideWebsiteServer(inj do.Injector) (*server.Server, error) {
	cnf := do.MustInvoke[*app.Config](inj)

	return website.NewServer(cnf, inj), nil
}

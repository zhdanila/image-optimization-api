package bootstrap

import (
	"github.com/fatih/structs"
	"github.com/samber/do/v2"
	"image-optimization-api/internal/app"
	"image-optimization-api/internal/app/provider"
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
	structs.DefaultTagName = `db`

	do.ProvideValue(b.inj, do.MustInvoke[*app.Config](b.inj).Env)
}

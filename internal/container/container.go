package container

import (
	"context"
	"github.com/samuelmachado/go-core/log"
)

// Components are a like service, but it doesn't include business case
type components struct {
	Log log.Logger
}

// Services hold the business case, and make the bridge between
type Services struct {
}

type Dependency struct {
	Components components
	Services   Services
}

func New(ctx context.Context) (context.Context, *Dependency, error) {
	cmp, err := setupComponents(ctx)
	if err != nil {
		return nil, nil, err
	}

	srv := Services{}

	dep := Dependency{
		Components: *cmp,
		Services:   srv,
	}

	return ctx, &dep, err

}

func setupComponents(_ context.Context) (*components, error) {

	log, err := log.NewLoggerZap(log.ZapConfig{
		Version:           "v0.1.0",
		DisableStackTrace: false,
	})
	if err != nil {
		return nil, err
	}

	return &components{
		Log: log,
	}, nil
}

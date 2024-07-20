package services

import (
	"context"
	"time"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
)

func NewTimer(cfg *dconfig.Config, repository *entities.Repository) (Timer, error) {
	timer, err := newAppTimer(cfg, repository)
	if err != nil {
		return nil, err
	}
	return &timerImpl{
		config: cfg,
		timer:  timer,
	}, nil
}

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type Timer interface {
	SetNow(ctx context.Context) context.Context
	GetNow(ctx context.Context) time.Time
}

type timerImpl struct {
	config *dconfig.Config
	timer  timer
}

func (t *timerImpl) SetNow(ctx context.Context) context.Context {
	return dcontext.SetTime(ctx, t.timer.now(ctx))
}

func (t *timerImpl) GetNow(ctx context.Context) time.Time {
	if now, err := dcontext.GetTime(ctx); err == nil {
		return now
	}
	return t.timer.now(ctx)
}

func newAppTimer(cfg *dconfig.Config, repository *entities.Repository) (timer, error) {
	location, err := time.LoadLocation(cfg.Sys.TZ)
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	if cfg.Sys.IsProd() {
		return &realTimerImpl{
			tz: location,
		}, nil
	}
	return &virtualTimerImpl{
		tz:           location,
		repositories: repository,
	}, nil
}

type timer interface {
	now(ctx context.Context) time.Time
}

type realTimerImpl struct {
	tz *time.Location
}

func (t *realTimerImpl) now(ctx context.Context) time.Time {
	return time.Now().In(t.tz)
}

type virtualTimerImpl struct {
	tz           *time.Location
	repositories *entities.Repository
}

func (t *virtualTimerImpl) now(ctx context.Context) time.Time {
	now := time.Now().In(t.tz)
	// Note: ここでユーザ固有のデバッグタイムに変換する
	// Todo: リポジトリから設定済のデバッグタイムを取得して適用する
	// userID := dcontext.GetAuthenticatedUserID(ctx)
	offset := 0
	return now.Add(time.Duration(offset) * time.Second)
}

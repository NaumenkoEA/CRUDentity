package repository

import (
	"context"
	"github.com/EgMeln/CRUDentity/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func (rep *Postgres) GetAllAdvert(e context.Context) ([]*model.Advert, error) {
	//TODO implement me
	panic("implement me")
}

type Adverts interface {
	AddAdvert(e context.Context, advert *model.Advert) error
	GetAllAdvert(e context.Context) ([]*model.Advert, error)
	GetByIDAdvert(e context.Context, id int) (*model.Advert, error)
	UpdateAdvert(e context.Context, id int, address string, price float32) error
	DeleteAdvert(e context.Context, id int) error
}

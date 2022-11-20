package service

import (
	"context"
	"github.com/EgMeln/CRUDentity/internal/model"
	"github.com/EgMeln/CRUDentity/internal/repository"
)

type AdvertService struct {
	conn repository.Adverts
}

func NewAdvertServicePostgres(rep *repository.Postgres) *AdvertService {
	return &AdvertService{conn: rep}
}

func (srv *AdvertService) Add(e context.Context, advert *model.Advert) error {
	return srv.conn.AddAdvert(e, advert)
}
func (srv *AdvertService) GetAll(e context.Context) ([]*model.Advert, error) {
	return srv.conn.GetAllAdvert(e)
}
func (srv *AdvertService) GetByID(e context.Context, id int) (*model.Advert, error) {
	return srv.conn.GetByIDAdvert(e, id)
}
func (srv *AdvertService) Update(e context.Context, id int, address string, price float32) error {
	return srv.conn.UpdateAdvert(e, id, address, price)
}
func (srv *AdvertService) Delete(e context.Context, id int) error {
	return srv.conn.DeleteAdvert(e, id)
}

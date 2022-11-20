package repository

import (
	"context"
	"fmt"
	"github.com/EgMeln/CRUDentity/internal/model"
)

func (rep *Postgres) AddAdvert(e context.Context, advert *model.Advert) error {
	_, err := rep.Pool.Exec(e, "INSERT INTO advert (id,address,price) VALUES ($1,$2,$3)", advert.ID, advert.Address, advert.Price)
	if err != nil {
		return fmt.Errorf("can't create advert %w", err)
	}
	return err
}

func (rep *Postgres) GetAllAdverts(e context.Context) ([]*model.Advert, error) {
	rows, err := rep.Pool.Query(e, "SELECT * FROM advert")
	if err != nil {
		return nil, fmt.Errorf("can't select all adverts %w", err)
	}
	defer rows.Close()
	var adverts []*model.Advert
	for rows.Next() {
		var advert model.Advert
		values, err := rows.Values()
		if err != nil {
			return adverts, err
		}
		advert.ID = int(values[0].(int32))
		advert.Address = values[1].(string)
		advert.Price = values[2].(float32)
		adverts = append(adverts, &advert)
	}
	return adverts, err
}

func (rep *Postgres) GetByIDAdvert(e context.Context, id int) (*model.Advert, error) {
	var advert model.Advert
	err := rep.Pool.QueryRow(e, "SELECT id,address, price from parking where id=$1", id).Scan(&advert.ID, &advert.Address, &advert.Price)
	if err != nil {
		return nil, fmt.Errorf("can't select advert %w", err)
	}
	return &advert, err
}

func (rep *Postgres) UpdateAdvert(e context.Context, id int, address string, price float32) error {
	_, err := rep.Pool.Exec(e, "UPDATE advert SET address =$1,price =$2 WHERE id = $3", address, price, id)
	if err != nil {
		return fmt.Errorf("can't update advert %w", err)
	}
	return err
}

func (rep *Postgres) DeleteAdvert(e context.Context, id int) error {
	row, err := rep.Pool.Exec(e, "DELETE FROM advert where id=$1", id)
	if err != nil {
		return fmt.Errorf("can't delete advert %w", err)
	}
	if row.RowsAffected() != 1 {
		return fmt.Errorf("nothing to delete%w", err)
	}
	return err
}

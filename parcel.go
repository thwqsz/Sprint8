package main

import (
	"database/sql"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	// реализуйте добавление строки в таблицу parcel, используйте данные из переменной p
	res, err := s.db.Exec("insert into parcel (address, client, status, created_at) values (?, ?, ?, ?)", p.Address, p.Client, p.Status, p.CreatedAt)
	if err != nil {
		return 0, err
	}

	num, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(num), nil
	// верните идентификатор последней добавленной записи
}

func (s ParcelStore) Get(number int) (Parcel, error) {
	p := Parcel{}
	// реализуйте чтение строки по заданному number
	// здесь из таблицы должна вернуться только одна строка
	row := s.db.QueryRow("select number, client, status, address, created_at from parcel where number = ?", number)
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
	if err != nil {
		return p, err
	}
	// заполните объект Parcel данными из таблицы
	return p, nil
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	var res []Parcel
	// реализуйте чтение строк из таблицы parcel по заданному client
	// здесь из таблицы может вернуться несколько строк
	rows, err := s.db.Query("select number, client, status, address, created_at from parcel where client = ? ", client)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var i_res Parcel
		err := rows.Scan(&i_res.Number, &i_res.Client, &i_res.Status, &i_res.Address, &i_res.CreatedAt)
		if err != nil {
			return res, err
		}
		res = append(res, i_res)
	}
	// заполните срез Parcel данными из таблицы

	return res, nil
}

func (s ParcelStore) SetStatus(number int, status string) error {
	// реализуйте обновление статуса в таблице parcel
	_, err := s.db.Exec("update parcel set status = ? where number = ?", status, number)
	if err != nil {
		return err
	}
	return nil
}

func (s ParcelStore) SetAddress(number int, address string) error {
	// реализуйте обновление адреса в таблице parcel
	// менять адрес можно только если значение статуса registered
	_, err := s.db.Exec("update parcel set address = ? where status = ? and number = ?", address, ParcelStatusRegistered, number)
	if err != nil {
		return err
	}
	return nil
}

func (s ParcelStore) Delete(number int) error {
	// реализуйте удаление строки из таблицы parcel
	// удалять строку можно только если значение статуса registered

	_, err := s.db.Exec("delete from parcel where status = ? and number = ?", ParcelStatusRegistered, number)
	if err != nil {
		return err
	}
	return nil
}

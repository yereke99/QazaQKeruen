package repository

import (
	"context"
	"qkeruen/models"
)

type SecurityDB struct {
	DB PgxIface
}

func NewSecurityService(ds PgxIface) SecurityDB {
	security := SecurityDB{
		DB: ds,
	}
	return security
}

func (pool *SecurityDB) Insert(data models.Security) error {
	q := `INSERT INTO security(
		    userId,
		    firstName,
		    lastName, 
		    A,
		    B,
		    firstNameD,
		    lastNameD, 
		    carNumber,
		    timeStart)VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := pool.DB.Exec(
		context.Background(),
		q,
		data.UserId,
		data.FirsrtName,
		data.LastName,
		data.From,
		data.To,
		data.FirsrtNameD,
		data.LastNameD,
		data.CarNumber,
		data.TimeStart,
	)
	if err != nil {
		return err
	}

	return nil
}

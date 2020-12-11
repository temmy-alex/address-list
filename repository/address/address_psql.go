package addressRepository

import (
	"address-list/models"
	"database/sql"
	"log"
)

type AddressRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b AddressRepository) GetAddressUsers(db *sql.DB, address models.Address, address_users []models.Address) ([]models.Address, error) {
	rows, err := db.Query("SELECT * FROM address")

	if err != nil {
		return []models.Address{}, err
	}

	for rows.Next() {
		rows.Scan(&address.ID, &address.Street, &address.City, &address.Zip, &address.UserID)
		address_users = append(address_users, address)
	}

	if err != nil {
		return []models.Address{}, err
	}

	return address_users, nil
}

func (b AddressRepository) GetAddressUser(db *sql.DB, address models.Address, id int) (models.Address, error) {
	rows := db.QueryRow("SELECT * FROM address WHERE id=$1", id)
	err := rows.Scan(&address.ID, &address.Street, &address.City, &address.Zip, &address.UserID)

	return address, err
}

func (b AddressRepository) GetInfoAddress(db *sql.DB, address models.Address, id int) (models.Address, error) {
	rows := db.QueryRow("SELECT b.street, b.city, b.zip, a.name FROM users AS a INNER JOIN address AS b ON a.id = b.user_id WHERE b.user_id=$1", id)
	err := rows.Scan(&address.ID, &address.Street, &address.City, &address.Zip, &address.UserID)

	log.Println(*rows)

	return address, err
}

func (b AddressRepository) AddAddressUser(db *sql.DB, address models.Address) (int, error) {
	err := db.QueryRow("INSERT INTO address (street, city, zip, user_id) VALUES ($1, $2, $3, $4) RETURNING id;",
		address.Street, address.City, address.Zip, address.UserID).Scan(&address.ID)

	if err != nil {
		return 0, err
	}

	return address.ID, nil
}

func (b AddressRepository) UpdateAddressUser(db *sql.DB, address models.Address) (int64, error) {
	result, err := db.Exec("UPDATE address SET street=$1, city=$2, zip=$3, $user_id=$4 WHERE id=$5 RETURNING id",
		&address.Street, &address.City, &address.Zip, &address.UserID, &address.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b AddressRepository) RemoveAddressUser(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM address WHERE id = $1", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

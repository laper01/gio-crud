package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	stmt, err := config.DB.Prepare("INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return lastInsertId > 0

}

func Detail(id int) entities.Category {
	var category entities.Category
	row := config.DB.QueryRow("SELECT * FROM categories WHERE id = ?", id)
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}
	return category

}

func Update(id int, category entities.Category) bool {
	stmt, err := config.DB.Prepare("UPDATE categories SET name = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return rowsAffected > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}

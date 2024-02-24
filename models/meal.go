package models

import "foodiesbackend/db"

type Meal struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Image        string `json:"image"`
	Summary      string `json:"summary"`
	Instructions string `json:"instructions"`
	Creator_id   int64  `json:"creator_id"`
}

type ResponseMeal struct {
	Id            int64  `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Image         string `json:"image"`
	Summary       string `json:"summary"`
	Instructions  string `json:"instructions"`
	Creator_name  string `json:"creator"`
	Creator_email string `json:"creator_email"`
}

func GetAllMeals() ([]ResponseMeal, error) {
	query := `SELECT meals.id, title, slug, image, summary, instructions, username, email  FROM meals INNER JOIN users ON meals.creator_id = users.id`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []ResponseMeal

	for rows.Next() {
		var meal ResponseMeal
		if err := rows.Scan(&meal.Id, &meal.Title, &meal.Slug, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_name, &meal.Creator_email); err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}

	return meals, nil
}

func GetMealById(id int64) (Meal, error) {
	query := `SELECT * FROM meals WHERE id = $1`
	row := db.DB.QueryRow(query, id)

	var meal Meal

	err := row.Scan(&meal.Id, &meal.Title, &meal.Slug, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_id)

	if err != nil {
		return meal, err
	}

	return meal, nil
}

func GetResponseMealById(id int64) (ResponseMeal, error) {
	query := `SELECT meals.id, title, slug, image, summary, instructions, username, email  FROM meals INNER JOIN users ON meals.creator_id = users.id WHERE meals.id = $1`
	row := db.DB.QueryRow(query, id)

	var meal ResponseMeal

	err := row.Scan(&meal.Id, &meal.Title, &meal.Slug, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_name, &meal.Creator_email)

	if err != nil {
		return meal, err
	}

	return meal, nil
}

func (m *Meal) Save() error {
	query := `INSERT INTO meals (title, slug, image, summary, instructions, creator_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(m.Title, m.Slug, m.Image, m.Summary, m.Instructions, m.Creator_id)
	err = result.Scan(&m.Id)

	return err
}

func (m *Meal) Update() error {
	query := `UPDATE meals SET title=$1, slug=$2, image=$3, summary=$4, instructions=$5 WHERE id=$6`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.Title, m.Slug, m.Image, m.Summary, m.Instructions, m.Id)

	if err != nil {
		return err
	}

	return err
}

func (m *Meal) Delete() error {
	query := `DELETE FROM meals WHERE id=$1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.Id)

	return err
}

func GetMealsByCreatorId(id int64) ([]ResponseMeal, error) {
	query := `SELECT meals.id, title, slug, image, summary, instructions, username, email  FROM meals INNER JOIN users ON meals.creator_id = users.id WHERE meals.creator_id = $1`

	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var meals []ResponseMeal

	for rows.Next() {
		var meal ResponseMeal
		if err := rows.Scan(&meal.Id, &meal.Title, &meal.Slug, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_name, &meal.Creator_email); err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}

	return meals, nil
}

package models

import "foodiesbackend/db"

type Meal struct {
	Id            int64  `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Image         string `json:"image"`
	Summary       string `json:"summary"`
	Instructions  string `json:"instructions"`
	Creator       string `json:"creator"`
	Creator_email string `json:"creator_email"`
}

func GetAllMeals() ([]Meal, error) {
	rows, err := db.DB.Query("SELECT * FROM meals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []Meal

	for rows.Next() {
		var meal Meal
		if err := rows.Scan(&meal.Id, &meal.Title, &meal.Slug, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator, &meal.Creator_email); err != nil {
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

	err := row.Scan(&meal.Id, &meal.Title, &meal.Slug, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator, &meal.Creator_email)

	if err != nil {
		return meal, err
	}

	return meal, nil
}

func (m *Meal) Save() error {
	query := `INSERT INTO meals (title, slug, image, summary, instructions, creator, creator_email) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(m.Title, m.Slug, m.Image, m.Summary, m.Instructions, m.Creator, m.Creator_email)
	err = result.Scan(&m.Id)

	return err
}

func (e *Meal) Update() error {
	query := `UPDATE meals SET title=$1, slug=$2, image=$3, summary=$4, instructions=$5, creator=$6, creator_email=$7 WHERE id=$8`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Slug, e.Image, e.Summary, e.Instructions, e.Creator, e.Creator_email, e.Id)

	if err != nil {
		return err
	}

	return err
}

func (e *Meal) Delete() error {
	query := `DELETE FROM meals WHERE id=$1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	return err
}

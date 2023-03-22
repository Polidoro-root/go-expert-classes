package main

import (
	"context"
	"database/sql"

	"github.com/Polidoro-root/go-expert-classes/15_sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	/* err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Description", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	*/

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{ID: "49290f3e-f5f6-4539-be1a-fe3f8d956356", Name: "Updated", Description: sql.NullString{String: "Updated"}})

	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "49290f3e-f5f6-4539-be1a-fe3f8d956356")

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}

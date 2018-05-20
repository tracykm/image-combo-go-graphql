package main

import (
	"database/sql"
	"log"

	"github.com/graphql-go/graphql"
)

type SrcImage struct {
	ID       int    `json:"id"`
	Url      string `json:"url"`
	Size     int    `json:"size"`
	Category string `json:"type"`
}

var srcType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SrcImage",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"size": &graphql.Field{
				Type: graphql.Int,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func insertSrcImage(db *sql.DB, url string, size int, category string) SrcImage {
	sqlStatement := `
	INSERT INTO src_images (url, size, category)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, url, size, category).Scan(&id)
	if err != nil {
		panic(err)
	}
	return SrcImage{id, url, size, category}
}

func querySrcImages(db *sql.DB) []*SrcImage {
	rows, err := db.Query("SELECT * FROM src_images")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bks := make([]*SrcImage, 0)
	for rows.Next() {
		bk := new(SrcImage)
		err := rows.Scan(&bk.ID, &bk.Url, &bk.Size, &bk.Category)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// for _, bk := range bks {
	// 	fmt.Print(bk.ID, bk.Email)
	// }
	return bks
}

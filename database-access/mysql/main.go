// Package main
// GO SQL database drivers https://github.com/golang/go/wiki/SQLDrivers
package main

import (
	"fmt"
	"github.com/denglitong/golang/database-access/mysql/basic"
	"log"
)

// showBasicCRUD show Create, Read, Update and Delete database operations.
func showBasicCRUD() {
	basic.OpenDB()
	defer basic.CloseDB()
	basic.PingDB()

	// basic query for multiple rows
	albums, err := basic.AlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Albums found: %v\n", albums)

	// basic query for single row
	alb, err := basic.AlbumById(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// basic add data
	albId, err := basic.AddAlbum(basic.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albId)

	alb.Price *= 2
	row, err := basic.UpdateAlbum(alb)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Album updated row: %v\n", row)
}

// main
// DB_USER=root DB_PASSWORD=12345678 go run .
func main() {
	showBasicCRUD()
}

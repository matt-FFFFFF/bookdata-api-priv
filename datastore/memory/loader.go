package memory

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/matt-FFFFFF/bookstore"
)

// LoadData reads a CSV file passed in as io.Reader and returns a slice of bookstore.Book pointers
func LoadData(r io.Reader) ([]*bookstore.Book, error) {
	reader := csv.NewReader(r)

	ret := make([]*bookstore.Book, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		averageRating, _ := strconv.ParseFloat(row[3], 64)
		numPages, _ := strconv.Atoi(row[7])
		ratings, _ := strconv.Atoi(row[8])
		reviews, _ := strconv.Atoi(row[9])

		if err != nil {
			log.Println(err)
		}
		book := &bookstore.Book{
			BookID:        uuid.New(),
			Title:         row[1],
			Authors:       []string{row[2]},
			AverageRating: averageRating,
			ISBN:          row[4],
			ISBN13:        row[5],
			LanguageCode:  row[6],
			NumPages:      numPages,
			Ratings:       ratings,
			Reviews:       reviews,
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, book)
	}
	return ret, nil
}

package model

import (
	"context"
	"google.golang.org/appengine/datastore"
)

// Wine model
type Wine struct {
	Name string
	Year int
}

// GetWines returns wine mode list
func GetWines(c context.Context) ([]Wine, error) {
	q := datastore.NewQuery("Wine").Limit(10)
	wines := make([]Wine, 0, 10)
	if _, err := q.GetAll(c, &wines); err != nil {
		return nil, err
	}

	return wines, nil
}
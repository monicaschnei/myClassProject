package util

import (
	"database/sql"
	"math/rand"
)

// create random Ids
func CreateRandomIds(min, max int64) sql.NullInt64 {
	randomId := min + rand.Int63n(max-min+1)
	return sql.NullInt64{Int64: randomId, Valid: true}
}

package api

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToPgtypeText(s string) (pgtype.Text, error) {
	var txt pgtype.Text
	err := txt.Scan(s)
	return txt, err
}

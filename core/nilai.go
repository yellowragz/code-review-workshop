package core

import (
	"context"

	"github.com/tokopedia/code-review-workshop/errors"
	"github.com/tokopedia/code-review-workshop/tracer"
)

type NilaiSiswa struct {
	NIM     int64  `json:"nim"      db:"nim"`
	MapelID int64  `json:"mapel_id" db:"mapel_id"`
	Nilai   int64  `json:"nili"     db:"nili"`
	Sekolah string `json:"sekolah"  db:"sekolah"`
	Kelas   string `json:"kelas"    db:"kelas"`
}

func GetNilaiSiswaByID(ctx context.Context, nim, mapel_id int64) (result []NilaiSiswa, err error) {
	span, ctx := tracer.StartSpanFromContext(ctx)
	defer span.Finish()

	err = stmt.GetNilaiSiswaByID.Select(&result, mapel_id, nim)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}

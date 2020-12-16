package module

import (
	"context"

	nilai "github.com/tokopedia/code-review-workshop/core/nilai"
	"github.com/tokopedia/tokopoints/errors"
	"github.com/tokopedia/tokopoints/tracer"
)

func GetNilaiSiswa(ctx context.Context, nim int64, sekolah, kelas string, mapel_id int64) (resp GetNilaiSiswaResponse, err error) {
	span, ctx := tracer.StartSpanFromContext(ctx)
	defer span.Finish()

	if nim != 0 && mapel_id != 0 {
		nilai, err := nilai.GetNilaiSiswaByID(ctx, nim, mapel_id)
		if err != nil {
			return resp, errors.AddTrace(err)
		}

		detailSekolah, err := nilai.GetDetailSekolah(ctx, nilai.Sekolah[0])
		if err != nil {
			return resp, errors.AddTrace(err)
		}

		resp = mappingGetNilaiSiswa(nilai, detailSekolah)
	} else if sekolah != "semua" {
		nilai, err := nilai.GetNilaiSiswaBySekolah(ctx, sekolah)
		if err != nil {
			return resp, errors.AddTrace(err)
		}

		detailSekolah, err := nilai.GetDetailSekolah(ctx, nilai.Sekolah[0])
		if err != nil {
			return resp, errors.AddTrace(err)
		}

		resp = mappingGetNilaiSiswa(nilai, detailSekolah)
	} else if sekolah == "semua" {
		nilai, err := nilai.GetNilaiSiswaSemuaSekolah(ctx)
		if err != nil {
			return resp, errors.AddTrace(err)
		}

		detailSekolah, err := nilai.GetDetailSekolah(ctx, nilai.Sekolah[0])
		if err != nil {
			return resp, errors.AddTrace(err)
		}

		resp = mappingGetNilaiSiswa(nilai, detailSekolah)
	}

	return resp, nil
}

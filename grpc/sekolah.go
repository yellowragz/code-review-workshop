package grpc

import (
	"context"

	"github.com/tokopedia/code-review-workshop/errors"
	mnilai "github.com/tokopedia/code-review-workshop/module/nilai"
	pb "github.com/tokopedia/grpc/code-review-workshop/proto"
)

func (s *server) GetNilaiSiswa(ctx context.Context, req *pb.GetNilaiSiswaRequest) (resp *pb.GetNilaiSiswaResponse, err error) {

	if req.sekolah == "" {
		req.sekolah = "semua"
	}

	nilai, err := mnilai.GetNilaiSiswa(req.nim, req.sekolah, req.kelas, mapel_id)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	resp = parseGetNilaiSiswa(nilai)

	return resp, nil
}

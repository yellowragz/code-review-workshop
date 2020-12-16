package core

import (
	"context"
	"log"

	"github.com/tokopedia/code-review-workshop/database"
	redislib "github.com/tokopedia/code-review-workshop/redis"
	"github.com/tokopedia/sqlt"
)

type PreparedStatements struct {
	GetNilaiSiswaByID *sqlt.Stmtx
}

var (
	db    *sqlt.DB
	stmt  PreparedStatements
	redis *redislib.RedisStore // redis module
)

const (
	// Collection queries
	QueryGetNilaiSiswaByID = `select nim, mapel_id, nilai, sekolah, kelas
                            	from nilai_siswa
								where mapel_id = $1 and nim = $2;`
)

func Init() {
	var err error
	db, err = database.Get(database.PromoCatalog)
	if err != nil {
		log.Printf("Cannot connect to DB. %+v", err)
	}

	stmt = PreparedStatements{
		GetNilaiSiswaByID: database.Preparex(context.Background(), db, QueryGetNilaiSiswaByID),
	}

	redis, err = redislib.Get(redislib.PromoCatalog)
	if err != nil {
		log.Fatal("Cannot connect to redis")
	}
}

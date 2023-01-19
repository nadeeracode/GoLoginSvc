package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/nadeeracode/GoLoginSvc/util"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewSever(config, store)
	require.NoError(t, err)

	return server

}

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())

}

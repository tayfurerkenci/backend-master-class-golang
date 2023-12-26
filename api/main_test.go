package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/tayfurerkenci/backend-master-class-golang/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	// config := util.Config{
	// 	TokenSymmetricKey:   util.RandomString(32),
	// 	AccessTokenDuration: time.Minute,
	// }

	server := NewServer(store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tayfurerkenci/backend-master-class-golang/util"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoToken(t *testing.T) {
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	// Create a new Paseto token with a symmetric key.
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Create a new token.
	token, err := maker.CreateToken(payload.Username, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Create a new maker with a different symmetric key.
	invalidMaker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Try to verify the token with the wrong key.
	payload, err = invalidMaker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorContains(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

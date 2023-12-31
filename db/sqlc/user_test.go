package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tayfurerkenci/backend-master-class-golang/util"
)

// createRandomAccount creates a new account with random data
// does not have 'Test' prefix, so it won't be run by 'go test'
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	fetchedUser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)

	require.Equal(t, user.Username, fetchedUser.Username)
	require.Equal(t, user.HashedPassword, fetchedUser.HashedPassword)
	require.Equal(t, user.FullName, fetchedUser.FullName)
	require.Equal(t, user.Email, fetchedUser.Email)

	// Check that the user's createdAt time is within 1 second of the fetchedUser's createdAt time.
	require.WithinDuration(t, user.CreatedAt, fetchedUser.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")

	require.WithinDuration(t, user.PasswordChangedAt, fetchedUser.PasswordChangedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tayfurerkenci/backend-master-class-golang/util"
)

// createRandomAccount creates a new account with random data
// does not have 'Test' prefix, so it won't be run by 'go test'
func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, account.ID, fetchedAccount.ID)
	require.Equal(t, account.Owner, fetchedAccount.Owner)
	require.Equal(t, account.Balance, fetchedAccount.Balance)
	require.Equal(t, account.Currency, fetchedAccount.Currency)

	// Check that the account's createdAt time is within 1 second of the fetchedAccount's createdAt time.
	require.WithinDuration(t, account.CreatedAt, fetchedAccount.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)

	// Check that the account's createdAt time is within 1 second of the updatedAccount's createdAt time.
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	// Create 10 random accounts
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

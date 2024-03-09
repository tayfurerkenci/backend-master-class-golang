package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tayfurerkenci/simple-bank/backend/util"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	fetchedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedEntry)

	require.Equal(t, entry.ID, fetchedEntry.ID)
	require.Equal(t, entry.AccountID, fetchedEntry.AccountID)
	require.Equal(t, entry.Amount, fetchedEntry.Amount)

	// Check that the Entry's createdAt time is within 1 second of the fetchedEntry's createdAt time.
	require.WithinDuration(t, entry.CreatedAt, fetchedEntry.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, entry.ID, updatedEntry.ID)
	require.Equal(t, entry.AccountID, updatedEntry.AccountID)
	require.Equal(t, arg.Amount, updatedEntry.Amount)

	// Check that the Entry's createdAt time is within 1 second of the updatedEntry's createdAt time.
	require.WithinDuration(t, entry.CreatedAt, updatedEntry.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	deletedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedEntry)
}

func TestListEntrys(t *testing.T) {
	account := createRandomAccount(t)
	// Create 10 random Entrys
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	Entrys, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Entrys, 5)

	for _, Entry := range Entrys {
		require.NotEmpty(t, Entry)
	}
}

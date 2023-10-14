package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tayfurerkenci/backend-master-class-golang/util"
)

func createRandomTransfer(t *testing.T, sender Account, receiver Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: sender.ID,
		ToAccountID:   receiver.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, sender.ID, transfer.FromAccountID)
	require.Equal(t, receiver.ID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	senderAccount := createRandomAccount(t)
	receiverAccount := createRandomAccount(t)
	createRandomTransfer(t, senderAccount, receiverAccount)
}

func TestGetTransfer(t *testing.T) {
	senderAccount := createRandomAccount(t)
	receiverAccount := createRandomAccount(t)
	transfer := createRandomTransfer(t, senderAccount, receiverAccount)
	fetchedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransfer)

	require.Equal(t, transfer.ID, fetchedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, fetchedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, fetchedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, fetchedTransfer.Amount)

	// Check that the Transfer's createdAt time is within 1 second of the fetchedTransfer's createdAt time.
	require.WithinDuration(t, transfer.CreatedAt, fetchedTransfer.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

func TestUpdateTransfer(t *testing.T) {
	senderAccount := createRandomAccount(t)
	receiverAccount := createRandomAccount(t)
	transfer := createRandomTransfer(t, senderAccount, receiverAccount)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, transfer.ID, updatedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, arg.Amount, updatedTransfer.Amount)

	// Check that the Transfer's createdAt time is within 1 second of the updatedTransfer's createdAt time.
	require.WithinDuration(t, transfer.CreatedAt, updatedTransfer.CreatedAt, 1*time.Second,
		"Time calculation is not within expected range.")
}

func TestDeleteTransfer(t *testing.T) {
	senderAccount := createRandomAccount(t)
	receiverAccount := createRandomAccount(t)
	transfer := createRandomTransfer(t, senderAccount, receiverAccount)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	deletedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedTransfer)
}

func TestListTransfers(t *testing.T) {
	senderAccount := createRandomAccount(t)
	receiverAccount := createRandomAccount(t)

	// Create 10 random Transfers
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, senderAccount, receiverAccount)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	Transfers, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Transfers, 5)

	for _, Transfer := range Transfers {
		require.NotEmpty(t, Transfer)
	}
}

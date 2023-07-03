package db

import (
	"context"
	"testing"
	"time"

	"github.com/jatin-dua/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	amount := util.RandomMoney()

	transfer, err := testQueries.CreateTransfer(context.Background(), CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	})

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.Equal(t, amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	amount := util.RandomMoney()

	transfer, err := testQueries.CreateTransfer(context.Background(), CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	})

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)

	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		amount := util.RandomMoney()

		transfer, err := testQueries.CreateTransfer(context.Background(), CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		})

		require.NoError(t, err)
		require.NotEmpty(t, transfer)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

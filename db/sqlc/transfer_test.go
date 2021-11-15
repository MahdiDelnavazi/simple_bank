package db

import (
	"context"
	"github.com/mahdidl/simple_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateTransfer(t *testing.T) Transfer {

	account1, account2 := createRandomAccount(t), createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.WithinDuration(t, transfer.CreatedAt, account1.CreatedAt, time.Second)

	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	CreateTransfer(t)
}

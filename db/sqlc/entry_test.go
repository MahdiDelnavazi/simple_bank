package db

import (
	"context"
	"github.com/mahdidl/simple_bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomEntry(t *testing.T) Entry {

	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestQueries_GetEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.ID, entry2.ID)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
	require.Equal(t, entry.Amount, entry2.Amount)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{AccountID: 32}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

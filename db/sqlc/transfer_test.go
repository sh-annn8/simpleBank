package db

import (
	"testing"
	"context"
	"time"
	"database/sql"

	"simpleBank/util"
	"github.com/stretchr/testify/require"
)



// 生成一个数据随机的测试用例，避免测试相互依赖
func createRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: util.RandomInt(1, 10),
		ToAccountID: util.RandomInt(1, 100),
		Amount: util.RandomAmount(),
	}

	entry, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.FromAccountID, entry.FromAccountID)
	require.Equal(t, arg.ToAccountID, entry.ToAccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	entry1 := createRandomTransfer(t)
	entry2, err := testQueries.GetTransfer(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.FromAccountID, entry2.FromAccountID)
	require.Equal(t, entry1.ToAccountID, entry2.ToAccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	entry1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID: entry1.ID,
		Amount: util.RandomAmount(),
	}

	err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	entry2, err := testQueries.GetTransfer(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.FromAccountID, entry2.FromAccountID)
	require.Equal(t, entry1.ToAccountID, entry2.ToAccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	entry1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetTransfer(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

/* TestListTransfers
   为数据库增加 10 条记录
   取出 5 条记录
   查看 5 条记录是否为空
*/
func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	// 取第 5-10 条记录
	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	entrys, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)

	for _, transfer := range entrys {
		require.NotEmpty(t, transfer)
	}
}


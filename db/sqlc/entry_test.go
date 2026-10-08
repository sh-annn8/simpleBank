package db

import (
	"testing"
	"context"
	"time"

	"simpleBank/util"
	"github.com/stretchr/testify/require"
)

var n = 10

// 生成一个数据随机的测试用例，避免测试相互依赖
func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomAmount(),
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
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

/* TestListEntrys
   为数据库增加 10 条记录
   取出 5 条记录
   查看 5 条记录是否为空
*/
func TestListEntrys(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	// 取第 5-10 条记录
	arg := ListEntrysParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}

	entrys, err := testQueries.ListEntrys(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)

	for _, entry := range entrys {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}

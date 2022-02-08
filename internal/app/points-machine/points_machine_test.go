package points_machine

import "testing"

func Test_Accrual(t *testing.T) {
	// test case
	t.Run("success accrual", func(t *testing.T) {
		userId := 1
		amount := uint(100)

		InitAccount(userId)
		Accrual(userId, amount)

		if accountsMap[userId] != amount {
			// terminates this test case
			t.Fatal("account amount is not equal to test amount")
		}
	})
}

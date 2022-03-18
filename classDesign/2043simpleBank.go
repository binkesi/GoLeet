package classdesign

// https://leetcode-cn.com/problems/simple-bank-system/

type Bank struct {
	balance []int64
}

func ConstructorBank(balance []int64) Bank {
	return Bank{balance: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	n := len(this.balance)
	if account1 <= n+1 && account2 <= n+1 && this.balance[account1-1] >= money {
		this.balance[account1-1] -= money
		this.balance[account2-1] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	n := len(this.balance)
	if account <= n+1 {
		this.balance[account-1] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	n := len(this.balance)
	if account <= n+1 && this.balance[account-1] >= money {
		this.balance[account-1] -= money
		return true
	}
	return false
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

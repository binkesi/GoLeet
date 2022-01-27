package weekgame

// https://leetcode-cn.com/contest/weekly-contest-263/problems/simple-bank-system/

type Bank struct {
	customers []int64
}

func Constructor(balance []int64) Bank {
	return Bank{customers: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 <= 0 || account2 <= 0 || account1 > len(this.customers) || account2 > len(this.customers) || money > this.customers[account1-1] {
		return false
	}
	this.customers[account1-1] -= money
	this.customers[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account <= 0 || account > len(this.customers) {
		return false
	} else {
		this.customers[account-1] += money
		return true
	}
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account <= 0 || account > len(this.customers) || money > this.customers[account-1] {
		return false
	}
	this.customers[account-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

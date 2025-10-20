package homework02

import (
	"fmt"

	"gorm.io/gorm"
)

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
// transactions 表（包含字段 id 主键，
// from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，
// 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

// Accounts 的表
type Accounts struct {
	gorm.Model
	Balance int64 `gorm:"not null;default:0"`
}

// Transactions 的表
type Transactions struct {
	gorm.Model
	ID            uint  `gorm:"primaryKey;autoIncrement"` // 主键，自增
	FromAccountID uint  `gorm:"not null"`                 // 转出账户ID
	ToAccountID   uint  `gorm:"not null"`                 // 转入账户ID
	Amount        int64 `gorm:"not null"`                 // 转账金额
}

// 需要先检查账户 A 的余额是否足够
func hasSufficientBalance(act *Accounts, amount int64) bool {
	return act.Balance >= amount

}

// 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元
func aToB(actA *Accounts, actB *Accounts, amount int64) {
	actA.Balance -= amount
	actB.Balance += amount
}

// 创建测试数据
func createTestData(db *gorm.DB) error {
	// 清空表
	db.Exec("DELETE FROM accounts")
	db.Exec("DELETE FROM transactions")

	// 创建测试账户
	accounts := []Accounts{
		{Balance: 200}, // 账户1: 2元
		{Balance: 500}, // 账户2: 5元
	}

	if err := db.Create(&accounts).Error; err != nil {
		return err
	}

	fmt.Println(" 测试数据创建完成")
	fmt.Printf("账户1余额: %d分\n", accounts[0].Balance)
	fmt.Printf("账户2余额: %d分\n", accounts[1].Balance)
	return nil
}

// 执行转账
func Run(db *gorm.DB) error {

	//db.AutoMigrate(&Accounts{})
	//db.AutoMigrate(&Transactions{})
	// 1. 创建事前数据
	//if err := createTestData(db); err != nil {
	//	return err
	//}
	//return nil
	// 2. 执行转账事务
	return transferMoney(db, 1, 2, 100) // 从账户1向账户2转账100分(1元)
}

func transferMoney(db *gorm.DB, fromID, toID uint, amount int64) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// 1. 查询转出账户
		var fromAccount Accounts
		err := tx.Where("id =?", fromID).First(&fromAccount).Error
		if err != nil {
			return fmt.Errorf("转出账户不存在: %w", err)
		}

		// 2. 在事务中，需要先检查账户 A 的余额是否足够，
		if !hasSufficientBalance(&fromAccount, 100) {
			return fmt.Errorf("账户A的余额不足够: %w", err)
		}

		// 3. 查询转入账户
		var toAccount Accounts
		err = tx.Where("id =?", toID).First(&toAccount).Error
		if err != nil {
			return fmt.Errorf("转出账户不存在: %w", err)
		}

		// 4. 更新余额
		err = tx.Model(&fromAccount).Update("balance", fromAccount.Balance-amount).Error
		if err != nil {
			return err
		}
		err = tx.Model(&toAccount).Update("balance", toAccount.Balance+amount).Error
		if err != nil {
			return err
		}

		// 5. 记录交易
		transactions := Transactions{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		err = tx.Create(&transactions).Error
		if err != nil {
			return err
		}

		fmt.Printf(" 转账成功: 账户%d -> 账户%d, 金额%d分\n", fromID, toID, amount)
		return nil

	})
}

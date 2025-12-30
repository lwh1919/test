package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	"log"
	"time"
)

func scanAndCancelOrders(db *sql.DB) {
	// 1. 计算当前时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 2. 执行核心SQL，查询需要关闭的订单
	// 使用 FOR UPDATE SKIP LOCKED 可以避免在分布式环境下重复处理（如果数据库支持）
	rows, err := db.Query(`
        SELECT id FROM orders 
        WHERE status = 'pending' AND action_time < ? 
        ORDER BY action_time 
        LIMIT 400
        FOR UPDATE SKIP LOCKED
    `, now)
	if err != nil {
		log.Println("查询订单失败:", err)
		return
	}
	defer rows.Close()

	var orderIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Println("扫描订单ID失败:", err)
			continue
		}
		orderIDs = append(orderIDs, id)
	}

	// 3. 逐条更新订单状态（在实际应用中，这里可能还需要释放库存等逻辑）
	for _, id := range orderIDs {
		_, err := db.Exec("UPDATE orders SET status = 'cancelled' WHERE id = ?", id)
		if err != nil {
			log.Printf("更新订单 %d 状态失败: %v\n", id, err)
		} else {
			log.Printf("订单 %d 已超时自动取消\n", id)
		}
	}
	log.Printf("本次扫描完成，共处理 %d 个订单\n", len(orderIDs))
}

func main() {
	// 连接Docker中的MySQL
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/your_database?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建Cron定时任务
	c := cron.New()
	// 每5秒执行一次
	c.AddFunc("*/5 * * * * *", func() { scanAndCancelOrders(db) })
	c.Start()

	// 保持主程序运行
	select {}
}

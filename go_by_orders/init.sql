USE your_database;

CREATE TABLE orders (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        user_id INT,
                        amount DECIMAL(10, 2),
                        status VARCHAR(20) NOT NULL DEFAULT 'pending', -- 状态: pending, paid, cancelled
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        action_time DATETIME
);
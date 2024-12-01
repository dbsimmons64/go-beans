package database

import (
	"database/sql"
	"log"
)

func Create_txn_table(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    txn_date DATE NOT NULL,
    who TEXT NOT NULL,
    description TEXT,
    payee TEXT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    category TEXT,
    inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	data := `
	INSERT INTO transactions (txn_date, who, description, payee, amount, category, inserted_at, updated_at)
VALUES 
('2024-11-25', 'Alice', 'Lunch meeting', 'Local Bistro', 25.50, 'Food', '2024-11-25 12:00:00', '2024-11-25 12:00:00'),
('2024-11-26', 'Bob', 'Office supplies', 'Stationery World', 49.99, 'Supplies', '2024-11-26 09:30:00', '2024-11-26 09:30:00'),
('2024-11-27', 'Charlie', 'Ride to airport', 'City Taxi Co.', 32.75, 'Transport', '2024-11-27 14:15:00', '2024-11-27 14:15:00'),
('2024-11-28', 'Diana', 'Monthly subscription', 'Streaming Service', 15.99, 'Entertainment', '2024-11-28 10:00:00', '2024-11-28 10:00:00'),
('2024-11-29', 'Edward', 'Dinner with clients', 'Steakhouse Grill', 120.00, 'Business', '2024-11-29 20:30:00', '2024-11-29 20:30:00');
`

	_, err = db.Exec(data)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
}

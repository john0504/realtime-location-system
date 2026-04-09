package db

import "log"

func InitTables() {
	log.Println("Initializing DB tables...")

	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS landmarks (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		latitude DOUBLE PRECISION NOT NULL,
		longitude DOUBLE PRECISION NOT NULL,
		radius INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)

	if err != nil {
		log.Println("Create table error:", err)
		return
	}

	log.Println("Table ensured")

	_, err = DB.Exec(`
	INSERT INTO landmarks (name, latitude, longitude, radius)
	VALUES 
	('台中火車站', 24.1367, 120.6850, 300),
	('逢甲夜市', 24.1789, 120.6465, 300)
	ON CONFLICT (name) 
	DO UPDATE SET 
		latitude = EXCLUDED.latitude,
		longitude = EXCLUDED.longitude,
		radius = EXCLUDED.radius;
	`)

	if err != nil {
		log.Println("Insert data error:", err)
		return
	}

	log.Println("Seed data ensured")
}

package types

import (
	"database/sql"
)

func CreateTables(db *sql.DB) error {
	stmt := `	
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) UNIQUE NOT NULL,
			email VARCHAR(255) NOT NULL,
			password CHAR(60),
			token TEXT,
			expires DATETIME
		);
		
		CREATE TABLE IF NOT EXISTS snippets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			user_name VARCHAR(255) NOT NULL,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			created DATETIME NOT NULL,
			status TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id) 
		);
		CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snippet_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		author_name TEXT NOT NULL,
		content TEXT NOT NULL,
		created DATETIME NOT NULL,
		FOREIGN KEY (snippet_id) REFERENCES snippets(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS reactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snippet_id INTEGER,
		comment_id INTEGER,
		user_id INTEGER NOT NULL,
		type TEXT NOT NULL,
		UNIQUE (snippet_id, comment_id, user_id, type),
		FOREIGN KEY (snippet_id) REFERENCES snippets(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS notifications (
		id INTEGER PRIMARY KEY,
		receiver_id INTEGER NOT NULL,
		author_name TEXT NOT NULL,
		action_type TEXT NOT NULL,
		snippet_id INTEGER NOT NULL,
		comment_id INTEGER DEFAULT 0,
		timestamp DATETIME NOT NULL,
		is_read INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snippet_id INTEGER NOT NULL,
		category TEXT NOT NULL,
		FOREIGN KEY (snippet_id) REFERENCES snippets(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS roles (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	role TEXT NOT NULL,
	FOREIGN KEY (username) REFERENCES users(username)
);
CREATE TABLE IF NOT EXISTS requests (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INT,
	username TEXT NOT NULL,
	status TEXT,
	FOREIGN KEY (username) REFERENCES users(username)
);
CREATE TABLE IF NOT EXISTS reports (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	snippet_id INT NOT NULL,
	snippet_title TEXT NOT NULL,
	moderator_id INT NOT NULL,
	report_type TEXT NOT NULL,
	complete TEXT NOT NULL,
	response TEXT NOT NULL,
	FOREIGN KEY (snippet_id) REFERENCES snippets(id)
);
	`

	_, err := db.Exec(stmt)

	return err
}

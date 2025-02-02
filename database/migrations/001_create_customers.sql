CREATE TABLE IF NOT EXISTS Customer (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT unique,
    email TEXT unique,
    password TEXT
);

import sqlite3

conn = sqlite3.connect('./database/database.db')
cursor = conn.cursor()

with open('./database/migrations/001_create_customers.sql', 'r') as f:
    create_sql = f.read()
cursor.executescript(create_sql)

with open('./database/migrations/002_insert_customers.sql', 'r') as f:
    insert_sql = f.read()
cursor.executescript(insert_sql)

conn.commit()
conn.close()


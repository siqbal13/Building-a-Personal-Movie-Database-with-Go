import sqlite3
import pandas as pd

# Path to the SQLite database
db_path = 'imdb_database.db'

# Connect to the SQLite database (it will be created if it does not exist)
conn = sqlite3.connect(db_path)
cursor = conn.cursor()

# Drop existing tables if they exist
drop_schema = """
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS directors_genres;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS movies_genres;
DROP TABLE IF EXISTS roles;
"""
cursor.executescript(drop_schema)
conn.commit()

# Create tables based on the inspected schema
schema = """
CREATE TABLE actors (
    id INTEGER PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    gender TEXT
);

CREATE TABLE directors (
    id INTEGER PRIMARY KEY,
    first_name TEXT,
    last_name TEXT
);

CREATE TABLE directors_genres (
    director_id INTEGER,
    genre TEXT,
    prob REAL,
    PRIMARY KEY (director_id, genre)
);

CREATE TABLE movies (
    id INTEGER PRIMARY KEY,
    name TEXT,
    year INTEGER,
    rank REAL
);

CREATE TABLE movies_genres (
    movie_id INTEGER,
    genre TEXT,
    PRIMARY KEY (movie_id, genre)
);

CREATE TABLE roles (
    role_id INTEGER PRIMARY KEY,
    actor_id INTEGER,
    movie_id INTEGER,
    role TEXT
);
"""

# Execute the schema creation
cursor.executescript(schema)
conn.commit()

print("Database schema created successfully")

# Function to load CSV data into SQLite table
def load_csv_to_sqlite(csv_file_path, table_name, conn):
    # Read the CSV file into a DataFrame
    df = pd.read_csv(csv_file_path, on_bad_lines='skip')
    
    # Write the data into the SQLite table
    df.to_sql(table_name, conn, if_exists='append', index=False)

# Paths to the CSV files (update the paths if necessary)
csv_files = {
    "actors": "IMDb/IMDB-actors.csv",
    "directors": "IMDb/IMDB-directors.csv",
    "directors_genres": "IMDb/IMDB-directors_genres.csv",
    "movies": "IMDb/IMDB-movies.csv",
    "movies_genres": "IMDb/IMDB-movies_genres.csv",
    "roles": "IMDb/IMDB-roles.csv"
}

# Load data from each CSV file into its respective table
for table_name, csv_file_path in csv_files.items():
    load_csv_to_sqlite(csv_file_path, table_name, conn)

# Verify by listing the number of records in each table
table_records = {}
for table_name in csv_files.keys():
    cursor.execute(f"SELECT COUNT(*) FROM {table_name}")
    table_records[table_name] = cursor.fetchone()[0]

# Commit and close the connection
conn.commit()
conn.close()

print("Data loaded into database successfully")
print("Record counts:", table_records)

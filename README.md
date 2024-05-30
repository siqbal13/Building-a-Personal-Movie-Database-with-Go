
# Personal Movie Database with Go

## Overview

This project demonstrates the use of Go with a cgo-free SQLite library to create a personal movie database. The data is sourced from the Northwestern University IMDb Data Files.

## Setup and Installation

1. **Download and Extract Data Files**:
   - Download the zip archive from the [Northwestern site](https://arch.library.northwestern.edu/concern/datasets/3484zh40n?locale=en).
   - Extract the files and place them in the `IMDb` directory of this project.

2. **Database Schema and Data Loading**:
   - The database schema is created based on the structure of the CSV files.
   - Data is loaded into the SQLite database from the CSV files using the `setup_database.py` script.

3. **Go Application**:
   - Install the required Go SQLite library: `go get modernc.org/sqlite`.
   - Run the Go application to perform queries on the database.

## Database Schema

The SQLite database schema consists of the following tables:
- `actors`
- `directors`
- `directors_genres`
- `movies`
- `movies_genres`
- `roles`

## Usage

Run the Python script to set up the database:

```sh
python setup_database.py
```

Run the Go application to interact with the database:

```sh
go run main.go
```

## Personal Collection Integration

To extend this database, you can add a table for your personal movie collection. This table could include the following columns:
- `movie_id` (reference to `movies` table)
- `location` (e.g., shelf, digital, etc.)
- `personal_rating` (your rating of the movie)

## Purpose and Use of Personal Movie Database

The personal movie database serves to organize and keep track of your movie collection. It allows for:
- Easy querying of movie information.
- Tracking where each movie is located.
- Storing personal ratings and notes on movies.

### Potential User Interactions

A useful Go movie application might include:
- A web interface for browsing the movie collection.
- Functionality to add, update, or delete movies from the personal collection.
- Filtering and searching based on different criteria (e.g., genre, rating, location).

### Advantages Over IMDb

While IMDb provides extensive movie data, this personal movie database allows for:
- Customization specific to your collection.
- Personal notes and ratings.
- A tailored interface that fits your preferences and needs.

## Future Enhancements

### Database Enhancements

- Add more tables and relationships, such as user reviews or watch history.
- Incorporate external data sources for additional movie metadata.

### Further Application Development

- Develop a recommendation system based on personal ratings.
- Integrate with third-party APIs to fetch the latest movie information.
- Enhance the user interface for better usability.

### Personal Movie Recommendation System

A recommendation system could leverage:
- Personal ratings and reviews.
- Viewing history and preferences.
- Machine learning algorithms to suggest movies you might enjoy.

By continuously enhancing the database and application, you can create a comprehensive and personalized movie management system that goes beyond what IMDb offers.

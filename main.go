package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "modernc.org/sqlite"
)

// Monte Carlo simulation to estimate the value of pi
func monteCarloPi(numPoints int) float64 {
	insideCircle := 0

	for i := 0; i < numPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			insideCircle++
		}
	}

	return 4.0 * float64(insideCircle) / float64(numPoints)
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Start pprof for profiling
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Open the SQLite database
	db, err := sql.Open("sqlite", "imdb_database.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test the connection and perform a simple query
	rows, err := db.Query("SELECT name FROM movies LIMIT 10")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	log.Println("Top 10 Movies:")
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		log.Println(name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error with rows: %v", err)
	}

	// Run Monte Carlo simulation
	numPoints := 1000000
	start := time.Now()
	piEstimate := monteCarloPi(numPoints)
	elapsed := time.Since(start)
	log.Printf("Monte Carlo simulation with %d points estimated pi as %.6f in %s\n", numPoints, piEstimate, elapsed)

	// Prevent the program from exiting immediately
	select {}
}

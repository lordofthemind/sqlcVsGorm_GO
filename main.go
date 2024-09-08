package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/repositories"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
	"github.com/lordofthemind/sqlcVsGorm_GO/pkgs"
	"golang.org/x/exp/rand"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BenchmarkResult holds the result of a benchmark.
type BenchmarkResult struct {
	Repository string
	Operation  string
	Duration   time.Duration
}

// Create a map to keep track of used emails
var usedEmails = map[string]bool{}

// createRandomAuthor generates a random author for testing.
func createRandomAuthor(rng *rand.Rand) (string, sql.NullString, string, sql.NullTime) {
	name := fmt.Sprintf("Author%d", rng.Intn(1000))
	bio := sql.NullString{String: fmt.Sprintf("Bio%d", rng.Intn(1000)), Valid: true}

	// Generate a unique email
	var email string
	for {
		email = fmt.Sprintf("author%d@example.com", rng.Intn(1000))
		if !usedEmails[email] {
			usedEmails[email] = true
			break
		}
	}

	dateOfBirth := sql.NullTime{Time: time.Now().AddDate(-rng.Intn(60), 0, 0), Valid: true}
	return name, bio, email, dateOfBirth
}

// benchmarkCreate runs the CreateAuthor benchmark.
func benchmarkCreate(repo repositories.AuthorRepository, repoName string, count int, rng *rand.Rand) BenchmarkResult {
	start := time.Now()
	for i := 0; i < count; i++ {
		name, bio, email, dateOfBirth := createRandomAuthor(rng)
		_, err := repo.CreateAuthor(context.Background(), name, bio, email, dateOfBirth)
		if err != nil {
			log.Fatalf("[%s] Failed to create author: %v", repoName, err)
		}
	}
	return BenchmarkResult{Repository: repoName, Operation: "CreateAuthor", Duration: time.Since(start)}
}

// benchmarkGet runs the GetAuthor benchmark.
func benchmarkGet(repo repositories.AuthorRepository, repoName string, count int) BenchmarkResult {
	start := time.Now()
	for i := 0; i < count; i++ {
		id := int32(rand.Intn(100)) // Using random int32 ID for testing purposes
		_, err := repo.GetAuthor(context.Background(), id)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("[%s] Failed to get author: %v", repoName, err)
		}
	}
	return BenchmarkResult{Repository: repoName, Operation: "GetAuthor", Duration: time.Since(start)}
}

// benchmarkList runs the ListAuthors benchmark.
func benchmarkList(repo repositories.AuthorRepository, repoName string) BenchmarkResult {
	start := time.Now()
	_, err := repo.ListAuthors(context.Background())
	if err != nil {
		log.Fatalf("[%s] Failed to list authors: %v", repoName, err)
	}
	return BenchmarkResult{Repository: repoName, Operation: "ListAuthors", Duration: time.Since(start)}
}

// benchmarkDelete runs the DeleteAuthor benchmark.
func benchmarkDelete(repo repositories.AuthorRepository, repoName string, count int) BenchmarkResult {
	start := time.Now()
	for i := 0; i < count; i++ {
		id := int32(rand.Intn(100)) // Using random int32 ID for testing purposes
		err := repo.DeleteAuthor(context.Background(), id)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("[%s] Failed to delete author: %v", repoName, err)
		}
	}
	return BenchmarkResult{Repository: repoName, Operation: "DeleteAuthor", Duration: time.Since(start)}
}

// benchmarkUpdate runs the UpdateAuthor benchmark.
func benchmarkUpdate(repo repositories.AuthorRepository, repoName string, count int, rng *rand.Rand) BenchmarkResult {
	start := time.Now()
	for i := 0; i < count; i++ {
		id := int32(rand.Intn(100)) // Using random int32 ID for testing purposes
		name, bio, email, dateOfBirth := createRandomAuthor(rng)
		err := repo.UpdateAuthor(context.Background(), id, name, bio, email, dateOfBirth)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("[%s] Failed to update author: %v", repoName, err)
		}
	}
	return BenchmarkResult{Repository: repoName, Operation: "UpdateAuthor", Duration: time.Since(start)}
}

// benchmarkGetAuthorsByBirthdateRange runs the GetAuthorsByBirthdateRange benchmark.
func benchmarkGetAuthorsByBirthdateRange(repo repositories.AuthorRepository, repoName string, startDate, endDate time.Time) BenchmarkResult {
	start := time.Now()
	_, err := repo.GetAuthorsByBirthdateRange(context.Background(), startDate, endDate)
	if err != nil {
		log.Fatalf("[%s] Failed to get authors by birthdate range: %v", repoName, err)
	}
	return BenchmarkResult{Repository: repoName, Operation: "GetAuthorsByBirthdateRange", Duration: time.Since(start)}
}

func performBenchmarks(sqlcRepo, gormRepo repositories.AuthorRepository) {
	// Define the number of test iterations
	testCount := 100 // Adjust the count as needed for your testing

	// Create a new random generator with a seed
	rng := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	// Create maps to store results for side-by-side comparison
	results := map[string]map[string]BenchmarkResult{
		"SQLC": {},
		"GORM": {},
	}

	// Define date range for GetAuthorsByBirthdateRange
	startDate := time.Now().AddDate(-5, 0, 0) // 5 years ago
	endDate := time.Now()

	// Run benchmarks for SQLC repository
	log.Println("Running benchmarks for SQLC repository...")
	results["SQLC"]["CreateAuthor"] = benchmarkCreate(sqlcRepo, "SQLC", testCount, rng)
	results["SQLC"]["GetAuthor"] = benchmarkGet(sqlcRepo, "SQLC", testCount)
	results["SQLC"]["ListAuthors"] = benchmarkList(sqlcRepo, "SQLC")
	results["SQLC"]["DeleteAuthor"] = benchmarkDelete(sqlcRepo, "SQLC", testCount)
	results["SQLC"]["UpdateAuthor"] = benchmarkUpdate(sqlcRepo, "SQLC", testCount, rng)
	results["SQLC"]["GetAuthorsByBirthdateRange"] = benchmarkGetAuthorsByBirthdateRange(sqlcRepo, "SQLC", startDate, endDate)

	// Run benchmarks for GORM repository
	log.Println("Running benchmarks for GORM repository...")
	results["GORM"]["CreateAuthor"] = benchmarkCreate(gormRepo, "GORM", testCount, rng)
	results["GORM"]["GetAuthor"] = benchmarkGet(gormRepo, "GORM", testCount)
	results["GORM"]["ListAuthors"] = benchmarkList(gormRepo, "GORM")
	results["GORM"]["DeleteAuthor"] = benchmarkDelete(gormRepo, "GORM", testCount)
	results["GORM"]["UpdateAuthor"] = benchmarkUpdate(gormRepo, "GORM", testCount, rng)
	results["GORM"]["GetAuthorsByBirthdateRange"] = benchmarkGetAuthorsByBirthdateRange(gormRepo, "GORM", startDate, endDate)

	// Log results side by side and determine the winner
	var sqlcTotal, gormTotal time.Duration
	for operation := range results["SQLC"] {
		sqlcDuration := results["SQLC"][operation].Duration
		gormDuration := results["GORM"][operation].Duration
		difference := gormDuration - sqlcDuration

		sqlcTotal += sqlcDuration
		gormTotal += gormDuration

		// Determine winner for each operation
		winner := "SQLC"
		if gormDuration < sqlcDuration {
			winner = "GORM"
		}

		log.Printf("Operation: %s\n", operation)
		log.Printf("  SQLC Duration: %v\n", sqlcDuration)
		log.Printf("  GORM Duration: %v\n", gormDuration)
		log.Printf("  Difference    : %v\n", difference)
		log.Printf("  Winner        : %s\n", winner)
		log.Println()
	}

	// Summarize overall results
	log.Println("Summary:")
	log.Printf("  Total SQLC Time: %v\n", sqlcTotal)
	log.Printf("  Total GORM Time: %v\n", gormTotal)
	if sqlcTotal < gormTotal {
		log.Println("Overall Winner: SQLC")
	} else {
		log.Println("Overall Winner: GORM")
	}
}

func main() {
	// Set up logging
	logFile, err := pkgs.SetUpLogger("SqlcVsGorm.log")
	if err != nil {
		log.Fatalf("Failed to set up logger: %v", err)
	}
	defer logFile.Close()

	// Set up SQLC database connection
	sqlDB, err := sql.Open("postgres", "postgresql://postgres:postgresSqlcVsGormSecret@localhost:5434/SqlcVsGorm_SQLC?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to SQLC DB: %v", err)
	}
	defer sqlDB.Close()

	// Create a new instance of *sqlcgen.Queries
	queries := sqlcgen.New(sqlDB)

	// Create the SQLC repository using the *sqlcgen.Queries instance
	sqlcRepo := repositories.NewSQLCRepository(queries)

	// Set up GORM database connection
	gormDB, err := gorm.Open(postgres.Open("postgresql://postgres:postgresSqlcVsGormSecret@localhost:5434/SqlcVsGorm_GORM"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to GORM DB: %v", err)
	}

	// Auto migrate GORM schema
	err = gormDB.AutoMigrate(&sqlcgen.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Create the GORM repository
	gormRepo := repositories.NewGORMRepository(gormDB)

	// Perform benchmarks using the repositories
	performBenchmarks(sqlcRepo, gormRepo)
}

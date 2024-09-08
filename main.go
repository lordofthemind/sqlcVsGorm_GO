package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/repositories"
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

// logBenchmarkResult logs the result of each benchmark operation.
func logBenchmarkResult(result BenchmarkResult) {
	log.Printf("[%s] %s took %v\n", result.Repository, result.Operation, result.Duration)
}

// createRandomAuthor generates a random author for testing.
func createRandomAuthor(rng *rand.Rand) (string, sql.NullString) {
	name := fmt.Sprintf("Author%d", rng.Intn(1000))
	bio := sql.NullString{String: fmt.Sprintf("Bio%d", rng.Intn(1000)), Valid: true}
	return name, bio
}

// benchmarkCreate runs the CreateAuthor benchmark.
func benchmarkCreate(repo repositories.AuthorRepository, repoName string, count int, rng *rand.Rand) {
	start := time.Now()
	for i := 0; i < count; i++ {
		name, bio := createRandomAuthor(rng)
		_, err := repo.CreateAuthor(context.Background(), name, bio)
		if err != nil {
			log.Fatalf("[%s] Failed to create author: %v", repoName, err)
		}
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "CreateAuthor", Duration: time.Since(start)})
}

// benchmarkGet runs the GetAuthor benchmark.
func benchmarkGet(repo repositories.AuthorRepository, repoName string, count int) {
	start := time.Now()
	for i := int64(1); i <= int64(count); i++ {
		_, err := repo.GetAuthor(context.Background(), i)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("[%s] Failed to get author: %v", repoName, err)
		}
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "GetAuthor", Duration: time.Since(start)})
}

// benchmarkList runs the ListAuthors benchmark.
func benchmarkList(repo repositories.AuthorRepository, repoName string) {
	start := time.Now()
	_, err := repo.ListAuthors(context.Background())
	if err != nil {
		log.Fatalf("[%s] Failed to list authors: %v", repoName, err)
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "ListAuthors", Duration: time.Since(start)})
}

// benchmarkDelete runs the DeleteAuthor benchmark.
func benchmarkDelete(repo repositories.AuthorRepository, repoName string, count int) {
	start := time.Now()
	for i := int64(1); i <= int64(count); i++ {
		err := repo.DeleteAuthor(context.Background(), i)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("[%s] Failed to delete author: %v", repoName, err)
		}
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "DeleteAuthor", Duration: time.Since(start)})
}

func performBenchmarks(sqlcRepo, gormRepo repositories.AuthorRepository) {
	// Define the number of test iterations
	testCount := 100 // Adjust the count as needed for your testing

	// Create a new random generator with a seed
	rng := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	// Run benchmarks for SQLC repository
	log.Println("Running benchmarks for SQLC repository...")
	benchmarkCreate(sqlcRepo, "SQLC", testCount, rng)
	benchmarkGet(sqlcRepo, "SQLC", testCount)
	benchmarkList(sqlcRepo, "SQLC")
	benchmarkDelete(sqlcRepo, "SQLC", testCount)

	// Run benchmarks for GORM repository
	log.Println("Running benchmarks for GORM repository...")
	benchmarkCreate(gormRepo, "GORM", testCount, rng)
	benchmarkGet(gormRepo, "GORM", testCount)
	benchmarkList(gormRepo, "GORM")
	benchmarkDelete(gormRepo, "GORM", testCount)
}

func main() {
	// Set up database connections for SQLC and GORM
	// SQLC connection
	sqlDB, err := sql.Open("postgres", "postgresql://postgres:postgresSqlcVsGormSecret@localhost:5434/SqlcVsGorm_SQLC?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to SQLC DB: %v", err)
	}
	defer sqlDB.Close()

	sqlcRepo := repositories.NewSqlcAuthorRepository(sqlDB)

	// GORM connection
	gormDB, err := gorm.Open(postgres.Open("postgresql://postgres:postgresSqlcVsGormSecret@localhost:5434/SqlcVsGorm_GORM"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to GORM DB: %v", err)
	}

	gormRepo := repositories.NewGormAuthorRepository(gormDB)

	// Perform benchmarks using the repositories
	performBenchmarks(sqlcRepo, gormRepo)
}

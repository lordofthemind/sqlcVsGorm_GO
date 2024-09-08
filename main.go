package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/repositories"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
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
func createRandomAuthor(rng *rand.Rand) (name, bio string) {
	return fmt.Sprintf("Author%d", rng.Intn(1000)), fmt.Sprintf("Bio%d", rng.Intn(1000))
}

func benchmarkCreate(repo repositories.AuthorRepository, repoName string, count int, rng *rand.Rand) {
	start := time.Now()
	for i := 0; i < count; i++ {
		name, bio := createRandomAuthor(rng)
		_, err := repo.CreateAuthor(context.Background(), name, bio)
		if err != nil {
			log.Fatalf("Failed to create author: %v", err)
		}
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "CreateAuthor", Duration: time.Since(start)})
}

func benchmarkGet(repo repositories.AuthorRepository, repoName string, count int) {
	start := time.Now()
	for i := int64(1); i <= int64(count); i++ {
		_, err := repo.GetAuthor(context.Background(), i)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("Failed to get author: %v", err)
		}
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "GetAuthor", Duration: time.Since(start)})
}

func benchmarkList(repo repositories.AuthorRepository, repoName string) {
	start := time.Now()
	_, err := repo.ListAuthors(context.Background())
	if err != nil {
		log.Fatalf("Failed to list authors: %v", err)
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "ListAuthors", Duration: time.Since(start)})
}

func benchmarkDelete(repo repositories.AuthorRepository, repoName string, count int) {
	start := time.Now()
	for i := int64(1); i <= int64(count); i++ {
		err := repo.DeleteAuthor(context.Background(), i)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("Failed to delete author: %v", err)
		}
	}
	logBenchmarkResult(BenchmarkResult{Repository: repoName, Operation: "DeleteAuthor", Duration: time.Since(start)})
}

func main() {
	// Create a new random generator with a seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Setup SQLC repository
	sqlDB, err := sql.Open("postgres", "postgresql://postgres:postgresSqlcVsGormSecret@localhost:5434/SqlcVsGorm_SQLC?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer sqlDB.Close()

	// Initialize SQLC repository
	sqlcRepo := repositories.NewSqlcAuthorRepository(sqlDB)

	// Setup GORM repository
	gormDB, err := gorm.Open(postgres.Open("postgresql://postgres:postgresSqlcVsGormSecret@localhost:5434/SqlcVsGorm_GORM"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate GORM schema
	err = gormDB.AutoMigrate(&sqlcgen.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	gormRepo := repositories.NewGormAuthorRepository(gormDB)

	// Define the number of test iterations
	testCount := 100 // Adjust the count as needed for your testing

	// Run benchmarks for SQLC repository
	benchmarkCreate(sqlcRepo, "SQLC", testCount, rng)
	benchmarkGet(sqlcRepo, "SQLC", testCount)
	benchmarkList(sqlcRepo, "SQLC")
	benchmarkDelete(sqlcRepo, "SQLC", testCount)

	// Run benchmarks for GORM repository
	benchmarkCreate(gormRepo, "GORM", testCount, rng)
	benchmarkGet(gormRepo, "GORM", testCount)
	benchmarkList(gormRepo, "GORM")
	benchmarkDelete(gormRepo, "GORM", testCount)
}

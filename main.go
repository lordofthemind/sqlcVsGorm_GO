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

// createRandomAuthor generates a random author for testing.
func createRandomAuthor(rng *rand.Rand) (string, sql.NullString) {
	name := fmt.Sprintf("Author%d", rng.Intn(1000))
	bio := sql.NullString{String: fmt.Sprintf("Bio%d", rng.Intn(1000)), Valid: true}
	return name, bio
}

// benchmarkCreate runs the CreateAuthor benchmark.
func benchmarkCreate(repo repositories.AuthorRepository, repoName string, count int, rng *rand.Rand) BenchmarkResult {
	start := time.Now()
	for i := 0; i < count; i++ {
		name, bio := createRandomAuthor(rng)
		_, err := repo.CreateAuthor(context.Background(), name, bio)
		if err != nil {
			log.Fatalf("[%s] Failed to create author: %v", repoName, err)
		}
	}
	return BenchmarkResult{Repository: repoName, Operation: "CreateAuthor", Duration: time.Since(start)}
}

// benchmarkGet runs the GetAuthor benchmark.
func benchmarkGet(repo repositories.AuthorRepository, repoName string, count int) BenchmarkResult {
	start := time.Now()
	for i := int64(1); i <= int64(count); i++ {
		_, err := repo.GetAuthor(context.Background(), i)
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
	for i := int64(1); i <= int64(count); i++ {
		err := repo.DeleteAuthor(context.Background(), i)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("[%s] Failed to delete author: %v", repoName, err)
		}
	}
	return BenchmarkResult{Repository: repoName, Operation: "DeleteAuthor", Duration: time.Since(start)}
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

	// Run benchmarks for SQLC repository
	log.Println("Running benchmarks for SQLC repository...")
	results["SQLC"]["CreateAuthor"] = benchmarkCreate(sqlcRepo, "SQLC", testCount, rng)
	results["SQLC"]["GetAuthor"] = benchmarkGet(sqlcRepo, "SQLC", testCount)
	results["SQLC"]["ListAuthors"] = benchmarkList(sqlcRepo, "SQLC")
	results["SQLC"]["DeleteAuthor"] = benchmarkDelete(sqlcRepo, "SQLC", testCount)

	// Run benchmarks for GORM repository
	log.Println("Running benchmarks for GORM repository...")
	results["GORM"]["CreateAuthor"] = benchmarkCreate(gormRepo, "GORM", testCount, rng)
	results["GORM"]["GetAuthor"] = benchmarkGet(gormRepo, "GORM", testCount)
	results["GORM"]["ListAuthors"] = benchmarkList(gormRepo, "GORM")
	results["GORM"]["DeleteAuthor"] = benchmarkDelete(gormRepo, "GORM", testCount)

	// Log results side by side
	for operation := range results["SQLC"] {
		sqlcDuration := results["SQLC"][operation].Duration
		gormDuration := results["GORM"][operation].Duration
		difference := gormDuration - sqlcDuration

		log.Printf("Operation: %s\n", operation)
		log.Printf("  SQLC Duration: %v\n", sqlcDuration)
		log.Printf("  GORM Duration: %v\n", gormDuration)
		log.Printf("  Difference    : %v\n", difference)
		log.Println()
	}
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

	// Auto migrate GORM schema
	err = gormDB.AutoMigrate(&sqlcgen.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	gormRepo := repositories.NewGormAuthorRepository(gormDB)

	// Perform benchmarks using the repositories
	performBenchmarks(sqlcRepo, gormRepo)
}

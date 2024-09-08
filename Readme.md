
# sqlcVsGorm

**sqlcVsGorm** is a project aimed at comparing the performance of two popular Go database libraries, **SQLC** and **GORM**, when interacting with PostgreSQL. The project benchmarks various database operations, such as CRUD operations and complex queries, to provide insights into the performance characteristics of each library.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [What is SQLC?](#what-is-sqlc)
- [What is GORM?](#what-is-gorm)
- [Why Compare SQLC and GORM?](#why-compare-sqlc-and-gorm)
- [Installation](#installation)
- [How It Works](#how-it-works)
- [Performance Benchmarking](#performance-benchmarking)
- [Contributing](#contributing)

## Overview

This project benchmarks **SQLC** and **GORM** by implementing the same set of database operations—such as inserting, updating, retrieving, and deleting records—using both libraries. The performance results help determine which library is more efficient under various scenarios.

In this project, the same set of database operations, such as inserting records, reading data, updating, and deleting, is implemented using both SQLC and GORM. By switching between the two implementations in `main.go`, we can compare their performance in various scenarios.

## Project Structure

The project is organized as follows:

```bash
.
├── Changelog.md
├── Makefile
├── Readme.md
├── Result.txt
├── go.mod
├── go.sum
├── internals
│   ├── repositories
│   │   ├── AuthorRepositoryInterface.go
│   │   ├── GormAuthorRepository.go
│   │   └── SqlcAuthorRepository.go
│   └── sqlc
│       ├── migrations
│       │   ├── 20240908133254_authors_table.down.sql
│       │   └── 20240908133254_authors_table.up.sql
│       ├── queries
│       │   └── authors.sql
│       ├── schema
│       │   └── schema.sql
│       └── sqlcgen
│           ├── authors.sql.go
│           ├── db.go
│           ├── models.go
│           └── querier.go
├── logs
│   ├── 20240908_165429_SqlcVsGorm.log
│   ├── 20240908_165439_SqlcVsGorm.log
│   └── 20240908_165450_SqlcVsGorm.log
├── main.go
├── pkgs
│   └── Logger.go
├── sqlc.yaml
└── tree.txt
```

### Key Components

1. **Top-Level Files:**
   - **Changelog.md:** Keeps a history of changes, updates, and versioning for the project.
   - **Makefile:** Contains commands for automating tasks like building and running the project.
   - **Readme.md:** Provides instructions, setup guides, and documentation for the project.
   - **Result.txt:** Likely stores results or performance comparisons between GORM and SQLC, useful for benchmarking.
   - **go.mod & go.sum:** Standard Go files. `go.mod` defines the module and dependencies, and `go.sum` ensures consistent builds by locking dependency versions.

2. **`internals` Directory:**
   - This directory houses the core logic of your project, broken into two key parts: repositories and SQLC-related files.

   - **`repositories`:**  
     - **AuthorRepositoryInterface.go:** Defines the common interface for repository operations on the `Author` entity, abstracting database access.
     - **GormAuthorRepository.go:** Implements the `AuthorRepositoryInterface` using GORM for interacting with the database.
     - **SqlcAuthorRepository.go:** Implements the `AuthorRepositoryInterface` using SQLC, offering an alternative database interaction strategy.

   - **`sqlc`:**
     - **migrations:** Contains SQL migration files to manage database schema changes.
       - `.up.sql`: Scripts for applying changes to the database (e.g., creating tables).
       - `.down.sql`: Scripts for rolling back database changes.
     - **queries/authors.sql:** Contains SQL queries for interacting with the `Author` table, used by SQLC to generate Go code.
     - **schema/schema.sql:** Contains the overall database schema for defining the structure of the `Author` table.
     - **sqlcgen:** Generated files by SQLC from the queries and schema.
       - `authors.sql.go`: Generated Go code for interacting with the `Author` table based on the queries.
       - `db.go`: SQLC's generated file to manage database connections.
       - `models.go`: Defines Go structs representing the database schema (like `Author`).
       - `querier.go`: Generated code to interface with the database using the SQLC methods.

3. **`logs` Directory:**
   - Stores log files generated during project execution, particularly during benchmarking (e.g., SQLC vs. GORM performance comparisons).

4. **`main.go`:**
   - The entry point of the Go application. This file contains the logic to benchmark SQLC and GORM performance, running various operations (e.g., querying, updating) using both libraries, and recording the results.

5. **`pkgs` Directory:**
   - **Logger.go:** Custom logging utility to handle logging throughout the application. Likely logs events like benchmark results, errors, and status messages.

6. **`sqlc.yaml`:**
   - SQLC configuration file that specifies settings for generating code from SQL queries, such as database connections and output directories.

7. **`tree.txt`:**
   - A simple text file capturing the directory structure of the project. It can be useful for documentation or quick reference.

---

### Key Points:
- **Repository Pattern:** The project uses repository patterns (`AuthorRepositoryInterface.go`) to abstract database operations, making it easy to switch between GORM and SQLC.
- **SQLC vs. GORM Comparison:** The project is designed to compare SQLC's performance against GORM by running similar queries and measuring the time taken.
- **Generated SQLC Code:** The SQLC-generated files (`sqlcgen`) are automatically created based on the SQL queries and schemas provided, allowing for type-safe and efficient database interactions.


## What is SQLC?

**SQLC** is a Go library that helps developers write type-safe SQL queries. Instead of relying on an ORM to abstract SQL queries, SQLC generates Go code based directly on your SQL statements. This approach provides the following benefits:

- **Type Safety**: SQLC parses your SQL files and generates Go types, ensuring compile-time safety.
- **Direct SQL Control**: You have full control over your SQL queries, without the abstraction layer of an ORM.
- **Performance**: Since SQLC doesn’t involve a runtime ORM, there is no overhead, making it a highly efficient way to interact with the database.

In this project, SQLC is used by writing SQL files that describe the database queries, and then running the `sqlc` tool to generate Go code that interacts with PostgreSQL.

## What is GORM?

**GORM** is one of the most popular ORMs in the Go ecosystem. It abstracts SQL into Go-friendly functions and methods, making it easier to work with databases in Go. It comes with a range of powerful features:

- **Migrations**: GORM supports automatic migrations, simplifying database schema management.
- **Associations**: GORM handles complex relationships between models such as `One-to-One`, `One-to-Many`, and `Many-to-Many`.
- **Hooks and Callbacks**: GORM provides lifecycle hooks that can be used to execute logic before or after database operations.

Despite these features, GORM’s abstraction adds overhead, which can sometimes impact performance. The project explores how this performance compares to SQLC.

## Why Compare SQLC and GORM?

SQLC and GORM represent two different philosophies of database interaction in Go. 

- **SQLC**: SQL-first, where you write the SQL and let the tool generate type-safe Go code.
- **GORM**: Code-first, where you work with Go structs and let the library handle SQL generation and database interaction.

This comparison is useful to understand:

- The performance trade-offs between using an ORM and raw SQL.
- How complex queries affect both libraries.
- Which approach is more maintainable for different types of applications.

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/lordofthemind/sqlcVsGorm_GO.git
    cd sqlcVsGorm
    ```

2. **Install dependencies**:
    - **SQLC**:
      ```bash
      go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
      ```
    - **GORM**:
      ```bash
      go get -u gorm.io/gorm
      go get -u gorm.io/driver/postgres
      ```

3. **Set up PostgreSQL**: 
   Create a PostgreSQL database and update the connection settings in the project configuration.

4. **Generate SQLC code**:
   Run SQLC to generate Go code from the SQL files:
    ```bash
    sqlc generate
    ```

## How It Works

This project compares the performance of two ORM/SQL approaches in Go: **SQLC** and **GORM**, using the same set of operations. The goal is to benchmark how these libraries perform when executing database queries, such as inserting, updating, retrieving, and deleting records.

### Key Operations:

- **Insert records** (Create)
- **Fetch records** (Get)
- **List all records** (List)
- **Update records** (Update)
- **Delete records** (Delete)
- **Fetch records within a range** (GetAuthorsByBirthdateRange)

### Switching Between SQLC and GORM

The code uses an interface (`AuthorRepository`) that both SQLC and GORM implementations adhere to. This allows switching between the two libraries without changing the underlying business logic:

```go
type AuthorRepository interface {
    CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (int32, error)
    GetAuthor(ctx context.Context, id int32) (sqlcgen.Author, error)
    ListAuthors(ctx context.Context) ([]sqlcgen.Author, error)
    DeleteAuthor(ctx context.Context, id int32) error
    UpdateAuthor(ctx context.Context, id int32, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error
    GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error)
}
```

The project can easily swap the implementation (SQLC or GORM) by creating the appropriate repository instance (`NewSQLCRepository` or `NewGORMRepository`).
## Performance Benchmarking

The benchmarks measure the time taken to perform operations using SQLC and GORM. This includes single record insertions, updates, deletions, and complex queries such as fetching authors within a date range.

The benchmarks are run using custom logic implemented in the `performBenchmarks` function, which logs the performance of each library side by side:

- **CreateAuthor**: Measure the time taken to insert records.
- **GetAuthor**: Measure the time taken to retrieve specific records.
- **ListAuthors**: Measure the time taken to retrieve all records.
- **UpdateAuthor**: Measure the time taken to update records.
- **DeleteAuthor**: Measure the time taken to delete records.
- **GetAuthorsByBirthdateRange**: Measure the time taken to fetch records within a specific date range.

Each operation is benchmarked for both SQLC and GORM repositories, and the total time is logged, allowing for side-by-side comparison of performance.

### Running the Benchmarks

You can run the performance benchmarks with the following command:

```bash
go run main.go
```

This will log the execution times for SQLC and GORM for each operation, allowing you to determine which approach performs better in terms of speed. The results will be logged in a file (e.g., `SqlcVsGorm.log`).
### Performance Results

From our tests, we observed the following key points:
- **SQLC** generally outperformed **GORM** in more complex operations like creating, updating, and deleting records. For instance, SQLC's execution time was significantly lower for `CreateAuthor`, `DeleteAuthor`, and `UpdateAuthor` operations, as seen in the tests, with differences reaching over 200 milliseconds in some cases.
- **GORM**, despite being slower in these operations, performed better in certain read-heavy tasks such as `GetAuthor` and `ListAuthors`, where the framework's optimizations for object retrieval are more evident.
- The total execution time for SQLC across all benchmarks was consistently lower than that of GORM, suggesting that SQLC handles most operations more efficiently by eliminating ORM-related overhead.

While SQLC's performance advantage comes from being closer to raw SQL, GORM provides ease of use with features like migrations and automatic relationship handling, which can save development time at the cost of some performance. For more detailed analysis and benchmark logs, refer to logs.


## Contributing

Contributions to this project are highly encouraged! If you have any ideas for enhancing the benchmarking process, improving the codebase, or adding additional complex queries for further comparison, feel free to submit a pull request.

Whether it’s fixing bugs, optimizing performance, or suggesting new features, we welcome all forms of contribution. Here’s how you can contribute:

1. **Fork the Repository**: Create your own fork of the repository to work on.
2. **Create a Branch**: Make a feature branch for your changes (e.g., `feature-add-benchmarks`).
3. **Make Changes**: Implement your improvements or additions.
4. **Submit a Pull Request**: Open a pull request detailing the changes you’ve made and why they improve the project.

For more detailed instructions on how to contribute, please refer to the project's guidelines in the `CONTRIBUTING.md` file (if available). If you’re unsure about the direction of your contribution, feel free to open an issue to discuss it with the maintainers first.
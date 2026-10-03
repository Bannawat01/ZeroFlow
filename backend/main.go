package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	http.HandleFunc("GET /projects/{projectID}/tasks", func(w http.ResponseWriter, r *http.Request) {
		projectID := r.PathValue("projectID")
		fmt.Fprintf(w, "Tasks for project %s", projectID)
	})

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	rows, err := pool.Query(ctx,
		"SELECT id, title, status FROM tasks WHERE project_id = $1",
		1,
	)
	if err != nil {
		log.Fatal("อ่านงานไม่สำเร็จ: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var title, status string

		if err := rows.Scan(&id, &title, &status); err != nil {
			log.Fatal("อ่านข้อมูลในแถวไม่สำเร็จ: ", err)
		}

		fmt.Printf("%d | %s | %s\n", id, title, status)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("อ่านผลลัพธ์ไม่สำเร็จ: ", err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Refresh struct {
	Environment    string    `json:"environment"`
	SLNo           string    `json:"sl_no"`
	DateOfRefresh  time.Time `json:"date_of_refresh"`
	CodeBase       string    `json:"code_base"`
	ChangeTicket   string    `json:"change_ticket"`
	FreeField1     string    `json:"free_field_1"`
	FreeField2     string    `json:"free_field_2"`
	FreeField3     string    `json:"free_field_3"`
	DelFlag        bool      `json:"del_flg"`
	RModTime       string    `json:"r_mod_time"`
	RCreTime       string    `json:"r_cre_time"`
}

func main() {
	router := gin.Default()

	// Set up database connection
	db, err := sql.Open("postgres", "your_database_connection_string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Endpoint to get refresh details
	router.GET("/api/refreshes", func(c *gin.Context) {
		var refreshes []Refresh

		rows, err := db.Query("SELECT * FROM refresh_table")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch refresh details"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var refresh Refresh
			err := rows.Scan(
				&refresh.Environment,
				&refresh.SLNo,
				&refresh.DateOfRefresh,
				&refresh.CodeBase,
				&refresh.ChangeTicket,
				&refresh.FreeField1,
				&refresh.FreeField2,
				&refresh.FreeField3,
				&refresh.DelFlag,
				&refresh.RModTime,
				&refresh.RCreTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			refreshes = append(refreshes, refresh)
		}

		c.JSON(http.StatusOK, refreshes)
	})

	// Endpoint to create a new refresh
	router.POST("/api/refreshes", func(c *gin.Context) {
		var newRefresh Refresh
		if err := c.ShouldBindJSON(&newRefresh); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Perform validation or additional logic as needed

		_, err := db.Exec("INSERT INTO refresh_table (environment, sl_no, date_of_refresh, code_base, change_ticket, "+
			"free_field_1, free_field_2, free_field_3, del_flg, r_mod_time, r_cre_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			newRefresh.Environment, newRefresh.SLNo, newRefresh.DateOfRefresh, newRefresh.CodeBase, newRefresh.ChangeTicket,
			newRefresh.FreeField1, newRefresh.FreeField2, newRefresh.FreeField3, newRefresh.DelFlag,
			newRefresh.RModTime, newRefresh.RCreTime)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create refresh"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Refresh created successfully"})
	})

	router.Run(":8080")
}

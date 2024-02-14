package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
)

type PatchDetails struct {
	DeployFile    string `json:"deploy_file"`
	SrcFileSize   int    `json:"src_file_size"`
	PatchID       int    `json:"patch_id"`
	ReleaseID     int    `json:"release_id"`
	ModificationTime string `json:"mod_time"`
}

func main() {
	// Initialize Oracle DB connection
	db, err := sql.Open("godror", "your_oracle_connection_string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new Gin router
	r := gin.Default()

	// Define the GET endpoint to fetch the patch list based on the provided filename
	r.GET("/patch-list", func(c *gin.Context) {
		// Get the filename from the query parameters
		filename := c.Query("filename")

		// Query the database to fetch the patch list
		rows, err := db.Query(`
			SELECT deploy_file, src_file_size, patch_id, release_id, mod_time
			FROM apdmrhel.dfs_adm_rep
			WHERE deploy_file = $1 AND status = 'D'
			ORDER BY mod_time DESC
		`, filename)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patch list"})
			return
		}
		defer rows.Close()

		var patchList []PatchDetails

		// Iterate over the rows and populate the patchList
		for rows.Next() {
			var patch PatchDetails
			err := rows.Scan(&patch.DeployFile, &patch.SrcFileSize, &patch.PatchID, &patch.ReleaseID, &patch.ModificationTime)
			if err != nil {
				log.Println(err)
				continue
			}
			patchList = append(patchList, patch)
		}

		// Return the patchList as JSON
		c.JSON(http.StatusOK, patchList)
	})

	// Run the server on port 8080
	r.Run(":8080")
}

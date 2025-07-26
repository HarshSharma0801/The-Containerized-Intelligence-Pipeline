package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type ComputeResponse struct {
	Time        float64 `json:"time"`
	Operation   string  `json:"operation"`
	ProcessedAt string  `json:"processedAt"`
}

func main() {
	// Set Gin mode from environment variable
	if ginMode := os.Getenv("GIN_MODE"); ginMode != "" {
		gin.SetMode(ginMode)
	}

	r := gin.Default()

	// Health check endpoint (handles both GET and HEAD)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Go server is running",
		})
	})
	r.HEAD("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/compute", func(c *gin.Context) {

		timeNow := time.Now()
		calculatePrime()
		timeSince := time.Since(timeNow)

		response := ComputeResponse{
			Time:        timeSince.Seconds(),
			Operation:   "prime_calculation",
			ProcessedAt: time.Now().Format(time.RFC3339),
		}

		c.JSON(http.StatusOK, response)
	})

	// Get port from environment variable
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8086"
	}

	r.Run(":" + port)
}

func calculatePrime() {
	const MAX int64 = 100000
	var start int64 = 0
	var totalPrime int64 = 0
	const CONCURRENCY = 10
	var wg sync.WaitGroup

	for i := 1; i <= CONCURRENCY; i++ {
		wg.Add(1)
		go doBatch(fmt.Sprintf("Worker-%d", i), &wg, &start, MAX, &totalPrime)
	}
	wg.Wait()
	fmt.Printf("Total primes found: %d\n", totalPrime)
}

func CheckPrime(x int64, totalPrime *int64) {
	if x < 2 {
		return
	}
	if x == 2 {
		atomic.AddInt64(totalPrime, 1)
		return
	}
	if x%2 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		if x%int64(i) == 0 {
			return
		}
	}
	atomic.AddInt64(totalPrime, 1)
}

func doBatch(name string, wg *sync.WaitGroup, start *int64, max int64, totalPrime *int64) {
	defer wg.Done()

	startingTime := time.Now()
	for {
		x := atomic.AddInt64(start, 1)
		if x > max {
			break
		}
		CheckPrime(x, totalPrime)
	}
	fmt.Printf("Batch %s processed up to %d in %s\n", name, max, time.Since(startingTime))
}

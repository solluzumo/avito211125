package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"net/http"
// 	"sort"
// 	"sync"
// 	"time"
// )

// type PullRequest struct {
// 	PullRequestId   string `json:"pull_request_id"`
// 	PullRequestName string `json:"pull_request_name"`
// 	AuthorId        string `json:"author_id"`
// }

// type User struct {
// 	Id       string
// 	TeamName string
// }

// // helper для расчета percentiles
// func percentile(durations []float64, p float64) float64 {
// 	if len(durations) == 0 {
// 		return 0
// 	}
// 	sort.Float64s(durations)
// 	k := (p / 100) * float64(len(durations)-1)
// 	f := int(math.Floor(k))
// 	c := int(math.Ceil(k))
// 	if f == c {
// 		return durations[f]
// 	}
// 	return durations[f]*(float64(c)-k) + durations[c]*(k-float64(f))
// }

// // фиксированный список пользователей и команд (200 пользователей, 20 команд)
// func fixedUsers() []User {
// 	users := make([]User, 200)
// 	for i := 0; i < 200; i++ {
// 		users[i] = User{
// 			Id:       fmt.Sprintf("u%d", i+1),
// 			TeamName: fmt.Sprintf("team%d", (i%20)+1),
// 		}
// 	}
// 	return users
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	baseURL := "http://localhost:8080/pullRequest/create"
// 	users := fixedUsers()

// 	const (
// 		totalRequests = 1000 // общее число запросов
// 		rps           = 5    // запросов в секунду
// 	)

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	var successCount int
// 	var durations []float64

// 	client := &http.Client{Timeout: 10 * time.Second}
// 	interval := time.Second / time.Duration(rps) // пауза между запросами

// 	for i := 0; i < totalRequests; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			author := users[rand.Intn(len(users))]
// 			prID := fmt.Sprintf("p%d-%d", time.Now().UnixNano(), i)
// 			pr := PullRequest{
// 				PullRequestId:   prID,
// 				PullRequestName: fmt.Sprintf("Feature update %d", rand.Intn(10000)),
// 				AuthorId:        author.Id,
// 			}

// 			body, _ := json.Marshal(pr)
// 			start := time.Now()
// 			resp, err := client.Post(baseURL, "application/json", bytes.NewBuffer(body))
// 			duration := time.Since(start).Seconds() * 1000 // мс

// 			mu.Lock()
// 			durations = append(durations, duration)
// 			if err == nil && resp != nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
// 				successCount++
// 			}
// 			mu.Unlock()

// 			if resp != nil {
// 				resp.Body.Close()
// 			}
// 		}(i)

// 		time.Sleep(interval) // контролируем RPS
// 	}

// 	wg.Wait()

// 	// статистика
// 	total := len(durations)
// 	sum := 0.0
// 	for _, d := range durations {
// 		sum += d
// 	}
// 	avg := sum / float64(total)

// 	fmt.Printf("Total requests: %d\n", total)
// 	fmt.Printf("Successful requests: %d\n", successCount)
// 	fmt.Printf("Success rate: %.3f%%\n", float64(successCount)/float64(total)*100)
// 	fmt.Printf("Average latency: %.2f ms\n", avg)
// 	fmt.Printf("p50 latency: %.2f ms\n", percentile(durations, 50))
// 	fmt.Printf("p95 latency: %.2f ms\n", percentile(durations, 95))
// 	fmt.Printf("p99 latency: %.2f ms\n", percentile(durations, 99))
// }
//

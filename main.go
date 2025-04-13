package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Host struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Metric struct {
	Name string `json:"name"`
}

// 関数型の定義
type HostsFunc func(string) ([]Host, error)
type MetricsFunc func(string, string) ([]Metric, error)

// モック可能な変数
var getHosts HostsFunc = fetchHosts
var getMetrics MetricsFunc = fetchMetrics

func fetchHosts(apiKey string) ([]Host, error) {
	req, err := http.NewRequest("GET", "https://api.mackerelio.com/api/v0/hosts", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", apiKey)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, status code: %d", string(body), resp.StatusCode)
	}

	var result struct {
		Hosts []Host `json:"hosts"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Hosts, nil
}

func fetchMetrics(apiKey, hostID string) ([]Metric, error) {
	url := fmt.Sprintf("https://api.mackerelio.com/api/v0/hosts/%s/metrics/names", hostID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", apiKey)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, status code: %d", string(body), resp.StatusCode)
	}

	var result struct {
		Names []string `json:"names"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	metrics := make([]Metric, len(result.Names))
	for i, name := range result.Names {
		metrics[i] = Metric{Name: name}
	}

	return metrics, nil
}

func main() {
	apiKey := flag.String("apikey", "", "Mackerel API key")
	hostID := flag.String("host", "", "Host ID (optional)")
	flag.Parse()

	if *apiKey == "" {
		apiKey = stringPtr(os.Getenv("MACKEREL_APIKEY"))
		if *apiKey == "" {
			log.Fatal("API key is required. Set it with -apikey flag or MACKEREL_APIKEY environment variable")
		}
	}

	if *hostID == "" {
		hosts, err := getHosts(*apiKey)
		if err != nil {
			log.Fatalf("Failed to get hosts: %v", err)
		}

		fmt.Println("Available hosts:")
		for _, host := range hosts {
			fmt.Printf("ID: %s, Name: %s\n", host.ID, host.Name)
		}
		return
	}

	metrics, err := getMetrics(*apiKey, *hostID)
	if err != nil {
		log.Fatalf("Failed to get metrics: %v", err)
	}

	fmt.Printf("Metrics for host %s:\n", *hostID)
	for _, metric := range metrics {
		fmt.Println(metric.Name)
	}
}

func stringPtr(s string) *string {
	return &s
}

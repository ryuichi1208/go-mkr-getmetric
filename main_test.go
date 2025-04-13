package main

import (
	"net/http"
	"testing"
)

func TestGetHosts(t *testing.T) {
	// オリジナルの関数を保存
	originalGetHosts := getHosts

	// テスト終了時に元に戻す
	defer func() { getHosts = originalGetHosts }()

	// 関数をモック化
	getHosts = func(apiKey string) ([]Host, error) {
		if apiKey != "test-api-key" {
			return nil, nil
		}
		return []Host{
			{ID: "host1", Name: "test-host-1"},
			{ID: "host2", Name: "test-host-2"},
		}, nil
	}

	// テスト実行
	hosts, err := getHosts("test-api-key")
	if err != nil {
		t.Fatalf("Failed to get hosts: %v", err)
	}

	// 結果の確認
	if len(hosts) != 2 {
		t.Errorf("Expected 2 hosts, got %d", len(hosts))
	}
	if hosts[0].ID != "host1" || hosts[0].Name != "test-host-1" {
		t.Errorf("Host 1 data incorrect: %+v", hosts[0])
	}
	if hosts[1].ID != "host2" || hosts[1].Name != "test-host-2" {
		t.Errorf("Host 2 data incorrect: %+v", hosts[1])
	}
}

func TestGetMetrics(t *testing.T) {
	// オリジナルの関数を保存
	originalGetMetrics := getMetrics

	// テスト終了時に元に戻す
	defer func() { getMetrics = originalGetMetrics }()

	// 関数をモック化
	getMetrics = func(apiKey, hostID string) ([]Metric, error) {
		if apiKey != "test-api-key" || hostID != "host1" {
			return nil, nil
		}
		return []Metric{
			{Name: "metric1"},
			{Name: "metric2"},
			{Name: "metric3"},
		}, nil
	}

	// テスト実行
	metrics, err := getMetrics("test-api-key", "host1")
	if err != nil {
		t.Fatalf("Failed to get metrics: %v", err)
	}

	// 結果の確認
	if len(metrics) != 3 {
		t.Errorf("Expected 3 metrics, got %d", len(metrics))
	}
	expectedMetrics := []string{"metric1", "metric2", "metric3"}
	for i, m := range metrics {
		if m.Name != expectedMetrics[i] {
			t.Errorf("Expected metric %s, got %s", expectedMetrics[i], m.Name)
		}
	}
}

// ヘルパー関数
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsOptions(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	opts := &StatsOptions{
		StartDate:   "2025-01-01",
		EndDate:     "2025-01-31",
		Aggregation: "day",
		Limit:       100,
		Offset:      10,
	}

	path := "/stats"
	fullPath, err := client.AddOptions(path, opts)
	assert.NoError(t, err)
	assert.Contains(t, fullPath, "start_date=2025-01-01")
	assert.Contains(t, fullPath, "end_date=2025-01-31")
	assert.Contains(t, fullPath, "aggregated_by=day")
	assert.Contains(t, fullPath, "limit=100")
	assert.Contains(t, fullPath, "offset=10")
}

func TestStatMetrics(t *testing.T) {
	metrics := StatMetrics{
		Blocks:           5,
		BounceDrops:      2,
		Bounces:          3,
		Clicks:           100,
		DeferredDrops:    1,
		Delivered:        950,
		InvalidEmails:    4,
		Opens:            300,
		Processed:        1000,
		Requests:         1000,
		SpamReportDrops:  1,
		SpamReports:      2,
		UniqueClicks:     85,
		UniqueOpens:      250,
		UnsubscribeDrops: 1,
		Unsubscribes:     3,
	}

	assert.Equal(t, 5, metrics.Blocks)
	assert.Equal(t, 2, metrics.BounceDrops)
	assert.Equal(t, 3, metrics.Bounces)
	assert.Equal(t, 100, metrics.Clicks)
	assert.Equal(t, 1, metrics.DeferredDrops)
	assert.Equal(t, 950, metrics.Delivered)
	assert.Equal(t, 4, metrics.InvalidEmails)
	assert.Equal(t, 300, metrics.Opens)
	assert.Equal(t, 1000, metrics.Processed)
	assert.Equal(t, 1000, metrics.Requests)
	assert.Equal(t, 1, metrics.SpamReportDrops)
	assert.Equal(t, 2, metrics.SpamReports)
	assert.Equal(t, 85, metrics.UniqueClicks)
	assert.Equal(t, 250, metrics.UniqueOpens)
	assert.Equal(t, 1, metrics.UnsubscribeDrops)
	assert.Equal(t, 3, metrics.Unsubscribes)
}

func TestGlobalStat(t *testing.T) {
	stat := GlobalStat{
		Date: "2025-01-01",
		Stats: StatMetrics{
			Delivered: 1000,
			Opens:     300,
			Clicks:    100,
		},
	}

	assert.Equal(t, "2025-01-01", stat.Date)
	assert.Equal(t, 1000, stat.Stats.Delivered)
	assert.Equal(t, 300, stat.Stats.Opens)
	assert.Equal(t, 100, stat.Stats.Clicks)
}

func TestStatItem(t *testing.T) {
	item := StatItem{
		Name: "newsletter",
		Type: "category",
		Metrics: StatMetrics{
			Delivered: 500,
			Opens:     150,
			Clicks:    50,
		},
	}

	assert.Equal(t, "newsletter", item.Name)
	assert.Equal(t, "category", item.Type)
	assert.Equal(t, 500, item.Metrics.Delivered)
	assert.Equal(t, 150, item.Metrics.Opens)
	assert.Equal(t, 50, item.Metrics.Clicks)
}

func TestGetGlobalStats(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[
			{
				"date": "2025-01-01",
				"stats": {
					"delivered": 1000,
					"opens": 300,
					"clicks": 100
				}
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	stats, err := client.GetGlobalStats(context.TODO(), nil)
	assert.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, "2025-01-01", stats[0].Date)
	assert.Equal(t, 1000, stats[0].Stats.Delivered)
	assert.Equal(t, 300, stats[0].Stats.Opens)
	assert.Equal(t, 100, stats[0].Stats.Clicks)
}

func TestGetGlobalStats_WithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "2025-01-01", r.URL.Query().Get("start_date"))
		assert.Equal(t, "2025-01-31", r.URL.Query().Get("end_date"))
		assert.Equal(t, "day", r.URL.Query().Get("aggregated_by"))
		if _, err := fmt.Fprint(w, `[]`); err != nil {
			t.Fatal(err)
		}
	})

	opts := &StatsOptions{
		StartDate:   "2025-01-01",
		EndDate:     "2025-01-31",
		Aggregation: "day",
	}
	_, err := client.GetGlobalStats(context.TODO(), opts)
	assert.NoError(t, err)
}

func TestGetGlobalStats_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetGlobalStats(context.TODO(), nil)
	assert.Error(t, err)
}

func TestGetCategoryStats(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/categories/stats", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "newsletter,promotion", r.URL.Query().Get("categories"))
		if _, err := fmt.Fprint(w, `[
			{
				"date": "2025-01-01",
				"stats": [
					{
						"name": "newsletter",
						"type": "category",
						"metrics": {
							"delivered": 500,
							"opens": 150,
							"clicks": 50
						}
					}
				]
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	categories := []string{"newsletter", "promotion"}
	stats, err := client.GetCategoryStats(context.TODO(), categories, nil)
	assert.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, "2025-01-01", stats[0].Date)
	assert.Len(t, stats[0].Stats, 1)
	assert.Equal(t, "newsletter", stats[0].Stats[0].Name)
	assert.Equal(t, "category", stats[0].Stats[0].Type)
}

func TestGetCategoryStats_WithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/categories/stats", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "newsletter", r.URL.Query().Get("categories"))
		assert.Equal(t, "2025-01-01", r.URL.Query().Get("start_date"))
		assert.Equal(t, "2025-01-31", r.URL.Query().Get("end_date"))
		if _, err := fmt.Fprint(w, `[]`); err != nil {
			t.Fatal(err)
		}
	})

	categories := []string{"newsletter"}
	opts := &StatsOptions{
		StartDate: "2025-01-01",
		EndDate:   "2025-01-31",
	}
	_, err := client.GetCategoryStats(context.TODO(), categories, opts)
	assert.NoError(t, err)
}

func TestGetCategoryStats_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/categories/stats", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	categories := []string{"newsletter"}
	_, err := client.GetCategoryStats(context.TODO(), categories, nil)
	assert.Error(t, err)
}

func TestGetCategorySums(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/categories/stats/sums", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[
			{
				"date": "2025-01-01",
				"stats": [
					{
						"name": "newsletter",
						"type": "category",
						"metrics": {
							"delivered": 1000,
							"opens": 300,
							"clicks": 100
						}
					}
				]
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	stats, err := client.GetCategorySums(context.TODO(), nil)
	assert.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, "2025-01-01", stats[0].Date)
	assert.Len(t, stats[0].Stats, 1)
	assert.Equal(t, "newsletter", stats[0].Stats[0].Name)
}

func TestGetCategorySums_WithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/categories/stats/sums", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "2025-01-01", r.URL.Query().Get("start_date"))
		assert.Equal(t, "day", r.URL.Query().Get("aggregated_by"))
		if _, err := fmt.Fprint(w, `[]`); err != nil {
			t.Fatal(err)
		}
	})

	opts := &StatsOptions{
		StartDate:   "2025-01-01",
		Aggregation: "day",
	}
	_, err := client.GetCategorySums(context.TODO(), opts)
	assert.NoError(t, err)
}

func TestGetCategorySums_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/categories/stats/sums", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetCategorySums(context.TODO(), nil)
	assert.Error(t, err)
}

func TestGetSubuserStats(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "subuser1,subuser2", r.URL.Query().Get("subusers"))
		if _, err := fmt.Fprint(w, `[
			{
				"date": "2025-01-01",
				"stats": [
					{
						"name": "subuser1",
						"type": "subuser",
						"metrics": {
							"delivered": 500,
							"opens": 150,
							"clicks": 50
						}
					}
				]
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	subusers := []string{"subuser1", "subuser2"}
	stats, err := client.GetSubuserStats(context.TODO(), subusers, nil)
	assert.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, "2025-01-01", stats[0].Date)
	assert.Len(t, stats[0].Stats, 1)
	assert.Equal(t, "subuser1", stats[0].Stats[0].Name)
	assert.Equal(t, "subuser", stats[0].Stats[0].Type)
}

func TestGetSubuserStats_WithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "subuser1", r.URL.Query().Get("subusers"))
		assert.Equal(t, "2025-01-01", r.URL.Query().Get("start_date"))
		assert.Equal(t, "2025-01-31", r.URL.Query().Get("end_date"))
		if _, err := fmt.Fprint(w, `[]`); err != nil {
			t.Fatal(err)
		}
	})

	subusers := []string{"subuser1"}
	opts := &StatsOptions{
		StartDate: "2025-01-01",
		EndDate:   "2025-01-31",
	}
	_, err := client.GetSubuserStats(context.TODO(), subusers, opts)
	assert.NoError(t, err)
}

func TestGetSubuserStats_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	subusers := []string{"subuser1"}
	_, err := client.GetSubuserStats(context.TODO(), subusers, nil)
	assert.Error(t, err)
}

func TestGetSubuserSums(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats/sums", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[
			{
				"date": "2025-01-01",
				"stats": [
					{
						"name": "subuser1",
						"type": "subuser",
						"metrics": {
							"delivered": 1000,
							"opens": 300,
							"clicks": 100
						}
					}
				]
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	stats, err := client.GetSubuserSums(context.TODO(), nil)
	assert.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, "2025-01-01", stats[0].Date)
	assert.Len(t, stats[0].Stats, 1)
	assert.Equal(t, "subuser1", stats[0].Stats[0].Name)
}

func TestGetSubuserSums_WithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats/sums", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "2025-01-01", r.URL.Query().Get("start_date"))
		assert.Equal(t, "day", r.URL.Query().Get("aggregated_by"))
		if _, err := fmt.Fprint(w, `[]`); err != nil {
			t.Fatal(err)
		}
	})

	opts := &StatsOptions{
		StartDate:   "2025-01-01",
		Aggregation: "day",
	}
	_, err := client.GetSubuserSums(context.TODO(), opts)
	assert.NoError(t, err)
}

func TestGetSubuserSums_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats/sums", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubuserSums(context.TODO(), nil)
	assert.Error(t, err)
}

func TestGetSubuserMonthlyStats(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats/monthly", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[
			{
				"date": "2025-01-01",
				"stats": [
					{
						"name": "subuser1",
						"type": "subuser",
						"metrics": {
							"delivered": 30000,
							"opens": 9000,
							"clicks": 3000
						}
					}
				]
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	stats, err := client.GetSubuserMonthlyStats(context.TODO(), nil)
	assert.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, "2025-01-01", stats[0].Date)
	assert.Len(t, stats[0].Stats, 1)
	assert.Equal(t, "subuser1", stats[0].Stats[0].Name)
	assert.Equal(t, 30000, stats[0].Stats[0].Metrics.Delivered)
}

func TestGetSubuserMonthlyStats_WithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats/monthly", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "2025-01-01", r.URL.Query().Get("start_date"))
		assert.Equal(t, "day", r.URL.Query().Get("aggregated_by"))
		if _, err := fmt.Fprint(w, `[]`); err != nil {
			t.Fatal(err)
		}
	})

	opts := &StatsOptions{
		StartDate:   "2025-01-01",
		Aggregation: "day",
	}
	_, err := client.GetSubuserMonthlyStats(context.TODO(), opts)
	assert.NoError(t, err)
}

func TestGetSubuserMonthlyStats_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/stats/monthly", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubuserMonthlyStats(context.TODO(), nil)
	assert.Error(t, err)
}
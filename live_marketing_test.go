//go:build live

package sendgrid

import (
	"context"
	"testing"
)

func TestLiveMarketingSenders(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	list, err := client.GetMarketingSenders(ctx)
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetMarketingSenders: %v", err)
	}
	assertNoUnknownFields[[]*MarketingSender](t, rec)
	if len(list) == 0 {
		t.Log("Sender が 0 件のため GetMarketingSender は未実施")
		return
	}

	if _, err := client.GetMarketingSender(ctx, list[0].ID); err != nil {
		t.Fatalf("GetMarketingSender: %v", err)
	}
	assertNoUnknownFields[OutputGetMarketingSender](t, rec)
}

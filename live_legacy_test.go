//go:build live

package sendgrid

import (
	"context"
	"fmt"
	"slices"
	"testing"
)

func TestLiveLegacyListLifecycle(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	created, err := client.CreateLegacyList(ctx, &InputCreateLegacyList{Name: liveName("legacy-list")})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("CreateLegacyList: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("作成した Legacy List の ID が空です")
	}
	what := fmt.Sprintf("Legacy List %d", created.ID)
	deleteList := deleteOnce(t, what, func() error {
		return client.DeleteLegacyList(context.Background(), created.ID, nil)
	})
	assertNoUnknownFields[OutputCreateLegacyList](t, rec)
	if created.ID == 0 || created.RecipientCount != 0 {
		t.Fatalf("作成結果が想定外です: %+v", created)
	}

	got, err := client.GetLegacyList(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacyList: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacyList](t, rec)
	if got.Name != created.Name {
		t.Errorf("name = %q, want %q", got.Name, created.Name)
	}

	lists, err := client.GetLegacyLists(ctx)
	if err != nil {
		t.Fatalf("GetLegacyLists: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacyLists](t, rec)
	if !slices.ContainsFunc(lists.Lists, func(l *LegacyList) bool { return l.ID == created.ID }) {
		t.Errorf("一覧に作成したリスト %d がありません", created.ID)
	}

	recipients, err := client.GetLegacyListRecipients(ctx, created.ID, &InputGetLegacyListRecipients{Page: 1, PageSize: 100})
	if err != nil {
		t.Fatalf("GetLegacyListRecipients: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacyListRecipients](t, rec)
	if len(recipients.Recipients) != 0 || recipients.RecipientCount != 0 {
		t.Errorf("作成直後のリストに受信者がいます: %d 件", recipients.RecipientCount)
	}

	newName := liveName("legacy-list-updated")
	updated, err := client.UpdateLegacyList(ctx, created.ID, &InputUpdateLegacyList{Name: newName})
	if err != nil {
		t.Fatalf("UpdateLegacyList: %v", err)
	}
	assertNoUnknownFields[OutputUpdateLegacyList](t, rec)
	if updated.Name != newName {
		t.Errorf("UpdateLegacyList の戻り値: name = %q, want %q", updated.Name, newName)
	}
	got, err = client.GetLegacyList(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacyList: %v", err)
	}
	if got.Name != newName {
		t.Errorf("更新後の再取得: name = %q, want %q", got.Name, newName)
	}

	if err := deleteList(); err != nil {
		t.Fatalf("DeleteLegacyList: %v", err)
	}
	_, err = client.GetLegacyList(ctx, created.ID)
	assertDeleted(t, rec, err, what)
}

func TestLiveLegacyListBulkDelete(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	var ids []int64
	deleteLists := deleteOnce(t, "Legacy List (bulk)", func() error {
		if len(ids) == 0 {
			return nil
		}
		return client.DeleteLegacyLists(context.Background(), (*InputDeleteLegacyLists)(&ids))
	})
	for i := range 2 {
		created, err := client.CreateLegacyList(ctx, &InputCreateLegacyList{Name: liveName(fmt.Sprintf("legacy-list-bulk%d", i))})
		skipIfUnavailable(t, rec, err)
		if err != nil {
			t.Fatalf("CreateLegacyList: %v", err)
		}
		if created.ID == 0 {
			t.Fatal("作成した Legacy List の ID が空です")
		}
		ids = append(ids, created.ID)
	}

	if err := deleteLists(); err != nil {
		t.Fatalf("DeleteLegacyLists: %v", err)
	}
	for _, id := range ids {
		_, err := client.GetLegacyList(ctx, id)
		assertDeleted(t, rec, err, fmt.Sprintf("Legacy List %d", id))
	}
}

func TestLiveLegacySegmentLifecycle(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	// list_id を指定しないセグメントはアカウント全体の受信者が対象になるため、必ず空の試験用リストに限定する
	list, err := client.CreateLegacyList(ctx, &InputCreateLegacyList{Name: liveName("legacy-segment-list")})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("CreateLegacyList: %v", err)
	}
	if list.ID == 0 {
		t.Fatal("試験用リストの ID が空です")
	}
	deleteOnce(t, fmt.Sprintf("Legacy List %d", list.ID), func() error {
		return client.DeleteLegacyList(context.Background(), list.ID, nil)
	})

	conditions := []*LegacySegmentCondition{
		{Field: "email", Value: "example.com", Operator: "contains", AndOr: ""},
	}
	created, err := client.CreateLegacySegment(ctx, &InputCreateLegacySegment{
		Name:       liveName("legacy-segment"),
		ListID:     list.ID,
		Conditions: conditions,
	})
	if err != nil {
		t.Fatalf("CreateLegacySegment: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("作成した Legacy Segment の ID が空です")
	}
	what := fmt.Sprintf("Legacy Segment %d", created.ID)
	deleteSegment := deleteOnce(t, what, func() error {
		return client.DeleteLegacySegment(context.Background(), created.ID)
	})
	assertNoUnknownFields[OutputCreateLegacySegment](t, rec)
	if created.ID == 0 || created.ListID != list.ID {
		t.Fatalf("作成結果が想定外です: %+v", created)
	}

	got, err := client.GetLegacySegment(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacySegment: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacySegment](t, rec)
	if got.Name != created.Name || got.ListID != list.ID {
		t.Errorf("取得結果が想定外です: name = %q, list_id = %d", got.Name, got.ListID)
	}

	segments, err := client.GetLegacySegments(ctx)
	if err != nil {
		t.Fatalf("GetLegacySegments: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacySegments](t, rec)
	if !slices.ContainsFunc(segments.Segments, func(s *LegacySegment) bool { return s.ID == created.ID }) {
		t.Errorf("一覧に作成したセグメント %d がありません", created.ID)
	}

	newName := liveName("legacy-segment-updated")
	updated, err := client.UpdateLegacySegment(ctx, created.ID, &InputUpdateLegacySegment{
		Name:       newName,
		ListID:     list.ID,
		Conditions: conditions,
	})
	if err != nil {
		t.Fatalf("UpdateLegacySegment: %v", err)
	}
	assertNoUnknownFields[OutputUpdateLegacySegment](t, rec)
	if updated.Name != newName {
		t.Errorf("UpdateLegacySegment の戻り値: name = %q, want %q", updated.Name, newName)
	}
	got, err = client.GetLegacySegment(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacySegment: %v", err)
	}
	if got.Name != newName || got.ListID != list.ID {
		t.Errorf("更新後の再取得: name = %q, list_id = %d", got.Name, got.ListID)
	}

	recipients, err := client.GetLegacySegmentRecipients(ctx, created.ID, &InputGetLegacySegmentRecipients{Page: 1, PageSize: 100})
	if err != nil {
		t.Fatalf("GetLegacySegmentRecipients: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacySegmentRecipients](t, rec)
	if len(recipients.Recipients) != 0 {
		t.Errorf("空のリストに限定したセグメントに受信者がいます: %d 件", len(recipients.Recipients))
	}

	if err := deleteSegment(); err != nil {
		t.Fatalf("DeleteLegacySegment: %v", err)
	}
	_, err = client.GetLegacySegment(ctx, created.ID)
	assertDeleted(t, rec, err, what)
}

func TestLiveLegacyCampaignLifecycle(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	// 送信先 (list_ids / segment_ids) と送信者を指定しない下書きにすることで、誤って送信されない状態にする
	created, err := client.CreateLegacyCampaign(ctx, &InputCreateLegacyCampaign{
		Title:   liveName("legacy-campaign"),
		Subject: "kenzo0107/sendgrid live test",
	})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("CreateLegacyCampaign: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("作成した Legacy Campaign の ID が空です")
	}
	what := fmt.Sprintf("Legacy Campaign %d", created.ID)
	deleteCampaign := deleteOnce(t, what, func() error {
		return client.DeleteLegacyCampaign(context.Background(), created.ID)
	})
	assertNoUnknownFields[OutputCreateLegacyCampaign](t, rec)
	if created.ID == 0 || created.Status != "Draft" {
		t.Fatalf("作成結果が想定外です: id = %d, status = %q", created.ID, created.Status)
	}

	got, err := client.GetLegacyCampaign(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacyCampaign: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacyCampaign](t, rec)
	if got.Title != created.Title {
		t.Errorf("title = %q, want %q", got.Title, created.Title)
	}

	campaigns, err := client.GetLegacyCampaigns(ctx, &InputGetLegacyCampaigns{Limit: 100})
	if err != nil {
		t.Fatalf("GetLegacyCampaigns: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacyCampaigns](t, rec)
	if !slices.ContainsFunc(campaigns.Result, func(c *LegacyCampaign) bool { return c.ID == created.ID }) {
		t.Errorf("一覧に作成したキャンペーン %d がありません", created.ID)
	}

	newTitle := liveName("legacy-campaign-updated")
	newSubject := "kenzo0107/sendgrid live test updated"
	updated, err := client.UpdateLegacyCampaign(ctx, created.ID, &InputUpdateLegacyCampaign{Title: newTitle, Subject: newSubject})
	if err != nil {
		t.Fatalf("UpdateLegacyCampaign: %v", err)
	}
	assertNoUnknownFields[OutputUpdateLegacyCampaign](t, rec)
	if updated.Title != newTitle || updated.Subject != newSubject {
		t.Errorf("UpdateLegacyCampaign の戻り値: title = %q, subject = %q", updated.Title, updated.Subject)
	}
	got, err = client.GetLegacyCampaign(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacyCampaign: %v", err)
	}
	if got.Title != newTitle || got.Subject != newSubject {
		t.Errorf("更新後の再取得: title = %q, subject = %q", got.Title, got.Subject)
	}

	schedule, err := client.GetLegacyCampaignSchedule(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetLegacyCampaignSchedule: %v", err)
	}
	assertNoUnknownFields[OutputGetLegacyCampaignSchedule](t, rec)
	if schedule.SendAt != 0 {
		t.Errorf("未スケジュールの下書きに send_at = %d が設定されています", schedule.SendAt)
	}

	if err := deleteCampaign(); err != nil {
		t.Fatalf("DeleteLegacyCampaign: %v", err)
	}
	_, err = client.GetLegacyCampaign(ctx, created.ID)
	assertDeleted(t, rec, err, what)
}

func TestLiveLegacySenderRead(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	senders, err := client.GetLegacySenders(ctx)
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetLegacySenders: %v", err)
	}
	assertNoUnknownFields[[]*LegacySender](t, rec)

	if len(senders) == 0 {
		t.Log("Sender が 0 件のため GetLegacySender は未実施")
		return
	}

	id := senders[0].ID
	if _, err := client.GetLegacySender(ctx, id); err != nil {
		t.Fatalf("GetLegacySender(%d): %v", id, err)
	}
	assertNoUnknownFields[OutputGetLegacySender](t, rec)
}

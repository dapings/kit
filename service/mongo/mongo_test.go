package mongo

import "testing"

func TestRewriteKey(t *testing.T) {
	key := RewriteKey("tags")
	if key != "tags" {
		t.Error(key)
		return
	}

	key = RewriteKey("tags.public.pd")
	if key != "tags_public_pd" {
		t.Error(key)
		return
	}

	key = RewriteKey("$tags")
	if key != "_tags" {
		t.Error(key)
		return
	}
}

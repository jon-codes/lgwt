package blogposts_test

import (
	"slices"
	"testing"
	"testing/fstest"

	blogposts "github.com/jon-codes/lgwt-unit/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 1
Description: Description 1
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        "Hello\nWorld",
	}
	assertPost(t, got, want)
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if got.Title != want.Title {
		t.Errorf("got %+v, want %+v", got, want)
	}
	if got.Description != want.Description {
		t.Errorf("got %+v, want %+v", got, want)
	}

	if !slices.Equal(got.Tags, want.Tags) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

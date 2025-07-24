package main

import (
	"fmt"
	"log"

	"github.com/Pavel-Casp/ProxiApi/internal/client/jsonplaceholder"
)

func main() {
	client := jsonplaceholder.NewClient()

	// Test Posts
	posts, err := client.GetPosts()
	if err != nil {
		log.Fatalf("Failed to get posts: %v", err)
	}
	fmt.Printf("Got %d posts\n", len(posts))

	// Test single Post
	post, err := client.GetPostByID(1)
	if err != nil {
		log.Fatalf("Failed to get post: %v", err)
	}
	fmt.Printf("Post 1: %s\n", post.Title)

	// Test Users
	users, err := client.GetUsers()
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}
	fmt.Printf("Got %d users\n", len(users))

	// Test Comments
	comments, err := client.GetCommentsByPostID(1)
	if err != nil {
		log.Fatalf("Failed to get comments: %v", err)
	}
	fmt.Printf("Got %d comments for post 1\n", len(comments))

	// Test Post creation
	newPost := jsonplaceholder.Post{
		UserID: 1,
		Title:  "New Post",
		Body:   "This is a new post created with Resty",
	}

	createdPost, err := client.CreatePost(newPost)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}
	fmt.Printf("Created post with ID: %d\n", createdPost.ID)
}

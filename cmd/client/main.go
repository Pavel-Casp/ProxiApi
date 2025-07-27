package main

import (
	"fmt"
	"log"

	"github.com/Pavel-Casp/ProxiApi/internal/client/jsonplaceholder"
	"github.com/Pavel-Casp/ProxiApi/internal/config"
)

func main() {
	log.Println("Starting ProxiApi client...")

	// Load configuration
	cfg := config.Load()
	log.Printf("Configuration loaded: %+v\n", cfg)

	// Initialize client
	client := jsonplaceholder.NewClient(cfg)

	// Test API endpoints
	if err := testEndpoints(client); err != nil {
		log.Fatalf("API test failed: %v", err)
	}

	log.Println("All tests completed successfully")
}

func testEndpoints(client *jsonplaceholder.Client) error {
	// Test getting posts
	posts, err := client.GetPosts()
	if err != nil {
		return fmt.Errorf("GetPosts failed: %w", err)
	}
	log.Printf("Retrieved %d posts\n", len(posts))

	// Test getting single post
	post, err := client.GetPostByID(1)
	if err != nil {
		return fmt.Errorf("GetPostByID failed: %w", err)
	}
	log.Printf("Retrieved post #1: %q\n", post.Title)

	// Test creating post
	newPost := jsonplaceholder.Post{
		UserID: 1,
		Title:  "Test Post",
		Body:   "This is a test post created by ProxiApi",
	}

	createdPost, err := client.CreatePost(newPost)
	if err != nil {
		return fmt.Errorf("CreatePost failed: %w", err)
	}
	log.Printf("Created new post with ID: %d\n", createdPost.ID)

	return nil
}

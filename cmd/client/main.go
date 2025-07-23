package main

import (
	"fmt"
	"log"

	"github.com/Pavel-Casp/ProxiApi/internal/client/jsonplaceholder"
)

func main() {
	client := jsonplaceholder.NewClient()

	// Пример получения всех постов
	posts, err := client.GetPosts()
	if err != nil {
		log.Fatalf("Failed to get posts: %v", err)
	}

	fmt.Println("First 3 posts:")
	for i, post := range posts {
		if i >= 3 {
			break
		}
		fmt.Printf("%d: %s\n", post.ID, post.Title)
	}

	// Пример получения конкретного поста
	post, err := client.GetPostByID(1)
	if err != nil {
		log.Fatalf("Failed to get post: %v", err)
	}

	fmt.Printf("\nPost with ID 1:\nTitle: %s\nBody: %s\n", post.Title, post.Body)

	// Пример создания поста
	newPost := jsonplaceholder.Post{
		UserID: 1,
		Title:  "New Post",
		Body:   "This is a new post created by the client",
	}

	createdPost, err := client.CreatePost(newPost)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	fmt.Printf("\nCreated post:\nID: %d\nTitle: %s\nBody: %s\n",
		createdPost.ID, createdPost.Title, createdPost.Body)
}

package jsonplaceholder

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient() *Client {
	return &Client{
		client: resty.New().
			SetBaseURL("https://jsonplaceholder.typicode.com").
			SetHeader("Content-Type", "application/json").
			SetTimeout(10 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(1 * time.Second).
			SetRetryMaxWaitTime(5 * time.Second).
			AddRetryCondition(
				func(r *resty.Response, err error) bool {
					return r.StatusCode() >= 500 || err != nil
				},
			).
			OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
				fmt.Printf("Request: %s %s\n", req.Method, req.URL)
				return nil
			}).
			OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
				fmt.Printf("Response: %d %s\n", resp.StatusCode(), resp.Time())
				return nil
			}),
	}
}

// Posts
func (c *Client) GetPosts() ([]Post, error) {
	var posts []Post
	_, err := c.client.R().
		SetResult(&posts).
		Get("/posts")
	return posts, err
}

func (c *Client) GetPostByID(id int) (*Post, error) {
	var post Post
	_, err := c.client.R().
		SetResult(&post).
		Get("/posts/" + fmt.Sprint(id))
	return &post, err
}

func (c *Client) CreatePost(post Post) (*Post, error) {
	var result Post
	_, err := c.client.R().
		SetBody(post).
		SetResult(&result).
		Post("/posts")
	return &result, err
}

// Users
func (c *Client) GetUsers() ([]User, error) {
	var users []User
	_, err := c.client.R().
		SetResult(&users).
		Get("/users")
	return users, err
}

func (c *Client) GetUserByID(id int) (*User, error) {
	var user User
	_, err := c.client.R().
		SetResult(&user).
		Get("/users/" + fmt.Sprint(id))
	return &user, err
}

// Comments
func (c *Client) GetComments() ([]Comment, error) {
	var comments []Comment
	_, err := c.client.R().
		SetResult(&comments).
		Get("/comments")
	return comments, err
}

func (c *Client) GetCommentsByPostID(postID int) ([]Comment, error) {
	var comments []Comment
	_, err := c.client.R().
		SetResult(&comments).
		Get("/posts/" + fmt.Sprint(postID) + "/comments")
	return comments, err
}

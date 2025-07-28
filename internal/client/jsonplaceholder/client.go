package jsonplaceholder

import (
	"fmt"
	"log"

	"github.com/Pavel-Casp/ProxiApi/internal/config"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
	cfg    *config.Config
}

func NewClient(cfg *config.Config) *Client {
	client := resty.New().
		SetBaseURL(cfg.API.BaseURL).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetTimeout(cfg.API.RequestTimeout).
		SetRetryCount(cfg.Retry.MaxRetries).
		SetRetryWaitTime(cfg.Retry.WaitTime).
		SetRetryMaxWaitTime(cfg.Retry.MaxWait).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() >= 500 || err != nil
			},
		)

	if cfg.Debug.Enabled {
		client.SetDebug(true).
			OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
				log.Printf("[DEBUG] Request: %s %s\nBody: %v\n",
					req.Method, req.URL, req.Body)
				return nil
			}).
			OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
				log.Printf("[DEBUG] Response: %d %s\nBody: %s\n",
					resp.StatusCode(), resp.Time(), resp.String())
				return nil
			})
	}

	return &Client{
		client: client,
		cfg:    cfg,
	}
}

func (c *Client) GetPosts() ([]Post, error) {
	var posts []Post
	_, err := c.client.R().
		SetResult(&posts).
		Get("/posts")

	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	return posts, nil
}

func (c *Client) GetPostByID(id int) (*Post, error) {
	var post Post
	_, err := c.client.R().
		SetResult(&post).
		Get(fmt.Sprintf("/posts/%d", id))

	if err != nil {
		return nil, fmt.Errorf("failed to get post %d: %w", id, err)
	}
	return &post, nil
}

func (c *Client) CreatePost(post Post) (*Post, error) {
	var result Post
	_, err := c.client.R().
		SetBody(post).
		SetResult(&result).
		Post("/posts")

	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return &result, nil
}

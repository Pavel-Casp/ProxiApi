package jsonplaceholder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const baseURL = "https://jsonplaceholder.typicode.com"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := baseURL + endpoint

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	return resp, nil
}

// Методы для работы с постами

func (c *Client) GetPosts() ([]Post, error) {
	resp, err := c.doRequest("GET", "/posts", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return posts, nil
}

func (c *Client) GetPostByID(id int) (*Post, error) {
	resp, err := c.doRequest("GET", "/posts/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &post, nil
}

func (c *Client) CreatePost(post Post) (*Post, error) {
	body, err := json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal post: %w", err)
	}

	resp, err := c.doRequest("POST", "/posts", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createdPost Post
	if err := json.NewDecoder(resp.Body).Decode(&createdPost); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &createdPost, nil
}

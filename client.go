package kakaorestapi

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const apiBaseURL = "https://dapi.kakao.com"

// Client struct of kakao rest API
type Client struct {
	apiKey     string
	httpClient []*http.Client
}

// NewClientWithProxies returns client instance with proxies pool.
func NewClientWithProxies(key string, proxies []string) (*Client, error) {
	var clients []*http.Client

	for _, proxy := range proxies {
		url, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(url),
			},
		}

		clients = append(clients, client)
	}

	return &Client{
		apiKey:     key,
		httpClient: clients,
	}, nil
}

// DefaultClient returns default client instance.
func DefaultClient(key string) (*Client, error) {
	return &Client{
		apiKey: key,
		httpClient: []*http.Client{
			&http.Client{},
		},
	}, nil
}

func (c *Client) get(url string, params map[string]interface{}) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, apiBaseURL+url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Authorization", "KakaoAK "+c.apiKey)

	values := req.URL.Query()

	for key, value := range params {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	req.URL.RawQuery = values.Encode()

	return c.fetchResponse(req)
}

func (c *Client) fetchResponse(req *http.Request) ([]byte, error) {
	client := c.getClient()
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return bytes, fmt.Errorf("failed to get response of URL: %s with HTTP code: %d", req.URL.String(), resp.StatusCode)
	}

	return bytes, nil
}

func (c *Client) getClient() *http.Client {
	if len(c.httpClient) == 1 {
		return c.httpClient[0]
	}

	rand.Seed(time.Now().UnixNano())

	return c.httpClient[rand.Intn(len(c.httpClient))]
}

package MIDDLEWARE

import (
	"fmt"
	r "gopkg.in/redis.v5"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"time"
)

var preffix = "_PAGE_CACHE_"

//Storage mecanism for caching strings in memory
type Storage struct {
	client *r.Client
}

//NewStorage creates a new redis storage
func NewStorage(url string) (*Storage, error) {
	var (
		opts *r.Options
		err  error
	)

	if opts, err = r.ParseURL(url); err != nil {
		return nil, err
	}

	return &Storage{
		client: r.NewClient(opts),
	}, nil
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	val, _ := s.client.Get(preffix + key).Bytes()
	return val
}

//Set a cached content by key
func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.client.Set(preffix+key, content, duration)
}

//Only use to Get Method
func Cached(storage *Storage, duration string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		re := regexp.MustCompile(`\?token=(\w*)\.(\w*).(\w*)`)
		negString := string(re.Find([]byte(r.RequestURI)))
		Key := strings.Replace(r.RequestURI, negString, "", 1)

		content := storage.Get(Key)
		if content != nil {
			fmt.Print("Cache Hit!\n")
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)

			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}

			w.WriteHeader(c.Code)
			content := c.Body.Bytes()

			if d, err := time.ParseDuration(duration); err == nil {
				fmt.Printf("New page cached: %s for %s\n", Key, duration)
				storage.Set(Key, content, d)
			} else {
				fmt.Printf("Page not cached. err: %s\n", err)
			}

			w.Write(content)
		}

	})
}

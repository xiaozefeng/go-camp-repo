package errg_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func Test_just_errors(t *testing.T) {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.baidu.com",
		"http://www.sina.com.cn",
		"https://u.geekbang.org",
	}
	for _, url := range urls {
		url := url
		fmt.Println("starting fetch, url:", url)
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("successfully fetched all URLS")
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(ctx context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

func Test_parallel(t *testing.T) {
	var Google = func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)
		var searches = []Search{Web, Image, Video}
		var results = make([]Result, len(searches))
		for i, search := range searches {
			i, search := i, search
			fmt.Println("starting search ...")
			g.Go(func() error {
				result, err := search(ctx, query)
				if err != nil {
					return err
				}
				results[i] = result
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, r := range results {
		fmt.Printf("%v\n", r)
	}
}


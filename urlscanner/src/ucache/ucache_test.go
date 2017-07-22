package ucache

import (
	"fmt"
	"testing"
)

func TestCacheClient(t *testing.T) {
	url := "http://google.com/test_url"
	Set(url)
	if Get(url) {
		fmt.Println("Present")
	}
}

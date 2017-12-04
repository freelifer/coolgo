package service

import (
	"net/url"
	"testing"
)

func Test_urlEncode(t *testing.T) {
	value := "https://uland.taobao.com/coupon/edetail?e=QrweHOsuj2sGQASttHIRqa59KcrVlGwiINX7EmZ0YliAnVXUwZx8IhEJwQnkdjDZ7aNoMgWMAfl%2FJ9rOMHOUiMPWjN00SJO8DfqEFBOhTcyABFd8ZHrxxMbGHnPVW3ogxXxVCXdmdxYzhukA9d4NUmPfrr0N2WBeCqEIqV4SxV7k92%2BM7h46c6J7%2BkHL3AEW&traceId=0ab013ac15116183416294882e"

	body := url.QueryEscape(value)

	src, _ := url.QueryUnescape(body)
	t.Log(body)

	t.Log(src)
}

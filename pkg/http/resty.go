package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type restyClient interface {
	AddRetryCondition(conditionFunc resty.RetryConditionFunc) *resty.Client
	NewRequest() *resty.Request
	SetCookie(cookie *http.Cookie) *resty.Client
	SetCookies(cookies []*http.Cookie) *resty.Client
	SetCookieJar(jar http.CookieJar) *resty.Client
	SetProxy(proxy string) *resty.Client
	RemoveProxy() *resty.Client
	SetRedirectPolicy(policies ...interface{}) *resty.Client
	SetTimeout(timeout time.Duration) *resty.Client
	GetCookies() []*http.Cookie
	DeleteCookies()
}

type cookieAwareRestyClient struct {
	*resty.Client
}

func newCookieAwareRestyClient(client *resty.Client) restyClient {
	return &cookieAwareRestyClient{client}
}

func (c *cookieAwareRestyClient) GetCookies() []*http.Cookie {
	return c.Cookies
}

func (c *cookieAwareRestyClient) DeleteCookies() {
	c.Cookies = []*http.Cookie{}
}

package rq

import "fmt"

func (r *Request) Get(URL string) (*Response, error) {
	return r.Execute("GET", URL)
}

func (r *Request) Getf(format string, a ...any) (*Response, error) {
	return r.Get(fmt.Sprintf(format, a...))
}

func (r *Request) Post(URL string) (*Response, error) {
	return r.Execute("POST", URL)
}

func (r *Request) Postf(format string, a ...any) (*Response, error) {
	return r.Post(fmt.Sprintf(format, a...))
}

func (r *Request) Put(URL string) (*Response, error) {
	return r.Execute("PUT", URL)
}

func (r *Request) Putf(format string, a ...any) (*Response, error) {
	return r.Put(fmt.Sprintf(format, a...))
}

func (r *Request) Patch(URL string) (*Response, error) {
	return r.Execute("PATCH", URL)
}

func (r *Request) Patchf(format string, a ...any) (*Response, error) {
	return r.Patch(fmt.Sprintf(format, a...))
}

func (r *Request) Delete(URL string) (*Response, error) {
	return r.Execute("DELETE", URL)
}

func (r *Request) Deletef(format string, a ...any) (*Response, error) {
	return r.Delete(fmt.Sprintf(format, a...))
}

package rq

import "net/http"

// Client is global HTTP client.
var Client = new(http.Client)

# sekiju/rq
> `rq` is a lightweight Go library for making HTTP requests.

## Usage
The `rq` library simplifies setting up and making HTTP requests. The `New` function configures request parameters such as the body, headers, and other options.

Methods like `Get` and `Post` execute the request. These methods take a URL as the first argument, and additional arguments formatted according to Go's `fmt.Sprintf` function. This allows for easy addition of query parameters or variables to the URL.

Here's an example of how to use it:
```go
body := map[string]interface{}{
"name":     "sekiju/rq",
"birthday": 1710018066000,
"is_child": true,
}

// Create a new request instance with a specified body
req := rq.New(rq.SetBody(body))

// Execute a POST request
res, err := req.Post("https://example.com/post")
```

In this example, a new request is created with a specified body and a POST request is executed to the specified URL. To add query parameters to the URL, use the following:

```go
// Execute a GET request with query parameters
res, err := rq.New().Get("https://example.com/get?param=%s", "value")
```

In this example, we execute a GET request to a URL that includes a query parameter. The value of the parameter is formatted using `Sprintf`.

## Request Parameters

* `SetHeader(key, value string)` - Sets a header for the request.
* `SetBody(interface{})` - Sets the body of the request.
* `SetBodyType(BodyType)` - Sets the `Content-Type` of the body and encodes it into header format. Default is `JsonBodyType`.
  * `BodyType` is an enum with the following values:
    * `RawBodyType` - No header, just encodes to a byte slice.
    * `JsonBodyType` - `application/json`, automatic marshalling to JSON.

## Credits

* https://httpbin.org/ - A simple HTTP request and response service used for tests.
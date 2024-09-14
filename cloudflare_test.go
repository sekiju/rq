package rq

import "testing"

func TestCloudflare(t *testing.T) {
	t.Run("GET Cloudflare", func(t *testing.T) {
		res, err := New().Get("https://unconnected.ranobe.plus/v2/books/swordmasters-youngest-son")
		if err != nil {
			t.Fatalf("Error making POST request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}

		var v any
		if err = res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}
	})
}

package order

import (
	"os"
	"testing"

	"github.com/zerodays/woocommerce-go/internal/backend"
)

var baseURL, consumerKey, consumerSecret string

func TestMain(m *testing.M) {
	baseURL = os.Getenv("BASE_URL")
	consumerKey = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")

	os.Exit(m.Run())
}

func TestClient_List(t *testing.T) {
	b := backend.New(baseURL, consumerKey, consumerSecret)
	client := New(b)
	orders, _, err := client.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("#orders: %d", len(orders))
	if len(orders) > 0 {
		t.Logf("%#v", orders[0])
	}
}

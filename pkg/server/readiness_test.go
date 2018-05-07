package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pmorie/osb-broker-lib/pkg/broker"
	"github.com/pmorie/osb-broker-lib/pkg/metrics"
	"github.com/pmorie/osb-broker-lib/pkg/rest"
	prom "github.com/prometheus/client_golang/prometheus"
)

func TestHasReadiness(t *testing.T) {
	cases := []struct {
		name         string
		broker       broker.Interface
		responseCode int
	}{
		{
			name:         "readiness not available",
			broker:       &FakeBroker{},
			responseCode: http.StatusNotFound,
		},
		{
			name: "readiness found",
			broker: &FakeBrokerWithReadiness{
				readinessFunc: func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}},
			responseCode: http.StatusOK,
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			// Prom. metrics
			reg := prom.NewRegistry()
			osbMetrics := metrics.New()
			reg.MustRegister(osbMetrics)

			api := &rest.APISurface{
				Broker:  tc.broker,
				Metrics: osbMetrics,
			}

			s := New(api, reg)
			fs := httptest.NewServer(s.Router)
			defer fs.Close()

			url := fs.URL + "/readiness"
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("unable to get readiness endpoint")
			}
			if resp.StatusCode != tc.responseCode {
				t.Fatalf("expected status code: %v but got: %v", tc.responseCode, resp.StatusCode)
			}
		})
	}
}

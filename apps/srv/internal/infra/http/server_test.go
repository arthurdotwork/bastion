package http_test

import (
	"context"
	"fmt"
	"math/rand"
	stdHTTP "net/http"
	"sync"
	"testing"
	"time"

	"github.com/arthurdotwork/bastion/internal/infra/http"
	"github.com/stretchr/testify/require"
)

func TestServer_Serve(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	addr := fmt.Sprintf("127.0.0.1:%d", rand.Intn(16384)+49152)
	srv := http.NewServer(addr)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := srv.Serve(ctx)
		require.NoError(t, err)
	}()

	require.Eventually(t, func() bool {
		livenessResp, err := stdHTTP.Get(fmt.Sprintf("http://%s/checks/liveness", addr))
		if err != nil {
			t.Log("error on liveness probe:", err)
			return false
		}

		if livenessResp.StatusCode != stdHTTP.StatusOK {
			t.Log("liveness probe failed:", livenessResp.StatusCode)
			return false
		}

		readinessResp, err := stdHTTP.Get(fmt.Sprintf("http://%s/checks/readiness", addr))
		if err != nil {
			t.Log("error on readiness probe:", err)
			return false
		}

		if readinessResp.StatusCode != stdHTTP.StatusOK {
			t.Log("readiness probe failed:", readinessResp.StatusCode)
			return false
		}

		return true
	}, time.Second*5, time.Millisecond*100)

	cancel()
	wg.Wait()
}

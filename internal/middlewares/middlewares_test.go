package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type writer struct {
	statusCode int
	sb         strings.Builder
}

func (w *writer) Header() http.Header {
	return nil
}

func (w *writer) Write(bytes []byte) (int, error) {
	return w.sb.Write(bytes)
}

func (w *writer) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

type testWithPrint struct {
	t *testing.T
}

func (t testWithPrint) printi(c byte) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte{c}); err != nil {
				t.t.Fatalf("error in printi: %c", c)
			}
			next.ServeHTTP(w, r)
		})
	}
}

type handler struct {
	statusCode int
}

func (h handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(h.statusCode)
}

var _successHandler http.Handler = handler{statusCode: 200}

func TestChain(t *testing.T) {
	tp := testWithPrint{t: t}

	mws := []Middleware{
		tp.printi('1'),
		tp.printi('2'),
		tp.printi('3'),
	}
	chain1 := Chain(mws)
	w1 := new(writer)

	mws[0] = tp.printi('9')
	mws = append(mws, tp.printi('4'))
	chain2 := Chain(mws)
	w2 := new(writer)

	chain1(_successHandler).ServeHTTP(w1, nil)
	chain2(_successHandler).ServeHTTP(w2, nil)

	assert.Equal(t, "123", w1.sb.String())
	assert.Equal(t, "9234", w2.sb.String())

	assert.Equal(t, 200, w1.statusCode)
	assert.Equal(t, 200, w2.statusCode)
}

type lg struct {
	t        *testing.T
	messages chan string
}

func (l *lg) Debug(message interface{}, args ...interface{}) { l.msg(message, args...) }
func (l *lg) Info(message string, args ...interface{})       { l.msg(message, args...) }
func (l *lg) Warn(message string, args ...interface{})       { l.msg(message, args...) }
func (l *lg) Error(message interface{}, args ...interface{}) { l.msg(message, args...) }
func (l *lg) Fatal(message interface{}, args ...interface{}) { l.msg(message, args...) }

func (l *lg) msg(message interface{}, args ...interface{}) {
	msgStr, ok := message.(string)
	assert.Equalf(l.t, true, ok, "message should be string")
	go func() {
		l.messages <- fmt.Sprintf(msgStr, args...)
	}()
}

func TestLogger(t *testing.T) {
	tlg := &lg{t: t, messages: make(chan string)}

	mw := Logger(tlg)

	canProceedHandler := make(chan struct{})
	defer close(canProceedHandler)

	exitChan := make(chan struct{})
	defer close(exitChan)

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		<-canProceedHandler
		w.WriteHeader(500)
		exitChan <- struct{}{}
	})

	mwh := mw(h)

	select {
	case msg := <-tlg.messages:
		t.Fatalf("message logged before start handling: %s", msg)
	case <-time.After(time.Millisecond * 300):
	}

	wr := &writer{}
	go mwh.ServeHTTP(wr, &http.Request{Method: http.MethodGet, RequestURI: "localhost"})

	select {
	case msg := <-tlg.messages:
		assert.Equalf(t, true, strings.Contains(msg, "localhost"), "message should contain request uri")
	case <-time.After(time.Millisecond * 200):
		t.Fatal("message is not logged yet")
	}

	assert.Empty(t, wr.statusCode, "handler must not proceed before closing the channel")

	canProceedHandler <- struct{}{}
	<-exitChan

	assert.Equal(t, wr.statusCode, 500)

	// check if not panic on nil request
	go mwh.ServeHTTP(wr, nil)
	canProceedHandler <- struct{}{}
	<-exitChan
}

func TestRecoverer(t *testing.T) {
	hPanic := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic("try to recover me")
	})

	hNormal := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
	})

	l := &lg{t: t, messages: make(chan string)}
	wr := &writer{}

	assert.NotPanics(t, func() {
		Recoverer(l)(hPanic).ServeHTTP(wr, nil)
	}, "must not panic")

	select {
	case msg := <-l.messages:
		assert.Equalf(t, true, strings.Contains(msg, "try to recover me"), "message must contain recovered value")
	case <-time.After(time.Millisecond * 200):
		t.Fatal("message is not logged yet")
	}

	Recoverer(l)(hNormal).ServeHTTP(wr, nil)

	select {
	case msg := <-l.messages:
		t.Fatalf("no panic -> no message, but got: %s", msg)
	case <-time.After(time.Millisecond * 300):
	}
}

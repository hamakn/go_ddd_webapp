package middleware

import (
	"net/http"

	"google.golang.org/appengine"
)

// ContextSetter is middleware to set appengine context
type ContextSetter struct {
	Namespace *string
}

// NewContextSetter returns new ContextSetter
func NewContextSetter() *ContextSetter {
	return &ContextSetter{}
}

func (c *ContextSetter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// generate appengine context to avoid:
	//   appengine: NewContext passed an unknown http.Request
	// refs:
	//   https://groups.google.com/forum/#!topic/google-appengine-go/Av7Lg956D6Y
	//   https://qiita.com/tenntenn/items/0b92fc089f8826fabaf1
	ctx := r.Context()                  // for gsc
	ctx = appengine.WithContext(ctx, r) // for gsc
	next(w, r.WithContext(ctx))
}

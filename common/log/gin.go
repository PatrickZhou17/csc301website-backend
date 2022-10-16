package log

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

const (
	ContextLogKey = "__context_log_key__"
)

const (
	requestStartTimeKey   = "start"
	requestEndTimeKey     = "end"
	requestDurationKey    = "duration"
	requestStatusKey      = "status"
	requestURLRawQueryKey = "query"
	requestClientIPKey    = "ip"
	requestUserAgentKey   = "user-agent"
	requestURLPath        = "path"

	nameKey    = "name"
	traceIDKey = "trace"
)

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				httpRequest = httpRequest[:len(httpRequest)-2]
				if brokenPipe {
					stdLogger.Errorf("[Broken Pipe]:\n[Request]:\n%s[Error]:\n%s", string(httpRequest), err)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				}
				if stack {
					stdLogger.Errorf("[Recovery from panic]:\n[Request]:\n%s[Error]:\n%s\n[Stack]:\n%s", string(httpRequest), err, string(debug.Stack()))
				} else {
					stdLogger.Errorf("[Recovery from panic]:\n[Request]:\n%s[Error]:\n%s", string(httpRequest), err)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// GetFromGin return a Entry, you should use Entry in api handler to log
func GetFromGin(c *gin.Context) *Logger {
	var (
		e  *Logger
		ok bool
	)
	ee, ok := c.Get(ContextLogKey)
	if ok {
		e, ok = ee.(*Logger)
	}
	if ok {
		return e
	} else {
		return stdLogger.WithFields(nil, nil)
	}
}

// GetFromGinWithName is same as GetFromGin and you can pass a name to it, then log will add a field called name
func GetFromGinWithName(c *gin.Context, name string) *Logger {
	e := GetFromGin(c)
	logger := e.logger.With(ZapString(nameKey, name))
	sugaredLogger := e.sugaredLogger.With(nameKey, name)
	e = &Logger{logger, sugaredLogger}
	c.Set(ContextLogKey, e)

	return e
}

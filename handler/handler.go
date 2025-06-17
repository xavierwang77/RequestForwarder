package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxyHandler(target string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		logger.Error("Failed to parse target URL", zap.Error(err))
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 处理响应错误（可选）
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		logger.Error("Failed to proxy request", zap.Error(err))
		http.Error(w, "Proxy error", http.StatusBadGateway)
	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

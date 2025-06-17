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

	// 自定义转发行为
	proxy.Director = func(req *http.Request) {
		originalPath := req.URL.Path

		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.Host = targetURL.Host

		// 拼接目标路径
		req.URL.Path = singleJoiningSlash(targetURL.Path, originalPath)

		// 日志输出即将转发的请求信息
		logger.Info("Forwarding request",
			zap.String("method", req.Method),
			zap.String("target", req.URL.String()),
			zap.String("originalPath", originalPath),
			zap.String("host", req.Host),
			zap.String("clientIP", req.RemoteAddr),
		)
	}

	// 处理响应错误（可选）
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		logger.Error("Failed to proxy request", zap.Error(err))
		http.Error(w, "Proxy error", http.StatusBadGateway)
	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

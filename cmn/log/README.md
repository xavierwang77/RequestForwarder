### logger 使用示例
```go
import (
	"go.uber.org/zap"
)

func main() {
    // get log instance
    logger := zap.L().Suger()
    
    logger.Infof("...")
    logger.Warnf("...")
    logger.Errorf("...")
}
```
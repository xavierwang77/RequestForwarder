### 日志读取器使用示例
```go
func main() {
	// 获取配置文件读取器
	configReader, err := config.GetConfigReader()
	if err != nil {
        // 处理错误
    }
	
	dbHost := configReader.GetString("dbms.host")
	
	// ...
}
```
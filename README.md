# ALAPI MCP Server

这是一个基于 [ALAPI](https://www.alapi.cn) 的 MCP (Model Control Protocol) 服务器实现，可以自动将 ALAPI 的 OpenAPI 规范转换为 MCP 工具。

## 功能特点

- 自动加载 ALAPI OpenAPI 规范
- 支持加载全部或指定 API (短视频解析、天气查询、每天60秒读懂世界、企业查询...)
- 符合 MCP 协议标准
- 统一的错误处理和响应格式
- 支持环境变量配置

## 环境要求

- Go 1.24.1 或更高版本
- 有效的 ALAPI Token
- 支持 MCP 的客户端（如 Claude Desktop、Continue、Cursor 等）

## 安装

### 构建安装
```bash
#Github
git clone https://github.com/alapi-sdk/mcp-alapi-cn.git

#CNB 国内加速地址
#git clone https://cnb.cool/alapi/mcp-alapi-cn.git

cd mcp-server
go mod tidy
go build
```

### 下载已构建的软件包

Github:

CNB(国内加速): 



## 配置和使用

### 环境变量配置

项目使用以下环境变量：

- `ALAPI_TOKEN`（必需）：ALAPI 的认证令牌,在 [token管理](https://www.alapi.cn/dashboard/data/token) 里面创建, **如果不设置，mcp 会启动不了**
- `ALAPI_ID`（可选）：指定要加载的 API ID，不设置则加载所有 API（实际的API_id 可通过 [ALAPI](https://www.alapi.cn)  官网查看，可在我的API里面查询）




### Cursor配置方式

1. 打开Cursor设置 > 扩展 > MCP工具
2. 添加新的MCP工具
3. 按照以下格式填写配置：

```json
{
    "mcpServers": {
        "mcp-alapi-cn": {
            "command": "C:\\Users\\Administrator\\实际目录\\mcp-alapi-cn.exe",
            "env": {
                "ALAPI_TOKEN": "xxxx",
                "ALAPI_API_ID": "0"
            }
        }
    }
}
```

### CherryStudio配置方式
1. 打开 CherryStudio 设置 -> MCP 服务器 
2. 添加MCP服务器
3. 配置说明：

名称： `MCP-ALAPI-CN`  
类型： `STDIO`    
命令： `C:\\Users\\Administrator\\实际目录\\mcp-alapi-cn.exe`
环境变量：
```
ALAPI_TOKEN=你的token
```

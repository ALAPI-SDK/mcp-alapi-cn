# 🔥 ALAPI MCP Server

[![smithery badge](https://smithery.ai/badge/@ALAPI-SDK/mcp-alapi-cn)](https://smithery.ai/server/@ALAPI-SDK/mcp-alapi-cn)
![License](https://img.shields.io/badge/license-MIT-green)

这是一个基于 [ALAPI](https://www.alapi.cn) 的 MCP (Model Context Protocol) 服务器实现，可以通过MCP协议直接调用ALAPI的接口，轻松将ALAPI丰富的API资源集成到各类支持MCP的AI应用中。

## ✨ 功能特点

- 🔍 **丰富API资源** - 提供IP查询、天气查询、热搜榜单、企业信息等100+种常用API服务
- 🧩 **MCP 协议支持** - 完全兼容 Model Context Protocol，轻松集成到各类AI应用
- 🔄 **自动加载** - 自动加载 ALAPI OpenAPI 规范，无需手动配置
- 🎯 **灵活调用** - 支持加载全部或指定API，按需使用（短视频解析、天气查询、每天60秒读懂世界、企业查询等）
- 🛡️ **统一响应** - 提供统一的错误处理和响应格式，简化开发流程
- ⚙️ **可配置性** - 支持环境变量配置，轻松适应不同使用场景

## 📖 使用指南

首先需要了解 [MCP](https://modelcontextprotocol.io/introduction) 协议，然后按照以下配置添加 ALAPI MCP 服务。

## 在 MCP 客户端中使用

ALAPI MCP Server 设计为与支持 MCP 协议的客户端配合使用，不支持独立的命令行配置。请按照以下步骤在您的 MCP 客户端中配置使用：

### 方法一：使用 Smithery 自动安装

通过 [Smithery](https://smithery.ai/server/@ALAPI-SDK/mcp-alapi-cn) 安装 ALAPI MCP Server，适用于 Claude Desktop 客户端：

```bash
npx -y @smithery/cli install @ALAPI-SDK/mcp-alapi-cn --client claude
```

### 方法二：直接使用预编译二进制文件（推荐，不用安装依赖）

1. 访问 [Releases 页面](https://github.com/ALAPI-SDK/mcp-alapi-cn/releases) 下载最新版本 
   > 国内加速下载地址：https://cnb.cool/alapi/mcp-alapi-cn/-/releases

2. 根据您的操作系统选择对应最新的版本：
   - Windows: `mcp-alapi-cn.exe`
   - macOS: `mcp-alapi-cn-macos`
   - Linux: `mcp-alapi-cn-linux`

3. 在支持 MCP 的客户端（如 Claude Desktop）中配置：
   - 打开 MCP 客户端设置
   - 添加新的 MCP 服务器
   - 指定下载的二进制文件路径
   - 配置必要的环境变量（如 `ALAPI_TOKEN`）

### 配置环境变量

#### `ALAPI_TOKEN` - ALAPI平台访问令牌（必填）

需要配置为您在 ALAPI 平台上创建的 [TOKEN](https://www.alapi.cn/dashboard/data/token)，用于接口认证。

### MCP 客户端配置示例

在 Claude Desktop 或其他支持 MCP 的客户端中，可以参考以下配置：

#### JSON 配置（使用本地二进制文件）

```json
{
  "mcpServers": {
    "alapi": {
      "path": "/path/to/mcp-alapi-cn",
      "env": {
        "ALAPI_TOKEN": "在 ALAPI 平台申请创建的 token"
      }
    }
  }
}
```


### 支持的 MCP 客户端

根据 [MCP 文档](https://modelcontextprotocol.io/clients)，以下客户端支持 MCP 工具集成：

- Claude Desktop App
- Claude Code
- Cursor
- Continue
- Cline
- 等其他支持 MCP 协议的客户端

> 注意：确保您的 MCP 客户端支持工具集成功能，请参考各客户端的具体文档进行配置。

## 🛠️ 支持的功能

以下是部分支持的功能，完整功能请以实际调用为准：

| 功能名称     | 描述                                                                          |
| ------------ | ----------------------------------------------------------------------------- |
| IP查询       | 查询IP归属地信息，支持IPv4和IPv6                                              |
| 天气查询     | 查询国内天气详情，包含天气信息、天气预警、天气指数、AQI等                     |
| 今日热榜     | 查询今日热榜数据，包含抖音热搜、头条热榜、知乎、36k、百度热搜、微博热搜等数据 |
| 文本审核     | 查询文本是否违规，特别适合AI内容审核                                          |
| 短网址       | 短网址缩短，将长网址缩短为短网址，方便短信发送、二维码生成等场景              |
| 二维码识别   | 支持一图多码，支持数十二维码，条形码的识别                                    |
| 历史上的今天 | 查询历史上的指定日期发生的大事                                                |
| 汇率查询     | 根据货币代码查询货币汇率和换算，支持全球170+个国家和地区货币查询              |
| 企业查询     | 根据企业名称或统一社会信用代码查询企业基本工商信息                            |
| 快递查询     | 免费查询快递物流轨迹，支持中通、申通、顺丰、极兔、百世、圆通等                |
| 星座运势     | 提供星座运势查询服务，支持查询十二星座今日或明日、本周、本月、本年的运势      |
| 黄金价格     | 查询当前黄金的实时价格，最高价，最低价，品牌黄金的价格                        |

## 🔗 相关链接

- [ALAPI官网](https://www.alapi.cn) - 获取API服务和Token
- [MCP协议文档](https://modelcontextprotocol.io) - 了解更多MCP协议相关信息

## 🙏 鸣谢

- [ALAPI](https://www.alapi.cn) - 提供强大的API支持
- [Model Context Protocol](https://modelcontextprotocol.io) - 提供协议标准



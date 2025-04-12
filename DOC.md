# ALAPI接口平台 MCP 接入

通过 MCP 协议提供与 ALAPI 接口平台交互，可快捷查询IP归属地、天气情况、数据热榜、文本审核、短网址缩短等操作

[前往云开发平台运行 MCP Server](https://tcb.cloud.tencent.com/dev#/ai?tab=mcp&p&mcp-template=mcp-alapi-cn)

---

## 环境变量

- 需要将 **ALAPI_TOKEN** 配置为您在 ALAPI 上创建的 [TOKEN](https://www.alapi.cn/dashboard/data/token)

## 🗺️ 功能清单

以下仅罗列部分功能，具体请以实际调用为准

| 命令名称                          | 功能描述                                          | 核心参数          |
|-------------------------------|-----------------------------------------------|---------------|
| user_apis                     | 获取已申请的接口列表                                    |               |
| /api/ip                       | 查询IP归属地信息                                     | ip            |
| /api/tianqi                   | 查询国内天气详情、包含天气信息、天气预警、天气指数、AQI 等               | city          |
| /api/tophub                   | 查询今日热榜数据，包含抖音热搜、头条热榜、知乎、36k、百度热搜、搜狗热搜、微博热搜等数据 | id,type       |
| /api/censor/text              | 查询文本是否违规,特别适合AI内容审核                           | text          |
| /api/url                      | 短网址缩短，将长网址缩短为短网址，方便短信发送、二维码生成等场景              | url           |
| /api/qr/decode                | 支持一图多码，支持数十二维码，条形码的识别                         | url           |
| /api/eventHistory             | 历史上的今天,查历史上的指定日期发生的大事                         | date          |
| /api/exchange                 | 汇率查询,根据货币代码查询货币汇率和换算，支持全球170+个国家和地区货币查询       | money,from,to |
| /api/enterprise/simple_search | 企业工商信息搜索,根据企业名称或统一社会信用代码查询企业基本工商信息            | keyword       |
| /api/kd                       | 快递查询V1,免费查询快递物流轨迹，支持中通、申通、顺丰、极兔、百世、圆通等        | numner,phone  |
| /api/star                     | 星座运势查询,提供星座运势查询服务，支持查询十二星座今日或明日、本周、本月、本年的运势。  | star          |
| /api/gold                     | 查询当前黄金的实时价格，最高价，最低价，品牌黄金的价格                   | market        |

## 仓库地址

[https://cnb.cool/alapi/mcp-alapi-cn](https://cnb.cool/alapi/mcp-alapi-cn)

---

## 🔌 使用方式

- [在云开发 Agent 中使用](https://docs.cloudbase.net/ai/mcp/use/agent)
- [在 MCP Host 中使用](https://docs.cloudbase.net/ai/mcp/use/mcp-host)
- [通过 SDK 接入](https://docs.cloudbase.net/ai/mcp/use/sdk)

---

[云开发 MCP 控制台](https://tcb.cloud.tencent.com/dev#/ai?tab=mcp)  

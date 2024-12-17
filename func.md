# 证书监控平台

## 1. 项目概述
> 本项目是一个基于 Golang 和 Gin 框架 开发的 SSL 证书自动续签平台，主要功能包括：
> * 域名管理
> * SSL 证书续签
> * DNS 账户管理
> * 证书推送至 CDN
> * 日志管理与查询
> * 定时任务与通知系统
---
## 2. 功能模块
    
### 2.1 用户管理
* 用户注册：支持管理员和普通用户两种角色。
* 用户登录：验证用户名与密码，返回 JWT 令牌。
* 获取用户列表：管理员可查看所有用户。
* 删除用户：管理员可删除用户。
* 角色管理：区分 admin 和 user 角色，控制操作权限。
---
### 2.2 域名管理
* 添加域名：添加需要管理的域名。
* 获取域名列表：列出所有域名及证书状态。
* 删除域名：删除指定域名。
* 证书状态查询：获取域名的 SSL 证书状态（有效、即将过期、已过期）。
---
### 2.3 ACME管理和配置
* 脚本配置：配置acme脚本路径，或者其他配置。
* 脚本更新：配置acme脚本更新规则。
* 核心方法：实现续签等功能。
---
### 2.4 DNS 账户管理
* 添加 DNS 账户：添加与 acme.sh 集成的 DNS 提供商信息。
* 获取 DNS 账户：列出所有 DNS 账户配置。
* 删除 DNS 账户：删除指定 DNS 配置。
* 支持的 DNS 提供商：例如 Cloudflare、阿里云 DNS 等。
---
### 2.5 SSL 证书管理
* 自动续签证书：使用 acme.sh 定时检查并自动续签到期的 SSL 证书。
* 手动触发续签：管理员可手动触发指定域名的证书续签。
* 证书推送：将续签后的 SSL 证书推送到 CDN 提供商（如 Cloudflare）。
* 证书下载：提供证书下载接口。
---
### 2.6 日志管理
* 记录操作日志：
* 用户操作（如添加/删除域名、DNS 账户）
* 系统任务（如证书续签、证书推送）
* 日志查询：
* 按日期、操作类型等筛选日志。
* 查看任务状态：记录任务的成功与失败状态。
---
### 2.7 定时任务与通知
* 定时任务：
* 每日检查证书到期状态。
* 自动续签即将过期的 SSL 证书。
* 通知系统：
* 通过邮件或 Webhook 提醒用户证书续签状态。
* 通知任务失败或错误信息。
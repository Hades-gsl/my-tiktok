# my-tiktok

使用 `Go` 编写的一个简易的抖音的后端项目，为第六届字节跳动青训营后端大作业。

## 介绍

使用 `kitex` 作为 `RPC` 框架，使用 `hertz` 作为 `web` 框架，使用 `gorm` 作为
`ORM` 框架，实现了简单的抖音后端。

## 项目结构

- config 一些常量和配置的定义
- db `mysql` 数据库相关代码
- idl 相关服务的 `RPC` 定义
- kitex_gen kitex生成的代码
- rdb `redis` 相关代码
- service 
    - chat 聊天服务
    - comment 评论服务
    - control 控制\网关服务
    - favorite 点赞服务
    - feed 视频流服务
    - publish 发布服务
    - relation 关注服务
    - user 用户服务
- build.sh 构建脚本
- kitex_gen.sh kitex生成代码脚本

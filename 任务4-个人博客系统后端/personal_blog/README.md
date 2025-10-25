// 本次作业要求你使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。 
项

// 1.项目初始化
// 创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理。
// go mod init personal_blog

// 安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）。
// go get -u github.com/gin-gonic/gin
// go get -u gorm.io/driver/mysql

// 2.数据库设计与模型定义
// 设计数据库表结构，至少包含以下几个表：
// users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
// posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
// comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
// 使用 GORM 定义对应的 Go 模型结构体。


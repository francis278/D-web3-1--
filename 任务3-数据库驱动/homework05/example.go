package homework05

import (
	"fmt"

	"gorm.io/gorm"
)

// 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//  要求 ：
//   使用Gorm定义 User 、 Post 和 Comment 模型，
//   其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
//   Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
//   编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//   编写Go代码，使用Gorm查询评论数量最多的文章信息。

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
//  为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//  为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
//  如果评论数量为 0，则更新文章的评论状态为 "无评论"。

// User 的表
type User struct {
	gorm.Model
	Name     string
	Posts    []Post    `gorm:"foreignKey:AuthorID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}

// Post 的表
type Post struct {
	gorm.Model
	Title    string
	Content  string
	Summary  string
	Status   string
	Comments []Comment `gorm:"foreignKey:PostID"`

	// 外键
	AuthorID uint `gorm:"not null;index"`
	Author   User `gorm:"foreignKey:AuthorID"`
}

// Comment 的表
type Comment struct {
	gorm.Model
	Content string
	Status  string

	// 外键
	UserID uint `gorm:"not null;index"`
	PostID uint `gorm:"not null;index"`
	User   User `gorm:"foreignKey:UserID"`
	Post   Post `gorm:"foreignKey:PostID"`
}

// 创建测试数据
func createTestData(db *gorm.DB) error {
	// 清空现有数据（按依赖顺序）
	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM users")

	// 1. 创建用户
	users := []User{
		{Name: "张三"},
		{Name: "李四"},
		{Name: "王五"},
		{Name: "赵六"},
	}
	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	// 2. 创建文章
	posts := []Post{
		{
			Title:    "Go语言入门指南",
			Content:  "这是一篇关于Go语言基础的教程...",
			Summary:  "学习Go语言的基本语法和特性",
			Status:   "published",
			AuthorID: users[0].ID, // 张三的文章
		},
		{
			Title:    "GORM使用技巧",
			Content:  "GORM是一个强大的Go语言ORM框架...",
			Summary:  "掌握GORM的高级用法",
			Status:   "published",
			AuthorID: users[1].ID, // 李四的文章
		},
		{
			Title:    "数据库设计原则",
			Content:  "良好的数据库设计是应用成功的关键...",
			Summary:  "学习数据库设计的最佳实践",
			Status:   "draft",
			AuthorID: users[0].ID, // 张三的草稿
		},
		{
			Title:    "Web开发实战",
			Content:  "使用Go语言构建Web应用...",
			Summary:  "从零开始构建完整的Web应用",
			Status:   "published",
			AuthorID: users[2].ID, // 王五的文章
		},
	}
	if err := db.Create(&posts).Error; err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	}

	// 3. 创建评论
	comments := []Comment{
		{
			Content: "写得很好，受益匪浅！",
			Status:  "active",
			UserID:  users[1].ID, // 李四评论
			PostID:  posts[0].ID, // 对张三的文章评论
		},
		{
			Content: "期待更多关于GORM的内容",
			Status:  "active",
			UserID:  users[2].ID, // 王五评论
			PostID:  posts[1].ID, // 对李四的文章评论
		},
		{
			Content: "作者讲得很详细，点赞！",
			Status:  "active",
			UserID:  users[3].ID, // 赵六评论
			PostID:  posts[0].ID, // 对张三的文章评论
		},
		{
			Content: "有些地方不太明白，能再解释一下吗？",
			Status:  "active",
			UserID:  users[0].ID, // 张三评论
			PostID:  posts[1].ID, // 对李四的文章评论
		},
		{
			Content: "实战案例很有参考价值",
			Status:  "active",
			UserID:  users[1].ID, // 李四评论
			PostID:  posts[3].ID, // 对王五的文章评论
		},
	}
	if err := db.Create(&comments).Error; err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	}

	fmt.Println("✅ 初始数据创建完成！")
	fmt.Printf("创建了 %d 个用户\n", len(users))
	fmt.Printf("创建了 %d 篇文章\n", len(posts))
	fmt.Printf("创建了 %d 条评论\n", len(comments))

	return nil
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	var postCount int64
	tx.Model(&Post{}).Where("author_id = ?", p.AuthorID).Count(&postCount)
	fmt.Printf("用户的文章数量统计 %d\n", postCount)
	return nil
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Post{})
	// db.AutoMigrate(&Comment{})
	// createTestData(db)

	//   编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// user := User{Name: "张三"}
	// db.Preload("Posts").Preload("Comments").Find(&user)
	// fmt.Printf("作者: %s\n", user.Name)

	// for _, post := range user.Posts {
	// 	fmt.Printf("文章: %s\n", post.Title)
	// }

	// for _, comment := range user.Comments {
	// 	fmt.Printf("评论信息: %s\n", comment.Content)
	// }

	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	// var posts []Post
	// db.Preload("Comments").
	// 	Preload("Author").
	// 	Preload("Comments.User").
	// 	Select("posts.*, count(1) as comment_count").
	// 	Joins("left join comments on comments.post_id = posts.id").
	// 	Group("posts.id").Order("comment_count desc").Find(&posts)

	// maxCommentCount := len(posts[0].Comments)
	// var results []Post

	// for _, post := range posts {
	// 	if len(post.Comments) == maxCommentCount {
	// 		results = append(results, post)
	// 	} else {
	// 		break
	// 	}
	// }

	// for _, result := range results {
	// 	fmt.Printf("文章: %s, 评论数量: %d\n", result.Title, maxCommentCount)
	// }

	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	// post := Post{
	// 	Title:    "Go并发编程：从Channel到Context",
	// 	Content:  "Go语言的并发模型是其最大的特色之一...",
	// 	Summary:  "全面解析Go语言并发编程的核心概念",
	// 	Status:   "published",
	// 	AuthorID: 1,
	// }
	//db.Create(&post)

	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	post := Post{}
	post.ID = 1

	var comments []Comment
	db.Debug().Where("post_id = ?", post.ID).Find(&comments)

	for _, comment := range comments {
		db.Delete(&comment) // 这样会触发 AfterDelete 钩子
	}
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {

	var commentCount int64
	tx.Debug().Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)
	if commentCount == 0 {
		fmt.Printf("无评论")
	} else {
		fmt.Printf("文章的评论数量 %d\n", commentCount)
	}
	return nil
}

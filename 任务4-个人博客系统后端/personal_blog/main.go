package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 添加这一行！
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

func main() {
	// db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	db, err := gorm.Open(mysql.Open("root:Fanzhf123!@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	// 自动迁移模型
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 用户认证与授权==============
	r := gin.Default()
	// 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，
	r.POST("/user", func(c *gin.Context) {
		Register(c, db)
	})
	// 登录时验证用户输入的用户名和密码。
	r.POST("/login", func(c *gin.Context) {
		Login(c, db)
	})
	// 实现文章的创建功能
	r.POST("/addpost", func(c *gin.Context) {
		AddPost(c, db)
	})
	// 实现文章的读取功能
	r.POST("/getpostlist", func(c *gin.Context) {
		GetPostList(c, db)
	})
	// 实现文章的读取功能
	r.GET("/getpostdetail/:id", func(c *gin.Context) {
		GetPostDetail(c, db)
	})
	// 实现文章的更新功能
	r.POST("/updatePost", func(c *gin.Context) {
		UpdatePost(c, db)
	})
	// 实现文章的删除功能
	r.POST("/deletePost", func(c *gin.Context) {
		DeletePost(c, db)
	})
	// 评论功能
	// 实现评论的创建功能，已认证的用户可以对文章发表评论。
	r.POST("/addComment", func(c *gin.Context) {
		AddComment(c, db)
	})
	// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
	r.POST("/readCommentList", func(c *gin.Context) {
		ReadCommentList(c, db)
	})

	// 错误处理与日志记录
	// 对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，
	// 返回合适的 HTTP 状态码和错误信息。

	err1 := r.Run(":8081")
	if err1 != nil {
		panic(err1)
	}
	//router.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

func Register(c *gin.Context, db *gorm.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context, db *gorm.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser User
	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// fmt.Println("Username:", user.Username)
	// fmt.Println("Password:", user.Password)
	// fmt.Println("storedUser.Password:", storedUser.Password)

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Login sucessful",
		"token":   tokenString,
		"user": gin.H{
			"id":       storedUser.ID,
			"username": storedUser.Username,
		},
	})
}

// 文章管理功能
func AddPost(c *gin.Context, db *gorm.DB) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
}

// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
func GetPostList(c *gin.Context, db *gorm.DB) {
	var posts []Post

	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取所有文章列表"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "文章的读取成功。",
		"posts":   posts,
	})
}

// 单个文章的详细信息
func GetPostDetail(c *gin.Context, db *gorm.DB) {
	var post Post

	if err := db.Where("ID=?", c.Param("id")).Find(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "单个文章的详细信息"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "单个文章的详细信息",
		"post":    post,
	})
}

// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func UpdatePost(c *gin.Context, db *gorm.DB) {

	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var postResult Post
	if err := db.Where("user_id=?", post.UserID).First(&postResult).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章的作者信息不存在"})
		return
	}

	result := db.Debug().Model(&Post{}).Where("ID=?", post.ID).Updates(Post{
		Title:   post.Title,
		Content: post.Content,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}
	// fmt.Println("post.Title:", post.Title)
	// fmt.Println("post.Content:", post.Content)
	c.JSON(http.StatusOK, gin.H{"message": "文章更新成功"})
}

// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func DeletePost(c *gin.Context, db *gorm.DB) {

	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var postResult Post
	if err := db.Where("user_id=?", post.UserID).First(&postResult).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章的作者信息不存在"})
		return
	}

	result := db.Debug().Delete(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "删除不存在"})
		return
	}
	// fmt.Println("post.Title:", post.Title)
	// fmt.Println("post.Content:", post.Content)
	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}

// 实现评论的创建功能，已认证的用户可以对文章发表评论。
func AddComment(c *gin.Context, db *gorm.DB) {

	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.First(&user, comment.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	var post Post
	if err := db.First(&post, comment.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "对文章发表评论成功"})
}

// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
func ReadCommentList(c *gin.Context, db *gorm.DB) {

	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(post)

	var postResult Post
	if err := db.Debug().Where("ID=?", post.ID).Find(&postResult).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "单个文章的详细信息取得失败"})
		return
	}

	var comment Comment
	if err := db.Debug().Where("post_id=?", post.ID).Find(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章的所有评论取得失败"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "文章的所有评论列表",
		"comment": comment,
	})
}

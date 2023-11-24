package types

import (
	"time"
)

type User struct {
	Id           int
	Email        string
	Username     string
	PasswordHash string
}
type Roles struct {
	Id       int
	Username string
	Role     string
}
type Request struct {
	Id       int
	Username string
	UserId   int
	Status   string
}
type Report struct {
	Id          int
	PostId      int
	PostTitle   string
	ReportType  string
	ModeratorId int
	Response    string
}

type Post struct {
	Id         int
	AuthorId   int
	AuthorName string
	Title      string
	Content    string
	Created    time.Time
	Time       string
	Categories []string
	Likes      int
	Dislikes   int
	Comments   []Comment
	Status     string
}

type Reaction struct {
	Id           int
	PostId       int
	CommentId    int
	UserId       int
	ReactionType string
}

type Comment struct {
	Id       int
	PostId   int
	UserId   int
	Content  string
	Created  time.Time
	Username string
	Time     string
	Likes    int
	Dislikes int
}

type CreateUserData struct {
	Id       int
	Email    string
	Username string
	Password string
	Token    *string
	Expired  *time.Time
}

type GetUserData struct {
	Username string
	Password string
}

type Err struct {
	StatusCode int
	StatusText string
}

type ErrText struct {
	Username string
	Email    string
	Pass1    string
	Pass2    string
}
type Config struct {
	ClientID     string
	ClientSecret string
	Endpoint     string
	Scopes       string
}
type UserService interface {
	CreateUser(user *CreateUserData) error
	CheckUserExists(user *CreateUserData) (bool, ErrText)
	CheckLogin(user *GetUserData) (int, error)
	AddToken(userid int, cookie string) error
	RemoveToken(token string) error
	GetUserByToken(token string) (user *User, err error)
	GetLoginId(user *CreateUserData) int
	GetAllUsers() []*Roles
	GetUserRole(username string) string
	CreateAdministrator()
	UpdateRole(*Roles)
	CreateRequest(*Request)
	GetAllRequests() []*Request
	UpdateRequestStatus(status string, userId string)
	GetRequestStatus(userId int) string
	
}

type CreatePost struct {
	AuthorId   int
	AuthorName string
	Title      string
	Content    string
	Categories []string
}

type PostService interface {
	GetAllPosts() ([]*Post, error)
	CreateNewPost(*CreatePost) (int, error)
	GetPostByID(id int) (*Post, error)
	CreateLike(*Reaction) error
	CreateLikeComment(*Reaction)
	CheckLike(*Reaction) bool
	CheckLikeComment(*Reaction) bool
	CountLikes(int, string) int
	Filter([]string) ([]*Post, error)
	CreateComment(*Comment)
	GetPosts(int) []*Post
	GetPostsByUserID(id int) ([]*Post, error)
	DeletePost(int)
	DeleteComment(int)
	ReportPost(*Report)
	DeleteReport(int)
	ReportComplete(int)
	GetAllReports() []*Report
	ReportResponse(response string, postId int)
	UpdatePostStatus(postId int)
	GetReportsByID(int)[]*Report
}

type UserRepo interface {
	CreateUserDB(user *User)
	GetUserNameDB(user string) error
	GetUserEmailDB(user string) error
	CheckLoginDB(user *GetUserData) (int, error)
	AddTokenDB(userid int, cookieToken string) error
	RemoveTokenDB(token string) error
	GetUserByToken(token string) (user *User, err error)
	GetUserIdDB(user *CreateUserData) int
	CreateUserRoleDB(user *Roles)
	GetAllUsers() []*Roles
	GetUserRole(username string) string
	CreateAdministrator()
	UpdateRole(*Roles)
	CreateRequest(*Request)
	GetAllRequests() []*Request
	UpdateRequestStatus(status string, userId string)
	GetRequestStatus(userId int) string
}

type PostRepo interface {
	GetAllPosts() ([]*Post, error)
	CreatePostDB(*Post) (int, error)
	GetPostByID(id int) (*Post, error)
	CreateLikeDB(*Reaction) error
	CreateLikeCommentDB(*Reaction)
	CheckLikeDB(*Reaction) bool
	CheckLikeCommentDB(*Reaction) bool
	UpdateLikeDB(*Reaction)
	UpdateLikeCommentDB(*Reaction)
	RemoveLikeDB(*Reaction)
	RemoveLikeCommentDB(*Reaction)
	CountLikes(int, string) int
	CountLikesComment(string, int) int
	Filter([]string) ([]*Post, error)
	CreateComment(*Comment)
	GetAllComments(int) []Comment
	GetPostsByUserId(int) []*Post
	GetPostsByUserID(id int) ([]*Post, error)
	DeletePost(int)
	DeleteComment(int)
	CreateReport(*Report)
	DeleteReport(int)
	ReportComplete(int)
	GetAllReports() []*Report
	ReportResponse(response string, postId int)
	UpdatePostStatus(postId int)
	GetReportsByID(int)[]*Report
}

type Category struct {
	Name string
}

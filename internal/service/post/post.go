package post

import (
	"forum/internal/types"
)

type PostService struct {
	repo types.PostRepo
}

func NewPostService(repo types.PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) GetAllPosts() ([]*types.Post, error) {
	return p.repo.GetAllPosts()
}

func (p *PostService) CreateNewPost(postData *types.CreatePost) (int, error) {
	post := &types.Post{
		AuthorId:   postData.AuthorId,
		AuthorName: postData.AuthorName,
		Title:      postData.Title,
		Content:    postData.Content,
		Categories: postData.Categories,
	}
	id, err := p.repo.CreatePostDB(post)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostService) GetPostByID(id int) (*types.Post, error) {
	return p.repo.GetPostByID(id)
}

func (p *PostService) Filter(categories []string) ([]*types.Post, error) {
	return p.repo.Filter(categories)
}

func (p *PostService) GetPostsByUserID(id int) ([]*types.Post, error) {
	return p.repo.GetPostsByUserID(id)
}

func (p *PostService) DeletePost(postId int) {
	p.repo.DeletePost(postId)
}

func (p *PostService) ReportPost(data *types.Report) {
	p.repo.CreateReport(data)
}

func (p *PostService) DeleteReport(postId int) {
	p.repo.DeleteReport(postId)
}

func (p *PostService) ReportComplete(postId int) {
	p.repo.ReportComplete(postId)
}

func (p *PostService) ReportResponse(response string, postId int) {
	p.repo.ReportResponse(response, postId)
}

func (p *PostService) GetAllReports() []*types.Report {
	return p.repo.GetAllReports()
}

func (p *PostService) GetReportsByID(userId int) []*types.Report {
	return p.repo.GetReportsByID(userId)
}

func (p *PostService) UpdatePostStatus(postId int) {
	p.repo.UpdatePostStatus(postId)
}

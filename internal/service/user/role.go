package user

import "forum/internal/types"

func (u *UserService) GetAllUsers() []*types.Roles {
	return u.repo.GetAllUsers()
}

func (u *UserService) GetUserRole(username string) string {
	return u.repo.GetUserRole(username)
}

func (u *UserService) UpdateRole(user *types.Roles) {
	u.repo.UpdateRole(user)
}

func (u *UserService) CreateAdministrator() {
	u.repo.CreateAdministrator()
}

func (u *UserService) CreateRequest(request *types.Request) {
	u.repo.CreateRequest(request)
}

func (u *UserService) GetAllRequests() []*types.Request {
	return u.repo.GetAllRequests()
}

func (u *UserService) UpdateRequestStatus(status string, username string) {
	u.repo.UpdateRequestStatus(status, username)
}

func (u *UserService) GetRequestStatus(userId int) string {
	return u.repo.GetRequestStatus(userId)
}

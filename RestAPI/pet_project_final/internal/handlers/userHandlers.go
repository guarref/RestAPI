package handlers

import (
	"context"
	"pet_project_final/internal/userService"
	"pet_project_final/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

// GetUsersUserIdTasks implements users.StrictServerInterface.
func (u *UserHandler) GetUsersUserIdTasks(ctx context.Context, request users.GetUsersUserIdTasksRequestObject) (users.GetUsersUserIdTasksResponseObject, error) {
	alltasks, err := u.Service.GetTasksUserId(request.UserId)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range alltasks {
		task := users.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

// DeleteUsersId implements users.StrictServerInterface.
func (u *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {

	id := uint(request.Id)

	err := u.Service.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}

// GetUsers implements users.StrictServerInterface.
func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {

	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (u *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := uint(request.Id)

	userToUpdate := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	patchedUser, err := u.Service.UpdateUserByID(id, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &patchedUser.ID,
		Email:    &patchedUser.Email,
		Password: &patchedUser.Password,
	}

	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (u *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

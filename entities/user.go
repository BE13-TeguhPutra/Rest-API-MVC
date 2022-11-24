package entities

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Phone     string
	Password  string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

type UserResponse struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func UserCoretoRespone(dataCore UserCore) UserResponse {
	return UserResponse{
		Name:    dataCore.Name,
		Email:   dataCore.Email,
		Phone:   dataCore.Phone,
		Address: dataCore.Address,
	}

}

func ListUserCoreToResponse(listCore []UserCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range listCore {
		dataResponse = append(dataResponse, UserCoretoRespone(v))
	}
	return dataResponse
}

func UserRequestToCore(req UserRequest) UserCore {
	return UserCore{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
	}
}

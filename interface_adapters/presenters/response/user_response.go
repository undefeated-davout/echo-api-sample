package response

import "undefeated-davout/echo-api-sample/entities"

type AddUserResponse struct {
	ID entities.UserID `json:"id"`
}

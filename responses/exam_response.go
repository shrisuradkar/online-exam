package responses

import "time"

type ExamResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	First_name   *string   `json:"first_name"`
	Last_name    *string   `json:"last_name"`
	Email        *string   `json:"email"`
	Phone        *string   `json:"phone"`
	Course       *string   `json:"course"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Candidate_id string    `json:"user_id"`
}

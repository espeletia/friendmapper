package handlers

import (
	"encoding/json"
	"hradec/internal/domain"
	"hradec/internal/handlers/models"
	"hradec/internal/usecases"
	"net/http"

	"go.uber.org/zap"
)

type UserHandler struct {
	userUsecase *usecases.UserUsecase
	authUsecase *usecases.AuthUsecase
}

func NewUserHandler(userusecase *usecases.UserUsecase, authUsecase *usecases.AuthUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userusecase,
		authUsecase: authUsecase,
	}
}

func (uu *UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Decode the JSON body to the `Viewport` struct
		var userData models.UserData
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		zap.L().Info("userData", zap.Any("userData", userData))

		defer r.Body.Close()

		usr, err := uu.userUsecase.CreateUser(ctx, mapModelUserDataToDomainUserData(userData), userData.Password)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		zap.L().Info("HIT", zap.Any("places", usr))
		// Convert the result to JSON and write to the response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(mapDomainUserToModelUser(*usr))
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}

	}
}

func (uu *UserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Decode the JSON body to the `Viewport` struct
		var creds models.LoginCreds
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		zap.L().Info("userData", zap.Any("userData", creds))

		defer r.Body.Close()

		usr, err := uu.authUsecase.Login(ctx, mapModelUserCreds(creds))
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		zap.L().Info("HIT", zap.Any("places", usr))
		// Convert the result to JSON and write to the response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(models.LoginResp{
			Token: usr,
		})
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}

	}
}

func (uu *UserHandler) GetUsersByUsernamePattern() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		usernamePattern := r.URL.Query().Get("username")
		if usernamePattern == "" {
			http.Error(w, "username pattern is required", http.StatusBadRequest)
			return
		}

		// Call the usecase to get users by username pattern
		users, err := uu.userUsecase.GetByUsernameSimilar(ctx, usernamePattern)
		if err != nil {
			http.Error(w, "failed to get users", http.StatusInternalServerError)
			return
		}

		// Map domain users to model users
		modelUsers := make([]models.User, len(users))
		for i, u := range users {
			modelUsers[i] = mapDomainUserToModelUser(u)
		}

		// Convert the result to JSON and write to the response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(modelUsers)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}

func mapModelUserCreds(usr models.LoginCreds) domain.LoginCreds {
	return domain.LoginCreds{
		Email:    usr.Email,
		Password: usr.Password,
	}
}

func mapModelUserDataToDomainUserData(usr models.UserData) domain.UserData {
	return domain.UserData{
		Username:    usr.Username,
		Email:       usr.Email,
		DisplayName: usr.DisplayName,
	}
}

func mapDomainUserToModelUser(usr domain.User) models.User {
	return models.User{
		ID:          usr.ID,
		DisplayName: usr.DisplayName,
		Email:       usr.Email,
		Username:    usr.Username,
	}
}

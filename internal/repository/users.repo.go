package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/oktaviandwip/musalabel/config"
	models "github.com/oktaviandwip/musalabel/internal/models/users"
)

type RepoUsersIF interface {
	CreateUser(data *models.User) (*config.Result, error)
	GetPassByEmail(email string) (*models.User, error)
	UpdateProfile(user *models.User) (*config.Result, error)
	// FetchProfile(userId string) (*config.Result, error)
	// FetchProfileForHeader(userId string) (*config.Result, error)

	// Update(data *models.User, user_id string) (*config.Result, error)
	// Delete(data *models.User) (*config.Result, error)
}

type RepoUsers struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db}
}

// Create User
func (r *RepoUsers) CreateUser(data *models.User) (*config.Result, error) {
	q := `INSERT INTO users (
					email, 
					password, 
					phone_number,
					role
				)
				VALUES(
					:email,
					:password,
					:phone_number,
					:role
				)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return nil, errors.New("email has been used")
		}
		return nil, err
	}

	return &config.Result{Message: "1 data user created"}, nil
}

// Login
func (r *RepoUsers) GetPassByEmail(email string) (*models.User, error) {
	var result models.User
	q := `SELECT * FROM users WHERE email = $1`

	if err := r.Get(&result, q, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("email not found")
		}
		return nil, err
	}

	return &result, nil
}

// Update Profile
func (r *RepoUsers) UpdateProfile(data *models.User) (*config.Result, error) {
	q := `
		UPDATE users
		SET
			image = $1,
			email = COALESCE(NULLIF($2, ''), email),
			phone_number = COALESCE(NULLIF($3, ''), phone_number),
			address = COALESCE(NULLIF($4, ''), address),
			display_name = COALESCE(NULLIF($5, ''), display_name),
			first_name = COALESCE(NULLIF($6, ''), first_name),
			last_name = COALESCE(NULLIF($7, ''), last_name),
			birthday = COALESCE(CAST(NULLIF($8, '') AS DATE), birthday),
			gender = COALESCE(NULLIF($9, ''), gender),
			updated_at = NOW()
		WHERE
			id = CAST($10 AS UUID)
		RETURNING id, image, email, phone_number, role, address, display_name, first_name, last_name, birthday, gender;
	`

	args := []interface{}{
		data.Image,
		data.Email,
		data.Phone_number,
		data.Address,
		data.Display_name,
		data.First_name,
		data.Last_name,
		data.Birthday,
		data.Gender,
		data.Id,
	}

	var user models.User
	err := r.QueryRowx(q, args...).StructScan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &config.Result{Data: user}, nil
}

// // get for header
// func (r *RepoUsers) FetchProfileForHeader(userId string) (*config.Result, error) {
// 	var result models.UserProfileHeader

// 	q := `select u.email , p.display_name , p.photo_profile
// 			from users u join profile p on u.user_id = p.user_id
// 			where u.user_id = ?`

// 	if err := r.Get(&result, r.Rebind(q), userId); err != nil {
// 		return nil, err
// 	}

// 	return &config.Result{Data: result}, nil
// }

// func (r *RepoUsers) FetchProfile(userId string) (*config.Result, error) {
// 	// Get data user
// 	userQuery := "SELECT email, phone_number FROM users WHERE user_id = ?"
// 	userData := models.UserData{}

// 	if err := r.Get(&userData, r.Rebind(userQuery), userId); err != nil {
// 		return nil, err
// 	}

// 	// Get data profile
// 	profileQuery := "SELECT photo_profile, address, display_name, first_name, last_name, gender, birthday FROM profile WHERE user_id = ?"
// 	profileData := models.Profile{}

// 	if err := r.Get(&profileData, r.Rebind(profileQuery), userId); err != nil {
// 		return nil, err
// 	}

// 	// Merge result
// 	data := map[string]interface{}{
// 		"photo_profile": profileData.Photo_profile,
// 		"email":         userData.Email,
// 		"phone_number":  userData.Phone_number,
// 		"address":       profileData.Address,
// 		"display_name":  profileData.Display_name,
// 		"first_name":    profileData.First_name,
// 		"last_name":     profileData.Last_name,
// 		"gender":        profileData.Gender,
// 		"birthday":      profileData.Birthday,
// 	}

// 	return &config.Result{Data: data}, nil
// }

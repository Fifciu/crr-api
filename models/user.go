package models

import (
	"database/sql/driver"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	u "github.com/filipjedrasik/crr-api/utils"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	TicketExpires NullTime  `json:"ticketExpires"`
	CreatedAt     time.Time `json:"createdAt"`
	LastVisit     time.Time `json:"lastVisit"`
	Token         string    `gorm:"-" json:"token"`
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (user User) TableName() string {
	return "user"
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Adres email jest wymagany"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Hasło jest wymagane"), false
	}

	if len(user.Name) < 4 {
		return u.Message(false, "Imię i nazwisko jest wymagane"), false
	}

	temp := &User{}

	err := GetDB().Table("user").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Nie mogłem połączyć się z bazą danych"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email jest już zajęty"), false
	}

	return u.Message(false, "Udało się zwalidować"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.LastVisit = time.Now().UTC()
	user.TicketExpires = NullTime{
		Valid: false,
	}

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Nie udało się utworzyć konta")
	}

	newTime, err := SaveVisit(user.ID)
	if err != nil {
		return u.Message(false, "Wystąpił błąd")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = ""
	user.LastVisit = newTime

	response := u.Message(true, "Konto zostało utworzone")
	response["user"] = user
	return response
}

func SaveVisit(userId uint) (time.Time, error) {
	newTime := time.Now().UTC()
	user := &User{}
	err := GetDB().Table("user").Where("id = ?", userId).First(user).Error
	if err != nil {
		return newTime, errors.New("Wystąpił błąd")
	}

	err = GetDB().Model(&user).Update("last_visit", newTime).Error
	if err != nil {
		return newTime, errors.New("Wystąpił błąd")
	}

	return newTime, nil
}

func Login(email, password string) map[string]interface{} {
	user := &User{}
	err := GetDB().Table("user").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email nie istnieje")
		}
		return u.Message(false, "Nie udało się połączyć z bazą danych")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Błędne hasło")
	}

	// Logged in
	newTime, err := SaveVisit(user.ID)
	if err != nil {
		return u.Message(false, "Wystąpił błąd")
	}

	user.Password = ""
	user.LastVisit = newTime

	// JWT
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	response := u.Message(true, "Zalogowano")
	response["user"] = user
	return response
}

func GetUser(u uint) *User {
	user := &User{}
	err := GetDB().Table("user").Where("id = ?", u).First(user).Error
	if err != nil || user.Email == "" {
		return nil
	}

	_, err = SaveVisit(u)
	if err != nil {
		return nil
	}

	user.Password = ""
	return user
}

func BuyTicket(u uint) time.Time {
	user := &User{}
	err := GetDB().Table("user").Where("id = ?", u).First(user).Error
	if err != nil {
		return time.Time{}
	}

	if user.TicketExpires.Time.IsZero() {
		user.TicketExpires = NullTime{
			Time:  time.Now().AddDate(0, 0, 30).UTC(),
			Valid: true,
		}
	} else {
		user.TicketExpires = NullTime{
			Time:  user.TicketExpires.Time.AddDate(0, 0, 30).UTC(),
			Valid: true,
		}
	}

	GetDB().Save(&user)

	return user.TicketExpires.Time
}

package serializers

import "github.com/iReflect/reflect-app/apps/user/models"

// UserAuthSerializer ...
type UserAuthSerializer struct {
	models.User
	Token string
}

// User ...
type User struct {
	ID        uint
	Email     string
	FirstName string
	LastName  string
	Active    bool
}

// Team ...
type Team struct {
	ID          uint
	Name        string
	Description string
	Active      bool
}

// TeamsSerializer ...
type TeamsSerializer struct {
	Teams []Team
}

// MembersSerializer ...
type MembersSerializer struct {
	Members []User
}

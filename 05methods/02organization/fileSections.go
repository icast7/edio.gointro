package main

// 1. List packages
import "fmt"

// 2. List constants
const (
	ConstExample = "const before variables"
)

// 3. List of variables
var (
	ExportedVariable    = 42
	nonExportedVariable = "so say we all"
)

// 4. Main type(s) for the file, minimize number of structs per file´
type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}

// 5. List of functions
func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Location: &UserLocation{
			City:    "PDX",
			Country: "USA",
		},
	}
}

// 6. List of methods (a.k.a. functions with receivers)
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

// 7. MAIN
func main() {
	us := User{
		FirstName: "John",
		LastName:  "Doe",
		Location: &UserLocation{
			City:    "PDX",
			Country: "USA",
		},
	}
	fmt.Println(us.Greeting())
	fmt.Printf(">>> %+v", nonExportedVariable)
}

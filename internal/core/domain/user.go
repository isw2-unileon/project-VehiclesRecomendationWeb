package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`    // El "-" oculta contraseña al enviar JSON
	Role     string `json:"role"` // "admin" o "user"
}

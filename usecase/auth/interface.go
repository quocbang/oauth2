package auth

type IOAuth2 interface {
	Login() error
}

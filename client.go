package ccxp

type Client interface {
	Login() error
}

type client struct {
	username  string
	password  string
	pwdstr    string
	acixstore string
}

func New(username, password string) Client {
	return &client{
		username: username,
		password: password,
	}
}

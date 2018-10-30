package core

type User interface {
	IsOwner() bool
	Owner() Owner
	Name() string
	Email() string
	PrimaryPublicKey() []byte
	SecondaryPublicKeys() [][]byte
	Level() int32
}

type Owner interface {
	Weight() int32
}

type SearchOpts struct {
	Name               string
	Email              string
	PrimaryPublicKey   []byte
	SecondaryPublicKey []byte
	Level              int32
	IsOwner            bool
}

type UserManager interface {
	GetUsers(SearchOpts) ([]User, error)
	GetUser(SearchOpts)
	CreateUser(User)
}

package material

type Material struct {
	AccessToken string
}

func New(token string) *Material {
	o := &Material{
		AccessToken: token}
	return o
}

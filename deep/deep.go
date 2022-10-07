package deep

type Deep struct {
	Name string
}

func (d *Deep) Exec() (string, error) {
	return "It was returned from deep level", nil
}

package server

type Command struct {
	id     int8
	client *Client
	args   []string
	db     string
}
type Procesor struct {
	db *map[string]map[string]string
}

func (p *Procesor) process(command chan<- Command) {}

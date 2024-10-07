package server
type Command struct{
  id string 
  client *Client
  args []string
  db string
}
type Procesor struct{
 db  *map[string]map[string]string
}
func (p *Procesor) process(chan command<-Command){}

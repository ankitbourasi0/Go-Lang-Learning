package todo

//service layer

// todos []string , here todos start from lowercase in GO, it is private.
// if i change it to Todos so it makes it Public so i can access it without any error.
type Service struct {
	todos []string
}

// to generate the things press left alt+enter
func NewService() *Service {
	return &Service{todos: make([]string, 0)}
}

/*
		method add takes a todo of type string, it append todo in todos list. created in Service strcut

	    (svc *Service) is called reciever, A receiver a Pointer to Service struct, is special paramter of a method
	    that gives a method access to the fields within the struct.
	    For example if you remove that syntax you won't able to acess svc or todos within the method
*/
func (svc *Service) Add(todo string) {
	svc.todos = append(svc.todos, todo)
}

func (svc *Service) GetAll() []string {
	return svc.todos
}

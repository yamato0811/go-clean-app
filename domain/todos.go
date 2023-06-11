package domain

type Todo struct {
	Id TodoId `json:"id"`
	Title TodoTitle `json:"title"`
	Completed TodoCompleted `json:"completed"`
}

type TodoId struct {
	Value int `json:"value"`
}

type TodoTitle struct {
	Value string `json:"value"`
}

type TodoCompleted struct {
	Value bool `json:"value"`
}
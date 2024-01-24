package model

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MemStore struct {
	data map[string]Todo
}

func (m MemStore) Add(id string, todo Todo) error {
	//TODO Реализовать проверку данных коректность
	m.data[id] = todo
	return nil
}

func (m MemStore) Get(id string) (Todo, error) {
	todo := m.data[id]
	return todo, nil
}

func (m MemStore) List() (map[string]Todo, error) {
	return m.data, nil
}
func (m MemStore) Update(id string, todo Todo) error {
	m.data[id] = todo
	return nil
}
func (m MemStore) Remove(id string) error {
	//TODO Сделать проверку на наличие данного элемента в базе
	delete(m.data, id)
	return nil
}

function TodoForm({ todos, setTodos }) {
  const handleSubmit = (event) => {
    event.preventDefault();
    const value = event.target.todo.value;
    const newTodo = {
      title: value,
      id: self.crypto.randomUUID(),
      is_completed: false,
    };

    setTodos((prevTodos) => [...prevTodos, newTodo]);
    const updatedTodoList = JSON.stringify([...todos, newTodo]);
    localStorage.setItem("todos", updatedTodoList);
    event.target.reset();
  };
  return (
    <form className="form" onSubmit={handleSubmit}>
      <label htmlFor="todo">
        <input
          type="text"
          name="todo"
          id="todo"
          placeholder="Write your next task"
        />
      </label>
      <button>
        <span className="visually-hidden"></span>
        <img
          width="24"
          height="24"
          src="https://img.icons8.com/android/24/FFFFFF/plus.png"
          alt="plus"
        />
      </button>
    </form>
  );
}
export default TodoForm;

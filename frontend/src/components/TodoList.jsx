// src/components/TodoList.jsx

function TodoList({ todos }) {
  // src/components/TodoList.jsx

  function Item({ item }) {
    return (
      <li id={item?.id} className="todo_item">
        <button className="todo_items_left">
          <svg>
            <circle cx="11.998" cy="11.998" fillRule="nonzero" r="9.998" />
          </svg>
          <p>{item?.title}</p>
        </button>
        <div className="todo_items_right">
          <button>
            <span className="visually-hidden">Edit</span>
            <svg>
              <path d="" />
            </svg>
          </button>
          <button>
            <span className="visually-hidden">Delete</span>
            <svg>
              <path d="" />
            </svg>
          </button>
        </div>
      </li>
    );
  }

  // src/components/TodoList.jsx

  function TODOList({ todos }) {
    return (
      <ol className="todo_list">
        {todos && todos.length > 0 ? (
          todos?.map((item, index) => (
            <Item key={index} item={item} setTodos={setTodos} />
          ))
        ) : (
          <p>Seems lonely in here, what are you up to?</p>
        )}
      </ol>
    );
  }
}
export default TodoList;

import React, { useEffect, useRef, useState } from "react";

function Item({ item, setTodos }) {
  const [editing, setEditing] = useState(false);
  const [editValue, setEditValue] = useState(item.title);
  const inputRef = useRef(null);

  useEffect(() => {
    if (editing && inputRef.current) {
      inputRef.current.focus();
      inputRef.current.setSelectionRange(
        inputRef.current.value.length,
        inputRef.current.value.length
      );
    }
  }, [editing]);

  const completeTodo = () => {
    setTodos((prevTodos) =>
      prevTodos.map((todo) =>
        todo.id === item.id
          ? { ...todo, is_completed: !todo.is_completed }
          : todo
      )
    );
    updateLocalStorage();
  };

  const handleEdit = () => {
    setEditing(true);
  };

  const handleInputChange = (event) => {
    setEditValue(event.target.value);
  };

  const handleInputSubmit = (event) => {
    event.preventDefault();
    setTodos((prevTodos) =>
      prevTodos.map((todo) =>
        todo.id === item.id ? { ...todo, title: editValue } : todo
      )
    );
    updateLocalStorage();
    setEditing(false);
  };

  const handleInputBlur = () => {
    setTodos((prevTodos) =>
      prevTodos.map((todo) =>
        todo.id === item.id ? { ...todo, title: editValue } : todo
      )
    );
    updateLocalStorage();
    setEditing(false);
  };

  const handleDelete = () => {
    setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== item.id));
    updateLocalStorage();
  };

  const updateLocalStorage = () => {
    const storedTodos = JSON.parse(localStorage.getItem("todos")) || [];
    localStorage.setItem("todos", JSON.stringify(storedTodos));
  };

  return (
    <tr>
      <td>
        {editing ? (
          <form className="edit-form" onSubmit={handleInputSubmit}>
            <input
              ref={inputRef}
              type="text"
              value={editValue}
              onBlur={handleInputBlur}
              onChange={handleInputChange}
            />
          </form>
        ) : (
          <button className="todo_items_left" onClick={completeTodo}>
            <svg
              fill={item.is_completed ? "#22C55E" : "#0d0d0d"}
              viewBox="0 0 24 24"
              width="24"
              height="24"
              xmlns="http://www.w3.org/2000/svg"
            >
              {item.is_completed ? (
                <path d="M12 0C5.37258 0 0 5.37258 0 12C0 18.6274 5.37258 24 12 24C18.6274 24 24 12C24 5.37258 18.6274 0 12 0ZM10.2 16.2L5.4 11.4L6.6 10.2L10.2 13.8L17.4 6.6L18.6 7.8L10.2 16.2Z" />
              ) : (
                <circle cx="12" cy="12" r="9.998" />
              )}
            </svg>
            <p
              style={
                item.is_completed ? { textDecoration: "line-through" } : {}
              }
            >
              {item?.title}
            </p>
          </button>
        )}
      </td>
      <td className="todo_items_right">
        <button onClick={handleEdit}>
          <img
            width="24"
            height="24"
            src="https://img.icons8.com/carbon-copy/100/FFFFFF/create-new.png"
            alt="edit"
          />
        </button>
        <button onClick={handleDelete}>
          <img
            width="24"
            height="24"
            src="https://img.icons8.com/carbon-copy/100/FFFFFF/filled-trash.png"
            alt="delete"
          />
        </button>
      </td>
    </tr>
  );
}

export default Item;

import React from "react";
import Item from "./TodoItem";

function TodoList({ todos, setTodos }) {
  return (
    <table className="todo_table">
      <thead>
        <tr>
          <th>Task</th>
          <th>Edit / Hapus</th>
        </tr>
      </thead>
      <tbody>
        {todos && todos.length > 0 ? (
          todos.map((item, index) => (
            <Item key={index} item={item} setTodos={setTodos} />
          ))
        ) : (
          <tr>
            <td colSpan="2">Seems lonely in here, what are you up to?</td>
          </tr>
        )}
      </tbody>
    </table>
  );
}

export default TodoList;

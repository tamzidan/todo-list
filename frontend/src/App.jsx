"use client";
import React from "react";
import TodoForm from "../src/components/TodoForm";
import TodoHeader from "../src/components/TodoHeader";
import TodoHero from "../src/components/TodoHero";
import TodoList from "../src/components/TodoList";
import "./App.css";

function App() {
  const [todos, setTodos] = React.useState([]);

  React.useEffect(() => {
    const storedTodos = localStorage.getItem("todos");
    if (storedTodos) {
      setTodos(JSON.parse(storedTodos));
    }
  }, []);

  const todos_completed = todos.filter(
    (todo) => todo.is_completed === true
  ).length;
  const total_todos = todos.length;

  return (
    <div className="wrapper">
      <TodoHeader />
      <TodoHero todos_completed={todos_completed} total_todos={total_todos} />
      <TodoForm todos={todos} setTodos={setTodos} />
      <TodoList todos={todos} setTodos={setTodos} />
    </div>
  );
}
export default App;

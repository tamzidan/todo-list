// src/app/page.js

"use client";
import React from "react";
import TodoForm from "./components/TodoForm";
import TodoHeader from "./components/TodoHeader";
import TodoHero from "./components/TodoHero";
import TodoList from "./components/TodoList";
// import TodoItem from "./components/TodoItem";

import "./App.css";

function App() {
  const [todos, setTodos] = React.useState([
    { title: "Some task", id: self.crypto.randomUUID(), is_completed: false },
    {
      title: "Some other task",
      id: self.crypto.randomUUID(),
      is_completed: true,
    },
    { title: "last task", id: self.crypto.randomUUID(), is_completed: false },
  ]);

  const todos_completed = todos.filter(
    (todo) => todo.is_completed === true
  ).length;
  const total_todos = todos.length;

  return (
    <div className="wrapper">
      <TodoHeader />
      <TodoHero todos_completed={todos_completed} total_todos={total_todos} />
      <TodoForm />
      <TodoList todos={todos} />
    </div>
  );
}
export default App;

"use client";

import { useState, useEffect } from "react";
import TaskTable from "@/components/TaskTable";
import TaskToolbar from "@/components/TaskToolbar";
import { Task } from "@/types";
import { getTasks } from "./api/tasks/route";
import AddTask from "./api/tasks/addTask";
import HandleDelete from "./api/tasks/deleteTask";
import UpdateTask from "./api/tasks/updateTask";

export default function Home() {
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    getTasks().then(setTasks);
  }, []);

  const onAdd = async (title: string, description: string) => {
    const newTask = await AddTask(title, description);
    setTasks((prev) => [...prev, newTask]);
  };

  const onDelete = async (id: string) => {
    await HandleDelete(id);
    setTasks((prev) => prev.filter((t) => t.id !== id));
  };

  const onUpdateTask = async (
    id: string,
    dto: { title?: string; description?: string; status?: boolean }
  ) => {
    try {
      const updated = await UpdateTask(id, dto);
      setTasks((prev) =>
        prev.map((t) => (t.id === id ? { ...t, ...updated } : t))
      );
    } catch (err) {
      console.error("Ошибка при обновлении задачи:", err);
    }
  };

  return (
    <div className="w-3/4 flex flex-col items-center justify-center m-8 gap-4">
      <TaskToolbar onAdd={onAdd} />
      <TaskTable
        data={tasks}
        setData={setTasks}
        onDelete={onDelete}
        onUpdateTask={onUpdateTask}
      />
    </div>
  );
}

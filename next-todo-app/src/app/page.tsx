"use client";

import { useState, useEffect } from "react";
import TaskTable from "@/components/TaskTable";
import TaskToolbar from "@/components/TaskToolbar";
import { getTasks } from "./api/tasks/route";
import { Task } from "@/types";
import HandleDelete from "./api/tasks/deleteTask";
import AddTask from "./api/tasks/addTask";
import UpdateTaskStatus from "./api/tasks/updateStatus";

export default function Home() {
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    getTasks().then(setTasks);
  }, []);

  const onDelete = async (id: string) => {
    await HandleDelete(id);
    setTasks((prev) => prev.filter((t) => t.id !== id));
  };

  const onAdd = async (title: string, description: string) => {
    const newTask = await AddTask(title, description);
    setTasks((prev) => [...prev, newTask]);
  };

  const onToggleStatus = async (id: string) => {
    const updatedTask = await UpdateTaskStatus(id);
    if (!updatedTask) return;

    setTasks((prev) =>
      prev.map((t) =>
        t.id === updatedTask.id
          ? {
              ...t,
              status: updatedTask.status,
              completed_at: updatedTask.completed_at,
            }
          : t
      )
    );
  };

  return (
    <div className="w-3/4 flex flex-col items-center justify-center m-8 gap-4">
      <TaskToolbar onAdd={onAdd} />
      <TaskTable
        data={tasks}
        onDelete={onDelete}
        onToggleStatus={onToggleStatus}
      />
    </div>
  );
}

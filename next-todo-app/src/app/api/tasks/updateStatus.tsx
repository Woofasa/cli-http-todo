import { Task } from "@/types";

export default async function UpdateTaskStatus(id: string): Promise<Task> {
  const res = await fetch(`http://localhost:3001/?id=${id}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to update task");
  }

  const updatedTask: Task = await res.json();
  return updatedTask;
}

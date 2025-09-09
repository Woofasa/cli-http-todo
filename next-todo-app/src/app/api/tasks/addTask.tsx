import { Task } from "@/types";

export default async function AddTask(
  title: string,
  description: string
): Promise<Task> {
  const res = await fetch("http://localhost:3001", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title, description }),
  });

  if (!res.ok) {
    throw new Error("Не удалось создать задачу");
  }

  const task: Task = await res.json();
  return task;
}

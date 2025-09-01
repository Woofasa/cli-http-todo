import { Task } from "@/types";

export async function getTasks(): Promise<Task[]> {
  const res = await fetch("http://localhost:3001/", { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch tasks");
  const data: Task[] = await res.json();
  return data;
}

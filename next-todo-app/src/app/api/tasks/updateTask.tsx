export default async function UpdateTask(
  id: string,
  dto: { title?: string; description?: string; status?: boolean }
) {
  const res = await fetch(`http://localhost:3001/?id=${id}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(dto),
  });

  if (!res.ok) {
    throw new Error("Failed to update task");
  }

  return res.json();
}

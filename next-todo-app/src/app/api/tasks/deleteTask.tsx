export default async function handleDelete(id: string) {
  try {
    const res = await fetch(`http://localhost:3001/?id=${id}`, {
      method: "DELETE",
    });

    if (!res.ok) {
      throw new Error("Failed to delete task");
    }
  } catch (err) {
    console.error(err);
  }
}

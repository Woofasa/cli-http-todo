import TaskTable from "@/components/TaskTable";
import TaskToolbar from "@/components/TaskToolbar";
import { getTasks } from "./api/tasks/route";

export default async function Home() {
  const tasks = await getTasks();
  return (
    <div className="w-3/4 flex flex-col items-center justify-center m-8 gap-4">
      <TaskToolbar></TaskToolbar>
      <TaskTable tasks={tasks}></TaskTable>
    </div>
  );
}

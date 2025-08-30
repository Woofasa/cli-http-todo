import TaskTable from "@/components/TaskTable";
import TaskToolbar from "@/components/TaskToolbar";

export default function Home() {
  return (
    <div className="w-3/4 flex flex-col items-center justify-center m-8 gap-4">
      <TaskToolbar></TaskToolbar>
      <TaskTable></TaskTable>
    </div>
  );
}

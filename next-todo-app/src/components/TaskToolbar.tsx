import { Input } from "./ui/input";
import { AddDialog } from "./AddDialog";

export default function TaskToolbar() {
  return (
    <div className="w-full flex justify-between">
      <div>
        <Input type="text" placeholder="Filter task..."></Input>
      </div>
      <AddDialog></AddDialog>
    </div>
  );
}

import { Input } from "./ui/input";
import { AddDialog } from "./AddDialog";

interface Props {
  onAdd: (title: string, description: string) => void;
}

export default function TaskToolbar({ onAdd }: Props) {
  return (
    <div className="w-full flex justify-between">
      <div>
        <Input type="text" placeholder="Filter task..." />
      </div>
      <AddDialog onAdd={onAdd} />
    </div>
  );
}

import { Input } from "./ui/input";
import { Button } from "./ui/button";

export default function TaskToolbar() {
  return (
    <div className="w-full flex justify-between">
      <div>
        <Input type="text" placeholder="Filter task..."></Input>
      </div>
      <Button>Add</Button>
    </div>
  );
}

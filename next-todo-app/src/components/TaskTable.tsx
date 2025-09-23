import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { RectangleEllipsis } from "lucide-react";
import { Task } from "@/types";

type Props = {
  data: Task[];
  setData: React.Dispatch<React.SetStateAction<Task[]>>;
  onDelete: (id: string) => void;
  onUpdateTask: (
    id: string,
    dto: { title?: string; description?: string; status?: boolean }
  ) => void;
};

export default function TaskTable({
  data,
  setData,
  onDelete,
  onUpdateTask,
}: Props) {
  const [editing, setEditing] = useState<{
    id: string;
    field: "title" | "description";
  } | null>(null);
  const [value, setValue] = useState("");

  const handleUpdateTask = async (
    id: string,
    dto: { title?: string; description?: string; status?: boolean }
  ) => {
    try {
      await onUpdateTask(id, dto);
    } catch (err) {
      console.error("Ошибка при обновлении задачи:", err);
    }
  };

  return (
    <div className="w-full border-2 border-accent rounded-b-xs">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Title</TableHead>
            <TableHead>Description</TableHead>
            <TableHead>Status</TableHead>
            <TableHead className="w-42">Created at</TableHead>
            <TableHead className="w-42">Completed at</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>

        <TableBody>
          {data.map((task) => (
            <TableRow key={task.id}>
              {/* Title */}
              <TableCell>
                {editing?.id === task.id && editing.field === "title" ? (
                  <div className="flex items-center gap-2">
                    <input
                      className="border px-2 py-1 text-sm rounded"
                      value={value}
                      onChange={(e) => setValue(e.target.value)}
                      autoFocus
                      onBlur={() => setEditing(null)}
                    />
                    <Button
                      size="sm"
                      onMouseDown={(e) => e.preventDefault()}
                      onClick={() => {
                        handleUpdateTask(task.id, { title: value });
                        setEditing(null);
                      }}
                    >
                      OK
                    </Button>
                  </div>
                ) : (
                  <div className="flex items-center">
                    <span className="flex-1 text-sm font-medium">
                      {task.title}
                    </span>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-7 w-7"
                      onClick={() => {
                        setEditing({ id: task.id, field: "title" });
                        setValue(task.title);
                      }}
                    >
                      <RectangleEllipsis
                        className="h-4 w-4"
                        strokeWidth={1.5}
                      />
                    </Button>
                  </div>
                )}
              </TableCell>

              {/* Description */}
              <TableCell>
                {editing?.id === task.id && editing.field === "description" ? (
                  <div className="flex items-center gap-2">
                    <input
                      className="border px-2 py-1 text-sm rounded w-full"
                      value={value}
                      onChange={(e) => setValue(e.target.value)}
                      autoFocus
                      onBlur={() => setEditing(null)}
                    />
                    <Button
                      size="sm"
                      onMouseDown={(e) => e.preventDefault()}
                      onClick={() => {
                        handleUpdateTask(task.id, { description: value });
                        setEditing(null);
                      }}
                    >
                      OK
                    </Button>
                  </div>
                ) : (
                  <div className="flex items-center">
                    <span className="flex-1 text-sm font-medium">
                      {task.description}
                    </span>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-7 w-7"
                      onClick={() => {
                        setEditing({ id: task.id, field: "description" });
                        setValue(task.description);
                      }}
                    >
                      <RectangleEllipsis
                        className="h-4 w-4"
                        strokeWidth={1.5}
                      />
                    </Button>
                  </div>
                )}
              </TableCell>

              {/* Status */}
              <TableCell>
                <Button
                  variant="outline"
                  size="sm"
                  className={`w-24 justify-center ${
                    task.status
                      ? "border-green-500 text-green-600 hover:bg-green-50"
                      : "border-red-500 text-red-600 hover:bg-red-50"
                  }`}
                  onClick={() =>
                    handleUpdateTask(task.id, { status: !task.status })
                  }
                >
                  {task.status ? "Opened" : "Closed"}
                </Button>
              </TableCell>

              {/* Dates */}
              <TableCell className="w-42 text-sm">
                {new Date(task.created_at).toLocaleString("ru-RU", {
                  year: "numeric",
                  month: "2-digit",
                  day: "2-digit",
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </TableCell>
              <TableCell className="w-42 text-sm">
                {task.completed_at
                  ? new Date(task.completed_at).toLocaleString("ru-RU", {
                      year: "numeric",
                      month: "2-digit",
                      day: "2-digit",
                      hour: "2-digit",
                      minute: "2-digit",
                    })
                  : "-"}
              </TableCell>

              {/* Actions */}
              <TableCell>
                <Button variant="ghost" onClick={() => onDelete(task.id)}>
                  X
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}

"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Task } from "@/types";
import { Button } from "./ui/button";

interface Props {
  data: Task[];
  onDelete: (id: string) => void;
  onToggleStatus: (id: string) => void;
}

export default function TaskTable({ data, onDelete, onToggleStatus }: Props) {
  return (
    <div className="w-full border-2 border-accent rounded-b-xs">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Title</TableHead>
            <TableHead>Description</TableHead>
            <TableHead>Status</TableHead>
            <TableHead>Created at</TableHead>
            <TableHead>Completed at</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {data.map((task) => (
            <TableRow key={task.id}>
              <TableCell>{task.title}</TableCell>
              <TableCell>{task.description}</TableCell>
              <TableCell>
                <Button onClick={() => onToggleStatus(task.id)}>
                  {task.status ? "Opened" : "Closed"}
                </Button>
              </TableCell>
              <TableCell>
                {new Date(task.created_at).toLocaleString()}
              </TableCell>
              <TableCell>
                {task.completed_at
                  ? new Date(task.completed_at).toLocaleString()
                  : "-"}
              </TableCell>
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

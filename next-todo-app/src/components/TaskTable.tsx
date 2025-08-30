import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

export default function TaskTable() {
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
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow>
            <TableCell>Помыть попу</TableCell>
            <TableCell>С мочалкой и мылом</TableCell>
            <TableCell>Opened</TableCell>
            <TableCell>21.02.2001</TableCell>
            <TableCell>-</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>Поиграть в лол</TableCell>
            <TableCell>Пострадать</TableCell>
            <TableCell>Opened</TableCell>
            <TableCell>21.02.2025</TableCell>
            <TableCell>-</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  );
}

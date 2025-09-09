"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

interface Props {
  onAdd: (title: string, description: string) => void;
}

export function AddDialog({ onAdd }: Props) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [error, setError] = useState("");

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    if (!title.trim()) {
      setError("Title cannot be empty");
      return;
    }

    onAdd(title, description);

    setError("");
    setTitle("");
    setDescription("");
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="default">Add</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <form onSubmit={handleSubmit}>
          <DialogHeader>
            <DialogTitle>Create a new task.</DialogTitle>
          </DialogHeader>

          <div className="grid gap-4">
            <div className="grid gap-3">
              <Label htmlFor="title"></Label>
              <Input
                id="title"
                placeholder="Помыть попу"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
              />
              {error && <p className="text-red-500 text-sm">{error}</p>}
            </div>
            <div className="grid gap-3">
              <Label htmlFor="description"></Label>
              <Input
                id="description"
                placeholder="В тазике с мочалкой"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
              />
            </div>
          </div>

          <DialogFooter className="mt-5">
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
            <DialogClose asChild>
              <Button type="submit" disabled={!title.trim()}>
                Apply
              </Button>
            </DialogClose>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}

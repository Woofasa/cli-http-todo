export interface Task {
  id: string;
  title: string;
  description: string;
  status: boolean;
  created_at: string;
  completed_at?: string;
}

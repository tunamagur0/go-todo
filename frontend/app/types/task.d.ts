export type TodoStatus = 0 | 1 | 2 | 3;

export interface Todo {
  id: string;
  content: string;
  status: TodoStatus;
  createdAt: Date;
  finishedAt: Date | null;
}

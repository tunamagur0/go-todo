export const enum TodoStatus {
  statusNew,
  statusWIP,
  statusDone,
  statusPending,
  statusUnknown,
}

export interface Todo {
  id: string;
  content: string;
  status: TodoStatus;
  createdAt: Date;
  finishedAt: Date | null;
}

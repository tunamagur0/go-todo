import { GetterTree, ActionTree, MutationTree } from 'vuex';
import { Todo, TodoStatus } from '@/types/task';
import { camelize } from '@/libs/camelize';
import { convertDate } from '@/libs/convertDate';
import '@nuxtjs/axios';

export const state = () => ({
  todos: [] as Todo[],
  selectedId: '',
  isDelete: false,
  isUpdate: false,
});

export type RootState = ReturnType<typeof state>;

export const getters: GetterTree<RootState, RootState> = {
  newTodos: (state) =>
    state.todos.filter((e) => e.status === TodoStatus.statusNew),
  wipTodos: (state) =>
    state.todos.filter((e) => e.status === TodoStatus.statusWIP),
  doneTodos: (state) =>
    state.todos.filter((e) => e.status === TodoStatus.statusDone),
  pendingTodos: (state) =>
    state.todos.filter((e) => e.status === TodoStatus.statusPending),
  selectedTodo: (state) => {
    const filtered = state.todos.filter((e) => e.id === state.selectedId);
    return filtered.length === 0 ? null : filtered[0];
  },
  isSelected: (state) => state.selectedId !== '',
  isDelete: (state) => state.isDelete,
  isUpdate: (state) => state.isUpdate,
  todos: (state) => state.todos,
};

export const mutations: MutationTree<RootState> = {
  UPDATE_TODOS: (state: RootState, payload: { todos: Todo[] }) =>
    (state.todos = payload.todos),
  UPDATE_TODO: (state: RootState, payload: { id: string; todo: Todo }) => {
    const index: number | undefined = state.todos.findIndex(
      (e: Todo) => e.id === payload.id
    );
    if (index !== undefined) {
      state.todos[index] = payload.todo;
    }
  },
  DELETE_TODO: (state: RootState, payload: { id: string }) =>
    (state.todos = state.todos.filter((e: Todo) => e.id !== payload.id)),
  ADD_TODO: (state: RootState, payload: { todo: Todo }) => {
    state.todos.push(payload.todo);
    return state.todos;
  },
  UPDATE_SELECT_DELETE: (state: RootState, payload: { status: boolean }) =>
    (state.isDelete = payload.status),
  UPDATE_SELECT_UPDATE: (state: RootState, payload: { status: boolean }) =>
    (state.isUpdate = payload.status),
  SELECT_TODO: (state: RootState, payload: { id: string }) =>
    (state.selectedId = payload.id),
  UNSELECT_TODO: (state: RootState) => (state.selectedId = ''),
};

export const actions: ActionTree<RootState, RootState> = {
  async fetchTodos() {
    const todos: Todo[] = await this.$axios
      .$get('/api/todos')
      .then((r: Object[]) => {
        return r.map((e: Object) => convertDate(camelize(e)) as Todo);
      })
      .catch((err) => {
        console.log(err.request.response);
        return [];
      });

    if (todos.length > 0) {
      this.app.$accessor.UPDATE_TODOS({ todos });
    }
  },
  async deleteTodo(_, id: string) {
    await this.$axios.$delete(`/api/todo/${id}`).catch((err) => {
      console.log(err.request.response);
      throw err;
    });
    this.app.$accessor.DELETE_TODO({ id });
  },
  async addTodo(_, payload: { content: string }) {
    await this.$axios.post('/api/create', payload).catch((err) => {
      console.log(err.request.response);
      throw err;
    });
    await this.dispatch('fetchTodos');
  },
  async updateContent(_, payload: { id: string; content: string }) {
    await this.$axios
      .post(`/api/todo/${payload.id}/content`, { content: payload.content })
      .catch((err) => {
        console.log(err.request.response);
        throw err;
      });

    await this.dispatch('fetchTodos');
  },
  async updateStatus(_, payload: { id: string; status: TodoStatus }) {
    await this.$axios
      .post(`/api/todo/${payload.id}/status`, { status: payload.status })
      .catch((err) => {
        console.log(err.request.response);
        throw err;
      });

    await this.dispatch('fetchTodos');
  },
};

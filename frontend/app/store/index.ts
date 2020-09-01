import { GetterTree, ActionTree, MutationTree } from 'vuex';
import { Todo, TodoStatus } from '@/types/task';
import { camelize } from '@/libs/camelize';
import { convertDate } from '@/libs/convertDate';
import '@nuxtjs/axios';

export const state = () => ({
  todos: [] as Todo[],
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
};

export const actions: ActionTree<RootState, RootState> = {
  async fetchTodos({ commit }) {
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
};

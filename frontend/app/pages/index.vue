<template>
  <div class="w-screen min-h-screen flex bg-blue-200">
    <div class="flex flex-row space-x-8 w-full m-3">
      <todo-list :todos="newTodo" title="New" bg-color="bg-orange-200" />
      <todo-list :todos="doneTodo" title="Done" bg-color="bg-green-200" />
      <todo-list :todos="pendingTodo" title="Pending" bg-color="bg-red-200" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import '@nuxtjs/axios';
import TodoList from '@/components/TodoList.vue';
import { camelize } from '@/libs/camelize';
import { convertDate } from '@/libs/convertDate';
import * as Task from '@/types/task';

export default Vue.extend({
  components: {
    TodoList,
  },
  async asyncData({ $axios }) {
    const res = await $axios
      .$get('/api/todos')
      .then((r: Task.Todo[]) => {
        return r.map((e: Task.Todo) => convertDate(camelize(e)));
      })
      .catch((err) => {
        console.log(err.request.response);
        return [];
      });

    return { todos: res };
  },
  computed: {
    newTodo() {
      return this.todos.filter((e) => e.status === Task.TodoStatus.statusNew);
    },
    doneTodo() {
      return this.todos.filter((e) => e.status === Task.TodoStatus.statusDone);
    },
    pendingTodo() {
      return this.todos.filter(
        (e) => e.status === Task.TodoStatus.statusPending
      );
    },
  },
});
</script>

<style>
/* Sample `apply` at-rules with Tailwind CSS
.container {
@apply min-h-screen flex justify-center items-center text-center mx-auto;
}
*/
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>

<template>
  <div class="w-screen min-h-screen flex bg-white">
    <div class="flex flex-row space-x-8 w-full m-3">
      <todo-list :todos="newTodos" title="New" bg-color="bg-orange-200" />
      <todo-list :todos="wipTodos" title="WIP" bg-color="bg-blue-200" />
      <todo-list :todos="doneTodos" title="Done" bg-color="bg-green-200" />
      <todo-list :todos="pendingTodos" title="Pending" bg-color="bg-red-200" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapGetters } from 'vuex';
import '@nuxtjs/axios';
import TodoList from '@/components/TodoList.vue';

export default Vue.extend({
  components: {
    TodoList,
  },
  async fetch() {
    await this.$store.dispatch('fetchTodos');
  },
  computed: {
    ...mapGetters(['newTodos', 'wipTodos', 'doneTodos', 'pendingTodos']),
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

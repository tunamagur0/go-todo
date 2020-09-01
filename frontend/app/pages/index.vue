<template>
  <div class="w-screen h-screen flex flex-col bg-white">
    <todo-form />
    <div class="flex flex-row space-x-8 w-full h-full p-3">
      <todo-list
        :todos="newTodos"
        title="New"
        bg-color="bg-orange-200"
        :status="statusNew"
      />
      <todo-list
        :todos="wipTodos"
        title="WIP"
        bg-color="bg-blue-200"
        :status="statusWIP"
      />
      <todo-list
        :todos="doneTodos"
        title="Done"
        bg-color="bg-green-200"
        :status="statusDone"
      />
      <todo-list
        :todos="pendingTodos"
        title="Pending"
        bg-color="bg-red-200"
        :status="statusPending"
      />
    </div>
    <delete-modal v-if="isSelected && isDelete" />
    <edit-modal v-if="isSelected && isUpdate" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapGetters } from 'vuex';
import '@nuxtjs/axios';
import { TodoStatus } from '@/types/task';
import TodoList from '@/components/TodoList.vue';
import TodoForm from '@/components/TodoForm.vue';
import DeleteModal from '@/components/DeleteModal.vue';
import EditModal from '@/components/EditModal.vue';

export default Vue.extend({
  components: {
    TodoList,
    TodoForm,
    DeleteModal,
    EditModal,
  },
  async fetch() {
    await this.$store.dispatch('fetchTodos');
  },
  computed: {
    ...mapGetters([
      'newTodos',
      'wipTodos',
      'doneTodos',
      'pendingTodos',
      'isDelete',
      'isUpdate',
      'isSelected',
    ]),
    statusNew() {
      return TodoStatus.statusNew;
    },
    statusWIP() {
      return TodoStatus.statusWIP;
    },
    statusDone() {
      return TodoStatus.statusDone;
    },
    statusPending() {
      return TodoStatus.statusPending;
    },
  },
});
</script>

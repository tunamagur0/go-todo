<template>
  <div
    class="flex flex-col flex-1 h-full"
    @dragover="onDragOver"
    @drop="onDrop"
  >
    <div class="text-center">
      <p class="text-xl uppercase">{{ title }}</p>
    </div>
    <div
      class="flex flex-col w-full h-full flex-grow space-y-8 border-gray-800 rounded-md border-2 p-2 overflow-y-scroll"
      :class="bgColor"
    >
      <todo v-for="todo in todos" :key="todo.id" :todo="todo" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import * as Task from '@/types/task';
import Todo from '@/components/Todo.vue';

export default Vue.extend({
  components: {
    Todo,
  },
  props: {
    todos: {
      type: Array as Vue.PropType<Array<Task.Todo>>,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    bgColor: {
      type: String,
      required: true,
    },
    status: {
      type: Number,
      required: true,
    },
  },
  methods: {
    onDrop(e: DragEvent) {
      e.preventDefault();
      const id = e.dataTransfer?.getData('id');
      if (id) {
        this.$store
          .dispatch('updateStatus', { id, status: this.$props.status })
          .catch((err) => {
            console.error(err);
          });
      }
    },
    onDragOver(e: DragEvent) {
      e.preventDefault();
    },
  },
});
</script>

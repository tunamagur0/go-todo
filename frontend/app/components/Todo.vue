<template>
  <div
    class="flex flex-row flex-shrink-0 max-w-sm rounded overflow-hidden shadow-lg text-center bg-gray-100"
  >
    <div class="w-8 my-auto mx-0 cursor-pointer flex-shrink-0">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          d="M4 6h16M4 12h16M4 18h16"
        />
      </svg>
    </div>
    <div class="flex flex-col items-center flex-grow">
      <p class="px-2 py-2 text-xl">{{ todo.content }}</p>
      <!-- <div class="flex flex-row text-center">
        <p class="px-2 py-2 text-sm">{{ createdString }}</p>
        <p v-if="done" class="px-2 py-2 text-sm">
          {{ finishedString }}
        </p>
      </div> -->
    </div>
    <div class="w-6 my-auto mx-2 cursor-pointer flex-shrink-0" @click="remove">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
        />
      </svg>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Todo, TodoStatus } from '@/types/task';

export default Vue.extend({
  props: {
    todo: {
      type: Object as Vue.PropType<Todo>,
      required: true,
    },
  },
  computed: {
    createdString(): string {
      return this.dateToString(this.todo.createdAt);
    },
    finishedString(): string {
      if (this.todo.finishedAt === null) return '';
      return this.dateToString(this.todo.finishedAt);
    },
    done(): boolean {
      return this.todo.status === TodoStatus.statusDone;
    },
  },
  methods: {
    dateToString(date: Date): string {
      const day = ['日', '月', '火', '水', '木', '金', '土'];
      return `${date.getFullYear()}/${date.getMonth()}/${date.getDate()}(${
        day[date.getDay()]
      }) ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;
    },
    remove() {
      this.$store.dispatch('deleteTodo', this.todo.id).catch((err) => {
        console.error(err);
      });
    },
  },
});
</script>

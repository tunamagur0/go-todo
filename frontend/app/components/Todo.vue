<template>
  <div class="max-w-sm rounded overflow-hidden shadow-lg">
    <p class="px-2 py-2 text-xl">{{ todo.content }}</p>
    <div class="flex flex-row">
      <p class="px-2 py-2 text-sm">{{ createdString }}</p>
      <p v-if="done" class="px-2 py-2 text-sm">
        {{ finishedString }}
      </p>
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
  },
});
</script>

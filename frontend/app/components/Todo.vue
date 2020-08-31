<template>
  <div>
    <h2>{{ todo.id }}</h2>
    <h2>{{ todo.content }}</h2>
    <h2>{{ todo.status }}</h2>
    <h2>{{ todo.createdAt }}</h2>
    <h2>{{ todo.finishedAt }}</h2>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Todo } from '~/types/task';

export default Vue.extend({
  props: {
    todo: {
      type: Object as Vue.PropType<Todo>,
      required: true,
    },
  },
});
</script>

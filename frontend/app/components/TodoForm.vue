<template>
  <div class="w-full h-10 p-1 px-3 text-center">
    <div class="h-full flex flex-row space-x-4 justify-center">
      <input
        v-model="content"
        type="text"
        class="border w-2/4 rounded focus:outline-none focus:shadow-outline"
      />
      <button
        class="w-1/4 outline-none text-white text-lg bg-blue-500 hover:bg-blue-700 border-gray-400 border-2 rounded-lg shadow-lg"
        @click="onClick"
      >
        CREATE
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  data() {
    return { content: '' };
  },
  methods: {
    onClick() {
      if (this.content.length === 0) return;
      if (this.content.length > 255) {
        alert('content is too long!');
        return;
      }

      this.$store
        .dispatch('addTodo', { content: this.content })
        .then(() => (this.content = ''))
        .catch((err) => {
          console.error(err);
        });
    },
  },
});
</script>

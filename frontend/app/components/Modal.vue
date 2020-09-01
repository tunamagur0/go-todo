<template>
  <div
    class="w-screen h-screen bg-black bg-opacity-25 absolute left-0 top-0 flex items-center justify-center"
  >
    <div class="w-1/2 modal bg-white flex flex-col p-3">
      <div class="w-full flex-grow"><slot /></div>
      <div class="w-full flex justify-end h-10 space-x-3">
        <button
          class="uppercase px-2 text-gray-500 text-lg bg-white hover:bg-gray-300 border-gray-400 border-2 rounded-lg shadow-lg"
          @click="$emit('on-cancel')"
        >
          cancel
        </button>
        <button
          class="uppercase px-2 text-white text-lg bg-blue-300 hover:bg-blue-500 border-gray-400 border-2 rounded-lg shadow-lg"
          @click="$emit('on-close')"
        >
          ok
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({});
</script>

<style scoped>
.modal {
  height: 50vh;
}
</style>

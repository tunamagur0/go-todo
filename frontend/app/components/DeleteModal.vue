<template>
  <div class="w-full h-full absolute">
    <modal @on-close="close" @on-cancel="cancel">
      <p class="text-xl">
        Are you sure you want to permanently delete this task?
      </p>
    </modal>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapGetters } from 'vuex';
import Modal from '@/components/Modal.vue';

export default Vue.extend({
  components: {
    Modal,
  },
  computed: mapGetters(['isSelected', 'isDelete']),
  methods: {
    close() {
      if (this.isSelected && this.isDelete) {
        this.$store
          .dispatch('deleteTodo', this.$store.state.selectedId)
          .catch((err) => {
            console.error(err);
          });
      }
      this.$store.commit('UNSELECT_TODO');
      this.$store.commit('UPDATE_SELECT_DELETE', { status: false });
    },
    cancel() {
      this.$store.commit('UNSELECT_TODO');
      this.$store.commit('UPDATE_SELECT_DELETE', { status: false });
    },
  },
});
</script>

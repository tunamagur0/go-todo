<template>
  <div class="w-full h-full absolute">
    <modal @on-close="close" @on-cancel="cancel">
      <div class="flex flex-row space-x-4 justify-center">
        <p>Edit content</p>
        <input
          v-model="content"
          type="text"
          class="border w-2/4 rounded focus:outline-none focus:shadow-outline"
        />
      </div>
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
  data() {
    return {
      content: '',
    };
  },
  computed: {
    ...mapGetters(['isSelected', 'isUpdate', 'selectedTodo']),
  },
  mounted() {
    this.content = this.selectedTodo!.content;
  },
  methods: {
    close() {
      if (this.isSelected && this.isUpdate) {
        if (this.content.length === 0) return;
        if (this.content.length > 255) {
          alert('content is too long!');
          return;
        }

        this.$store
          .dispatch('updateContent', {
            id: this.$store.state.selectedId,
            content: this.content,
          })
          .catch((err) => {
            console.error(err);
          });
      }
      this.$store.commit('UNSELECT_TODO');
      this.$store.commit('UPDATE_SELECT_UPDATE', { status: false });
    },
    cancel() {
      this.$store.commit('UNSELECT_TODO');
      this.$store.commit('UPDATE_SELECT_UPDATE', { status: false });
    },
  },
});
</script>

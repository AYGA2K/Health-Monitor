<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'App'
});
</script>
<style lang="scss">
.body--light {
  background-color: white;
}

.body--dark {
  background-color: black;
}
</style>

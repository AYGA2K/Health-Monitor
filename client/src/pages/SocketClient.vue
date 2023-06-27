<template>
  <p>State: {{ connected }}</p>
</template>

<script>
import { state } from "./../socketGo";

export default {
  name: "ConnectionState",

  computed: {
    connected() {
      return state.data;
    }
  }
}
</script>

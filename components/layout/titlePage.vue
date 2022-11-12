<template>
  <v-col class="warpper_titlepage">
    <h1 class="font-40">{{ title }}</h1>
    <h1 class="font-16">{{ details }}</h1>
  </v-col>

</template>

<script>
export default {
  props: {
    title: {
      type: [String],
      default: " ",
    },
    details: {
      type: [String],
      default: " ",
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
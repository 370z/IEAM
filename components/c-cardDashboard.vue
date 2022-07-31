<template>
  <v-card elevation="2">
    <div>{{ type }}</div>
    <div style="text-align: center">{{ number }}</div>
    <div style="text-align: right">{{ currency }}</div>
  </v-card>
</template>

<script>
export default {
  props: {
    'type': {
      type: String,
      default: "ประเภท",
    },
    'number': {
      type: Number,
      default: "000000",
    },
    'currency': {
      type: String,
      default: "สกุลเงิน",
    },
  },
  // data() {},
};
</script>

<style lang="scss" scoped>
.v-card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 150px;
  padding: 15px 45px;
}
</style>
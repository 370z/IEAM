<template>
  <v-card elevation="5">
    <v-icon :style="`color:${color}`" size="50px">{{ icon }}</v-icon>
    <div class="data mt-2">
      <span class="type">{{ type }}</span>
      <span class="number">{{ number.toLocaleString("en-US") }}</span>
      <span class="currency">{{ currency }}</span>
    </div>
  </v-card>
</template>

<script>
export default {
  props: {
    type: {
      type: String,
      default: "ประเภท",
    },
    number: {
      type: Number,
      default: "000000",
    },
    currency: {
      type: String,
      default: "สกุลเงิน",
    },
    color: {
      type: String,
      default: "",
    },
    icon: {
      type: String,
      default: "",
    },
  },
  // data() {},
};
</script>

<style lang="scss" scoped>
.v-card {
  height: 180px;
  padding: 10px 30px;
  border-radius: 15px;
  display: flex;
  flex-direction: column;
  justify-content: center;

  .v-icon {
    justify-content: start;
  }

  .data {
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    .type {
      font-size: 20px;
      // font-weight: bold;
    }

    .number {
      font-size: 25px;
      font-weight: bold;
      letter-spacing: 1px;
      margin: -5px 0;
    }
  }
}
</style>
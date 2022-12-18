<template>
  <v-card elevation="5">
    <v-icon :style="`color:${color}`" size="50px">{{ icon }}</v-icon>
    <div class="data mt-2">
      <span class="type">{{ dashbord_data.nametype }}</span>
      <span class="number">{{
        dashbord_data.sum ? dashbord_data.sum.toLocaleString("en-US") : "0"
      }}</span>
      <span class="currency">บาท</span>
    </div>
  </v-card>
</template>

<script>
export default {
  props: {
    dashbord_data: {
      type: Object,
      default: () => ({}),
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
  min-width: 250px;

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
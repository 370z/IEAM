<template>
  <v-card class="content-card" elevation="5">
    <div class="warpoer_status">
      <v-icon :style="`color:${color}`" size="50px"
        >mdi-text-box-plus-outline</v-icon
      >
      <div class="data_status flex-column">
        <h3>อนุมัติเรียบร้อย</h3>
        <span>sdasdasdas</span>
      </div>
    </div>
    <div class="warpoer_status">
      <v-icon :style="`color:${color}`" size="50px"
        >mdi-text-box-plus-outline</v-icon
      >
      <div class="data_status flex-column">
        <h3>อนุมัติเรียบร้อย</h3>
        <span>sdasdasdas</span>
      </div>
    </div>
    <div class="warpoer_status">
      <v-icon :style="`color:${color}`" size="50px"
        >mdi-text-box-plus-outline</v-icon
      >
      <div class="data_status flex-column">
        <h3>อนุมัติเรียบร้อย</h3>
        <span>sdasdasdas</span>
      </div>
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
.warpoer_status {
  display: flex;

  .v-icon {
    margin-right: 10px;
  }
  .data_status {
    display: flex;
    flex-direction: row;
  }
}
</style>
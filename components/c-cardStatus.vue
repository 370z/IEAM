<template>
  <v-card class="content-card box_warpoer_status" elevation="5">
    <div
      class="warpoer_status"
      style="
        background: linear-gradient(169deg,#03a9f4 0%,rgba(0, 212, 255, 1) 100%);
      "
    >
      <v-icon size="50px">mdi-format-list-checkbox</v-icon>
      <div class="data_status flex-column">
        <h5>เอกสารทั้งหมด</h5>
        <span style="font-size: 40px">{{ data.total || "0" }}</span>
      </div>
    </div>
    <div
      class="warpoer_status"
      style="background: linear-gradient(169deg, #66bb6a 0%, #7cf982 100%)"
    >
      <v-icon size="50px">mdi-format-list-checks</v-icon>
      <div class="data_status flex-column">
        <h5>อนุมัติเรียบร้อย</h5>
        <span style="font-size: 40px">{{ data.approve || "0" }}</span>
      </div>
    </div>
    <div
      class="warpoer_status"
      style="
        background: linear-gradient(169deg,#fe2828 0%,#ff6060 100%);
      "
    >
      <v-icon size="50px">mdi-playlist-remove</v-icon>
      <div class="data_status flex-column">
        <h5>ไม่ถูกอนุมัติ</h5>
        <span style="font-size: 40px">{{ data.cancel || "0" }}</span>
      </div>
    </div>
  </v-card>
</template>

<script>
export default {
  props: {
    data: {
      type: Object,
      default: {},
    },
  },
};
</script>

<style lang="scss" scoped>
.box_warpoer_status {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  padding: 30px 30px;
}
.warpoer_status {
  display: flex;
  padding: 20px 20px;
  color: #ffffff;
  border-radius: 20px !important;

  .v-icon {
    margin-right: 10px;
    color: #ffffff;
  }
  .data_status {
    display: flex;
    flex-direction: row;
  }
}
</style>
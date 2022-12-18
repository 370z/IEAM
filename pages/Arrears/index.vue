<template>
  <div class="content-size">
    <layout-title-page
      title="จัดการรายการที่ค้าง"
      details="แสดงรายการที่ค้างต่างๆ เช่น ค้างคืนเงิน ค้างแนบหลักฐานการยืมเงิน"
    ></layout-title-page>
    <v-col class="content">
      <v-card class="content-card" elevation="6">
        <el-tabs v-model="activeName" type="card">
          <el-tab-pane label="ค้างคืนเงิน" name="first"
            ><table-history-notreturn :datatable="notreturn"
          /></el-tab-pane>
          <el-tab-pane label="ค้างแนบใบเสร็จ" name="second"
            ><table-history-notbill :datatable="notreturnbill"
          /></el-tab-pane>
        </el-tabs>
      </v-card>
    </v-col>
  </div>
</template>
  
  <script>
export default {
  data() {
    return {
      activeName: "first",
      notreturnbill: [],
      notreturn: [],
    };
  },
  mounted() {
    var id = this.$auth.state.user.accountId
    this.getnotbill(id);
    this.getnotreturn(id);
  },
  methods: {
    getnotbill(id) {
      this.$axios
        .get(`${process.env.BASE_URL}/notreturnbill/Getnotreturnbill`, {
          params: { id: id },
        })
        .then((response) => {
          this.notreturnbill = response.data?.data;
        });
    },
    getnotreturn(id) {
      this.$axios
        .get(`${process.env.BASE_URL}/notreturn/getnotreturn`, {
          params: { id: id },
        })
        .then((response) => {
          this.notreturn = response.data?.data;
        });
    },
  },
};
</script>
  
<style lang="scss" scoped>
</style>
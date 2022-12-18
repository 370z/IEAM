<template>
  <div class="content-size">
    <layout-title-page
      title="ยื่นเอกสาร เบิก/ยืม/คืน"
      details="การขออนุญาติในการ เบิก/ยืม/คืน"
    ></layout-title-page>
    <!-- <v-col class="content">
      <v-row> </v-row>
    </div> -->
      <c-form />
  </div>
</template>

<script>
export default {
  middleware: 'auth',
    mounted() {
    this.lv = [1];
    if (!this.lv.includes(this.$auth.state.user.level)) {
      this.$router.push("/dashboard");
    }
  },
  data() {
    return {
      items: [
        {
          title: "รายรับ",
          total: 1000000,
          currency: "บาท",
        },
        {
          title: "รายจ่าย",
          total: 30000,
          currency: "บาท",
        },
        {
          title: "การเบิก",
          total: 1000,
          currency: "บาท",
        },
        {
          title: "การยืม",
          total: 40000,
          currency: "บาท",
        },
        {
          title: "การคืน",
          total: 43040,
          currency: "บาท",
        },
      ],
    };
  },
};
</script>

<style>
.name_class{
  width: 100%;
  display: flex;
  /* justify-content: end; */
}

</style>
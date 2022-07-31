<template>
  <div class="content-size">
    <layout-title-page
      title="ยื่นเอกสาร เบิก/ยืม/คืน"
      details="การขออนุญาติในการ เบิก/ยืม/คืน"
    ></layout-title-page>
    <!-- <div class="content">
      <v-row> </v-row>
    </div> -->
    <div class="name_class">
      <button>sdfgdd</button>
    </div>
  </div>
</template>

<script>
export default {
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
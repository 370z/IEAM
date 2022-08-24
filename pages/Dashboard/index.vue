<template>
  <div class="content-size">
    <layout-title-page
      title="Dashboard"
      details="ดูภาพรวมของตัวระบบ"
    ></layout-title-page>
    <div class="content">
      <v-row>
        <v-col cols="2.5" v-for="item in items" :key="item.title">
          <c-card-dashboard
            :type="item.title"
            :number="item.total"
            :currency="item.currency"
          ></c-card-dashboard>
        </v-col>
      </v-row>
      <!-- <v-row>
        <v-col> 1</v-col>
        <v-col>2</v-col>
      </v-row> -->
      <v-row class="flex-column">
        <v-col>
          <v-card>
            <v-data-table
              :headers="headers"
              :items="desserts"
              :loading="false"
              loading-text="Loading... Please wait"
              no-data-text="ไม่พบข้อมูล"
            />
          </v-card>
        </v-col>
      </v-row>
      <!-- {{process.env.BASE_URL}} -->
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
      desserts:[],
      headers: [
        {
          text: "ลำดับ",
          align: "center",
          value: "id",
          sortable: false,
          width: "5%",
        },
        {
          text: "เลขใบแบบฟอร์ม",
          value: "calories",
          sortable: false,
          width: "10%",
        },
        { text: "รายการ", value: "name", sortable: false, width: "%" },
        { text: "จำนวนเงิน", value: "protein", sortable: false, width: "15%" },
        {
          text: "ประเภท",
          align: "center",
          value: "type",
          sortable: false,
          width: "10%",
        },
        {
          text: "วัน",
          align: "center",
          value: "iron",
          sortable: false,
          width: "10%",
        },
        {
          text: "เพิ่มเติม",
          align: "center",
          value: "actions",
          sortable: false,
          width: "10%",
        },
      ],
    };
  },

};
</script>

<style>
</style>
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
            :color="item.color"
            :icon="item.icon"
          ></c-card-dashboard>
        </v-col>
      </v-row>
      <!-- <v-row>
        <v-col cols="4">
          <c-card-status></c-card-status>
        </v-col>
        <v-col cols="8">
          <v-card class="content-card" elevation="6">
            <h1>dsfsd</h1>
            <h1>dsfsd</h1>
            <h1>dsfsd</h1>
          </v-card></v-col
        >
      </v-row> -->
      <v-row class="flex-column">
        <v-col>
          <v-card class="content-card" elevation="6">
            <h2 class="mb-5">ประวัติทำรายการย้อนหลัง</h2>
            <v-data-table
              :headers="headers"
              :items="desserts"
              :loading="false"
              loading-text="Loading... Please wait"
              no-data-text="ไม่พบข้อมูล"
            >
              <template v-slot:[`item.id`]="{ index }">
                {{ index + 1 }}
              </template>
              <template v-slot:[`item.updatedapprove`]="{ item }">
                {{ relativeTime(item.updatedapprove) }}
              </template>
               <template v-slot:[`item.isapprove`]="{ item }">
            <span
              :style="`color: ${
                item.isapprove == 0
                  ? '#FFC83B'
                  : item.isapprove == 1
                  ? '#48A451'
                  : '#CD2126'
              }`"
            >
              {{
                item.isapprove == 0
                  ? "รอการอนุมัติ"
                  : item.isapprove == 1
                  ? "อนุมัติเรียบร้อย"
                  : "ไม่อนุมัติ"
              }}
            </span>
          </template>
            </v-data-table>
          </v-card>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script>
import moment from "moment";
import cCardStatus from "../../components/c-cardStatus.vue";
export default {
  components: { cCardStatus },
  data() {
    return {
      items: [
        {
          title: "รายรับ",
          total: 1000000,
          currency: "บาท",
          color: "#4caf4f",
          icon: "mdi-bank-transfer-in",
        },
        {
          title: "รายจ่าย",
          total: 30000,
          currency: "บาท",
          color: "#ff3b2f",
          icon: "mdi-bank-transfer-out",
        },
        {
          title: "การเบิก",
          total: 1000,
          currency: "บาท",
          color: "#f9ae00",
          icon: "mdi-clipboard-list",
        },
        {
          title: "การยืม",
          total: 40000,
          currency: "บาท",
          color: "#ffe000",
          icon: "mdi-text-box-minus-outline",
        },
        {
          title: "การคืน",
          total: 43040,
          currency: "บาท",
          color: "#009fe3",
          icon: "mdi-text-box-plus-outline",
        },
      ],
      desserts: [],
      headers: [
        {
          text: "ลำดับ",
          align: "center",
          value: "id",
          sortable: false,
          width: "5%",
        },
        {
          text: "วัน",
          align: "center",
          value: "updatedapprove",
          sortable: false,
          width: "20%",
        },
        {
          text: "เลขใบแบบฟอร์ม",
          value: "id_form",
          sortable: false,
          width: "20%",
        },
        { text: "รายการ", value: "title", sortable: false, width: "%" },
        {
          text: "ประเภท",
          align: "center",
          value: "type.nametype",
          sortable: false,
          width: "10%",
        },
        {
          text: "สถานะ",
          align: "center",
          value: "isapprove",
          sortable: false,
          width: "15%",
        },
      ],
    };
  },
  mounted() {
    this.gethistory();
  },
  methods: {
    async gethistory() {
      await this.$axios
        .get(`${process.env.BASE_URL}/history/approve`)
        .then((response) => {
          this.desserts = response.data?.data;
          console.log(this.desserts);
        });
    },
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L LT");
      return result;
    },
  },
};
</script>
<style>
</style>
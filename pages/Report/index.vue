<template>
  <div class="content-size">
    <layout-title-page
      title="รายงาน"
      details="การจัดการรายงาน รายรับ รายจ่าย เบิก ยืม คืน"
    ></layout-title-page>
    <v-col class="content">
      <v-card class="content-card" elevation="6">
        <v-col
          ><h1 class="font-20">รายการทั้งหมด ({{ desserts.length }})</h1></v-col
        >
        <v-row style="align-items: baseline" class="input-report">
          <v-col>
            <el-date-picker v-model="search.year" type="year" placeholder="ปี">
            </el-date-picker>
          </v-col>
          <v-col>
            <!-- {{pickdate(search.year).format()}}
            <el-date-picker
              v-model="search"
              type="daterange"
              range-separator="-"
              start-placeholder="เริ่มต้น"
              end-placeholder="สิ้นสุด"
            >
            </el-date-picker> -->
            <el-date-picker
              popper-class="month-picker"
              v-model="search.month"
              type="month"
              format="MMMM"
              placeholder="เดือน"
            >
            </el-date-picker>
          </v-col>
          <v-col>
            <el-select
              v-model="search.type"
              class="w-full select-report"
              circle
              placeholder="ประเภทเอกสาร"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </v-col>
          <v-col>
            <v-btn
              color="#E39257"
              dark
              block
              depressed
              rounded
              @click="getform()"
            >
              <v-icon left>mdi-magnify </v-icon>ค้นหา</v-btn
            >
          </v-col>
          <v-spacer></v-spacer>
          <v-col>
            <v-btn color="blue" dark block depressed rounded @click="pdf()">
              <v-icon left>mdi-printer </v-icon>พิมพ์</v-btn
            >
          </v-col>
        </v-row>
        <v-col>
          <v-data-table
            :headers="headers"
            :items="desserts"
            :loading="loading"
            loading-text="Loading... Please wait"
            no-data-text="ไม่พบข้อมูล"
          >
            <template v-slot:[`item.id`]="{ index }">
              {{ index + 1 }}
            </template>
            <template v-slot:[`item.name`]="{ item }">
              {{ item.user.firstname + " " + item.user.lastname }}
            </template>
            <template v-slot:[`item.total`]="{ item }">
              {{ item.total.toLocaleString("en-US") }}
            </template>
          </v-data-table>
        </v-col>
      </v-card>
    </v-col>
  </div>
</template>

<script>
import moment from "moment";
import printreport from "./print/printreport";
moment.locale("th");

export default {
  data() {
    return {
      loading: false,
      search: {
        year: moment(),
        month: moment(),
        type: null,
      },
      dialog: false,
      options: [
        {
          value: null,
          label: "ทั้งหมด",
        },
        {
          value: 1,
          label: "รายรับ",
        },
        {
          value: 2,
          label: "รายจ่าย",
        },
        {
          value: 3,
          label: "การเบิก",
        },
        {
          value: 4,
          label: "การยืม",
        },
        {
          value: 5,
          label: "การคืน",
        },
      ],
      items: [
        { value: null, text: "ทั้งหมด" },
        { value: "0", text: "รอการอนุมัติ" },
        { value: "1", text: "อนุมัติเรียบร้อย" },
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
          text: "เลขใบแบบฟอร์ม",
          value: "number_form",
          sortable: false,
          width: "10%",
          align: "center",
        },
        {
          text: "ชื่อผู้ยืนเอกสาร",
          value: "name",
          sortable: false,
          width: "20%",
        },
        { text: "รายการ", value: "title", sortable: false },
        {
          text: "ประเภท",
          align: "center",
          value: "type.nametype",
          sortable: false,
          width: "15%",
        },
        {
          text: "จำนวนเงิน",
          value: "total",
          sortable: false,
          width: "10%",
          align: "center",
        },
      ],
    };
  },
  mounted() {
    this.getform();
    this.lv = [3];
    if (!this.lv.includes(this.$auth.state.user.level)) {
      this.$router.push("/dashboard");
    }
  },
  methods: {
    async getform() {
      this.loading = true;
      const data = {
        type: this.search.type,
        month: this.search.month
          ? this.pickdate(this.search.month).format()
          : null,
        year: this.search.year
          ? this.pickdate(this.search.year).format()
          : null,
      };
      await this.$axios
        .post(`${process.env.BASE_URL}/report/all`, data)
        .then((response) => {
          this.desserts = response.data?.data;
          this.loading = false;
        });
    },
    pickdate(time) {
      return moment(time);
    },
    async pdf() {
      await pdfMake
        .createPdf(await printreport.context(this.desserts, this.search))
        .print();
    },
  },
};
</script>

<style lang="scss">
.input-report {
  .el-input__inner {
    border-radius: 100px !important;
  }
}

.month-picker {
  .el-date-picker__header {
    display: none;
  }
}
</style>
<template>
    <v-card>
      <v-card-title>
        <v-spacer></v-spacer>
        <v-menu bottom left>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
              v-bind="attrs"
              v-on="on"
              @click="$emit('closedialog', true)"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </template>
        </v-menu>
      </v-card-title>
      <v-card-text>
        <v-col class="d-flex"
          ><h1>
            {{ detail.type.nametype ? detail.type.nametype : "" }}
          </h1></v-col
        >
        <v-row class="d-flex">
          <v-col>
            <p>หัวเรื่อง : {{ detail.title }}</p>
            <p>เลขใบฟอร์ม : {{ detail.number_form }}</p>
          </v-col>
          <v-col class="ml-10">
            <p>วันที่ : {{ relativeTime(detail.CreatedAt) }}</p>
            <p>จำนวนเงิน : {{ detail.total.toLocaleString("en-US") }} บาท</p>
          </v-col>
        </v-row>
        <v-col>
          <v-data-table
            :headers="headers"
            :items="detail.list"
            :loading="false"
            hide-default-footer
            loading-text="Loading... Please wait"
            no-data-text="ไม่พบข้อมูล"
          >
            <template v-slot:[`item.total`]="{ item }">
              {{ item.total.toLocaleString("en-US") }}
            </template>
            <template v-slot:[`item.unit_price`]="{ item }">
              {{ parseInt(item.unit_price).toLocaleString("en-US") }}
            </template>
            <template v-slot:[`item.number`]="{ item }">
              {{ parseInt(item.number).toLocaleString("en-US") }}
            </template>
          </v-data-table>
        </v-col>
        <v-col>
          <p>เอกสารแนบ :</p>
          <el-tabs type="card">
            <el-tab-pane label="หลักฐาน">
              <v-img
                class="ma-auto"
                :lazy-src="detail.img_url"
                max-width="50%"
                :src="detail.img_url"
              ></v-img
            ></el-tab-pane>
            <el-tab-pane label="ใบเสร็จ" v-if="detail.typeId == 4"
              ><v-img
                class="ma-auto"
                :lazy-src="detail.img_url_bill"
                max-width="50%"
                :src="detail.img_url_bill"
              ></v-img
            ></el-tab-pane>
            <el-tab-pane label="สลิปโอนเงิน" v-if="detail.typeId == 3 ||detail.typeId == 4 "
              ><v-img
                class="ma-auto"
                :lazy-src="detail.img_urltransfer"
                max-width="50%"
                :src="detail.img_urltransfer"
              ></v-img
            ></el-tab-pane>
          </el-tabs>
        </v-col>
      </v-card-text>
    </v-card>
</template>

<script>
import moment from "moment";

export default {
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    detail: {
      type: Object,
      default: {},
    },
  },
  data() {
    return {
      headers: [
        {
          text: "รายการ",
          value: "title",
          sortable: false,
          width: "50%",
        },
        {
          text: "จำนวน",
          value: "number",
          sortable: false,
          align: "center",
        },
        { text: "หน่วย", value: "unit", sortable: false, align: "center" },
        {
          text: "ราคาต่อหน่วย",
          value: "unit_price",
          sortable: false,
          align: "center",
        },
        {
          text: "รวม",
          align: "center",
          value: "total",
          sortable: false,
        },
      ],
    };
  },

  methods: {
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L");
      return result;
    },
  },
};
</script>

<style>
</style>
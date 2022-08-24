<template>
  <div class="content-size">
    <layout-title-page
      title="การจัดการบัญชี โอนเงิน"
      details="สถานะการโอนเงินของแบบฟอร์มแต่ละประเภท"
    ></layout-title-page>
    <v-card class="content-card" elevation="6">
      <v-data-table
        :headers="headers"
        :items="datatable"
        :search="search"
        :loading="false"
        loading-text="Loading... Please wait"
        no-data-text="ไม่พบข้อมูล"
      >
        <template v-slot:top>
          <v-row>
            <v-col cols="3"><h1 class="font-20">รายการทั้งหมด (100)</h1></v-col>
            <v-spacer />
            <v-col cols="2">
              <v-select
                v-model="search"
                :menu-props="{ offsetY: true }"
                :items="items"
                item-value="value"
                item-text="text"
                dense
                outlined
                rounded
                hide-details
              ></v-select>
            </v-col>
          </v-row>
        </template>
        <template v-slot:[`item.id`]="{ index }">
          {{ index + 1 }}
        </template>
        <template v-slot:[`item.istransfermoney`]="{ item }">
          <span
            :style="`color: ${
              item.istransfermoney == 0 ? '#FFC83B' : '#30C03E'
            }`"
          >
            {{
              item.istransfermoney == 0 ? "รอการดำเนินการ" : "ดำเนินการสำเร็จ"
            }}
          </span>
        </template>
        <template v-slot:[`item.actions`]="{ item }">
          <v-btn
            color="primary"
            v-if="item.istransfermoney == 0"
            outlined
            rounded
            small
          >
            <v-icon left @click="del(item)"> mdi-file-image </v-icon>แนบรูปภาพ
            <!-- <input
              ref="uploader"
              class="d-none"
              type="file"
              accept="image/*"
              @change="onFileChanged"
          /> -->
          </v-btn>
          <v-btn
            color="primary"
            v-if="item.istransfermoney == 1"
            outlined
            rounded
            small
            @click="view(item)"
          >
            <v-icon left>mdi-image-search </v-icon>ดูหลักฐาน</v-btn
          >
        </template>
      </v-data-table>
    </v-card>
    <v-dialog v-model="dialog" max-width="70%">
      <layout-dialog-transfermomey
        :detail="transfermdetail"
        @dialogview-click="dialog = !dialog"
      />
    </v-dialog>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      search: null,
      dialog: false,
      items: [
        { value: null, text: "ทั้งหมด" },
        { value: "0", text: "รอการดำเนินการ" },
        { value: "1", text: "ดำเนินการสำเร็จ" },
      ],
      datatable: [],
      transfermdetail: {},
      headers: [
        {
          text: "ลำดับ",
          align: "center",
          value: "id",
          sortable: false,
          width: "10%",
        },
        {
          text: "เลขใบแบบฟอร์ม",
          value: "id_form",
          sortable: false,
          width: "10%",
        },
        { text: "ชื่อ-นามสกุล", value: "name", sortable: false, width: "15%" },
        {
          text: "ประเภท",
          align: "center",
          value: "type",
          sortable: false,
          width: "15%",
        },
        { text: "จำนวนเงิน", value: "total", sortable: false, width: "15%" },
        {
          text: "สถานะการโอน",
          align: "center",
          value: "istransfermoney",
          sortable: false,
          width: "10%",
        },
        {
          text: "หลักฐาน",
          align: "center",
          value: "actions",
          sortable: false,
          width: "15%",
        },
      ],
    };
  },
  mounted() {
    this.gettransfermdetail();
  },
  methods: {
    async gettransfermdetail() {
      await axios
        .get(`${process.env.BASE_URL}/transfermoney/all`)
        .then((response) => {
          console.log(response.data.data);
          this.datatable = response.data?.data;
        });
    },
    view(item) {
      this.transfermdetail = item;
      this.dialog = !this.dialog;
      console.log(this.dialog);
    },
  },
};
</script>

<style>
</style>
<template>
  <div>
    <v-data-table
      :headers="headers"
      :items="datatable"
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
      <template v-slot:[`item.number_form`]="{ item }">
        <span
          style="cursor: pointer; color: #2096f3; text-decoration: underline"
          @click="view(item)"
        >
          {{ item.number_form }}
        </span>
      </template>
      <template v-slot:[`item.transfer`]="{ item }">
        <span
          :style="`color: ${
            item.typeId == 5 || item.isapprove != 1
              ? ''
              : item.istransfermoney
              ? '#48A451'
              : '#CD2126'
          }`"
        >
          {{
            item.typeId == 5 || item.isapprove != 1
              ? "-"
              : item.istransfermoney
              ? "เรียบร้อย"
              : "รอดำเนินการ"
          }}
        </span>
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
      <template
        v-slot:[`item.print`]="{ item }"
        v-if="(this.$auth.state.user.level == 1)"
      >
        <v-btn
          :disabled="item.isapprove != 1"
          class="ml-2"
          color="blue"
          style="color: #ffffff"
          small
          rounded
          elevation="false"
          @click="pdf(item)"
        >
          <v-icon>mdi-printer </v-icon>พิมพ์</v-btn
        >
      </template>
    </v-data-table>
    <v-dialog v-model="dialog" max-width="70%">
      <layout-dialog-transfermomey
        :detail="detail"
        @closedialog="dialog = !dialog"
        :show="dialog"
      />
    </v-dialog>
  </div>
</template>

<script>
import moment from "moment";
import printform from "./print/printform";
export default {
  props: {
    datatable: {
      type: Array,
      default: [],
    },
  },
  data() {
    return {
      detail: null,
      dialog: false,
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
          width: "15%",
        },
        {
          text: "เลขใบแบบฟอร์ม",
          value: "number_form",
          sortable: false,
          width: "10%",
          align: "center",
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
          width: "10%",
        },
        {
          text: "การโอน",
          align: "center",
          value: "transfer",
          sortable: false,
          width: "10%",
        },
        {
          text: "",
          align: "center",
          value: "print",
          sortable: false,
          width: "10%",
        },
      ],
    };
  },
  methods: {
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L LT");
      return result;
    },
    view(item) {
      this.detail = item;
      this.dialog = !this.dialog;
    },
    async pdf(item) {
      this.detail = item;
      await pdfMake
        .createPdf(await printform.context(this.detail))
        .print();
    },
  },
};
</script>

<style>
</style>
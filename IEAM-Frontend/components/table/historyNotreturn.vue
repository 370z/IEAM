<template>
  <div class="history_notbill">
    <!-- <pre>{{datatable[0].returnborrow}}</pre> -->
    <el-table :data="datatable" style="width: 100%" height="350">
      <el-table-column
        prop="number_form"
        label="เลขฟอร์ม"
        width="120"
        align="center"
      >
      </el-table-column>
      <el-table-column prop="title" label="รายการ"> </el-table-column>
      <el-table-column
        prop="total"
        label="จำนวนเงิน"
        width="150"
        align="center"
      >
      </el-table-column>
      <el-table-column label="วันที่ยืน" width="120" align="center">
        <template slot-scope="scope">
          {{ relativeTime(scope.row.CreatedAt) }}
        </template>
      </el-table-column>
      <el-table-column label="สถานะ" width="120" align="center">
        <template slot-scope="scope">
          <span
            :style="`color: ${
              scope.row.returnborrow?.isapprove == null
                ? ''
                : scope.row.returnborrow?.isapprove == 0
                ? '#FFC83B'
                : '#CD2126'
            }`"
          >
            {{
              scope.row.returnborrow?.isapprove == null
                ? "-"
                : scope.row.returnborrow?.isapprove == 0
                ? "รอดำเนินการ"
                : "ไม่อนุมัติ"
            }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="คืนเงิน" width="200" align="center">
        <template slot-scope="scope">
          <el-button
            size="small"
            type="primary"
            round
            plain
            @click="edit(scope.row.id_form)"
            :disabled="scope.row.returnborrow?.isapprove == 0"
            >
            {{
                scope.row.returnborrow?.isapprove == 2
                ? "ส่งอีกครั้ง"
                : "ยืนเอกสาร"
            }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
    
    <script>
import moment from "moment";
export default {
  props: {
    datatable: {
      type: Array,
      default: [],
    },
  },
  data() {
    return {
      dialog: false,
    };
  },
  methods: {
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L");
      return result;
    },
    edit(item) {
      this.$router.push(`arrears/${item}`);
    },
  },
};
</script>
    
    <style lang="scss" >
.history_notbill {
  thead tr {
    th {
      color: rgba(0, 0, 0, 0.578) !important;
      text-align: center !important;
    }
  }
}
</style>
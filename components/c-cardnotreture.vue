<template>
  <div>
    <v-card class="content-card notreture" elevation="6">
      <v-row
        ><h2 class="mb-5">
          รายการค้างคืนเงิน ( {{ datatable.length }} )
        </h2></v-row
      >

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
          width="120"
          align="center"
        >
        </el-table-column>
        <el-table-column
          prop="user.firstname"
          label="ชื่อ-นามสกุล"
          width="200"
          align="center"
        >
        </el-table-column>
        <el-table-column
          prop="user.phonenumber"
          label="เบอร์โทร"
          width="120"
          align="center"
        >
        </el-table-column>
      </el-table>
    </v-card>
  </div>
</template>

<script>
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
    };
  },
};
</script>

<style lang="scss" >
.notreture {
  thead tr {
    th {
      color: rgba(0, 0, 0, 0.578) !important;
      text-align: center !important;
    }
  }
}
</style>
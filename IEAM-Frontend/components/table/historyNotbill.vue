<template>
  <div class="history_notbill">
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
      <el-table-column
        label="วันที่ยืน"
        width="120"
        align="center"
      >
      <template slot-scope="scope">
        {{relativeTime(scope.row.CreatedAt)}}
      </template>
      </el-table-column>
      <el-table-column label="แนบใบเสร็จ" width="200" align="center">
        <template slot-scope="scope">
          <el-upload
            action="#"
            accept="image/*"
            list-type="picture"
            :on-change="addfile"
            :auto-upload="false"
            :show-file-list="false"
          >
            <el-button
              size="small"
              type="primary"
              round
              plain
              @click="view(scope.row)"
              >แนบรูปภาพ
            </el-button>
          </el-upload>
        </template>
      </el-table-column>
    </el-table>
    <v-dialog v-model="dialog" max-width="30%">
      <v-card class="dialogtransfer">
        <img
          class="el-upload-list__item-thumbnail"
          :src="selectedFile.url"
          :lazy-src="selectedFile.url"
          alt=""
        />
        <v-row class="space-center">
          <v-col cols="4">
            <v-btn
              width="100%"
              class="mx-2"
              color="primary"
              rounded
              depressed
              @click="change_status()"
              >บันทึกรายการ</v-btn
            >
          </v-col>
          <v-col cols="4">
            <v-btn
              width="100%"
              class="mx-2"
              color="error"
              rounded
              depressed
              outlined
              @click="dialog = !dialog"
              >ยกเลิก</v-btn
            >
          </v-col>
        </v-row>
      </v-card>
    </v-dialog>
  </div>
</template>
  
  <script>
import { storage } from "~/plugins/firebase";
import env from "~/env";
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
      selectedFile: {
        url: null,
      },
      dialog: false,
      billdetail: {},
    };
  },
  methods: {
    addfile(files) {
      this.selectedFile = files;
      this.dialog = true;
    },
    uploadFirebase(file) {
      if (file.status == "ready") {
        return new Promise((resolve, reject) => {
          let metadata = {
            contentType: file.raw.type,
          };

          let task = storage
            .ref()
            .child(`bill/bill-${file.raw.uid}`)
            .put(file.raw, metadata);

          task
            .then((snapshot) => {
              snapshot.ref.getDownloadURL().then((url) => {
                resolve(url);
              });
            })
            .catch((error) => {
              reject(error);
              this.$message.error("อัพโหลดรูปไม่สำเร็จ");
            });
        });
      }
    },
    async change_status() {
      const url = await this.uploadFirebase(this.selectedFile);
      const data = {
        id_form: this.billdetail.id_form,
        img_urlbill: url,
      };
      await this.$axios
        .patch(`${env.BASE_URL}/approve/addbill`, data)
        .then((response) => {
          alert("เพิ่มหลักฐานการชำระเงินเรียบร้อย");
          window.location.reload(true);
        })
        .catch((err) => {
          alert(err);
        });
    },
    view(item) {
      this.billdetail = item;
    },
    relativeTime(time) {
      // moment.locale("th");
      var result = moment().format("L");
      return result;
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
.dialogtransfer {
  display: flex;
  flex-direction: column;
  padding: 20px;

  img {
    margin: 25px;
  }
}
</style>
<template>
  <div class="content-size">
    <layout-title-page
      title="จัดการ โอนเงิน"
      details="สถานะการโอนเงินของแบบฟอร์มแต่ละประเภท"
    ></layout-title-page>
    <v-col class="content">
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
              <v-col cols="3"
                ><h1 class="font-20">รายการทั้งหมด ( {{datatable.length}} )</h1></v-col
              >
              <v-spacer />
              <v-col cols="3">
                <v-text-field
                  v-model="search"
                  append-icon="mdi-magnify"
                  placeholder="Search"
                  dense
                  outlined
                  rounded
                  hide-details
                ></v-text-field>
              </v-col>
            </v-row>
          </template>
          <template v-slot:[`item.id`]="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:[`item.user`]="{ item }">
            {{ item.user.firstname + " " + item.user.lastname }}
          </template>
          <template v-slot:[`item.total`]="{ item }">
            {{ item.total.toLocaleString("en-US")}}
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
            <el-upload
              v-if="item.istransfermoney == 0"
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
                @click="view(item), (viewdetail = false)"
                >แนบรูปภาพ
              </el-button>
            </el-upload>
            <v-btn
              color="primary"
              v-if="item.istransfermoney == 1"
              outlined
              rounded
              small
              @click="view(item), (dialog1 = !dialog1), (viewdetail = true)"
            >
              <v-icon left>mdi-image-search </v-icon>ดูหลักฐาน</v-btn
            >
          </template>
        </v-data-table>
      </v-card>
      <v-dialog v-model="dialog" max-width="70%">
        <layout-dialog-transfermomey
          :detail="transfermdetail"
          @closedialog="dialog = !dialog"
          :show="dialog"
        />
      </v-dialog>
      <v-dialog v-model="dialog1" max-width="30%">
        <v-card class="dialogtransfer">
          <img
            class="el-upload-list__item-thumbnail"
            :src="selectedFile.url || transfermdetail.img_urltransfer"
            :lazy-src="selectedFile.url"
            alt=""
          />
          <v-row class="space-center">
            <v-col cols="4" v-if="!viewdetail">
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
                @click="dialog1 = !dialog1"
                >ยกเลิก</v-btn
              >
            </v-col>
          </v-row>
        </v-card>
      </v-dialog>
    </v-col>
  </div>
</template>

<script>
import { storage } from "~/plugins/firebase";
import env from "~/env";
export default {
  data() {
    return {
      isSelecting: false,
      selectedFile: {
        url: null,
      },
      viewdetail: false,
      search: null,
      dialog: false,
      dialog1: false,
      loading: false,
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
          value: "number_form",
          sortable: false,
          align: "center",
          width: "10%",
        },
        {
          text: "ชื่อ-นามสกุล",
          value: "name",
          sortable: false,
          width: "15%",
          value: "user",
        },
        {
          text: "ประเภท",
          align: "center",
          sortable: false,
          width: "15%",
          value: "type.nametype",
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
    this.lv = [3];
    if (!this.lv.includes(this.$auth.state.user.level)) {
      this.$router.push("/dashboard");
    }
  },
  methods: {
    addfile(files) {
      this.selectedFile = files;
      this.dialog1 = true;
    },

    uploadFirebase(file) {
      if (file.status == "ready") {
        return new Promise((resolve, reject) => {
          let metadata = {
            contentType: file.raw.type,
          };

          let task = storage
            .ref()
            .child(`Transfermoney/transfermoney-${file.raw.uid}`)
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
        id_form: this.transfermdetail.id_form,
        istransfermomey: true,
        img_urltransfer: url,
      };
      await this.$axios
        .patch(`${env.BASE_URL}/approve/transfermomey`, data)
        .then((response) => {
          alert("เพิ่มหลักฐานการชำระเงินเรียบร้อย");
          window.location.reload(true);
        })
        .catch((err) => {
          alert(err);
        });
    },

    async gettransfermdetail() {
      await this.$axios
        .get(`${env.BASE_URL}/wer/approve`)
        .then((response) => {
          this.datatable = response.data?.data;
        });
    },
    view(item) {
      this.transfermdetail = item;
    },
  },
};
</script>

<style lang="scss" scoped>
.dialogtransfer {
  display: flex;
  flex-direction: column;
  padding: 20px;

  img {
    margin: 25px;
  }
}
</style>
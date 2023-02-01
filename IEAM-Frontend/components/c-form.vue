<template>
  <v-col>
    <v-card class="content-card" elevation="6">
      <v-form ref="form" @submit.prevent="onsubmit">
        <v-row>
          <v-col cols="12">
            <label for="">หัวข้อรายการ</label>
            <v-text-field
              v-model="list.title"
              placeholder="กรอกชื่อรายการ"
              dense
              outlined
              :rules="rules.req"
            ></v-text-field>
          </v-col>
          <v-col>
            <label for="">ประเภท</label>
            <v-select
              :disabled="mode == 'r'"
              :items="type"
              item-value="id"
              item-text="name"
              v-model="list.typeId"
              placeholder="เลือกประเภท"
              dense
              outlined
              :rules="rules.req"
            ></v-select>
          </v-col>
          <v-col>
            <label for="">ยอดรวม</label>
            <v-text-field
              v-model="list.total"
              :disabled="mode == 'r'"
              type="number"
              placeholder="กรอกยอดรวม"
              dense
              outlined
              :rules="rules.req"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row v-if="mode != 'r'">
          <v-col>
            <label for="">รายการ</label>
          </v-col>
          <v-col cols="12">
            <v-data-table
              :headers="headers"
              :items="list.list"
              :loading="false"
              hide-default-footer
              loading-text="Loading... Please wait"
              no-data-text="ไม่พบข้อมูล"
            >
              <template v-slot:[`item.actions`]="{ item }">
                <div>
                  <v-menu>
                    <template v-slot:activator="{ on, attrs }">
                      <v-icon v-bind="attrs" v-on="on"
                        >mdi-dots-vertical</v-icon
                      >
                    </template>
                    <v-list>
                      <v-list-item link>
                        <v-list-item-title @click="del(item)"
                          ><v-icon>mdi-delete</v-icon> ลบ</v-list-item-title
                        >
                      </v-list-item>
                    </v-list>
                  </v-menu>
                </div>
              </template>
            </v-data-table>
          </v-col>
        </v-row>
        <v-row v-if="mode != 'r'">
          <v-col style="cursor: pointer">
            <v-btn color="primary" depressed small @click="dialog = true"
              >เพิ่มรายการ</v-btn
            >
          </v-col>
        </v-row>
        <v-divider class="my-5" />
        <v-row>
          <v-col cols="12">
            <label for="">หลักฐาน</label>
          </v-col>
          <v-col cols="4">
            <el-upload
              action="#"
              :limit="1"
              :disabled="this.mode == 'v' ? true : false"
              :file-list="selectedFile"
              :on-change="addfile"
              list-type="picture-card"
              :auto-upload="false"
              accept="image/*"
              v-if="selectedFile"
            >
              <i slot="default" class="el-icon-plus"></i>
              <div slot="file" slot-scope="{ file }" style="height: 100%">
                <img
                  class="el-upload-list__item-thumbnail"
                  :src="file.url || list.img_url"
                  alt=""
                />
                <span v-if="mode != 'v'" class="el-upload-list__item-actions">
                  <span
                    class="el-upload-list__item-delete"
                    @click="handleRemove(file)"
                  >
                    <i class="el-icon-delete"></i>
                  </span>
                </span>
              </div>
            </el-upload>
          </v-col>
        </v-row>
        <v-row class="space-center">
          <v-btn
            class="mx-2"
            min-width="10%"
            color="primary"
            type="submit"
            rounded
            depressed
            >บันทึกรายการ</v-btn
          >
          <v-btn
            class="mx-2"
            min-width="10%"
            color="error"
            rounded
            depressed
            outlined
            @click="$router.go(-1)"
            >กลับ</v-btn
          >
        </v-row>
      </v-form>
    </v-card>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span>{{ this.mode != "e" ? "เพิ่มรายการ" : "แก้ไขรายการ" }}</span>
          <v-spacer></v-spacer>
          <v-menu bottom left>
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on" @click="dialog = false">
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </template>
          </v-menu>
        </v-card-title>
        <v-card-text>
          <v-form ref="formlist" v-model="isValid" @submit.prevent="onaddlist">
            <v-row>
              <v-col cols="12">
                <span>ขื่อรายการ</span>
                <v-text-field
                  v-model="detaillist.title"
                  outlined
                  dense
                  placeholder="ขื่อรายการ"
                  :rules="rules.req"
                />
              </v-col>
              <v-col cols="6">
                <span>จำนวน</span>
                <v-text-field
                  v-model="detaillist.number"
                  outlined
                  dense
                  placeholder="จำนวน"
                  @change="total()"
                  type="number"
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col cols="6">
                <span>หน่วย</span>
                <v-text-field
                  v-model="detaillist.unit"
                  outlined
                  dense
                  placeholder="หน่วย"
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col>
                <span>ราคาต่อหน่วย</span>
                <v-text-field
                  v-model="detaillist.unit_price"
                  outlined
                  dense
                  placeholder="ราคาต่อหน่วย"
                  suffix="บาท"
                  type="number"
                  @change="total()"
                  :rules="rules.req"
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <span>ยอดรวม</span>
                <v-text-field
                  v-model="detaillist.total"
                  disabled
                  outlined
                  dense
                  placeholder="ยอดรวม"
                  suffix="บาท"
                  type="number"
                  :rules="rules.req"
                  hide-details=""
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="space-center">
              <v-btn class="ml-5" color="primary" type="submit">ยืนยัน</v-btn>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <c-loader :isloader="loading" />
  </v-col>
</template>

<script>
import env from "~/env";
import { storage } from "~/plugins/firebase";
export default {
  props: {
    mode: {
      type: String,
      default: "c",
    },
  },
  data() {
    return {
      loading:false,
      type: [
        { id: 3, name: "การเบิก" },
        { id: 4, name: "การยืม" },
      ],
      isValid: null,
      dialog: false,
      isSelecting: false,
      selectedFile: [],
      list: {
        number_form: null,
        title: null,
        typeId: null,
        total: null,
        list: [],
        img_url: "",
      },
      detaillistbackup: {},
      detaillist: {
        title: "",
        number: "",
        unit: "",
        unit_price: "",
        total: 0,
      },
      headers: [
        {
          text: "รายการ",
          align: "center",
          value: "title",
          sortable: false,
        },
        {
          text: "จำนวน",
          value: "number",
          sortable: false,
        },
        { text: "หน่วย", value: "unit", sortable: false },
        { text: "ราคาต่อหน่วย", value: "unit_price", sortable: false },
        {
          text: "ยอดรวม",
          align: "center",
          value: "total",
          sortable: false,
          width: "10%",
        },
        {
          text: " ",
          align: "center",
          value: "actions",
          sortable: false,
          width: "5%",
        },
      ],
      rules: {
        req: [(v) => !!v || "กรุณาใส่ข้อมูล"],
      },
    };
  },

  watch: {
    dialog() {
      if (!this.dialog) {
        this.detaillist = {
          title: "",
          number: "",
          unit: "",
          unit_price: 0,
          total: 0,
        };
        this.$refs.formlist.reset();
      }
    },
  },
  mounted() {
    if (this.$auth.state.user.level == 3) {
      this.type = [
        { id: 1, name: "รายรับ" },
        { id: 2, name: "รายจ่าย" },
      ];
    }
    if (this.mode != "c") {
      this.$axios
        .get(`${env.BASE_URL}/form/detail`, {
          params: this.$route.params,
        })
        .then((response) => {
          if (this.mode == "r") {
            this.list.total = response.data?.data[0].total;
            this.list.title = response.data?.data[0].title;
          } else {
            this.list = response.data?.data[0];
            this.selectedFile.push({ url: this.list.img_url });
          }
        });
    }
    if (this.mode == "r") {
      this.id_form = this.$route.params;
      this.type = [{ id: 5, name: "การคืน" }];
      this.list.typeId = 5;
      // this.list.totall =
    }
  },
  methods: {
    onaddlist() {
      if (this.$refs.formlist.validate()) {
        this.list.list.push(this.detaillist);
        this.dialog = !this.dialog;
      }
    },
    addfile(files, fileList) {
      this.selectedFile.push(fileList[0]);
    },
    handleRemove(file) {
      for (let index = 0; index < this.selectedFile.length; index++) {
        if (this.selectedFile[index].uid === file.uid) {
          this.selectedFile.splice(index, 1);
        }
      }
    },
    total() {
      this.detaillist.total =
        parseInt(this.detaillist.unit_price) * parseInt(this.detaillist.number);
    },

    del(item) {
      this.list.list.splice(this.list.list.indexOf(item), 1);
    },
    uploadFirebase(file) {
      // for (let index = 0; index <= file.length; index++) {
      if (file.status == "ready") {
        return new Promise((resolve, reject) => {
          let metadata = {
            contentType: file.raw.type,
          };

          let task = storage
            .ref()
            .child(`Documents/documents-${file.raw.uid}`)
            .put(file.raw, metadata);

          task
            .then((snapshot) => {
              snapshot.ref.getDownloadURL().then((url) => {
                this.list.img_url = url;
                resolve(url);
              });
            })
            .catch((error) => {
              reject(error);
              this.$message.error("อัพโหลดรูปไม่สำเร็จ");
            });
        });
        // }
      }
    },
    async onsubmit() {
      if (this.$refs.form.validate()) {
        this.loading = true;
        if (this.selectedFile[0]) {
          await this.uploadFirebase(this.selectedFile[0]);
        }
        if (this.mode != "e") {
          this.mode == "r"
            ? await this.$axios
                .post(`${env.BASE_URL}/form/addreturn`, {
                  id_form: parseInt(this.$route.params.id),
                  returnborrow: {
                    ...this.list,
                    total: parseInt(this.list.total),
                    userId: this.$auth.state.user.accountId,
                  },
                })
                .then((response) => {
                  alert("เพิ่มข้อมูลสำเร็จ(รอการอนุมัติ)");
                  this.$router.replace("/arrears");
                })
            : await this.$axios
                .post(`${env.BASE_URL}/form/add`, {
                  ...this.list,
                  total: parseInt(this.list.total),
                  userId: this.$auth.state.user.accountId,
                })
                .then((response) => {
                  alert("เพิ่มข้อมูลสำเร็จ(รอการอนุมัติ)");
                  this.$router.replace("/accountIE" || "/dashboard");
                });
        } else {
          this.uploadFirebase(this.selectedFile[0]);
          await this.$axios
            .put(`${env.BASE_URL}/form/update`, {
              ...this.list,
              total: +this.list.total,
              userId: this.$auth.state.user.accountId,
            })
            .then((response) => {
              alert("แก้ไขข้อมูลสำเร็จ(รอการอนุมัติ)");
              this.$router.replace("/accountIE");
              this.$router.go(-1);
            });
        }
      }
      this.loading = false
    },
  },
};
</script>

<style>
</style>
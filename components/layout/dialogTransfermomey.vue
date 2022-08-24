<template>
  <v-card>
    <v-card-title>
      <span>{{ detail.type }}</span>
      <v-spacer></v-spacer>
      <v-menu bottom left>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on"  @click="$emit('dialogview-click')">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </template>
      </v-menu>
    </v-card-title>
    <v-card-text>
      <v-row class="d-flex">
        <v-col>
          <p>หัวเรื่อง : {{ detail.title }}</p>
          <p>เลขใบฟอร์ม : {{ detail.id_form }}</p>
        </v-col>
        <v-col class="ml-10">
          <p>วันที่ : {{ detail.CreatedAt }}</p>
          <p>จำนวนเงิน : {{ detail.total }} บาท</p>
        </v-col>
      </v-row>
      <v-col>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">รายการ</th>
                <th class="text-left">จำนวน</th>
                <th class="text-left">หน่วย</th>
                <th class="text-left">ราคา/หน่วย</th>
                <th class="text-left">รวม</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="item in detail.list"
                :key="item.id_listincomeexpense || item.id_listtransfer"
              >
                <td>{{ item.title ? item.title : "-" }}</td>
                <td>{{ item.number ? item.number : "-" }}</td>
                <td>{{ item.unit ? item.unit : "-" }}</td>
                <td>{{ item.unit_price ? item.unit_price : "-" }}</td>
                <td>{{ item.total ? item.total : "-" }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
      <v-col>
        <p>เอกสารแนบ :</p>
        <v-img
          class="ma-auto"
          lazy-src="https://picsum.photos/id/11/10/6"
          max-width="50%"
          src="https://picsum.photos/id/11/500/300"
        ></v-img>
      </v-col>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    detail: {
      type: Object,
      default: {},
    },
  },
};
</script>

<style>
</style>
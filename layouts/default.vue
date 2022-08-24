<template>
  <v-app>
    <v-navigation-drawer
    
      mini-variant
      mini-variant-width="70"
      app
      style="border-right: 5px solid #eb3c41; drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25));"
    >
      <v-list-item class="px-5">
        <v-img :src="require('../assets/images/login_logo.png')" />
      </v-list-item>

      <v-divider class="ma-auto" style="width: 70%"></v-divider>
      <v-list>
        <v-list-item-group class="groupManu">
          <v-list-item
            v-for="item in items"
            :key="item.title"
            :to="item.to"
            class="manu"
            active-class="activeManu"
            exact
          >
            <v-tooltip right>
              <template v-slot:activator="{ on, attrs }">
                <v-icon v-bind="attrs" v-on="on">
                  {{ item.icon }}
                </v-icon>
              </template>
              <span>{{ item.title }}</span>
            </v-tooltip>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      elevation="1"
      dense
      app
      style="border-bottom: 2px solid #eb3c41; background: #ffffff"
    >
      <v-toolbar-title>
        <span style="font-size: 15px">{{ title }}</span>
      </v-toolbar-title>
      <v-spacer />
      ............<v-icon to="/login">mdi-login-variant</v-icon>
    </v-app-bar>

    <!-- <v-container class="container--fluid" style="margin-top: 54px"> -->
    <v-main>
      <Nuxt />
    </v-main>
    <!-- </v-container> -->
  </v-app>
</template>

<script>
export default {
  name: "DefaultLayout",
  data() {
    return {
      mini: true,
      items: [
        {
          icon: "mdi-layers",
          title: "Dashboard",
          to: "/Dashboard",
        },
        {
          icon: "mdi-cash-sync",
          title: "การจัดการบัญชี โอนเงิน",
          to: "/Transfermoney",
        },
        {
          icon: "mdi-cash-multiple",
          title: "การจัดการบัญชี รายรับ/รายจ่าย",
          to: "/AccountIE",
        },
        {
          icon: "mdi-clipboard-check",
          title: "การจัดการอนุมัติ เบิก/ยืม/คืน",
          to: "/AccountWBR/approve",
        },
        {
          icon: "mdi-clipboard",
          title: "การจัดการบัญชี เบิก/ยืม/คืน ",
          to: "/AccountWBR",
        },
        {
          icon: "mdi-clipboard-plus",
          title: "เพิ่มแบบฟอร์ม เบิก/ยืม/คืน",
          to: "/Documents",
        },
      ],
      title: "ระบบบริหารจัดการบัญชีรายรับรายจ่าย ด้วยแพลตฟอร์มออนไลน์ ",
    };
  },
  created() {
    console.log(this.$route.path);
  },
};
</script>

<style lang="scss" scoped>
.groupManu {
  .manu::before {
    width: 0px;
  }

  .manu a {
    text-decoration: none;
  }

  .activeManu::before {
    width: 5px;
    background: $color_blue;
    border-radius: 0px 5px 5px 0px;
    opacity: 1;
    box-shadow: $shadow_layouts;
    transition: ease-in-out 0.3s;
  }
  .activeManu > i {
    color: $color_blue;
  }
}
</style>
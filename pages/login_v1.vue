<template>
  <div class="background">
    <div class="login__container">
      <img src="~/assets/images/login_logofull.png" width="40%" />
      <div class="my-5">
        <p>ระบบบริหารจัดการบัญชีรายรับรายจ่าย</p>
        <p>ด้วยแพลตฟอร์มออนไลน์</p>
      </div>

      <form class="login__form" @submit.prevent="login">
        <v-row justify="center" align="center">
          <v-col justify="center" align-self="center" cols="7">
                        <v-text-field
              v-model="formres.serial"
              label="SerialNumber "
            ></v-text-field>
          </v-col>
          <v-col cols="7">
            <c-login-input
              v-model="password"
              :type="i_type ? 'password' : 'text'"
              placeholder="รหัสผ่าน"
              required
            >
            </c-login-input>
          </v-col>
          <v-col cols="7" align="right" class="mr-4"> <nuxt-link to="/">ลืมรห้สผ่าน</nuxt-link> </v-col>
          <v-col cols="7">
            <c-login-button class="login__button" type="submit"
              >เข้าสู่ระบบ</c-login-button
            >
          </v-col>
        </v-row>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: "",
      password: "",
      i_type: true,
    };
  },
  methods: {
    login() {




      if (username == "admin" && password == "admin") {
        setTimeout(() => {
                this.$router.push("/Dashbord");
              }, 800);
              
      }
    },
  },
};
</script>

<style lang="scss" >
.background {
  width: 100%;
  height: 100%;
  background-image: url("~/assets/images/login_bg.png");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
}

.login__container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;

  p {
    margin: 0;
  }
}
.login__container > .v-input__slot {
  box-shadow: inset 0px 4px 4px rgb(0 0 0 / 25%);
}
</style>
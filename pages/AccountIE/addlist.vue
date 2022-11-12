<template>
  <div class="content-size">
    <layout-title-page
      title="เพิ่มรายการ"
      details="เพิ่มรายการบัญชีรายรับ-รายจ่าย"
    />
    <c-form type="add"/>
  </div>
</template>

<script>
export default {
  mounted() {
    this.lv =  [2];
    if(!this.lv.includes(this.$auth.state.user.level)){
      this.$router.push("/dashboard");
    }
  },
};
</script>

<style>
</style>
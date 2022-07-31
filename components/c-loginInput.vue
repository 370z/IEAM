<template>
  <div class="login__input-container">
    <input v-bind="$attrs" v-on:input="$emit('input', $event.target.value)" >
    <!-- <v-text-field
      solo
      rounded
      append-icon="mdi-eye"
     
    ></v-text-field> -->
    <template v-for="(_, slot) in $slots">
      <template :slot="slot">
        <slot :name="slot"></slot>
      </template>
    </template>
  </div>
</template>

<script>
export default {
  created() {
    console.log({ "----> attrs" : this.$attrs ,"----> props": this.$props ,"----> slots":this.$slots})
  }
};
</script>

<style scoped>

.login__input-container input {
  height: 40px;
  width: 100%;
  background: #ffffff;
  box-shadow: inset 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 100px;
  padding-left: 16px;
  font-size: 15px;
}
/* .v-input__control .v-input__slot {
  height: 40px;
  width: 100%;
  background: #ffffff;
  box-shadow: inset 0px 4px 4px rgba(0, 0, 0, 0.25) !important;
  border-radius: 100px;
  padding-left: 16px;
} */
</style>





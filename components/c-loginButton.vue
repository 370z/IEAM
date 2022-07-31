<template>
  <div class="login__button">
    <button>
        <slot></slot>
    </button>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.login__button button {
  width: 100%;
  height: 40px;

  background: linear-gradient(
    90.94deg,
    #192786 2.82%,
    rgba(47, 66, 192, 0.87) 67.61%
  );
  box-shadow: 1px 4px 8px -1px rgba(0, 0, 0, 0.25);
  border-radius: 5px;
  color: #ffffff;
  font-size: 15px;
}
</style>
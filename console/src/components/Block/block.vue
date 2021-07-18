<template>
  <div v-text="text" ></div>
</template>

<script>
export default {
  props: {
    text: {
      type: String,
      default: ''
    },
  }
}
</script>

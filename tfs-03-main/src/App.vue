<template>
  <main>
    <h1> Demo SFCs </h1>
    <ComponentDemo />
  </main>
</template>

<script>
import ComponentDemo from "./components/ComponentDemo.vue";
export default {
  name: "App",
  components: { ComponentDemo },
}
</script>

import Vue from "vue";
export default new Vue({
  methods: {
    err(e: Error) {
      alert(e);
    }
  }
});

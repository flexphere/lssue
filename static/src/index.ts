import Vue from "vue";
import AppComponent from "./components/App/app.vue";

let v = new Vue({
  el: "#app",
  template: `<app-component />`,
  components: {
    AppComponent
  }
});

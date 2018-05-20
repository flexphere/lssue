import Vue from "vue";
import DraggableComponent from "vuedraggable";
import TicketComponent from "../Ticket/ticket.vue";
import ebus from "../../lib/EBus";
import ws from "../../lib/WS";
import api from "../../lib/API";
import * as model from "../../lib/model";

export default Vue.extend({
  props: ["selected_board", "boards", "categories", "labels"],
  data() {
    return {
      visible: true
    };
  },
  created() {
    ebus.$on("settings", () => {
      this.visible = true;
    });
  },
  methods: {
    close() {
      this.visible = false;
    }
  }
});

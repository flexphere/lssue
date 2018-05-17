import Vue from "vue";
import DraggableComponent from "vuedraggable";
import TicketComponent from "../Ticket/ticket.vue";
import ebus from "../../lib/EBus";
import ws from "../../lib/WS";
import api from "../../lib/API";
import * as model from "../../lib/model";

export default Vue.extend({
  components: {
    DraggableComponent,
    TicketComponent
  },
  props: ["selected_board", "pipe"],
  methods: {
    newTicket() {
      ebus.$emit("newTicket", this.selected_board, this.pipe.id);
    },
    async movedTicket() {
      try {
        let param = <model.TicketSortParam>{
          board_name: this.selected_board,
          pipe_id: this.pipe.id,
          tickets: this.pipe.tickets.map((t: model.Ticket) => t.id)
        };
        await api.ticket.sort(param);
        ws.$emit("send", ws.UPDATE_TICKET);
      } catch (e) {
        alert(e);
      }
    }
  }
});

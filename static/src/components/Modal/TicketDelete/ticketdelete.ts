import Vue from "vue";
import ws from "../../../lib/WS";
import api from "../../../lib/API";
import ebus from "../../../lib/EBus";
import * as model from "../../../lib/model";

export default Vue.extend({
  data() {
    return {
      form: <model.TicketDeleteParam>{
        board_name: "",
        id: 0
      }
    };
  },
  created() {
    ebus.$on("deleteTicket", async (board_name: string, ticket_id: number) => {
      try {
        this.form.board_name = board_name;
        this.form.id = ticket_id;
        await api.ticket.delete(this.form);
        ws.$emit("send", ws.UPDATE_TICKET);
      } catch (e) {
        alert(e);
      }
    });
  }
});

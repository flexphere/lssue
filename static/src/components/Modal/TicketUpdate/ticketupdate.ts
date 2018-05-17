import Vue from "vue";
import DatepickerComponent from "vuejs-datepicker";
import ws from "../../../lib/WS";
import api from "../../../lib/API";
import ebus from "../../../lib/EBus";
import * as model from "../../../lib/model";

export default Vue.extend({
  props: ["pipes", "categories", "labels", "issues", "selectedBoard"],
  components: {
    DatepickerComponent
  },
  data() {
    return {
      formVisible: false,
      form: <model.TicketUpdateParam>{
        board_name: "",
        id: 0,
        title: "",
        due: "",
        memo: "",
        pipe_id: 1,
        category_id: 1,
        label_ids: []
      }
    };
  },
  created() {
    ebus.$on("updateTicket", (board_name: string, ticket: model.Ticket) => {
      this.form.board_name = board_name;
      this.form = Object.assign(this.form, ticket);
      this.formVisible = true;
    });
  },
  methods: {
    async submit() {
      try {
        await api.ticket.update(this.form);
        ws.$emit("send", ws.UPDATE_TICKET);
        this.closeModal();
      } catch (e) {
        alert(e);
      }
    },
    async deleteTicket() {
      ebus.$emit("deleteTicket", this.form.board_name, this.form.id);
      this.closeModal();
    },
    closeModal() {
      this.formVisible = false;
    }
  }
});

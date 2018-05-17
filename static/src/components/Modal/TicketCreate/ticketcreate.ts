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
      form: <model.TicketCreateParam>{
        board_name: "",
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
    ebus.$on("newTicket", (board_name: string, pipe_id: number) => {
      this.form.board_name = board_name;
      this.form.pipe_id = pipe_id;
      this.formVisible = true;
    });
  },
  methods: {
    async submit() {
      try {
        console.log(this.form);
        await api.ticket.create(this.form);
        ws.$emit("send", ws.UPDATE_TICKET);
        this.closeModal();
      } catch (e) {
        alert(e);
      }
    },
    closeModal() {
      this.formVisible = false;
    }
  }
});

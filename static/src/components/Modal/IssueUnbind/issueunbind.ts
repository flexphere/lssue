import Vue from "vue";
import SelectComponent from "vue-select";
import ws from "../../../lib/WS";
import api from "../../../lib/API";
import ebus from "../../../lib/EBus";
import * as model from "../../../lib/model";

export default Vue.extend({
  components: {
    SelectComponent
  },
  data() {
    return {
      form: <model.IssueUnbindParam>{
        board_name: "",
        id: 0,
        ticket_id: 0
      }
    };
  },
  created() {
    ebus.$on(
      "unbindIssue",
      async (board_name: string, ticket_id: number, issue_id: number) => {
        try {
          this.form.board_name = board_name;
          this.form.ticket_id = ticket_id;
          this.form.id = issue_id;
          await api.issue.unbind(this.form);
          ws.$emit("send", ws.UPDATE_TICKET);
        } catch (e) {
          alert(e);
        }
      }
    );
  }
});

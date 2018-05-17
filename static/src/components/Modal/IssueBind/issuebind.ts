import Vue from "vue";
import SelectComponent from "vue-select";
import ws from "../../../lib/WS";
import api from "../../../lib/API";
import ebus from "../../../lib/EBus";
import * as model from "../../../lib/model";

export default Vue.extend({
  props: ["pipes", "categories", "labels", "issues", "selectedBoard"],
  components: {
    SelectComponent
  },
  data() {
    return {
      formVisible: false,
      selectedIssue: null,
      form: <model.IssueBindParam>{
        board_name: "",
        id: 0,
        ticket_id: 0
      }
    };
  },
  created() {
    ebus.$on("bindIssue", (board_name: string, ticket: model.Ticket) => {
      this.form.board_name = board_name;
      this.form.ticket_id = ticket.id;
      this.formVisible = true;
    });
  },
  methods: {
    async submit() {
      try {
        this.form.id = (<any>this.selectedIssue).value;
        await api.issue.bind(this.form);
        this.closeModal();
        ws.$emit("send", ws.UPDATE_TICKET);
      } catch (e) {
        alert(e);
      }
    },
    closeModal() {
      this.formVisible = false;
    }
  },
  computed: {
    issue_options(): Array<any> {
      return this.issues.map((issue: model.Issue) => {
        return {
          label: issue.title,
          value: issue.id
        };
      });
    }
  }
});

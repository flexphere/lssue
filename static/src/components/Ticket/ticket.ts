import Vue from "vue";
import * as moment from "moment";
import ebus from "../../lib/EBus";
import IssueComponent from "../Issue/issue.vue";

export default Vue.extend({
  components: { IssueComponent },
  props: ["ticket", "selected_board"],
  methods: {
    updateTicket() {
      ebus.$emit("updateTicket", this.selected_board, this.ticket);
    },
    bindIssue() {
      ebus.$emit("bindIssue", this.selected_board, this.ticket);
    }
  },
  computed: {
    overdue(): boolean {
      return (
        this.ticket.due &&
        moment(this.ticket.due).isSameOrBefore(moment(), "day")
      );
    },
    duedate(): string {
      return moment(this.ticket.due).format("YYYY-MM-DD");
    }
  }
});

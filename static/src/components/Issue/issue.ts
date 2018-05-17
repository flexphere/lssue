import Vue from "vue";
import * as moment from "moment";
import ebus from "../../lib/EBus";

export default Vue.extend({
  props: ["ticket", "issue", "selected_board"],
  methods: {
    unbindIssue() {
      ebus.$emit(
        "unbindIssue",
        this.selected_board,
        this.ticket.id,
        this.issue.id
      );
    }
  },
  computed: {
    isopen(): boolean {
      return true;
    }
  }
});

import Vue from "vue";
import PipeComponent from "../Pipe/pipe.vue";
import TicketCreateComponent from "../Modal/TicketCreate/ticketcreate.vue";
import TicketUpdateComponent from "../Modal/TicketUpdate/ticketupdate.vue";
import TicketDeleteComponent from "../Modal/TicketDelete/ticketdelete.vue";
import IssueBindComponent from "../Modal/IssueBind/issuebind.vue";
import IssueUnbindComponent from "../Modal/IssueUnbind/issueunbind.vue";
import api from "../../lib/API";
import ws from "../../lib/WS";
import * as model from "../../lib/model";

export default Vue.extend({
  components: {
    PipeComponent,
    TicketCreateComponent,
    TicketUpdateComponent,
    TicketDeleteComponent,
    IssueBindComponent,
    IssueUnbindComponent
  },
  data() {
    return {
      selected_board: "",
      boards: <model.Board[]>[],
      pipes: <model.Pipe[]>[],
      tickets: <model.Ticket[]>[],
      categories: <model.Category[]>[],
      labels: <model.Label[]>[],
      issues: <model.Issue[]>[]
    };
  },
  computed: {
    default_param(): model.CommonParam {
      return {
        board_name: this.selected_board
      };
    }
  },
  beforeMount() {
    this.websocketHandler();
    this.getBoardList();
  },
  methods: {
    websocketHandler() {
      ws.$on("receive", (msg: model.WS_Message) => {
        switch (msg) {
          case ws.UPDATE_TICKET:
            this.getTickets();
            break;
          case ws.UPDATE_LABEL:
            this.getLabels();
            break;
          case ws.UPDATE_CATEGORY:
            this.getCategories();
            break;
          case ws.UPDATE_ISSUE:
            this.getIssues();
            break;
        }
      });
    },
    async getBoardList() {
      try {
        this.boards = await api.board.list();
        this.selected_board = <string>this.boards[1];
        this.getBoard();
      } catch (e) {
        alert(e);
      }
    },
    async getBoard() {
      try {
        await this.getPipes();
        await this.getCategories();
        await this.getLabels();
        await this.getIssues();
        await this.getTickets();
        ws.$emit("connect", this.selected_board);
      } catch (e) {
        alert(e);
      }
    },
    async getPipes() {
      try {
        let pipes = await api.pipe.list(this.default_param);
        this.pipes = pipes;
      } catch (e) {
        alert(e);
      }
    },
    async getCategories() {
      try {
        let categories = await api.category.list(this.default_param);
        this.categories = categories;
      } catch (e) {
        alert(e);
      }
    },
    async getLabels() {
      try {
        let labels = await api.label.list(this.default_param);
        this.labels = labels;
      } catch (e) {
        alert(e);
      }
    },
    async getIssues() {
      try {
        let issues = await api.issue.list(this.default_param);
        this.issues = issues;
      } catch (e) {
        alert(e);
      }
    },
    async getTickets() {
      try {
        let tickets = await api.ticket.list(this.default_param);
        for (let i = 0; i < this.pipes.length; i++) {
          this.$set(this.pipes[i], "tickets", []);
        }

        for (let t of tickets) {
          let category = this.categories.find(c => c.id === t.category_id);
          t.category = category ? category : <model.Category>{};

          t.label_ids = t.label_ids ? t.label_ids : [];
          t.labels = this.labels.filter(l => t.label_ids.includes(l.id));

          t.issue_ids = t.issue_ids ? t.issue_ids : [];
          t.issues = this.issues.filter(i => t.issue_ids.includes(i.id));

          let pipe = this.pipes.find(pipe => pipe.id === t.pipe_id);
          if (pipe) pipe.tickets.push(t);
        }

        this.tickets = tickets;
      } catch (e) {
        alert(e);
      }
    }
  }
});

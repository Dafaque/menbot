import Table from "../framework/table.js";

class ChatUsersView extends Table {
    constructor(data) {
        super();
        this.setTitle("Chat Users");
        this.table = new Table();
        this.chat = data;
        this.users = [];
        this.fetchUsers();
    }


    async fetchUsers() {
        fetch(`/api/chats/${this.chat.id}/users`).then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error("Failed to fetch users");
        }).then(data => {
            if (data.UserListResponseJSONResponse) {
                this.data = data.UserListResponseJSONResponse;
                this.render();
            }
        });
    }
}

export default ChatUsersView;
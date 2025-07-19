import Table from "../framework/table.js";

class ChatRolesListView extends Table {
    constructor(data) {
        super();
    }


    async fetchRoles() {
        this.data = [];
        fetch(`/api/chats/${this.chat.id}/roles`).then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error("Failed to fetch roles: "+ response.statusText);
        }).then(data => {
            if (data.RoleListResponseJSONResponse) {
                this.data = data.RoleListResponseJSONResponse;
                this.render();
            }
        });
    }

    onSelected = (row) => {
        window.app.router.navigate("/chats/details/roles/list/users", row);
    }

    appear = (data) => {
        if (data) {
            this.chat = data;
            this.setTitle(`Chat ${data.tg_chat_name} Roles List`);
        }
        this.fetchRoles();
    }
}

export default ChatRolesListView;
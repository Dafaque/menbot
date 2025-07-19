import Table from "../framework/table.js";

class ChatRolesListView extends Table {
    constructor(data) {
        super();
        this.setTitle("Chat Roles List");
        this.chat = data;
        this.fetchRoles();
    }


    async fetchRoles() {
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
}

export default ChatRolesListView;
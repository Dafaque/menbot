import { Table } from "kibodo";

class ChatsView extends Table {
    constructor() {
        super();
        this.setTitle("Chats");
        this.error = null;
    }

    async fetchChats() {
        this.data = [];
        fetch("/api/chats")
            .then(response => response.json())
            .then(this.onChatsLoaded.bind(this), this.onChatsError.bind(this))
            .catch(this.onChatsError.bind(this));
    }

    onChatsLoaded(chats) {
        if (!chats.ChatListResponseJSONResponse) {
            this.onChatsError("Failed to fetch chats");
            return;
        }
        this.data = chats.ChatListResponseJSONResponse;
        this.render();
    }
    onChatsError(error) {
        this.error = error;
        this.render();
    }
    
    onSelected = (row) => {
        window.app.router.navigate("/chats/details", row);
    }

    appear = () => {
        this.fetchChats();
    }
}

export default ChatsView;
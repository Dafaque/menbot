import { Table } from "kibodo";
import ChatDetailsView from "./chat-details";

export default class ChatsView extends Table {
    constructor() {
        super(
            {
                url: "/api/chats",
                attr: "ChatListResponseJSONResponse",
            }
        );
        this.title = "Chats";
        this.error = null;
    }
    
    onSelected (row: any) {
        window.app.push(new ChatDetailsView(row)).then(() => {
            this.fetchData();
        });
    }

}

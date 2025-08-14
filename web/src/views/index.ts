import { Menu } from "kibodo";
import ChatsView from "./chats";

export default class IndexView extends Menu {
    constructor() {
        super();
        this.title = "Index";
        this.addItem("Chats", this.gotoChats);
    }
    gotoChats() {
        window.app.push(new ChatsView());
    }
}

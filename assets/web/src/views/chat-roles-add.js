import Form from "../framework/form.js";

class ChatRolesAddView extends Form {
    constructor(data) {
        super();
    }

    onSave = (values) => {
        fetch(`/api/chats/${this.chat.id}/roles`, {
            method: "POST",
            body: JSON.stringify({
                name: values.name,
            }),
        }).then(response => {
            if (response.ok) {
                this.goBack();
            }
        });
    }

    appear = (data) => {
        if (data) {
            this.chat = data;
        }
        this.setTitle(`Add Role to Chat ${this.chat.tg_chat_name}`);
        this.addField("name", "Name", "text");
        this.render();
    }
}

export default ChatRolesAddView;
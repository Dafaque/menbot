import Form from "../framework/form.js";

class ChatRolesAddView extends Form {
    constructor(data) {
        super();
        this.setTitle("Add Role");
        this.data = data;
        this.addField("name", "Name", "text");
    }

    onSave = (values) => {
        fetch(`/api/chats/${this.data.id}/roles`, {
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
}

export default ChatRolesAddView;
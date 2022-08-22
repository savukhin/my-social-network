export class SendedMessage {
    token = "";
    text = "";
    time = new Date();
    author_id = 0;
    chat_id = 0;

    constructor(text="", time=new Date(), authorId=0) {
        this.text = text;
        this.time = time;
        this.author_id = authorId;
    }

}

export class Message {
    text = "";
    time = new Date();
    author_id = 0;
    chat_id = 0;

    constructor(text="", time=new Date(), authorId=0) {
        this.text = text;
        this.time = time;
        this.author_id = authorId;
    }

}
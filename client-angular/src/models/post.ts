import { PhotoAttachement } from "./photoAttachement";

export class Post {
    id = 0;
    text = "";
    author_id = 0;
    chat_id = 0;
    created_at = new Date();
    updated_at = new Date();
	photos: PhotoAttachement[] = []
}

import { PhotoAttachement } from "./photoAttachement";

export class LikePost {
    post_id = 0
	user_id = 0
}

export class Post {
    id = 0;
    text = "";
    author_id = 0;
    chat_id = 0;
    created_at = new Date();
    updated_at = new Date();
	photos: PhotoAttachement[] = []
    likes: LikePost[] = []
    has_current_user_like = false;

    current_likes = 0
}

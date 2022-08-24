import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { environment } from 'src/environments/environment';
import { Chat, ChatDTO } from 'src/models/chat';
import { Message } from 'src/models/message';
import { Post } from 'src/models/post';
import { AuthService } from './auth.service';

@Injectable()
export class ContentService {

  constructor(private auth: AuthService, private http: HttpClient) { }

  getUserPosts (user_id: number) {
    let observer = this.http.get<Post[]>(
      `${environment.serverUrl}/api/posts/user_posts/${user_id}`, 
      {observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          let posts = response.body as Post[]
          for (let i = 0; i < posts.length; i++) {
            posts[i].created_at = new Date(posts[i].created_at)
            posts[i].updated_at = new Date(posts[i].updated_at)
          }
          return posts
        }
        return false;
      })
    )
  }

  createPosts (text: string) {
    if (this.auth.user == undefined)
      return false
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return false
      
    let observer = this.http.post<Post>(
      `${environment.serverUrl}/api/posts/create_post`, 
      { text },
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as Post
        }
        return false;
      })
    )
  }

  
}

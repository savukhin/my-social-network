import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { environment } from 'src/environments/environment';
import { Chat } from 'src/models/chat';
import { Message } from 'src/models/message';
import { AuthService } from './auth.service';

@Injectable()
export class ContentService {

  constructor(private auth: AuthService, private http: HttpClient) { }

  getChats () {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.get<Chat>(
      `${environment.serverUrl}/api/chats`, 
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as Chat
        }
        return false;
      })
    )
  }

  getPersonalChat (user_id: number) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<Chat>(
      `${environment.serverUrl}/api/chat/by_user/${user_id}`, 
      {},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as Chat
        }
        return false;
      })
    )
  }

  getPersonalChatMessages (offset: number, count: number, chat_id: number) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<{messages: Message[]}>(
      `${environment.serverUrl}/api/chat/getMessages`, 
      {offset, count, chat_id},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as {messages: Message[]}
        }
        return false;
      })
    )
  }

  sendMessage (chat_id: number, text: string) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<{message: string}>(
      `${environment.serverUrl}/api/chat/sendMessage`, 
      {chat_id, text},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as {message: string}
        }
        return false;
      })
    )
  }
}

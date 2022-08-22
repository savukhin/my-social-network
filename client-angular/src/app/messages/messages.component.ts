import { Component, OnInit } from '@angular/core';
import { Chat } from 'src/models/chat';
import { Message } from 'src/models/message';
import { User } from 'src/models/user';
import { ContentService } from '../services/backend-api/content.service';

class Item {
  message: Message = new Message();
  author: User = new User();

  constructor(message:Message, author:User) {
    this.message = message;
    this.author = author;
  }
}

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class MessagesComponent implements OnInit {
  chats: Chat[] = [
    
  ]

  items: Item[] = [
    new Item(
      new Message("Hi! How are you doing?", new Date()),
      new User(1, "Kirill Klimonov", false),
    )
  ];

  constructor(private content: ContentService) { }

  ngOnInit(): void {
    this.content.getChats()?.subscribe(response => {
      console.log(response);
      if (response != false) {
        for (let i = 0; i < response.length; i++) {
          let msg = response[i].last_message

          if (msg != undefined)
            msg.time = new Date(msg.time)
          
          response[i].last_message = msg

          this.chats.push(response[i])
        }
      }
      
    })
  }

}

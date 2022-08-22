import { AfterViewInit, ChangeDetectorRef, Component, ElementRef, Renderer2, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Chat } from 'src/models/chat';
import { Message, SendedMessage } from 'src/models/message';
import { User } from 'src/models/user';
import { AuthService } from '../services/backend-api/auth.service';
import { ContentService } from '../services/backend-api/content.service';
import { WebsocketService } from '../services/backend-api/websocket/websocket.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements AfterViewInit {
  @ViewChild('block') block: ElementRef<HTMLDivElement> = {} as ElementRef;
  // @ViewChild('chatContent') myScrollContainer: ElementRef<HTMLDivElement> = {} as ElementRef;
  @ViewChild('chatContent') myScrollContainer: any;
  @ViewChild('targetView') targetView: ElementRef<HTMLDivElement> = {} as ElementRef;
  @ViewChild('messageArea') messageArea: ElementRef<HTMLTextAreaElement> = {} as ElementRef;
  contentWidth = 0;

  chat: Chat = new Chat(0, "Kirill Klimonov",
    {
        1: new User(1, "Saveliy Karpukhin", true),
        2: new User(2, "Kirill Klimonov", false),
    },
    [
        new Message("I was really cool party, thanks!", new Date(new Date().getSeconds() - 24 * 60 * 60), 1),
        new Message("I glad you've enjoyed :)", new Date(new Date().getSeconds() - 24 * 60 * 60 + 14), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
    ]
  )

  constructor(
      private route: ActivatedRoute, 
      private router: Router, 
      private cdref: ChangeDetectorRef, 
      private content: ContentService, private websocket: WebsocketService, private auth: AuthService) {


  }


  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      const userId = params['user'];
      if (userId == undefined)
        this.router.navigateByUrl('404', {skipLocationChange: true})
      
      this.content.getPersonalChat(userId)?.subscribe(response => {
        if (response == false) {
          this.router.navigateByUrl('/login')
          console.log("Can't get chat");
          return 
        }

        this.chat = Chat.fromDTO(response)
        console.log(this.chat);
        
        this.cdref.detectChanges()

        this.websocket.conect("ws://127.0.0.1:4201/ws/chat_id=" + this.chat.id)

        this.content.getPersonalChatMessages(0, 10, this.chat.id)?.subscribe(response => {
          if (response == false)
            return

          this.chat.messages = response.messages
          
          for (let i = 0; i < this.chat.messages.length; i++) {
            this.chat.messages[i].time = new Date(this.chat.messages[i].time)
          }

          this.cdref.detectChanges()
          })

        this.websocket.message.subscribe(value => {
            value.time = new Date(value.time)
            this.chat.messages.push(value)
        })
      })
    });
  }

  sendMessageClick() {
    let token = this.auth.getToken()
    if (!this.auth.user || !token)
      return

    let message = new SendedMessage()
    
    message.token = token
    message.author_id = this.auth.user.id
    message.text = this.messageArea.nativeElement.value
    message.chat_id = this.chat.id
    console.log(message);
    

    this.websocket.send(message)
  }

  onResize(): void {
    this.contentWidth = this.block.nativeElement.scrollWidth;
    this.cdref.detectChanges();
  }

  ngAfterViewInit(): void {
    this.contentWidth = this.block.nativeElement.scrollWidth;
    this.cdref.detectChanges();
    // this.scrollToBottom();
    let x = this.targetView.nativeElement as HTMLElement;
    setTimeout(() => {
      this.scroll(x);
    }, 20);
    this.targetView.nativeElement.scrollIntoView();
  }

  scroll(el: HTMLElement) {
    el.scrollIntoView();
  }
}

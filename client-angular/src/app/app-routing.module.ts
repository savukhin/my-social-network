import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ChatComponent } from './chat/chat.component';
import { FriendsComponent } from './friends/friends.component';
import { MessagesComponent } from './messages/messages.component';
import { UserPageComponent } from './user-page/user-page.component';

const routes: Routes = [
    {
        path: 'user',
        component: UserPageComponent
    },
    {
      path: 'messages',
      component: MessagesComponent
    },
    {
      path: 'chat/:id',
      component: ChatComponent
    },
    {
      path: 'friends',
      component: FriendsComponent
    },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }